package funcpoint

import "fmt"

type HelloFunc func(string) string

func SayHello(to string) string {
	return fmt.Sprintf("Hello, %s!", to)
}

// Func point 基礎概念
//
// 1. func可作為參數傳遞
// 2. 在變數後使用"()"，可執行該方法並傳入參數
func BasicExample() string {
	var hf HelloFunc

	hf = SayHello

	return hf("world")
}
