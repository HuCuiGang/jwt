package jwt

type AlgType  string

const (
	HS256 AlgType = "HS256"
	RS256 AlgType = "RS256"

)

type Header struct {
	Typ string
	Alg AlgType
}

type Token struct {
	Header Header
	Payload map[string]interface{}
	Signature string
	tokenStr string
}

