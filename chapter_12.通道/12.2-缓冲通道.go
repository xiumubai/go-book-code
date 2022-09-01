/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:12:11
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:13:25
 * @Description: 
 */
package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {

	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("Pushed two messages onto channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)

}


// 程序解读

// ➢ 创建一个长度为2的缓冲通道。
// ➢ 向通道发送两条消息。此时没有可用的接收者，因此消息被缓冲。
// ➢ 关闭通道，这意味着不能再向它发送消息。➢ 程序打印一条消息，指出通道包含两条消息，再休眠1s。
// ➢ 将通道作为参数传递给函数receiver。
// ➢ 函数receiver使用range迭代通道，并将通道中缓冲的消息打印到控制台。