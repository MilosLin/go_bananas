package funcpoint

import "fmt"

// 中間件範例
// 於既定的流程中，串入自訂的方法
type MWExampe struct {
	Context   *context
	callStack []middenware
}

type context struct {
	header string
	form   string
	code   int
}

// 中間件型態
type middenware func() string

// 使用Use方法加入middenware
func (M *MWExampe) Use(fn func() string) {
	M.callStack = append(M.callStack, fn)
}

// 執行原本的流程，輸出會因插入的中間件而改變
func (M *MWExampe) Do() (ret string) {
	ret += "Do() start;"

	for _, v := range M.callStack {
		ret += v()
	}
	ret += "Do() end;"
	return
}

// 預設的方法
func (M *MWExampe) PrintHeader() string {
	return fmt.Sprintf("PrintHeader:%v;", M.Context.header)
}

func (M *MWExampe) PrintForm() string {
	return fmt.Sprintf("PrintForm:%v;", M.Context.form)
}

func (M *MWExampe) PrintCode() string {
	return fmt.Sprintf("PrintCode:%v;", M.Context.code)
}
