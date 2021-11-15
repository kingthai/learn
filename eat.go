package main

import (
	"fmt"
	"time"
)

func philosopher(chair chan int, chopstick [5]chan int, i int)  {
	for {
		chair <- 1        //scramble for one chair
		chopstick[i] <- i //scramble for left chopstick
		chopstick[(i+1)%5] <- i //scramble for right chopstick

		fmt.Printf("Philosopher(%d) is eating\n", i)

		<-chopstick[(i+1)%5] //release right chopstick
		<-chopstick[i] //release left chopstick
		<-chair        //release chair

		time.Sleep(1*time.Second)
	}
}


func main()  {
	chairs := make(chan int, 4)
	//5 philosophers scramble for 4 chairs

	var chopsticks [5]chan int
	for i:=range chopsticks {chopsticks[i] = make(chan int,1)}
	//create 5 chopsticks

	for i:=0; i<5;i++{go philosopher(chairs, chopsticks,i)}

	time.Sleep(100000000000)
}
