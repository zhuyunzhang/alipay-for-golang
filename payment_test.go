package payment

import (
	"fmt"
	"net/url"
	"testing"
)

var (
	appID   = "2017102509511228"
	subject = "千游棋牌房卡"

	// RSA2(SHA256) 公钥
	publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvm0tmATQpi9TqGm08MUxyhazd1ijLqGjP+OVrir6TdpTYQr5Jz7TSpKx9Rb6DoTntW9fFtvNIStL4/Ibge6FyUv47bzwaZeYbukdaNrFPSZ6xPrBMLe+sFWDH4qrGJaf4qe+r9wHkH4iXS1gv+p77sVaCxFE6MFYzwIS+UV+1+6rIOMGe+6se8VAOtxZsVUqGr9spBmvHD5RVmEZ+A/OHYxnM+h2pYqs6VyifcXCb9mBvHnrTtTMt4c9fC+w+VGinYtQ/w2nQ55jzwD0YRFAJrzu/Dt8HM4JSuzupCmYuOOw1OeLP+72nPlyonEkDISui48MCXNTD9Kink0Ap62meQIDAQAB
-----END PUBLIC KEY-----`)
	//	publicKey = []byte(`-----BEGIN PUBLIC KEY-----
	//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnxj/9qwVfgoUh/y2W89L6BkRAFljhNhgPdyPuBV64bfQNN1PjbCzkIM6qRdKBoLPXmKKMiFYnkd6rAoprih3/PrQEB/VsW8OoM8fxn67UDYuyBTqA23MML9q1+ilIZwBC2AQ2UBVOrFXfFl75p6/B5KsiNG9zpgmLCUYuLkxpLQIDAQAB
	//-----END PUBLIC KEY-----`)

	// 私钥
	privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAvm0tmATQpi9TqGm08MUxyhazd1ijLqGjP+OVrir6TdpTYQr5Jz7TSpKx9Rb6DoTntW9fFtvNIStL4/Ibge6FyUv47bzwaZeYbukdaNrFPSZ6xPrBMLe+sFWDH4qrGJaf4qe+r9wHkH4iXS1gv+p77sVaCxFE6MFYzwIS+UV+1+6rIOMGe+6se8VAOtxZsVUqGr9spBmvHD5RVmEZ+A/OHYxnM+h2pYqs6VyifcXCb9mBvHnrTtTMt4c9fC+w+VGinYtQ/w2nQ55jzwD0YRFAJrzu/Dt8HM4JSuzupCmYuOOw1OeLP+72nPlyonEkDISui48MCXNTD9Kink0Ap62meQIDAQABAoIBAQCDxrnOgUaCOi4CVWRJWazi1GLNLCGAm4qFI9Do/gTI61TXyugwKGG/MVTE7mmMspxGChQblW+WjIev7lLqz0z1yAUzP5R+/dEWr2sSrJHhh0deGMioFf93tzVOFn/gPBvYlYx31iPF8gOFx2k0Zefti9SL5B9zkpkKZ70JkvX8JpN/ksWwkvW4Vvggi5FRNWpsMvUlgzlVEV9S/QaaXIP/H0exkqFzyb60Ta/3baxLxC+5JEnVR5/wTBQRoubNFIoN1fVxbwFLmOQJsw/p9N6ndraUYG6lyAh7HzSS8ZCxk57azv0QrYvZ074bYgSP7MCSZcOaX6LTHjx6gnRLQEYhAoGBAOgPDSd2h9GERnWXrbootgeEvxN/tTckjHmzS0N7ibtbtiisKGIe6lcWC+q7eAzlFmc1zm/BxbJ+MrJftkWDFh0aZnLhsztMiZmurzWL/dPfeYdc9wWBKW/0V1yIsPQQEfGavtry9JzJEcYZtJaGNujSU41K27DEKrHkOFmcYQVXAoGBANISkVZDoCuZoXhczifLDlak0/NO+Wfjw8pWBXtQ7cPYiGYVcU7J2SzJt5A3i/X5ZW0uF7WR7YVLMKnU3BZW38P/U/5zfUSpwjKm8KWbHH8KImfJC0bALuHYDbzQhQqElA8TPrn3Y5adIXn2VYC75hZKPmWAHuXKf53bo5V9nQCvAoGALtlzD2eHrg3lZ7IymaXEciojpV0gCbzwO1WuOTaErPWsBfQTxxN0vhYuW7pzVy8c4cNkJP3a2tlNhON72fbIDSIaUtEsLSmbkhJJPyc/HHo+f8yN9meIJRkQVhcmmm9wH/Xc2Dk6lzikxPjFk6oPBdwIkDrvtXU1JWrs0XKXx8sCgYEAgKys52kO4AX/mOlHwaooQzw9M2ipblBeKj7cprdgnDiy+8yglgEfjMaWMLlFespjrSexkB8tnRr8WNqwYOKft79a5J47GebdtAb7moTTRKGhh27nAFRRozai24fiJholUsKYBnMZRjVDPyB7KRpvCjI53BRJWLnbx4a0wamqlLMCgYEAoV5qNDqoJq/4g3zs466KXRNckwOwqNS8dcWmBwGJjFXvs7o1U9zydjVTGlXNmc4Ih8KV4dA/raOjWo7Q0QkTy5jlu8rBz4B6vIplOC2v/Elh/Ol6pGl+NcuPkYV9bKoY+ns8bhXB7/sz7vmq1XaNGQHImPkkDC+uLRst6kgplxo=
-----END RSA PRIVATE KEY-----`)
)

func TestMent(t *testing.T) {

	// 这是我从支付宝callback的http body部分读取出来的一段参数列表。
	//paramerStr := `discount=0.00&payment_type=1&subject=%E7%BC%B4%E7%BA%B3%E4%BF%9D%E8%AF%81%E9%87%91&trade_no=2015122121001004460085270336&buyer_email=xxaqch%40163.com&gmt_create=2015-12-21+13%3A13%3A28&notify_type=trade_status_sync&quantity=1&out_trade_no=a378c684be7a4f99be1bf3b56e6d38fd&seller_id=2088121529348920&notify_time=2015-12-21+13%3A17%3A45&body=%E7%BC%B4%E7%BA%B3%E4%BF%9D%E8%AF%81%E9%87%91&trade_status=TRADE_SUCCESS&is_total_fee_adjust=N&total_fee=0.01&gmt_payment=2015-12-21+13%3A13%3A28&seller_email=172886370%40qq.com&price=0.01&buyer_id=2088002578894463&notify_id=5104b719303162e2b79d577aeaa5494jjs&use_coupon=N&sign_type=RSA&sign=YeshUpQO1GsR4KxQtAlPzdlqKUMlTfEunQmwmNI%2BMJ1T2qzd9WuA6bkoHYMM8BpHxtp5mnFM3rXlfgETVsQcNIiqwCCn1401J%2FubOkLi2O%2Fmta2KLxUcmssQ0OnkFIMjjNQuU9N3eIC1Z6SzDkocK092w%2Ff3un4bxkIfILgdRr0%3D`
	//paramerStr := `gmt_create=2017-10-30+10%3A19%3A03&charset=utf-8&seller_email=18715462198%40163.com&subject=fangka&sign=aiLqPMifmCANphNJyiM%2BOnCFos3Rsr6d2oz6mzFtGxNYwaIReuRa%2B5Ex8mvwyRUtzPerXsBBiZL0Oj4tUVRKXihOWU2VFzKfdl3yGQNF9jheIS7%2BCdSg4S%2BGcTRsDBmpwxQ6sr%2Fx%2Fw%2BpbAqbT8LHw5we5P9WoyY%2FvF14lPb6RZtOaHwpnbzGJWN7wUxSlrBTg2e5PsuK0c%2B8exzC1vdsBeoYDLGtiE4RGkjdTUH38srxiVfA%2B5RloZQ3icKlpc7HXbNbIoQP6UoCtWEnfoZUdxlPPDbG%2BrhJMir%2FFtf10M00ncHqQvUqAonENLVebf2LUE9G6k%2BYte1iAbJBYe6pRA%3D%3D&buyer_id=2088602280237831&invoice_amount=0.01&notify_id=354a0b10097fd5e4a862a71340d807cmem&fund_bill_list=%5B%7B%22amount%22%3A%220.01%22%2C%22fundChannel%22%3A%22ALIPAYACCOUNT%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.01&app_id=2017102509511228&buyer_pay_amount=0.01&sign_type=RSA2&seller_id=2088421492900715&gmt_payment=2017-10-30+10%3A19%3A04&notify_time=2017-10-30+10%3A19%3A04&version=1.0&out_trade_no=105053DINGDAN600001078&total_amount=0.01&trade_no=2017103021001004830214635738&auth_app_id=2017102509511228&buyer_logon_id=wpj***%40163.com&point_amount=0.00`
	paramerStr := `gmt_create=2017-10-30 10:19:03&charset=utf-8&seller_email=18715462198@163.com&subject=fangka&sign=aiLqPMifmCANphNJyiM+OnCFos3Rsr6d2oz6mzFtGxNYwaIReuRa+5Ex8mvwyRUtzPerXsBBiZL0Oj4tUVRKXihOWU2VFzKfdl3yGQNF9jheIS7+CdSg4S+GcTRsDBmpwxQ6sr/x/w+pbAqbT8LHw5we5P9WoyY/vF14lPb6RZtOaHwpnbzGJWN7wUxSlrBTg2e5PsuK0c+8exzC1vdsBeoYDLGtiE4RGkjdTUH38srxiVfA+5RloZQ3icKlpc7HXbNbIoQP6UoCtWEnfoZUdxlPPDbG+rhJMir/Ftf10M00ncHqQvUqAonENLVebf2LUE9G6k+Yte1iAbJBYe6pRA==&buyer_id=2088602280237831&invoice_amount=0.01&notify_id=354a0b10097fd5e4a862a71340d807cmem&fund_bill_list=[{"amount":"0.01","fundChannel":"ALIPAYACCOUNT"}]&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.01&app_id=2017102509511228&buyer_pay_amount=0.01&sign_type=RSA2&seller_id=2088421492900715&gmt_payment=2017-10-30 10:19:04&notify_time=2017-10-30 10:19:04&version=1.0&out_trade_no=105053DINGDAN600001078&total_amount=0.01&trade_no=2017103021001004830214635738&auth_app_id=2017102509511228&buyer_logon_id=wpj***@163.com&point_amount=0.00`
	// 字符串转换成 map[string][]string
	values, _ := url.ParseQuery(string(paramerStr))
	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	//
	for k, v := range values {
		//if k == "sign" || k == "sign_type" {
		//	//不要'sign'和'sign_type'
		//	continue
		//}
		m[k] = v[0]
	}
	sign := values["sign"][0]
	strPreSign, _ := genAlipaySignString(m)
	fmt.Println("------------------------------------------------------------")
	//pass, _err := RSA2Verify([]byte(strPreSign), []byte(sign), publicKey)
	pass, _err := RSAVerify([]byte(strPreSign), []byte(sign), publicKey)
	if !pass {
		fmt.Println("verify sig not pass. error:", _err)
	}
	fmt.Println("test end ...")
}
