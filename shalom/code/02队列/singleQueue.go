package main

import (
	"errors"
	"fmt"
	"os"
)

type singleQueue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

//	添加数据
func (this *singleQueue) AddQueue(num int) (err error) {
	if this.rear == this.maxSize-1 {
		return errors.New("Queue is full")
	}
	this.rear++
	this.array[this.rear] = num
	return nil
}

//	取数据
func (this *singleQueue) GetQueue() (num int, err error) {
	if this.front == this.rear {
		return -1, errors.New("Queue is empty")
	}

	this.front++
	num = this.array[this.front]
	return num, err
}

//	显示所有数据
func (this *singleQueue) ShowQueue() error {
	if this.front == this.rear {
		return errors.New("Queue is empty")
	}
	fmt.Print("该数组中所有元素：")
	for i,v:=range this.array{

		fmt.Printf("array[%d] = %d\t",i,v)
	}
	fmt.Println()
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("this.array[%d] = %d\n", i, this.array[i])
	}
	return nil
}

func main() {
	singleQueue := singleQueue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			fmt.Println("输入你要添加的数：")
			fmt.Scanln(&num)
			err := singleQueue.AddQueue(num)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功！")
			}
		case "get":
			val, err := singleQueue.GetQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("获取的数据：", val)
			}
		case "show":
			err := singleQueue.ShowQueue()
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
