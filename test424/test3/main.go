package main

//3)编写一个方法算该矩形的面积(可以接收长len，和宽width)，将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印。

type MethodUtils struct {
}

func (mu MethodUtils) getArea(len, width float64) float64 {
	return width * len
}

func main() {
	println(MethodUtils{}.getArea(2.0, 3.5))
}
