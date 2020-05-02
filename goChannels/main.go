package main

import (
	"fmt"
	"sync"
)
var (
	wg = sync.WaitGroup{}
)

func main() {
	// goChannelOne()
	// goChannelTwo()
	goChannelThree()
}

func goChannelOne() {
	channel := make(chan int)
	wg.Add(2)
	go func() {
		x := <- channel
		fmt.Println(x)
		wg.Done()
	}()
	go func() {
		i := 42
		channel <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()
}

func goChannelTwo() {
	chann := make(chan int)
	for x := 0; x < 5; x++ {
		wg.Add(2)
		go func() {
			i := <- chann
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			chann <- 42
			wg.Done()
		}()
		wg.Wait()
	}
}

func goChannelThree() {
	// second parameter is a buffer where it will be able to store
	//50 integers
	chann := make(chan int, 50)
	wg.Add(2)
	// recieving 42 from the channel
	// recieve only channel
	go func(chann <-chan int) {
		// for a := range chann {
		// 	fmt.Println(a)
		// }
		for {
			if i, ok := <- chann; ok {
				fmt.Println(i)
				continue
			} 
			break
		}
		wg.Done()
	}(chann)
	// sending 42 to the channel
	// send only channel
	go func(chann chan<- int) {
		chann <- 42
		chann <- 27
		close(chann)
		wg.Done()
	}(chann)
	wg.Wait()
}