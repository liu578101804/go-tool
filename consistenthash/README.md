# consistenthash
一致性哈希算法



```
go get github.com/liu578101804/go-tool/consistenthash
```

一致哈希算法在分布式上经常使用，其原理借助了哈希算法的一致性。

- 把所有节点的key通过哈希算法算出一个在环上的数字，放到一个map数组里面，这些数据便组成一个环形
- 把要写入的数据key通过哈希算法算出一个固定的数字，在节点环里面去找相邻的节点进行读写数据。



使用场景：
- 分布式节点写入数据时
