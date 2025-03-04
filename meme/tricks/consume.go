package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
		ch2 <- "send numbers finish"
		close(ch2)
	}()
	go func() {
		for val := range ch1 {
			fmt.Println(val)
			time.Sleep(time.Second)
		}
		ch3 <- "print finish"
		close(ch3)
	}()
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)

}
