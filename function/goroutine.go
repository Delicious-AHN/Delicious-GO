package function

import (
	"fmt"
	"sync"
	"time"
)

func Go_routine() {
	go1()
	go2()
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func go1() {
	say("SYNC")

	go say("ASYNC1")
	go say("ASYNC2")
	go say("ASYNC3")

	time.Sleep(time.Second * 3)
}

func go2() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("HELLO")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}
