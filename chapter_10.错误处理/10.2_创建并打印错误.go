/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-30 17:33:39
 * @Description: 第一个go程序
 */
 package main

 import (
	 "errors"
	 "fmt"
 )
 
 func main() {
	 err := errors.New("Something went wrong")
	 if err != nil {
		 fmt.Println(err)
	 }
 }
 