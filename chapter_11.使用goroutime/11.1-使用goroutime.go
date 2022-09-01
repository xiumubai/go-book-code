/*
 * @Author: 朽木白
 * @Date: 2022-09-01 10:54:42
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 10:54:45
 * @Description: 
 */
package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("slowFunc() finished")
}

func main() {
	go slowFunc()
	fmt.Println("I am now shown straightaway!")
	time.Sleep(time.Second * 3)
}
