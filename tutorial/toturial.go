package tutorial

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

func pointerTest() {
	data := 10
	a := new(int) //same var b *int
	var b *int
	a = &data // เก็บ address ของ data ไว้ใน value ของ a
	// & ใส่ด้านหน้าเพื่อ get ค่า address ของตัวแปลนั้นๆออกมา
	fmt.Print(data, a)
	fmt.Print(*a) //เรียกดูค่าของตัวแปล address ที่มัน point ไป
	fmt.Print(b)
}

func doublePointerTest() {
	data := 10
	a := new(int)  //same var b *int
	a = &data      // เก็บ address ของ data ไว้ใน value ของ a
	b := new(*int) //double pointer same ,var c **int
	var c **int    //double pointer เก็บค่า address ของ pointer อีกที
	b = &a
	// & ใส่ด้านหน้าเพื่อ get ค่า address ของตัวแปลนั้นๆออกมา
	fmt.Print(data, a, b, c)
	fmt.Print(*a) //เรียกดูค่าของตัวแปล address ที่มัน point ไป
}
func double(n []int) []int {
	// ด้วย slice เป็น การ copy ค่าเป็น pass by referance ทำให้ ค่า n ที่รับมาที่ parameter เป็นค่าเดี๋ยวกับที่ส่งเข้ามา ถ้าแก้ไขค่า n
	//ภายใน func นี้ต่อให้ไม่ return s1 จากภายนอกที่ส่งเข้ามาค่าจะเปลี่ยนไปด้วย
	temp := make([]int, len(n))
	for i, v := range n {
		temp[i] = v * 2
	}
	return temp
}

func test() {
	s1 := []int{1, 2, 3}
	s2 := double(s1)

	fmt.Print(s1)
	fmt.Print(s2)
}

type Player struct {
	//ใน struct แต่ละ filed มี address เป็นของตัวเอง
	Username string
	Level    int
}

func (p Player) upLevel() {
	p.Level += 1
}

func (p *Player) upLevelPointer() {
	p.Level += 1
}

func usePointerInStruct() {
	//ใน struct แต่ละ filed มี address เป็นของตัวเอง
	// p1 address 101
	p1 := Player{
		// Username address 1001
		Username: "player1",
		// Level address 1002
		Level: 1,
	}
	// p1 address 102
	p2 := Player{
		// Username address 2001
		Username: "player1",
		// Level address 2002
		Level: 10,
	}
	p1.upLevel()

	p2.upLevel()
	// เมื่อใช้ค่า p1 p2 เลยไม่เปลี่ยน ✗

	// ถ้าอยากให้ค่าเปลี่ยนด้วยให้ประกาศ struct ตอนสร้างเป็น point ดังนี้
	p3 := &Player{
		// Username address 1001
		Username: "player1",
		// Level address 1002
		Level: 1,
	}
	// p1 address 102
	p4 := &Player{
		// Username address 2001
		Username: "player1",
		// Level address 2002
		Level: 10,
	}
	p3.upLevelPointer()
	p4.upLevelPointer()

	// เมื่อใช้ค่า p3 p4 เลยเปลี่ยน ✓
}
