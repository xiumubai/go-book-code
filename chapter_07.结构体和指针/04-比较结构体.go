/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 16:12:33
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  false,
	}
	if a == b {
		fmt.Println("a and b are the same")
	} else {
		fmt.Println("a and b not same")	
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}