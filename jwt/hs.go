package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"crypto/hmac"
	"io"
	"fmt"
)

func NewHS256(privateKey []byte) IJwt {

	jwtM := Jwt{}

	//HMac
	jwtM.SetSignFunc(func(bytes []byte) string {
		h := hmac.New(sha256.New, privateKey)
		io.WriteString(h, string(bytes))
		return fmt.Sprintf("%x",h.Sum(nil))
	})

	//base64
	jwtM.SetEncodeFunc(func(bytes []byte) string {
		return base64.URLEncoding.EncodeToString(bytes)
	})

	return &jwtM
}
