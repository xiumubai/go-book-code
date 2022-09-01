/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:27:42
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:27:52
 * @Description: 
 */
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
