/*
 * @Author: 朽木白
 * @Date: 2022-09-01 11:30:03
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 11:38:15
 * @Description: 
 */
// <-位于关键字chan左边时，表示通道在函数内是只读的；
// <-位于关键字chan右边时，表示通道在函数内是只写的；
// 没有指定<-时，表示通道是可读写的。


package main

import (
	"fmt"
	"time"
)

func fn(message <-chan string) // 左边只读
func fn(message chan<- string) // 右边只写
func fn(message chan string) // 可读可写

func main() {}
