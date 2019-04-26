package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type ValNode struct {
	row  int
	list int
	val  interface{}
}

//	01稀疏数组，黑白棋存档
//	思想：为了节省空间，将不重要的数据省略，提取并记录重要的数据
//	生活中有很多例子：因为人脑的惰性，某一品类的事物，一般只能记忆3种，例如：人们会常常记得班级前三名和最后三名，中间的几乎难以记住；再比如：对于牛奶，人们只能够记住“伊利”，“蒙牛”，“三元”，其他的品牌就难以记住，人的大脑好像自动就设置了这样的程序，实在令人惊讶
func main() {
	var chessMap [11][11]int
	chessMap[3][5] = 1
	chessMap[6][4] = 2

	for _, v := range chessMap {
		for _, val := range v {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}

	var sparseArr []ValNode
	valnode := ValNode{
		11,
		11,
		0,
	}
	sparseArr = append(sparseArr, valnode)

	for i, v := range chessMap {
		for j, val := range v {
			if val != 0 {
				valnode = ValNode{
					i,
					j,
					val,
				}
				sparseArr = append(sparseArr, valnode)
			}
		}
	}

	f, err := os.Create("chessMap.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range sparseArr {
		f.WriteString(string(v.row) + string(v.list) + string(v.val.(int)) + "\n")
	}

	fl, err := os.Open("chessMap.data")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(fl)

	var newChessMap [][]int
	buf, _ := r.ReadBytes('\n')

	newChessMap = make([][]int, buf[0])

	for i := 0; i < len(newChessMap); i++ {
		newChessMap[i] = make([]int, buf[1])
	}

	for {
		buf, err := r.ReadBytes('\n')

		if err == io.EOF {
			break
		}
		if buf[0] != 11 && buf[1] != 11 {
			newChessMap[buf[0]][buf[1]] = int(buf[2])
		}
	}

	for _, v := range newChessMap {
		for _, val := range v {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
