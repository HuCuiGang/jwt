package jwt

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T){
	jwt := NewJwt("111")
	payload := make(map[string]interface{})
	payload["sub"] = "1234567890"
	payload["name"] = "join"
	payload["admin"] = true
	tokenStr,err := jwt.CreateToken(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("生产tokenStr：",tokenStr)
	token, err := TokenFormatString(tokenStr)
	if err != nil {
		fmt.Println(token)
		return
	}
	fmt.Printf("token:%v ,type:%T \n",token,token)

	err = jwt.VerificationSignature(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("验证通过~")
}