package funcpoint

import "fmt"

type ArithOp func(int, int) int

// callback 範例
// calculate() 方法只要求傳入的func與指定的宣告式相符即可執行
//
// http://stackoverflow.com/questions/3601796/can-we-have-function-pointers-in-google-go
func CallBackExample() (ret string) {
	ret += fmt.Sprintf("calculate(Plus):%v\n", calculate(Plus))
	ret += fmt.Sprintf("calculate(Minus):%v\n", calculate2(Minus))
	ret += fmt.Sprintf("calculate(Multiply):%v\n", calculate(Multiply))
	return
}

// 接收的func宣告式可寫成type，或是直接寫成 func的宣告式
func calculate(fp func(int, int) int) int {
	return fp(3, 2)
}

func calculate2(fp ArithOp) int {
	return fp(3, 2)
}

// This is the same function but uses the type/fp defined above
//
// func calculate (fp ArithOp) {
//     ans := fp(3,2)
//     fmt.Printf("\n%v\n", ans)
// }

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
