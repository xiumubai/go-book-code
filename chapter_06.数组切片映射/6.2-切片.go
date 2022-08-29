/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:57:22
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses"
	// 添加单个元素
	cheeses = append(cheeses, "Camembert")
	// 添加多个元素
	cheeses = append(cheeses, "Reblochon", "Picodon")
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	// 删除元素
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	// 复制数组
	var smellyCheeses = make([]string, 4)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)
}