package simulateHashmap

import (
	"fmt"
	"log"
)

const BucketCount = 16

type HashMap struct {
	Buckets [BucketCount]*LinkNode
}

func CreateHashMap() *HashMap {
	myMap := &HashMap{}
	for i := 0; i < BucketCount; i++ {
		myMap.Buckets[i] = CreateLink()
	}

	return myMap
}

func (myMap *HashMap) AddKeyValue(key string, value string) {

	//1.将key散列成0-BucketCount的整数作为Map的数组下标
	var mapIndex = HashCode(key)
	//2.获取对应数组头结点
	var link = myMap.Buckets[mapIndex]
	log.Print(link.Data.Key == "" && link.NextNode == nil)
	//3.在此链表添加结点
	if link.Data.Key == "" && link.NextNode == nil {
		//如果当前链表只有一个节点，说明之前未有值插入  修改第一个节点的值 即未发生哈希碰撞
		link.Data.Key = key
		link.Data.Value = value

		fmt.Printf("node key:%v add to buckets %d first node\n", key, mapIndex)
	} else {
		//发生哈希碰撞
		index := link.AddNode(KV{key, value})
		fmt.Printf("node key:%v add to buckets %d %dth node\n", key, mapIndex, index)
	}
}

//按键取值
func (myMap *HashMap) GetValueForKey(key string) string {
	//1.将key散列成0-BucketCount的整数作为Map的数组下标
	var mapIndex = HashCode(key)
	//2.获取对应数组头结点
	var link = myMap.Buckets[mapIndex]

	var value string

	//遍历找到key对应的节点
	head := link
	for {
		if head.Data.Key == key {

			value = head.Data.Value
			break
		} else {

			head = head.NextNode
		}
	}

	return value
}

func HashCode(key string) int {

	var sum = 0
	for i := 0; i < len(key); i++ {

		sum += int(key[i])
	}

	return sum % BucketCount
}
