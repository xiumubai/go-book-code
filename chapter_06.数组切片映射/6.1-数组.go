/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:51:42
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func main() {
	var cheeses [2]string
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}