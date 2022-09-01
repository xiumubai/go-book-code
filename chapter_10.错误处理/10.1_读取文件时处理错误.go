/*
 * @Author: 朽木白
 * @Date: 2022-08-29 14:06:26
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-30 17:31:25
 * @Description: 第一个go程序
 */
 package main

 import (
	 "fmt"
	 "io/ioutil"
 )
 
 func main() {
	 var file []byte
	 var err error
	 file, err = ioutil.ReadFile("foo.txt")
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 fmt.Printf("%s", file)
 }
 