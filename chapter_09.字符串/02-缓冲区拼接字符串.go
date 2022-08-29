/*
 * @Author: 朽木白
 * @Date: 2022-08-29 17:16:17
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 17:16:21
 * @Description: 
 */
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())
}
