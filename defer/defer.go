// [มีการใช้งาน Defer]
package main

import (
	"fmt"
)

func main() {
	printFirst()
	defer printThird()  // -> ลำดับที่ 3
	defer printFourth() // -> ลำดับที่ 2
	defer printFinish() // -> ลำดับที่ 1
	// หากมีหลาย defer มันจะทำตัวที่ประกาศไว้ต่ำที่สุดก่อน
	printSecond()
}
func printFirst() {
	fmt.Println("First")
}
func printSecond() {
	fmt.Println("Second")
}
func printFinish() {
	fmt.Println("Close")
}
func printThird() {
	fmt.Println("Third")
}
func printFourth() {
	fmt.Println("Fourth")
}
