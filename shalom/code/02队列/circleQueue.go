package main

import (
	"fmt"
	"os"
	"errors"
)

type circleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

//	是否为空
func (this *circleQueue) isEmpty() bool {
	return this.tail == this.head
}

//	是否满了
func (this *circleQueue) isFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//	显示总元素个数
func (this *circleQueue) Size() (size int) {
	size = (this.tail + this.maxSize - this.head) % this.maxSize
	return
}

//	入队列
func (this *circleQueue) Push(num int) (err error) {
	if this.isFull() {
		return errors.New("Queue is full ! Add fail !")
	}

	this.array[this.tail] = num
	this.tail = (this.tail + 1) % this.maxSize
	return nil
}

//	出队列
func (this *circleQueue) Pop() (val int, err error) {
	if this.isEmpty() {
		return 0, errors.New("Queue is empty !")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

//	查看队列所有元素
func (this *circleQueue) ShowAll() error {
	if this.isEmpty() {
		return errors.New("Queue is empty !")
	}
	tem := this.Size()
	for i := 0; i < tem; i++ {
		fmt.Printf("array[%d] = %d\t", i, this.array[(this.head+i)%this.maxSize])
	}
	fmt.Println()
	return nil
}

func main() {
	circleQueue := circleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var num int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")
		fmt.Print("输入你的选择：")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Print("输入你要添加的数：")
			fmt.Scanln(&num)
			err := circleQueue.Push(num)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功！")
			}
		case "get":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("获取的数据：", val)
			}
		case "show":
			err := circleQueue.ShowAll()
			if err != nil {
				fmt.Println(err)
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("请输入正确的指令:")
		}
	}
}
