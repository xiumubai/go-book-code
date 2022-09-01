/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:43:46
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:44:22
 * @Description: 
 */
package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Time's up!")
		stop <- true
	}()

	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
