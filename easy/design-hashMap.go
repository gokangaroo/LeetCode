package main

import (
	"container/list"
	"fmt"
)

func main() {
	hashMap := HMConstructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // returns 1
	fmt.Println(hashMap.Get(3)) // returns -1 (not found)
	hashMap.Put(2, 1)
	fmt.Println(hashMap.Get(2)) // returns 1
	hashMap.Remove(2)
	fmt.Println(hashMap.Get(2)) // returns -1 (not found)
}

type MyHashMap struct {
	slot []*list.List
}

type KV struct {
	k int
	v int
}

var slot_num = 10000

/** Initialize your data structure here. */
func HMConstructor() MyHashMap {
	return MyHashMap{make([]*list.List, slot_num)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if this.slot[key/slot_num] == nil {
		this.slot[key/slot_num] = list.New()
		this.slot[key/slot_num].PushBack(KV{key, value})
		return
	}
	for xx := this.slot[key/slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value.(KV).k == key {
			xx.Value = KV{key, value}
			return
		}
	}
	this.slot[key/slot_num].PushBack(KV{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if this.slot[key/slot_num] == nil {
		return -1
	}
	for xx := this.slot[key/slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value.(KV).k == key {
			return xx.Value.(KV).v
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if this.slot[key/slot_num] == nil {
		return
	}
	for xx := this.slot[key/slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value.(KV).k == key {
			this.slot[key/slot_num].Remove(xx)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
