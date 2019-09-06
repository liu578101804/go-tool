package jwt

import (
	"testing"
)

func TestNewHS256(t *testing.T) {

	jwt := NewHS256([]byte("hello"))
	jwt.SetHeader("hs256")
	jwt.WriteBody(map[string]interface{}{
		"code": 123,
		"msg": "hello",
	})

	jwtStr := "eyJ0eXBlIjoiaHMyNTYiLCJhbGciOiJKV1QifQ==.eyJjb2RlIjoxMjMsIm1zZyI6ImhlbGxvIn0=.020bb7670d2486c1e7edc2511773ab8004c8190d2946c535754c046ba9a24451"

	createJwt,err := jwt.CreateJwtString()
	println(createJwt)

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