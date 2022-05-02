package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计英文、数字、空格和其他字符数量
//说明:统计一个文件中含有的英文、数字、空格及其它字符数量。

type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

func main() {
	//思路：打开一个文件，创建一个Reader，每读取一行，就遍历该行有多少个这样的数字
	//将其保存到一个结构体中
	fileName := "E:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	//定义这个CharCount实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str, 进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("字符的个数为=%v    数字的个数为%v   空格的个数为%v   其他字符的个数为%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}
