/*
 * @Author: 朽木白
 * @Date: 2022-09-01 23:20:58
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 23:21:25
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
