package main

import (
	"errors"
	"fmt"
)

func f1(a int) (int,error)  {
	if(a == 10){
		return -1,errors.New("error input")
	}
	return a,nil
}

type argsError struct {
	code int
	msg string
}

func (e argsError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}

func f2(a int) (int,error)  {
	if(a == 10){
		return -1,argsError{500,"error input"}
	}
	return a,nil
}
