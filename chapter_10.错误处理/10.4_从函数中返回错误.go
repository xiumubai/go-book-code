/*
 * @Author: 朽木白
 * @Date: 2022-08-30 17:37:25
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-08-30 17:37:34
 * @Description: 
 */
package main

import (
	"fmt"
)

// Half takes an integer and returns half the value
func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	n, err := Half(19)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
