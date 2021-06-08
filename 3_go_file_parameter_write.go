/*
*
*
*
*
 */

// basic
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args) > 1 {
		fmt.Println(strings.Join(args[1:], " "))
		fp, _ := os.Create("tmp.txt") // 파일 생성
		w := bufio.NewWriter(fp)      // 쓰기 버퍼 생성

		fmt.Fprint(w, strings.Join(args[1:], " ")) // 쓰기 버퍼에 입력

		w.Flush() // 쓰기 버퍼 Flush
	} else {
		fmt.Println("Please give parameter!!")
	}

}
