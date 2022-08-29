/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 16:33:34
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}
// 给结构体Movie声明方法summary
func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

func main() {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}

	fmt.Println(m.summary())
}
