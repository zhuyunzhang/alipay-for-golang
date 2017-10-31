package payment

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"sort"
)

/**
 * 函数目的：获得从参数列表拼接而成的待签名字符串
 * mapBody：是我们从HTTP request body parse出来的参数的一个map
 * 返回值：sign是拼接好排序后的待签名字串。
 */
func genAlipaySignString(mapBody map[string]interface{}) (string, error) {
	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	index := 0
	for _, k := range sorted_keys {
		//if k == "fund_bill_list" {
		//	billList, _ := json.Marshal(mapBody[k])
		//	mapBody[k] = string(billList)
		//}
		value := fmt.Sprintf("%v", mapBody[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value
		}
		//最后一项后面不要&
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}
	fmt.Println("生成的待签名---->", signStrings)
	return signStrings, nil
}

func aliPaySign(mReq map[string]interface{}, priKey []byte) string {
	signStrings, _ := genAlipaySignString(mReq)
	// 开始签名
	block, _ := pem.Decode(priKey)
	if block == nil {
		fmt.Println("rsaSign private_key error")
		return ""
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("无法还原私钥")
		return ""
	}
	result, _ := Rsa2Sign(signStrings, privateKey)
	return result
}

/**
 * RSA2签名
 * @param $data 待签名数据
 * @param $private_key_path 商户私钥文件路径
 * return 签名结果
 */
func Rsa2Sign(origData string, privateKey *rsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

/**
 * RSA签名验证
 * src:待验证的字串，sign:支付宝返回的签名
 * pass:返回true表示验证通过
 * err :当pass返回false时，err是出错的原因
 */
func RSAVerify(src []byte, sign, publicKey []byte) (pass bool, err error) {
	//步骤1，加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	//步骤2，计算代签名字串的SHA1哈希
	t := sha256.New()
	io.WriteString(t, string(src))
	digest := t.Sum(nil)
	//步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	data, _ := base64.StdEncoding.DecodeString(string(sign))
	fmt.Printf("base decoder: %v\n", string(sign))
	fmt.Println("------------------------------------------------------------")
	//步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, digest, data)
	if err != nil {
		return false, err
	}
	return true, nil
}
