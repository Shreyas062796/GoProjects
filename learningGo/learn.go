package main

import (
       "fmt"
       "strconv"
)
var (
        test string = "cool"
        newstring string = "test"
        nicestring string = "new"
        
)
func findMax(val1,val2 int) int{
        if(val1 > val2) {
	        return(val1)
	} else if(val2 > val1) {
	        return(val2)
	}
	return(0)
}

// func main() {
        // intToNum(10)
//      fmt.Println("hello")
//      fmt.Println(nicestring)
//      var x float64 = 10
//      k := 3
//      a := 1
//      j := 1.1
//      b := 1.2
//      max := findMax(k,a)
//      maxdub := findMaxDouble(j,b)
//      fmt.Println(maxdub)
//      fmt.Println(max)
//      fmt.Println(x)
//      var wid int64 = 100
//      var len int64 = 32
//      fmt.Println(len * wid)
// }

func findMaxDouble(val1,val2 float64) float64{
     if(val1 > val2) {
                return(val1)
        }else if(val2 > val1) {
                return(val2)
        }
         return(0.0)
}

func intToNum(num int) {
        var newstr string = strconv.Itoa(num)
        fmt.Printf("%T : %s\n",newstr,newstr)
}