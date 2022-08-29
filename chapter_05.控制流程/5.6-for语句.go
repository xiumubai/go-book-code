/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:08:02
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func main() {
	forA()
	forB()
	forRange()
}

func forA() {
	i := 0

	for i < 3 {
		i++
		fmt.Println("forA i is", i)
	}

}

func forB() {

	for i := 0; i < 3; i++{
		fmt.Println("forB i is", i)
	}
}

func forRange() {
	nums := []int{1,2,3}
	for i, n:= range nums {
		fmt.Println("The index of the loop is", i)
		fmt.Println("The value from the slice is", n)
	} 
}