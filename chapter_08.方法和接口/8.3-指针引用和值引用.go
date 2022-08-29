/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 16:45:19
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)


type Triangle struct {
	base   float64
	height float64
}
// 若果不带*就是值引用，接受到的f为3
// 若果带*就是指针引用，接受到的f为4
func (t *Triangle) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 1}
	t.changeBase(4)
	fmt.Println(t.base)
}
