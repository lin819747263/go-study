package main

import "fmt"

func range1()  {
	nums:=[]int{12,16,46}
	sum:=0
	for _,num := range nums{
		sum +=num
	}
	fmt.Println(sum)


	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for _,c:= range "linguoshuai"{
		fmt.Println(c)
	}
}