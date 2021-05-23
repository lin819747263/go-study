package main

import "fmt"

func map1()  {
	m:= make(map[string]int)
	m["lgs"]=10
	fmt.Println(m)

	n:=map[string]int{"444":12,"6555":456}
	delete(n,"65555555")
	fmt.Println(n)
}
