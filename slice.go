package main

import "fmt"

func slice1()  {
	s:=make([]string,3)
	s = append(s, "123","6666","8888","66666")
	fmt.Println(s)

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println(c)
}
