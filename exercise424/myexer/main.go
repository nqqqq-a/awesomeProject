package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type StudentSlice []Student

func (stu StudentSlice) Len() int {
	return len(stu)
}

func (stu StudentSlice) Less(i, j int) bool {
	return stu[i].Score < stu[j].Score
}

func (stu StudentSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var stus StudentSlice
	for i := 0; i < 10; i++ {
		stu := Student{
			fmt.Sprintf("名字~%d", rand.Intn(100)),
			rand.Intn(100),
			math.Trunc(rand.Float64()*1e2 + 0.5),
		}
		stus = append(stus, stu)
	}

	fmt.Println("排序前+", stus)
	fmt.Println("-----------")
	sort.Sort(stus)
	fmt.Println("排序后：", stus)
}
