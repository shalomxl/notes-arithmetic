package main

import (
	"fmt"

)

type HeroNode struct {
	num      int
	name     string
	nickName string
	next     *HeroNode
}
//	类似于你因为一个标记想起了一件事，又因为这件事的某个细节想起了另外一件事，单链表的数据储存结构
//	傻瓜式添加，从最后一个坑添加
func (head *HeroNode) AddinTail( /*head *HeroNode,*/ newHeroNode *HeroNode) {
	tem := head
	for {
		if tem.next == nil {
			break
		}
		tem = tem.next
	}
	tem.next = newHeroNode
}

//	按照序号添加
func (head *HeroNode) Add(newHeroNode *HeroNode) {
	tem := head
	flag := true
	for {
		if tem.next == nil {
			break
		} else if tem.next.num == newHeroNode.num {
			fmt.Println("========------------->")
			flag = false
			break
		} else if tem.next.num > newHeroNode.num {
			break
		}
		tem = tem.next
	}

	if !flag {
		fmt.Printf("对不起，编号%d已存在", newHeroNode.num)
		return
	} else {
		newHeroNode.next = tem.next
		tem.next = newHeroNode
	}
}

//	显示链表所有信息
func (head *HeroNode) Show() {
	tem := head

	for {
		if tem.next == nil {
			break
		}
		fmt.Printf("[%d , %s , %s]==>", tem.next.num, tem.next.name, tem.next.nickName)
		tem = tem.next
	}
}

//	删除
func (head *HeroNode) Delete(num int) {
	tem := head
	flag := false

	for {
		if tem.next == nil {
			break
		}else if tem.next.num == num{
			flag = true
			break
		}
		tem = tem.next
	}
	if flag{
		tem.next = tem.next.next
	}else {
		fmt.Println("sorry , 要删除的id不存在")
	}
}

func main() {

	
	//1. 先创建一个头结点,
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		num:      1,
		name:     "宋江",
		nickName: "及时雨",
	}

	hero2 := &HeroNode{
		num:      2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		num:      3,
		name:     "林冲",
		nickName: "豹子头",
	}
	//
	//hero4 := &HeroNode{
	//	num:      3,
	//	name:     "吴用",
	//	nickName: "智多星",
	//}

	//head.AddinTail(hero1)
	//head.AddinTail(hero2)
	head.Add(hero1)
	head.Add(hero2)
	head.Add(hero3)

	head.Show()

	fmt.Println("")
	head.Delete(2)
	head.Show()
}
