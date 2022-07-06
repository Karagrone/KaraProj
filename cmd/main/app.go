package main

import "fmt"

type Math interface {
	Sum() int
	Min() int
}

type Nums struct {
	n1, n2 int
}

func (nums Nums) Sum() int {
	return nums.n1 + nums.n2
}

func (nums Nums) Min() int {
	return nums.n1 - nums.n2
}

func main() {
	res := Nums{
		n1: 10,
		n2: 8,
	}

	fmt.Println(Math(res))
}
