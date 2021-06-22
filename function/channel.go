package function

import "fmt"

func Go_channel() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
	go_channel()
	buf()
}

func go_channel() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func buf() {
	c := make(chan int, 10)
	c <- 101
	fmt.Println(<-c)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
