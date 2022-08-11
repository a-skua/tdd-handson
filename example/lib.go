package example

import (
	"fmt"
)

// 2つの整数の和を返す
func Add(a, b int) int {
	return sum(a, b)
}

// 3つの整数の和を返す
func Add3(a, b, c int) int {
	return sum(a, b, c)
}

func sum(ns ...int) (result int) {
	for _, n := range ns {
		result += n
	}
	return
}

// 新しいユーザーIDを発行する
// ユーザーIDは数字10桁で、数字はランダムに生成される
func NewUserID(rnd func() int) string {
	return fmt.Sprintf("%010d", rnd()%10000000000)
}
