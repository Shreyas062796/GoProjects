package main

import (
	"fmt"
	// "time"
	"sync"
	"runtime"

)
var (
	wg = sync.WaitGroup{}
	mutex = sync.RWMutex{}
	counter int = 0
)

func main() {
	// go doSomething()
	// time.Sleep(100*time.Millisecond)
	// msg := "hello"
	// go func(message string) {
	// 	fmt.Println(message)
	// }(msg)
	// msg = "GoodBye"
	// time.Sleep(100*time.Millisecond)
	// printMessage()
	// driverFunction()
	// driverFunctionWithMutex()
	fmt.Println("max processes :",runtime.GOMAXPROCS(-1))
}

func driverFunction() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go SayHello()
		go increment()
	}
	wg.Wait()
}

func driverFunctionWithMutex() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mutex.RLock()
		go sayHelloWithMutex()
		mutex.Lock()
		go incrementWithMutex()
	}
	wg.Wait()
}
func SayHello() {
	fmt.Println("Hello ", counter)
	wg.Done()
}

func sayHelloWithMutex() {
	fmt.Println("Hello ", counter)
	mutex.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func incrementWithMutex() {
	counter++
	mutex.Unlock()
	wg.Done()
}

func printMessage() {
	msg := "hello"
	wg.Add(1)
	go func(message string) {
		fmt.Println(message)
		wg.Done()
	}(msg)
	msg = "test"
	wg.Wait()
}

func doSomething() {
	fmt.Println("do something")
}