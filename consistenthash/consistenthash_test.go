package consistenthash

import (
	"testing"
	"fmt"
)


func TestNew(t *testing.T) {
	var node string
	hashMap := NewConsistent(2,nil)

	for i := 0; i< 100; i++ {
		go func() {

			//添加节点
			hashMap.Add("127.0.0.1","127.0.0.2","127.0.0.3","127.0.0.4")

			node ,_ = hashMap.Get("127.0.0.2")
			fmt.Println("127.0.0.2=>"+node)
			node ,_ = hashMap.Get("127.0.0.4")
			fmt.Println("127.0.0.4=>"+node)

			fmt.Println(len(hashMap.List()))
			fmt.Println(hashMap.List())

			//移除节点
			hashMap.Remove("127.0.0.1","127.0.0.2")

			node ,_ = hashMap.Get("127.0.0.2")
			fmt.Println("127.0.0.2=>"+node)
			node ,_ = hashMap.Get("127.0.0.4")
			fmt.Println("127.0.0.4=>"+node)

			fmt.Println(len(hashMap.List()))
			fmt.Println(hashMap.List())

		}()
	}




}