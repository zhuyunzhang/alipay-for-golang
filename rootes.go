package payment

import (
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"strconv"

	"net/url"

	"github.com/labstack/echo"
	"qianuuu.com/lib/echotools"
	"qianuuu.com/lib/logs"
	"qianuuu.com/lib/values"
	"qianuuu.com/player/cache"
	"qianuuu.com/player/login/config"
	"qianuuu.com/player/login/plugin/cfgm"
	"qianuuu.com/player/login/plugin/gapi"
	"qianuuu.com/player/usecase"
)

var (
	uc        *usecase.Usecase
	cacheUser cache.UserCache
	gapic     *gapi.Client
)

// Routes 微信支付路由功能
func Routes(e *echo.Echo, usecase *usecase.Usecase, cache cache.UserCache, apic *gapi.Client) {
	// TODO: 初始化微信配置参数

	uc = usecase
	cacheUser = cache
	gapic = apic
	// e.Get("/pay/products", UserProducts)
	e.Get("/pay/alipay/users/:id/:gid", UserPay)
	//e.Get("/pay/alipay/check/:tradeno", CheckNotify)同步回调支付宝订单查询接口
	e.Post("/pay/alipay/notify", UserPayBackURL)
}

func getOps() config.Alipay {
	return config.Opts.Alipay
}

//      ┏┛ ┻━━━━━┛ ┻┓
//      ┃　　　　　　 ┃
//      ┃　　　━　　　┃
//      ┃　┳┛　  ┗┳　┃
//      ┃　　　　　　 ┃
//      ┃　　　┻　　　┃
//      ┃　　　　　　 ┃
//      ┗━┓　　　┏━━━┛
//        ┃　　　┃   神兽保佑
//        ┃　　　┃   代码无BUG！
//        ┃　　　┗━━━━━━━━━┓
//        ┃　　　　　　　    ┣┓
//        ┃　　　　         ┏┛
//        ┗━┓ ┓ ┏━━━┳ ┓ ┏━┛
//          ┃ ┫ ┫   ┃ ┫ ┫

// UserPay 用户支付接口
func UserPay(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	uid := t.ParamInt("id")
	gid := t.ParamString("gid")

	logs.Info("alipay uid: %v gid: %v", uid, gid)

	// TODO: 权限配置
	alipay := getOps()

	// 更新策略
	//goods := cfgm.GetConfig("products")
	//if goods == nil {
	//	logs.Info("没有找到 products 配置文件")
	//	return t.BadRequest("支付失败")
	//}

	products, err := uc.ConfigFile(uid, "products")
	if err != nil {
		return t.BadRequest("获取配置文件失败")
	}
	confg := products.GetString("value")
	goods, _ := values.NewValueMapArray([]byte(confg))
	if goods == nil {
		return t.BadRequest("获取配置文件失败")
	}

	_ = alipay

	// fmt.Println(channels)
	// return t.OK("")

	var good values.ValueMap
	for _, g := range goods {
		if g.GetString("id") == gid {
			good = g
		}
	}
	if good == nil {
		return t.BadRequest("无效的商品类型")
	}
	//alipay.priKey商品应用私钥
	var priKey = []byte(fmt.Sprintf(`-----BEGIN RSA PRIVATE KEY-----
%s
-----END RSA PRIVATE KEY-----`, alipay.priKey))
	fee := good.GetInt("fee")

	// 本地数据库中创建订单
	order, err := uc.CreateOrder(uid, fee, "alipay", "client", good.GetInt("count"), good.GetInt("gift"), uid)
	fmt.Println(order)
	if err != nil {
		return t.BadRequest("支付请求失败！")
	}
	ProductCode := "QUICK_MSECURITY_PAY"
	Biz := map[string]interface{}{
		"subject":      alipay.Subject,
		"out_trade_no": order.TradeNo,
		"total_amount": fmt.Sprintf("%.2f", float64(fee)/100),
		"product_code": ProductCode,
	}
	biz, _ := json.Marshal(Biz)
	m := make(map[string]interface{})
	m["app_id"] = alipay.AppID
	m["biz_content"] = string(biz)
	m["charset"] = "utf-8"
	m["method"] = "alipay.trade.app.pay"
	m["sign_type"] = "RSA2"
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = alipay.NotifyURL

	ret := map[string]interface{}{
		"app_id":         alipay.AppID,
		"method":         "alipay.trade.app.pay",
		"charset":        "utf-8",
		"sign_type":      "RSA2",
		"sign":           aliPaySign(m, priKey),
		"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
		"version":        "1.0",
		"notify_url":     alipay.NotifyURL,
		"biz_content":    string(biz),
		"iosbiz_content": url.QueryEscape(string(biz)),
	}
	return t.OK(ret)
}



//异步回调支付宝验证签名接口
func UserPayBackURL(c echo.Context) error {
	t := echotools.NewEchoTools(c)

	content, err := t.BodyText()
	if err != nil {
		return t.BadRequest(err.Error())
	}
	alipay := getOps()
	fmt.Println(content)
	values_m, _err := url.ParseQuery(string(content))
	if _err != nil {
		fmt.Println("error parse parameter, reason:", _err)
		return t.BadRequest(err.Error())
	}
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	for k, v := range values_m {
		if k == "sign" || k == "sign_type" {
			//不要'sign'和'sign_type'
			continue
		}
		m[k] = v[0]
		fmt.Println(v)
	}
	sign := values_m["sign"][0]
	strPreSign, _ := genAlipaySignString(m)
	//publicKey指的是支付宝的公钥，不是商品的应用公钥
	publicKey := []byte(fmt.Sprintf(`-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`, alipay.PublicKey))
	pass, _err := RSAVerify([]byte(strPreSign), []byte(sign), publicKey)
	if !pass {
		return t.BadRequest("签名验证失败！")
	}
	outtradeno := values_m["out_trade_no"][0]
	tradestatus := values_m["trade_status"][0]
	gmtcreate := values_m["gmt_create"][0]
	// 根据 orderid 查询数据记录
	if err != nil {
		return t.BadRequest("未找到订单！")
	}
	// tradestatus 是否为 TRADE_SUCCESS
	if tradestatus == "TRADE_SUCCESS" {
		//验证成功编写成功的方法

	} else {
		return t.BadRequest("交易失败！")
	}
	return c.String(http.StatusOK, "success")
}
