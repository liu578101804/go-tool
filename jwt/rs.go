package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func NewRS256() IJwt {

	jwtM := Jwt{}

	//Sha256
	jwtM.SetSignFunc(func(bytes []byte) string {
		h := sha256.New()
		h.Write(bytes)
		return fmt.Sprintf("%x",h.Sum(nil))
	})

	//base64
	jwtM.SetEncodeFunc(func(bytes []byte) string {
		return base64.URLEncoding.EncodeToString(bytes)
	})

	return &jwtM
}
