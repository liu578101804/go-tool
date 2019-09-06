package jwt

import (
	"testing"
)

func TestNewHS256(t *testing.T) {

	jwt := NewHS256([]byte("hello"))
	jwt.SetHeader("hs256","jwt")
	jwt.WriteBody(map[string]interface{}{
		"code": 123,
		"msg": "hello",
	})

	jwtStr := "eyJUeXBlIjoiaHMyNTYiLCJBbGciOiJqd3QifQ==.eyJjb2RlIjoxMjMsIm1zZyI6ImhlbGxvIn0=.36c2b24a09d744bceb16eb5aff8cf4a00cd63fb56c5ffcb225eb319c2cc0bf05"

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