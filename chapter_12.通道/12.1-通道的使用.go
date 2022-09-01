/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:07:20
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:07:25
 * @Description: 
 */
package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "slowFunc() finished"
}

func main() {
	c := make(chan string)
	go slowFunc(c)

	msg := <-c
	fmt.Println(msg)
}
