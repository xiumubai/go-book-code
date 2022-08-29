/*
 * @Author: 朽木白
 * @Date: 2022-08-29 17:07:19
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 17:14:02
 * @Description: 
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	t := "After a backslash, certain single character escapes represent special values\nn is a line feed or new line \n\t t is a tab"
	s := `After a backslash, certain single character escapes represent special values
		n is a line feed or new line
		t is a tab`
	
	// 拼接
	a := "Oh sweet ignition" + " be my fuse"
	a += "hahahha"
	n := 1
	m := a + strconv.Itoa(n)
	
	fmt.Println(m)
	fmt.Println(s, t)
}