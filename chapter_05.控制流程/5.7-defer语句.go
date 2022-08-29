/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:18:57
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer one")
	defer fmt.Println("defer two")
	fmt.Println("hello go") // 先执行
}

// 打印顺序

// hello go
// defer tow
// defer one

// defer是一个很有用的Go语言功能，它能够让您在函数返回前执行另一个函数。
// 函数在遇到return语句或到达函数末尾时返回。defer语句通常用于执行清理操作或确保操作（如网络调用）完成后再执行另一个函数。