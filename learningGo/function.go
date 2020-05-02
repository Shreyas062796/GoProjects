package main

import (
	"fmt"
	"strconv"
)
var (
	z int = 10
)
type Object struct {
	Age int
	Name string
	Company string
}
func changeVal(num int) int{
	return(20)
}
func main() {
	fmt.Println("gelo")
	fmt.Println(z)
	z = changeVal(z)
	fmt.Println(z)
	var x int64 = 10
	fmt.Println(x)
	var a int64 = 12
	var b int64 = 2
	fmt.Println(strconv.FormatInt(a,2))
	fmt.Println(strconv.FormatInt(b,2))
	fmt.Println("or :",a | b)
	fmt.Println("and :",a & b)
	fmt.Println("xor :",a ^ b)
	fmt.Println(strconv.FormatInt(x,2))
	fmt.Println(strconv.FormatInt(x << 3,2))
	fmt.Println(strconv.FormatInt(x >> 3,2))
	var newstr [4]string
	for i := 0; i < len(newstr); i++ {
		newstr[i] = strconv.Itoa(i+1)
	}
	fmt.Println(newstr)
	a := []int{1,2,3}
	for i := 0; i < len(a); i++ {
		a[i] *= 4
	}
	fmt.Println(a)
	intToNum(24)
	array := []int{}
	for i := 0; i < 20; i++ {
		array = append(array,i)
	}
	fmt.Println(array)
	testing maps
	testMap := map[int]int{}
	testMap[12] = 13
	testMap[51] = 14

	value, ok := testMap[13]
	if !ok {
		fmt.Println("status:",ok)
	} else {
		fmt.Println("value", value)
	}
	delete(testMap,12)
	for key, value := range testMap {
		_, ok := testMap[key]
		if ok {
			fmt.Println("ok:", ok)
			fmt.Println("key:",key,"value:",value)
		}
	}
	// testing structs
	testObj := Object{
		Age : 13,
		Name : "shreyas",
		Company : "some company",
	}
	fmt.Println(testObj.Name, ":", testObj.Age)
	// switch case testing
	var x interface{} = "string"
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case float64:
		fmt.Println("x is a floating point number")
	case string:
		fmt.Println("x is a string")
	default: 
		fmt.Println("x is another type")
	}
	//While loop
	i := 0
	for {
		i++
		fmt.Println("i :", i)
		if i == 10 {
			break
		}
	}
	for i := 0;i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i,"is odd")
	}
}