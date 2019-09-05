package encryption

import "encoding/base64"

func Base64Encrypt(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decrypt(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

func Base64UrlEncrypt(src []byte) string  {
	return base64.URLEncoding.EncodeToString(src)
}

func Base64UrlDecrypt(src string) ([]byte, error)  {
	return base64.URLEncoding.DecodeString(src)
}

