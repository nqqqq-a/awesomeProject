package main

import "fmt"

func sum(n1, n2 int) int {
	//defer会将后面的语句先压入独立的栈(defer栈)
	//当函数执行完之后，才会开始释放栈（且此时的变量会也被拷贝压入栈中，变量在之后的变化，不会影响这个独立栈内的变量值）
	//相当于冻结了这个语句，最后释放出来，执行结果按照其原先的变量情况
	defer fmt.Println("ok1, n1 = ", n1)
	defer fmt.Println("ok2, n2 = ", n2)
	res := n1 + n2
	fmt.Println(res)
	return res

	//defer的价值在于 defer file.close()、 defer connect.close() 在我们打开资源的时候，就可以这样写，防止最后我们忘记关闭资源
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)
}
