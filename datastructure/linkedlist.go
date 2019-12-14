package datastructure

import (
	"reflect"
	"fmt"
	"errors"
)

//定义节点类型
type ElementType interface {}

//节点结构体
type Node struct{
	Val ElementType //节点值
	Next *Node      //下个节点指针
}

//链表结构体
type LinkedList struct {
	Head *Node //链表的头结点
}

//初始化一个节点
func NewNode(element ElementType,next *Node) *Node{
	return &Node{element,next}
}

//初始化链表
func NewLinkedList() *LinkedList{
	//头结点Val记录链表长度
	head := &Node{0,nil}
	return &LinkedList{head}
}

//判断链表是否为空
func (list *LinkedList) IsEmpty() bool{
	return list.Head.Next == nil
}

//链表的长度
func (list *LinkedList) Length()  int{
	return int(reflect.ValueOf(list.Head.Val).Int())
}

//链表末尾添加元素
func (list *LinkedList) Append(node *Node)  {
	current := list.Head
	for{
		if current.Next == nil{
			break
		}
		current  = current.Next
	}
	current.Next = node
	list.incSize()
}

//链表头部添加元素
func (list *LinkedList) Prepend(node *Node)  {
	current := list.Head
	node.Next = current.Next
	current.Next = node
	list.incSize()
}

//头结点Val++
func (list *LinkedList) incSize()  {
	v := int(reflect.ValueOf(list.Head.Val).Int())
	list.Head.Val = v+1
}

//头结点Val--
func (list *LinkedList) descSize()  {
	v := int(reflect.ValueOf(list.Head.Val).Int())
	list.Head.Val = v-1
}

//查找链表节点
func (list *LinkedList) Find(v ElementType) (*Node,bool){
	//链表是否为空
	empty := list.IsEmpty()
	if empty{
		fmt.Println("List is empty!")
		return nil,false
	}
	current := list.Head
	for current.Next != nil{
		current  = current.Next
		if current.Val == v{
			return current,true
		}
	}
	return nil,false
}

//删除链表指定节点
func (list *LinkedList) Delete(v ElementType) error{
	empty := list.IsEmpty()
	if empty{
		return errors.New("List is empty!")
	}
	current := list.Head
	for current.Next != nil{
		if current.Next.Val == v{
			current.Next = current.Next.Next
			//删除成功 并调整链表size
			list.descSize()
			return nil
		}
		current = current.Next
	}
	return nil
}
