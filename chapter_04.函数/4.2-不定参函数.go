/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 14:36:29
 * @Description: 第一个go程序
 */
 package main

 import (
	 "fmt"
 )

//  计算任意整数的和
 func getSum(nums...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
 }
 
 func main() {

	s := getSum(1,2,3,4,5,6)
	
	 fmt.Println(s)
 }