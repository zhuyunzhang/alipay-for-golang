// package payment

// import (
// 	"fmt"
// 	"net/http"
// 	"sort"

// 	"strconv"

// 	"io/ioutil"

// 	"net/url"

// 	//"crypto/md5"
// 	//"encoding/hex"
// 	//"encoding/json"
// 	//"strings"

// 	"time"

// 	"qianuuu.com/lib/logs"
// 	"qianuuu.com/player/login/config"
// )

// func getRequest(url string, params map[string]interface{}) (string, error) {
// 	ps := ""
// 	i := 0

// 	sorted_keys := make([]string, 0)
// 	for k, _ := range params {
// 		sorted_keys = append(sorted_keys, k)
// 	}

// 	sort.Strings(sorted_keys)

// 	for _, k := range sorted_keys {
// 		i++
// 		ps = ps + fmt.Sprintf("%s=%v", k, params[k])
// 		if i != len(params) {
// 			ps += "&"
// 		}
// 	}
// 	fmt.Println(url + "?" + ps)
// 	resp, err := http.Get(url + "?" + ps)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(body), nil
// }
// func QueryPayOrder(tradeno string, timestamp time.Time, alipay config.Alipay) (string, error) {

// 	url := "https://openapi.alipay.com/gateway.do?timestamp=" + timestamp.Format("2006-01-02 15:04:05")+"&method=alipay.trade.query&app_id="+alipay.AppID+"&"
// 	body, err := GetRequest(url)
// 	logs.Info("pay content: %s", body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(body), err
// }
// func GetRequest(url string) ([]byte, error) {
// 	//这里是[]byte接收而不是用string接收，就是[]byte是引用类型，string是值类型
// 	logs.Info("get request: %s", url)
// 	defer logs.Info("get request end: %s", url)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return []byte(""), err
// 	}
// 	defer func() {
// 		_ = resp.Body.Close()
// 	}()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	//fmt.Println(string(body))
// 	if err != nil {
// 		return []byte(""), err
// 	}
// 	return body, nil
// }
