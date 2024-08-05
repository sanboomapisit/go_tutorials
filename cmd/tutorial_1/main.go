package main

import (
	"fmt"
	"time"
)

func buy1() {
	time.Sleep(1 * time.Second)
	fmt.Println("helloworld 1!!")
}
func buy2() {
	time.Sleep(1 * time.Second)
	fmt.Println("helloworld 2!!")
}
func buy3() {
	time.Sleep(1 * time.Second)
	fmt.Println("helloworld 3!!")
}
func buy4() {
	time.Sleep(1 * time.Second)
	fmt.Println("helloworld 4!!")
}

func printSchoolAddress(schoolAddress string) {
	println(schoolAddress)
}

func getSchoolAddress() string {
	return "กรุงเทพ"
}

func getSchoolAddress2() (int, string) {
	code := 1993
	address := "กรุงเทพ"
	return code, address
}

func main() {
	var a, b, c string
	var a0, b0, c0 bool = false, false, true
	var a1, b1, c1 = false, 10, "เรียนภาษา Go"
	// Shorthand
	a2, b2, c2 := false, 10, "เรียนภาษา Go"
	start := time.Now()
	buy1()
	go buy2()
	go buy3()
	buy4()
	printSchoolAddress("Test")
	schoolAddress := getSchoolAddress()
	resultCode, resultAddress := getSchoolAddress2()
	fmt.Println(schoolAddress)
	fmt.Println(resultCode, resultAddress)
	fmt.Println(a, b, c)
	fmt.Println(a0, b0, c0)
	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)

	// [For Loop]
	for i := 1; i < 9; i++ {
		println(i)
	}

	i := 1
	// [Do While Loop]
	for {
		if i == 2 {
			println(i)
			break
		}
		i++
	}

	// [While Loop]
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j = j + 1
	}

	// [if else แบบปกติ]
	x := 10
	if x < 10 {
		println("x มีค่าน้อยกว่า 10")
	} else {
		println("x มีค่ามากว่า หรือ เท่ากับ 10")
	}
	// [if else if]
	x1 := 10
	if x1 < 10 {
		println("x มีค่าน้อยกว่า 10")
	} else if x1 == 10 {
		println("x มีค่าเท่ากับ 10")
	} else {
		println("x มีค่ามากกว่า 10")
	}

	s := 2
	h := 3
	k := 0
	// [ท่า process ก่อน check เงื่อนไข]
	if k = s + h; k == 5 {
		println("k มีค่าเท่ากับ 5")
	}

	// [การประกาศ Array เปล่าๆ ว่างๆ ]
	var q [5]int //นี้ก็หมายความว่า a เป็น type int ที่มี Block การใส่ข้อมูล 5 ช่อง
	// [การใส่ค่าใน Array แต่ละช่อง]
	q[0] = 1
	q[1] = 2
	q[2] = 3

	// [การประกาศ Array แบบมีค่าเริ่มต้น]
	var name = [3]string{"Chaiyarin", "Atikom", "Kritsana"}
	fmt.Println("array :", name, "lenght = ", len(name))

	// Slice
	// คล้าย array แต่ไม่จำเป็นที่จะต้อง “กำหนดขนาดล่วงหน้า”
	// [การประกาศตัวแปร Slice]
	nameSlice := []string{}
	// [การใส่ค่าใน Slice]
	nameSlice = append(nameSlice, "Chaiyarin")
	nameSlice = append(nameSlice, "Atikom")
	nameSlice = append(nameSlice, "Kritsana")
	fmt.Println("slice :", nameSlice, "lenght = ", len(nameSlice))

	// Map คือ Data Structure รูปแบบหนึ่ง ที่เก็บข้อมูลในลักษณะ Key Value
	m := make(map[string]int)
	m["chaiyarin"] = 1
	m["atikom"] = 2
	m["kritsana"] = 3
	fmt.Println("map :", m, "lenght = ", len(m))

	delete(m, "atikom")
	fmt.Println("หลังลบ map :", m, "lenght = ", len(m))

	fmt.Println("ใช้เวลา :", time.Since(start), ": วินาที")
}
