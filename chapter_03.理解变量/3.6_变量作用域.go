/** 
每对大括号都表示一个块

每个内部块都可访问其外部块
*/

package main

import (
	"fmt"
)

var s = "Hello world"

func main() {
	fmt.Printf("Printing `s` variable from outer block %v\n", s)
	b := true
	if b {
		fmt.Printf("Printing `b` variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Printing `i` variable from outer block %v\n", i)
		}
	}
}
