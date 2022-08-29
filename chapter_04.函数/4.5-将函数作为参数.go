/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:49:33
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 14:49:42
 * @Description: 
 */
/* recursive function */
package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))

}
