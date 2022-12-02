package ffun

import "fmt"

func Add(num1, num2 int) int {
	return num1 + num2
}

func Asd() string {
	var n []int
	n = append(n, 1.0)
	n = append(n, 2.0)
	fmt.Println(n)

	m := n
	o := n
	fmt.Println("m", m)
	fmt.Println("o", o)

	m = append(m[0:1], 999)
	o = append(o[0:1], 888, 777)
	fmt.Println("n", n)
	fmt.Println("m", m)
	fmt.Println("o", o)
	return "123"
}
