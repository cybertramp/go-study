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

	// IF 문
	k := 2

	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {  //같은 라인
		fmt.Println("Two")
	} else {   //같은 라인
		fmt.Println("Other")
	}

	// 초기값 지정및 조건을 지정할 수 있음
	if val := 2; val == 2 {
		fmt.Println(val)
	}
	 
	// 아래 처럼 사용하면 Scope 벗어나 에러
	// val++
	var name string
    var category = 2
 
    switch category {
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
    println(name)

	// switch문에 비교할 변수를 미리 선언하지 않아도 됨
	// IF문 단순화 가능
	score := 40
	switch {
    case score >= 90:
        println("A")
    case score >= 80:
        println("B")
    case score >= 70:
        println("C")
    case score >= 60:
        println("D")
    default:
        println("No Hope")
    }
	var v interface{} = "hello"
	// Type Check
	switch v.(type) {
		case int:
			println("int")
		case bool:
			println("bool")
		case string:
			println("string")
		default:
			println("unknown")
	}   
}
