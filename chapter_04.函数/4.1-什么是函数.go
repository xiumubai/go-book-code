/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:22:22
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 14:28:49
 * @Description: 
 */
package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

// 关键字func指出这是一个函数的开头位置

// 如果函数签名声明了返回值，则函数体必须以终止语句结束

// 返回多个值

func getPrice() (int, string) {
	s := 12
	w := "5kg"

	return s, w
}

// 调用getPrice
func main() {
	q, p := getPrice()

	fmt.Println(q, p)
}