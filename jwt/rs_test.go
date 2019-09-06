package jwt

import (
	"testing"
)

func TestNewRS256(t *testing.T) {

	jwt := NewRS256()
	jwt.SetHeader("rs256")
	jwt.WriteBody(map[string]interface{}{
		"code": 123,
		"msg": "hello",
	})

	jwtStr := "eyJ0eXBlIjoicnMyNTYiLCJhbGciOiJKV1QifQ==.eyJjb2RlIjoxMjMsIm1zZyI6ImhlbGxvIn0=.6ffa9e290bf336d8f70c079692665f52a2e338a3db179ef55851de1fa4045241"

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