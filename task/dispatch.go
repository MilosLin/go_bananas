package task

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/MilosLin/go_bananas/core/logger"
)

/**
 * 自行解析外部帶入的flag參數
 *
 * command format: ./go_bananas task --name={task_name} --argu="-p1=123 -p2=abc"
 *
 */
func Dispatch(name, argu *string) {
	switch *name {
	case "tasktemplate": //task寫法範例
		run(NewTaskTemplate(), argu)
	case "parseflagstr": //flag字串解析
		run(NewParseFlagStr(), argu)
	case "logexample": //zap用法
		run(NewLogExample(), argu)
	case "mysqlcuid": //mysql CUID範例
		run(NewMysqlCUID(), argu)
	case "mysqlinsertbenchmark": //mysql insert benchmark
		run(NewMysqlInsertBenchmark(), argu)
	default:
		logger.Fatal("Undefine task", zap.String("task", *name), zap.String("argu", *argu))
	}
}

func run(t Itask, argu *string) {
	if err := t.Run(argu); err != nil {
		logger.Fatal(err.Error(), zap.Stack(""))
	}
}

/**
 * 使用corba自動解析的[]string做參數使用
 *
 * command format: ./go_bananas task --name={task_name} argu1 argu2 ..
 *
 */
func corbaDispatch(name string, argu []string) {
	switch name {
	case "corbaArgu":
		fmt.Printf("name=%v argu=%v", name, argu)
		fmt.Println("corbaArgu")
	default:
		fmt.Printf("dispatch default:%s.", name)
	}
}
