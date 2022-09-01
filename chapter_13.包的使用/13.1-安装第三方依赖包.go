/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:50:38
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 12:00:47
 * @Description: 
 */

// 需要执行 go install golang.org/x/example/stringutil

package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	s := "ti esrever dna ti pilf nwod gniht ym tup I"
	fmt.Println(stringutil.Reverse(s))
}
