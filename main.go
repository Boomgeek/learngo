package main

import (
	"fmt"
	"reflect"
	"time"
)

type Weight float64 // custom
type Xx struct {
	x int
}

func main() {

	a := 1.
	s := reflect.TypeOf(a).String()
	fmt.Println(s)

	w := Weight(70) // conversion
	fmt.Println(reflect.TypeOf(w).String())

	i := 1234
	var pi *int = &i // point to i
	fmt.Printf("%v : %v \n", pi, &i)
	fmt.Printf("%v : %v \n", *pi, i)

	i = 4444
	fmt.Printf("%v : %v \n", *pi, i)

	*pi = 2
	fmt.Printf("%v : %v \n", *pi, i) //

	p := &i     // point to i
	*p = *p * 2 // dereferencing
	fmt.Println(*p)

	var mm Xx
	mm.x = 9
	ps := &mm
	fmt.Println(reflect.TypeOf(ps).String() + ": " + reflect.TypeOf((*ps)).String())
	fmt.Println(ps.x == (*ps).x)

	var b [5]int // fixed size
	b[3] = 3     // assignment
	fmt.Println(b)

	c := [...]int{1, 3: 2, 3, 20: -1}
	fmt.Println(c)

	cc := []int{1, 3: 2, 3, 20: -1}
	fmt.Println(cc)

	var d [4][3]int
	fmt.Println(d)

	e := &[32]byte{}
	fmt.Println(e)

	f := []int{1, 2, 10: 3}
	fmt.Println(f)

	sss := make([]int, 0, 9) // สร้าง slice ที่มีความจุ 9

	for i := 0; i < 15; i++ { // เพิ่มข้อมูลเกินความจุที่กำหนดไว้
		sss = append(sss, i)
		fmt.Printf("len=%d cap=%d value=%v\n", len(sss), cap(sss), sss)
	}

	myBytes := []byte("Hello")
	fmt.Println(myBytes)
	// myString := string(myBytes[:])
	// kk := myBytes[:]
	fmt.Println(string(myBytes))

	dst := make([]byte, len(myBytes))
	fmt.Println(dst)
	_ = copy(dst, myBytes)
	fmt.Println(string(dst))

	x := dst[2:5] // elem. 2,3,4
	y := dst[:4]  // elem. < 4
	fmt.Println(string(x))
	fmt.Println(string(y))

	m := make(map[string]int)
	m["key1"] = 42
	fmt.Println("map: ", m)
	fmt.Println("map: ", m["key1"])

	ma := map[string]int{"foo": 1, "bar": 2} // initialize
	fmt.Println("map: ", ma)
	fmt.Println("map: ", ma["bar"])

	_, contains1 := m["key1"]
	_, contains2 := m["key2"]

	fmt.Println("contains1: ", contains1)
	fmt.Println("contains2: ", contains2)

	fmt.Println("length: ", len(ma))
	delete(ma, "foo")
	fmt.Println("ma: ", ma)
	for k, v := range ma {
		fmt.Printf("k: %v, v: %v \n", k, v)
	}
	fmt.Println("length: ", len(ma))

	sx := []string{"a", "b", "c"}
	for idx, val := range sx {
		fmt.Printf("k: %v, v: %v \n", idx, val)
	}

	wd := time.Now().Weekday()
	fmt.Println(wd)
	fmt.Println(reflect.TypeOf(wd).String())
	switch wd {
	case time.Sunday:
		fallthrough
	case 1:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	expensiveFunc := memoizedExpensiveFunc()

	fmt.Println(expensiveFunc(5))  // ครั้งแรกคำนวณ
	fmt.Println(expensiveFunc(5))  // ใช้ค่าจากการเก็บไว้
	fmt.Println(expensiveFunc(10)) // คำนวณอีกรอบ

	pppp := &Person{name: "Bob", age: 4}
	fmt.Println(*pppp)

	ffff := Person{name: "Bob", age: 4}
	fmt.Println(ffff)

	// p.age = 30
	// p.Aging(1)

}

func memoizedExpensiveFunc() func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if val, ok := cache[n]; ok {
			fmt.Println("From cache")
			return val
		}

		result := n * n
		cache[n] = result
		fmt.Println("Calculated")
		return result
	}
}

type Person struct {
	name string
	age  int
}

func (p *Person) Axxxx(abc int) {
	p.age = p.age + abc
}
