package funcpoint

import "fmt"

type MyFunc func(string) (string, error)

func A(argu string) (string, error) {
	return fmt.Sprintf("A():%+v", argu), nil
}

func B(argu string) (string, error) {
	return fmt.Sprintf("B():%+v", argu), nil
}

func C(argu string) (string, error) {
	return fmt.Sprintf("C():%+v", argu), nil
}

// 使用func point對流程做拆解
//
// note:
// 1. 若要將func當指標傳遞，需先宣告一type，宣告式需與傳入的func一致
func FlowControlExample(flow int) (ret string) {
	CallStack := []MyFunc{}

	switch flow {
	case 1: //流程 1
		CallStack = append(CallStack, A)
		CallStack = append(CallStack, B)
	case 2: //流程 2
		CallStack = append(CallStack, A)
		CallStack = append(CallStack, C)
	default: //預設
		CallStack = append(CallStack, A)
		CallStack = append(CallStack, B)
		CallStack = append(CallStack, C)
	}

	for _, v := range CallStack {
		out, _ := v("input")
		ret += out
	}
	return
}
