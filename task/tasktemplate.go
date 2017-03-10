package task

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MilosLin/go_bananas/core/module/validator"
)

// struct可增加驗證規則，設定值參考官網
//
// https://godoc.org/gopkg.in/go-playground/validator.v9
type TaskTemplate struct {
	RequireArgu string `validate:"required"`
}

// New TaskTemplate
func NewTaskTemplate() *TaskTemplate {
	return &TaskTemplate{}
}

// run task TaskTemplate
//
// task --name=tasktemplate --argu="-RequireArgu=teststringcan'twithspace"
func (instance *TaskTemplate) Run(argu *string) error {
	// 自行解析字串
	var argus = flag.NewFlagSet("", flag.ContinueOnError)
	argus.StringVar(&instance.RequireArgu, "RequireArgu", "", "RequireArgu")
	argus.Parse(strings.Fields(*argu))

	// 驗證參數，若錯誤則直接回傳
	if err := validator.Struct(instance); err != nil {
		return err
	}

	fmt.Printf("%+v\n", instance)
	return nil
}
