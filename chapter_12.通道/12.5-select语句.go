/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:32:57
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:34:58
 * @Description: 
 */
package main

import "time"
import "fmt"

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	// select语句只执行第一个满足条件的语句
	select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
	}

}
