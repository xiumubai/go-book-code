/*
 * @Author: 朽木白
 * @Date: 2022-09-01 14:21:45
 * @LastEditors: 1547702880@qq.com
 * @LastEditTime: 2022-09-01 14:21:53
 * @Description: 
 */
package example02

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("George")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
