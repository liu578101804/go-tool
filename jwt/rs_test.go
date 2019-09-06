package jwt

import (
	"testing"
)

func TestNewRS256(t *testing.T) {

	jwt := NewRS256()
	jwt.SetHeader("rs256","jwt")
	jwt.WriteBody(map[string]interface{}{
		"code": 123,
		"msg": "hello",
	})

	jwtStr := "eyJUeXBlIjoicnMyNTYiLCJBbGciOiJqd3QifQ==.eyJjb2RlIjoxMjMsIm1zZyI6ImhlbGxvIn0=.8a9d732e4a6f49081100471dfa0c2974231d6b3eb822dbc05bbba1268770bce8"

	createJwt,err := jwt.CreateJwtString()
	if err !=  nil{
		t.Error(err)
	}

	if createJwt != jwtStr {
		t.Error("create jwt string error")
	}

	if !jwt.CheckJwtString(jwtStr) {
		t.Error("jwt string error")
	}

	println("success")
}