package encryption

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"errors"
)

//ECB加密
func DESEncryptECB(src, key string) (string,error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "",err
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return "",errors.New("need a multiple of the blockSize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blockSize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out),nil
}

//ECB解密
func DESDecryptECB(src, key string) (string,error) {
	data, err := hex.DecodeString(src)
	if err != nil {
		return "",err
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "",err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return "",errors.New("input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out),nil
}

//CBC加密
func DESEncryptCBC(src, key string) (string,error) {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte )
	if err != nil {
		return "",err
	}
	data = PKCS5Padding(data , block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode .CryptBlocks(out, data)
	return fmt.Sprintf("%X", out),nil
}

//CBC解密
func DESDecryptCBC(src, key string) (string,error) {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		return "",err
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "",err
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext),nil
}


//明文补码算法
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}


//解密
func ThrDESDeCryptCBC(crypt,key []byte)[]byte{
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//创建切片
	context := make([]byte,len(crypt))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,key[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context,crypt)
	//去补码
	context = PKCS5UnPadding(context)
	return context
}

//加密
func ThrDESEnCryptCBC(origData,key []byte)[]byte{
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//补码
	origData = PKCS5Padding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block,key[:8])
	//创建明文长度的数组
	crypt := make([]byte,len(origData))
	//加密明文
	blockMode.CryptBlocks(crypt,origData)
	return crypt
}