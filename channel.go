package main

import (
	"fmt"
	"time"
)

func channel1()  {
	msg:=make(chan string)
	go func() {
		msg <- "ping"
	}()
	msgR:=<-msg


	fmt.Println(msgR)
}

func channel2()  {
	msg:=make(chan string,2)
	go func() {
		msg <- "ping"
		msg <- "ping2"
		msg <- "ping3"
		msg <- "ping3"
		msg <- "ping3"
	}()
	msgR:=<-msg
	fmt.Println(msgR)
	time.Sleep(time.Second)
	msgR1:=<-msg


	fmt.Println(msgR,msgR1)
}
func channel3()  {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}


func channel4()  {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "99999")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}

func ping(ping chan <- string, msg string)  {
	ping<-msg
}

func pong(ping <-chan string, pong chan <- string)  {
	msg:= <- ping
	pong <- msg
}

func channel5()  {
	c1:=make(chan string)
	c2:=make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for  i:=0; i< 2; i++{
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)

		case msg2:= <- c2:
			fmt.Println(msg2)
		}
	}

}

func channel6()  {
	c1:=make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <- time.After(time.Second * 1):
		fmt.Println("time out")
	}

}
