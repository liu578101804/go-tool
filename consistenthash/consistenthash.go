package consistenthash

//一致性哈希算法

import (
	"hash/crc32"
	"sort"
	"strconv"
	"errors"
	"sync"
)

var errEmpty = errors.New("hash环没有数据")

//用来存Key，uint32数组
type Uint32Slice []uint32

//获取长度
func (s Uint32Slice) Len() int {
	return len(s)
}

//判断大小
func (s Uint32Slice) Less(i,j int) bool {
	return s[i] < s[j]
}

//交换值
func (s Uint32Slice) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

//定义一个计算hash的算法
type HashFunc func(data []byte) uint32

type Consistent struct {
	//哈希算法
	hashFunc 		HashFunc
	//虚拟节点
	virtualNode 	int
	//排序后的key
	keys 		Uint32Slice
	//哈希环
	hashMap 	map[uint32]string

	//map 读写锁
	sync.RWMutex
}

//创建初始化
func NewConsistent(virtualNode int, fn HashFunc) *Consistent {
	c := &Consistent{
		virtualNode: virtualNode,
		hashFunc: fn,
		hashMap: make(map[uint32]string),
	}
	//默认算法用crc32
	if c.hashFunc == nil {
		c.hashFunc = crc32.ChecksumIEEE
	}
	return c
}

//是否为空
func (c *Consistent) IsEmpty() bool {
	return len(c.keys) == 0
}

//重新排序
func (c *Consistent) updateSortedKeys() {
	tmpKeys := c.keys[:0]
	for k := range c.hashMap {
		tmpKeys = append(tmpKeys, k)
	}
	sort.Sort(tmpKeys)
	c.keys = tmpKeys
}

//生成key
func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

//生成哈希key
func (c *Consistent) hashKey(key string) uint32 {
	if len(key) < 64 {
		var strCatch [64]byte
		copy(strCatch[:], key)
		return c.hashFunc(strCatch[:len(key)])
	}
	return c.hashFunc([]byte(key))
}

//添加节点
func (c *Consistent) Add(keys ...string) {
	c.Lock()
	defer c.Unlock()

	for _,key := range keys {
		//复制虚拟节点
		for i := 0; i< c.virtualNode; i++{
			c.hashMap[c.hashKey(c.generateKey(key,i))] = key
		}
	}
	c.updateSortedKeys()
}

//删除节点
func (c *Consistent) Remove(keys ...string) {
	c.Lock()
	defer c.Unlock()

	for _,key := range keys {
		//复制虚拟节点
		for i := 0; i< c.virtualNode; i++{
			delete(c.hashMap, c.hashKey(c.generateKey(key,i)))
		}
	}
	c.updateSortedKeys()
}

//获取节点
func (c *Consistent) Get(key string) (string,error) {
	c.Lock()
	defer c.Unlock()

	if c.IsEmpty(){
		return "",errEmpty
	}
	hash := c.hashFunc([]byte(key))

	//使用二分法查找最优节点，第一个节点hash值大于对象hash值就是最优节点
	idx := sort.Search(len(c.keys), func(i int)bool {
		return c.keys[i] >= hash
	})

	//如果查找结果大于哈希数组的最大索引
	//表示此时该对象hash值位于最后一个节点之后，那么放入第一个节点中
	if idx == len(c.keys) {
		idx = 0
	}
	return c.hashMap[c.keys[idx]],nil
}

//列表
func (c *Consistent) List() map[uint32]string {
	return c.hashMap
}
