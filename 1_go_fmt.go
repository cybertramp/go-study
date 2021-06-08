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
func main(){
	// 사전에 선언 안할 때는 var 사용
	// 사전에 대입 까지 할때는 var 없고 := 사용, 의미는 선언과 동시에 대입
	var input_tmp string
	message := "hello world!"

	fmt.Print(message,"\n")
	fmt.Println(message)
	fmt.Printf("%s\n", message)

	fmt.Println("Input Your Name.")
	if _, err := fmt.Scan(&input_tmp); err!= nil {
		panic(err)
	}
	fmt.Printf("Input: %s\n", input_tmp)
}
