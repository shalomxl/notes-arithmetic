package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type InsertSort struct {
	arr []int
}

func Init(n int) *InsertSort {

	var insertsort InsertSort
	rand.Seed(time.Now().UnixNano())
	rand.Intn(10000)
	for i := 0; i < n; i++ {
		m := rand.Intn(10000)
		insertsort.arr = append(insertsort.arr, m)
	}
	return &insertsort
}

func (this *InsertSort) InsertSort() {

	for i := 1; i < len(this.arr); i++ {
		insertval := this.arr[i]
		indexpre := i - 1
		for indexpre >= 0 && insertval < this.arr[indexpre] {
			this.arr[indexpre+1] = this.arr[indexpre]
			indexpre--
		}
		if indexpre+1 != i {
			this.arr[indexpre+1] = insertval
		}
	}
	fmt.Println(this.arr)
}

func main() {
	val := flag.Int("value", 1000, "Init val")
	flag.Parse()
	var insertsort = Init(*val)
	insertsort.InsertSort()
}
