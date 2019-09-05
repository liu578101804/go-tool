package encryption

import (
	"encoding/hex"
	"crypto/md5"
)

func Md5Encrypt(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
