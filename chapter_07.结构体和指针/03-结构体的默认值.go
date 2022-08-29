/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-29 16:08:38
 * @Description: 第一个go程序
 */
package main

import (
	"fmt"
)


type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

func main() {
	fmt.Printf("%+v\n", NewAlarm("07:00"))
}
