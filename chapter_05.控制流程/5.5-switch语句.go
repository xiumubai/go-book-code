/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:00:10
 * @Description: 第一个go程序
 */
 package main

 import (
	 "fmt"
 )
 
 func main() {
	i := 2

	switch i {
		case 5:
			fmt.Println("Two!")
		case 3:
			fmt.Println("Three!")
		case 4:
			fmt.Println("Four!")
		default:
			fmt.Println("I don't recognize that letter!")
	}
 }