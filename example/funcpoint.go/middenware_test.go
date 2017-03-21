package funcpoint

import "testing"

// 自訂中間件
func Test_MiddenWare_1(t *testing.T) {
	txt := &context{
		header: "header",
		form:   "form",
		code:   200,
	}
	M := MWExampe{
		Context: txt,
	}
	M.Use(func() string {
		return M.Context.header + " custom func;"
	})

	expect := "Do() start;header custom func;Do() end;"
	result := M.Do()

	if expect != result {
		t.Fatalf("Failed! Test_MiddenWare_1() out:%s", result)
	} else {
		t.Log("Pass! Test_MiddenWare_1()")
	}
}

// 使用預設方法
func Test_MiddenWare_2(t *testing.T) {
	txt := &context{
		header: "header",
		form:   "form",
		code:   200,
	}
	M := MWExampe{
		Context: txt,
	}
	M.Use(M.PrintCode)
	M.Use(M.PrintHeader)
	M.Use(M.PrintForm)

	expect := "Do() start;PrintCode:200;PrintHeader:header;PrintForm:form;Do() end;"
	result := M.Do()

	if expect != result {
		t.Fatalf("Failed! Test_MiddenWare_2() out:%v", result)
	} else {
		t.Log("Pass! Test_MiddenWare_2()")
	}
}
