# jwt

简单的jwt实现

### jwt是什么

JSON Web token ,通过数字签名的方式，以JSON对象为载体，在不同的服务终端之间安全的传输

### jwt的作用

jwt最常见的场景就是授权认证，一旦用户登录，后续每个请求都将包换jwt，系统在每次处理用户的请求之前，都要先进行jwt安全校验，通过之后再进行处理

### jwt的组成

jwt由3部门组成，用 . 拼接

这三部分分别是：

- Header

```
{
    "typ":"JWT",
    "alg":"HS256"
}
```

- Payload

```
{
    "sub":"12345679890",
    "name":"john",
    "admin":true
}
```

Signature

```
var encodedString = base64urlEncode(header)+'.'+
base64UrlEncode(payload);
var signature = HMACSHA256(encodedString,'secret')
```

JWT Token: `Header Base64` + `.` + `Payload Base64` + `.` + `Signature`

### How To Use

- Install

`go get github.com/HuCuiGang/jwt`

- Use

生成带签名的的token

```
    jwt := NewJwt("111")
    payload := make(map[string]interface{}) //map in Store user information
    payload["sub"] = "1234567890" 
    payload["name"] = "join" 
    payload["admin"] = true 
    tokenStr,err := jwt.CreateToken(payload)
    if err != nil { 
        fmt.Println(err)
        return 
    }
    fmt.Println("生成tokenStr：",tokenStr)
```

将token转换为structure

```
    token, err := TokenFormatString(tokenStr)
	if err != nil {
		fmt.Println(token)
		return
	}
	fmt.Printf("token:%v ,type:%T \n",token,token)
```

验证token签名
```
    err = jwt.VerificationSignature(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("验证通过~")
```






