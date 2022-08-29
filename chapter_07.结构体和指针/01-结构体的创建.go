/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 16:02:12
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)


type Movie struct {
	Name   string
	Age float32
}

func main() {
	m := Movie{
		Name:   "Citizen-Kane",
		Age: 10,
	}
	fmt.Println(m.Name, m.Age)
}
