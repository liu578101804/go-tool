package encryption

import (
	"fmt"
	"crypto/sha1"
	"crypto/sha256"
)

func Sha1Encrypt(data []byte) string {
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x\n",bs)
}

func Sha2Encrypt(data []byte) string  {
	h := sha256.New()
	h.Write(data)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x\n",bs)
}
