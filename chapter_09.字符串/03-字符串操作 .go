/*
 * @Author: 朽木白
 * @Date: 2022-08-29 17:16:17
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 17:30:38
 * @Description: 
 */
package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "sdhfjkdshfjsf"
	fmt.Println(len(s))

	// 将字符串转换为小写
	fmt.Println(strings.ToLower("VERY IMPORTANT MESSAGE"))

	// 在字符串中查找子串
	fmt.Println(strings.Index("surface", "face")) // 索引3
	fmt.Println(strings.Index("moon", "aer")) // 没找到-1

	// 删除字符串中的空格
	fmt.Println(strings.TrimSpace("   I don't need all this space     "))

}
