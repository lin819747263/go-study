package main

import "fmt"

func array1()  {
	var a [5]int
	a[1]=100
	fmt.Println("emp:", a)

	b := [3] string{"666","888","101010"}
	fmt.Println("emp:", b)
}
