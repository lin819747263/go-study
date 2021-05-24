package main

import (
	"fmt"
	"time"
)

func channel1() {
	msg := make(chan string)
	go func() {
		msg <- "ping"
	}()
	msgR := <-msg

	fmt.Println(msgR)
}

func channel2() {
	msg := make(chan string, 2)
	go func() {
		msg <- "ping"
		msg <- "ping2"
		msg <- "ping3"
		msg <- "ping3"
		msg <- "ping3"
	}()
	msgR := <-msg
	fmt.Println(msgR)
	time.Sleep(time.Second)
	msgR1 := <-msg

	fmt.Println(msgR, msgR1)
}
func channel3() {
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

func channel4() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "99999")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func channel5() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func channel6() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("time out")
	}

}

func channel7() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func channel8() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive All job")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func channel9() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
