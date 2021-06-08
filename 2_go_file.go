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
	"os"
	"bufio"
)
func main(){
	message := "hello world!"
	fmt.Println(message)
	
	fp, _ := os.Create("tmp.txt")		// 파일 생성
	w := bufio.NewWriter(fp)			// 쓰기 버퍼 생성

	fmt.Fprint(w, message)				// 쓰기 버퍼에 입력

	w.Flush()							// 쓰기 버퍼 Flush
}
