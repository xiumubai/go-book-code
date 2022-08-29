/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 15:58:59
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)

func main() {
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	fmt.Println(players["cook"])
	fmt.Println(players)
	// 删除元素
	delete(players, "cook")
	fmt.Println(players)
}