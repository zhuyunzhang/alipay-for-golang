package payment

//公共参数
type AliPayParameters struct {
	AppID       string `json:"app_id"`       //应用ID
	Method      string `json:"method"`       //接口名称
	Charset     string `json:"charset"`      //编码格式
	SignType    string `json:"sign_type"`    //签名字符串
	Sign        string `json:"sign"`         //签名，生成签名时忽略
	Timestamp   string `json:"timestamp"`    //时间
	Version     string `json:"version"`      //接口版本
	NotifyUrl   string `json:"notify_url"`   //异步回调
	BizContent  string `json:"biz_content"`  //业务参数集合
	Subject     string `json:"subject"`      //商品名称
	OutTradeNo  string `json:"out_trade_no"` //总价
	TotalAmount string `json:"total_amount"` //签名，生成签名时忽略
	ProductCode string `json:"product_code"` //签名类型，生成签名时忽略

}

type TradeNotification struct {
	AppId             string `json:"app_id"`              // 开发者的app_id
	AuthAppId         string `json:"auth_app_id"`         // App Id
	NotifyId          string `json:"notify_id"`           // 通知校验ID
	NotifyType        string `json:"notify_type"`         // 通知类型
	NotifyTime        string `json:"notify_time"`         // 通知时间
	TradeNo           string `json:"trade_no"`            // 支付宝交易号
	TradeStatus       string `json:"trade_status"`        // 交易状态
	TotalAmount       string `json:"total_amount"`        // 订单金额
	ReceiptAmount     string `json:"receipt_amount"`      // 实收金额
	InvoiceAmount     string `json:"invoice_amount"`      // 开票金额
	BuyerPayAmount    string `json:"buyer_pay_amount"`    // 付款金额
	SellerId          string `json:"seller_id"`           // 卖家支付宝用户号
	SellerEmail       string `json:"seller_email"`        // 卖家支付宝账号
	BuyerId           string `json:"buyer_id"`            // 买家支付宝用户号
	BuyerLogonId      string `json:"buyer_logon_id"`      // 买家支付宝账号
	FundBillList      string `json:"fund_bill_list"`      // 支付金额信息
	Charset           string `json:"charset"`             // 编码格式
	PointAmount       string `json:"point_amount"`        // 集分宝金额
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	OutBizNo          string `json:"out_biz_no"`          // 商户业务号
	GmtCreate         string `json:"gmt_create"`          // 交易创建时间
	GmtPayment        string `json:"gmt_payment"`         // 交易付款时间
	GmtRefund         string `json:"gmt_refund"`          // 交易退款时间
	GmtClose          string `json:"gmt_close"`           // 交易结束时间
	Subject           string `json:"subject"`             // 总退款金额
	Body              string `json:"body"`                // 商品描述
	RefundFee         string `json:"refund_fee"`          // 总退款金额
	Version           string `json:"version"`             // 接口版本
	SignType          string `json:"sign_type"`           // 签名类型
	Sign              string `json:"sign"`                // 签名
	PassbackParams    string `json:"passback_params"`     // 回传参数
	VoucherDetailList string `json:"voucher_detail_list"` // 优惠券信息
}
