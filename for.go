package main

import "fmt"

var count = 0

func for1() {
	i:=1
	for i<3 {
		fmt.Println(i)
		i++
	}
}

func for2() {
	for j:=0;j<6;j++ {
		fmt.Println(j)
	}
}

func for3() {
	for{
		count+=1
		fmt.Println(count)
		if count > 10 {
			break
		}
	}

}

func if1()  {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
