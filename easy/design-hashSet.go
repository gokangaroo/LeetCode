package main

import (
	"container/list"
	"fmt"
)

func main() {
	hashSet := HSConstructor()
	hashSet.Add(1)
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(1)) // return true
	fmt.Println(hashSet.Contains(3)) // return false
	hashSet.Add(2)
	fmt.Println(hashSet.Contains(2)) // return true
	hashSet.Remove(2)
	fmt.Println(hashSet.Contains(2)) // return false
}

type MyHashSet struct {
	slot []*list.List
}

var slot_num = 10000

/** Initialize your data structure here. */
func HSConstructor() MyHashSet {
	return MyHashSet{make([]*list.List, slot_num)}
}

func (this *MyHashSet) Add(key int) {
	if this.slot[key%slot_num] == nil {
		this.slot[key%slot_num] = list.New()
	}
	for xx := this.slot[key%slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value == key {
			return
		}
	}
	this.slot[key%slot_num].PushBack(key)
}

func (this *MyHashSet) Remove(key int) {
	if this.slot[key%slot_num] == nil {
		return
	}
	for xx := this.slot[key%slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value == key {
			this.slot[key%slot_num].Remove(xx)
			return
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.slot[key%slot_num] == nil {
		return false
	}
	for xx := this.slot[key%slot_num].Front(); xx != nil; xx = xx.Next() {
		if xx.Value == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
