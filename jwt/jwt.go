package jwt

import (
	"encoding/json"
	"strings"
	"fmt"
)

type IJwt interface {
	//设置头部
	SetHeader(string)
	//设置签名算法
	SetSignFunc(SignFunc)
	//设置编码算法
	SetEncodeFunc(EncodeFunc)

	//写入body
	WriteBody(map[string]interface{})

	//生成jwt
	CreateJwtString() (string,error)
	//验证jwt
	CheckJwtString(string) bool
}

type Header struct {
	Type 	string 	`json:"type"`
	Alg 	string	`json:"alg"`
}

//签名算法
type SignFunc func([]byte) string
//编码算法
type EncodeFunc func([]byte) string

type Jwt struct {
	Header 		Header
	Body 		map[string]interface{}

	signFun 	SignFunc
	encodeFun 	EncodeFunc
}

func (j *Jwt) SetHeader(headerType string){
	j.Header =  Header{
		Type: headerType,
		Alg: "JWT",
	}
}

func (j *Jwt) SetSignFunc(signFunc SignFunc) {
	j.signFun = signFunc
}

func (j *Jwt) SetEncodeFunc(encodeFunc EncodeFunc) {
	j.encodeFun = encodeFunc
}

func (j *Jwt) WriteBody(body map[string]interface{}) {
	j.Body = body
}

func (j *Jwt) CreateJwtString() (string,error) {
	//编码header
	headerByte,err := json.Marshal(j.Header)
	if err != nil {
		return "",err
	}
	headerStr := j.encodeFun(headerByte)

	//编码body
	bodyByte,err := json.Marshal(j.Body)
	if err != nil {
		return "",err
	}
	bodyStr := j.encodeFun(bodyByte)

	//签名
	signByte := j.signFun([]byte(string(headerStr)+"."+string(bodyStr)))

	return fmt.Sprintf("%s.%s.%s",headerStr,bodyStr,signByte),nil
}

func (j *Jwt) CheckJwtString(input string) bool  {
	arr := strings.Split(input,".")
	//格式是否正确
	if len(arr) != 3 {
		return false
	}
	//签名
	signByte := j.signFun([]byte(string(arr[0])+"."+string(arr[1])))
	if string(signByte) != arr[2] {
		return false
	}
	return true
}