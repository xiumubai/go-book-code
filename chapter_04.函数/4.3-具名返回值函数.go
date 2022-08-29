/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 14:39:10
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func sayhi() (x, y string) {
	x = "hello"
	y = "go"

	return
}

func main() {
	fmt.Println(sayhi())
}