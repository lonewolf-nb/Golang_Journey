package main

import (
	"fmt"
	"math"
)
var age int
var name string = "John"
var name1 = "Jane"

const (
	Tuesday = 2
	Wednesday = 3
	Thursday int = 4
)
func main(){
	count := 10
	fmt.Println("Random")
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(age)
	fmt.Println(count)

	var maxInt int64 = math.MaxInt64
fmt.Println(maxInt + 1) // Negative due to overflow

var smallFloat float64 = 1.0e-323
fmt.Println(smallFloat / math.MaxFloat64) // Output: 0 due to underflow

for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

numbers := []int{1,2,3,4,5}
for index, value := range numbers{
	fmt.Println("Index: %d, Value: %d\n",index,value);
}

for i := range 10 {
    fmt.Println(10 - i)
}
fmt.Println("Liftoff!")
}