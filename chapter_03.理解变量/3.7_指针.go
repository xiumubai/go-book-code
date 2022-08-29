/*
 * @Author: 朽木白
 * @Date: 2022-08-29 13:54:35
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 13:59:24
 * @Description: 
 */
package main

import (
	"fmt"
)

func printMemory(x *int) {
	// 指针指向的值
	fmt.Println(*x)	
	// 指针
	fmt.Println(x)
	return
}



func main() {

	x := 1
	fmt.Println(&x)
	printMemory(&x)

}