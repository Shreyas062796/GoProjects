package main

import (
	"fmt"
	"os"
	"encoding/json"
	"sort"
	"math/rand"
	crypto "crypto/rand"
	"time"
)

type Message struct {
    Name string `json:"name"`
	Body string `json:"body"`
    Time int64 `json:"time"`
}

func main() {
	os.Setenv("test","10")
	fmt.Println(os.Getenv("test"))
	var testArr []int = []int{1,4,6,7,3,4,8}
	var testArrTwo []int = []int{5,1,2,3,6,7,8}
	fmt.Println(reverseArray(testArr))
	sort.Ints(testArr)
	sort.Ints(testArrTwo)
	fmt.Println(testArr)
	fmt.Println(testArrTwo)
	if(len(testArr) != len(testArrTwo)) {
		fmt.Println("The length of the two arrays are different")
		os.Exit(1)
	}
	for i := 0; i < len(testArr); i++ {
		if(checkIfExists(testArr[i],testArrTwo)) {
			fmt.Printf("%d is in both arrays\n",testArr[i])
		}
	}
	m := &Message{Name : "Alice", Body : "Hello",Time : 1294706395881547000}
	b, err := json.Marshal(m)
	if(err != nil) {
		fmt.Println("converting string to json output error")
		fmt.Println(err)
		os.Exit(4)
	}
	res := Message{}
	fmt.Println(string(b))
	json.Unmarshal([]byte(b),&res)
	fmt.Println(res.Name)
	fmt.Println(res.Body)
	fmt.Println(res.Time)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomInt(3,10))
	fmt.Println("your crypto key is ",generateRandomCryptoKey())
	mapTest := map[string]int{"test":1}
	for key,value := range mapTest {
		fmt.Println(key,":",value)
	}
}

func checkIfExists(val int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if(arr[i] == val) {
			return true
		}
	}
	return false
}

func reverseArray(arr []int) []int {
	var temp int
	for i := 0;i < len(arr)/2; i++ {
		temp = arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}
	return(arr)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func generateRandomCryptoKey() [256]uint8 {
	key := [256]byte{}
	_, err := crypto.Read(key[:])
	if(err != nil) {
		panic(err)
	}
	return(key)
}