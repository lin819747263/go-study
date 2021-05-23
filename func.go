package main

import "fmt"

func fun1(a int) int {
	return a
}


func fun2(a int, b int) (int,int) {
	return a,b
}


func fun3(a ...int) {
	fmt.Println(a)
}

func fun4() func() int{
	return func() int {
		return 123
	}
}
