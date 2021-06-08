/*
*
*
*
*
 */

// basic
package main

import (
	"fmt"
)

func main() {

	// 변수
	var a int			// 초기 값을 지정하지 않으면 Zero Value 기본 적용
						// 즉, 숫자형에는 0, bool 타입에는 false, 그리고 string 형에는 "" (빈문자열)을 할당
	var b int = 11
	c := 11
	fmt.Printf("%d %d %d\n", a,b,c)

	var a1, a2, a3 int = 11, 11, 11
	var (
		b1 = 1
		b2 = 2
		b3 = 3
	)
	fmt.Printf("%d %d %d\n", a1, a2, a3)
	fmt.Printf("%d %d %d\n", b1, b2, b3)
	
	// 상수
	const d int = 10
	const e string = "hi"
	fmt.Printf("%d %s\n", d,e)

	const (
		name="TEST"
		job=false
	)
	fmt.Printf("%s %s\n", name, job)

	// Go Keyword
	// 예약어라 사용 못함
	/*
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	*/

	// Data Type
	var k1 bool
	var k2 string
	var k3 int
	var k4 int8
	var k5 int16
	var k6 int32
	var k7 int64
	
	var k8 uint
	var k9 uint8
	var k10 uint16
	var k11 uint32
	var k12 uint64
	var k13 uintptr

	var k14 float32 
	var k15 float64 
	var k16 complex64 	// 복소수
	var k17 complex128	// 복소수

	var k18 byte		// uint8과 동일, byte code에 사용
	var k19 rune		// int32와 동일, unicode point에 사용
	fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n",k1,k2,k3,k4,k5,k6,k7,k8,k9)
	fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n",k10,k11,k12,k13,k14,k15,k16,k17,k18,k19)
	
	// backtick string => Raw String Literal
	string1 := `test
	\n test \n test`
	fmt.Println(string1)

	// Quotation Marks => Interpreted String Literal
	string2 := "test\n test \n test"
	fmt.Println(string2)

	// Type Conversion
	var i int = 100
	var u uint = uint(i)
	fmt.Printf("%d %p\n",i,&i)
	fmt.Printf("%d %p\n",u,&u)

}
