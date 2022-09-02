/*
 * @Author: 朽木白
 * @Date: 2022-09-01 23:20:58
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-02 17:22:04
 * @Description: 
 */
package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}

// 运行 go mod init 18.1_hello_web
// 运行 go build 18.1_hello_web
// 运行 go run init 18.1_hello_web
// 浏览器打开 http://localhost:8000/ 就可以看到Hello World了