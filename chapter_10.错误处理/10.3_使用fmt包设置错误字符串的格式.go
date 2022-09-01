/*
 * @Author: 朽木白
 * @Date: 2022-08-30 17:34:37
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-30 17:34:46
 * @Description: 
 */
package main

import (
	"fmt"
)

func main() {
	name, role := "Richard Jupp", "Drummer"
	err := fmt.Errorf("The %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)
	}
}
