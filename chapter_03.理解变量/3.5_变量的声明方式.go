/*
 * @Author: 朽木白
 * @Date: 2022-08-29 11:18:19
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 11:53:19
 * @Description: 
 */
package main

import (
	"fmt"
	"strconv"
	"reflect"
)


// 变量的声明在函数外可以省略类型
var a = "this is golang! "

func main() {

	// 变量声明
	var b bool  = true
	fmt.Println(reflect.TypeOf(b)) // bool

	// 类型转换
	var s string = strconv.FormatBool(b)
	fmt.Println(reflect.TypeOf(s)) // string


	// 零值
	var i int 
	fmt.Println(i) // 0 

	// 简短变量声明-不能在函数外面使用
	h := "hello gopher"
	fmt.Println(h, a)
}