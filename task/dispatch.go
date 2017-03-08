package task

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MilosLin/go_bananas/core/log"
)

/**
 * 自行解析外部帶入的flag參數
 *
 * command format: ./go_bananas task --name={task_name} --argu="-p=123 -p2=abc"
 *
 */
func Dispatch(name, argu string) {

	switch name {
	case "parseflagstr": //flag字串解析
		var p1 string
		var p2 int
		f := flag.NewFlagSet("ask", flag.ExitOnError)
		f.StringVar(&p1, "p1", "default", "paramter 1")
		f.IntVar(&p2, "p2", 5, "paramter 2")
		f.Parse(strings.Fields(argu))
		fmt.Printf("\np1=%v \np2=%v\n", p1, p2)
		//io, _, e := zap.Open("/home/work/go/src/github.com/MilosLin/go_bananas/test.log")

		//fmt.Printf("open error %v\n", e)

		log.Inst("root").Debug("debug")
		log.Inst("root").Info("Info")
		log.Inst("root").Warn("Warn")
		log.Inst("root").Error("Error")
		//log.Fatal("Fatal")
		/*logger, _ := zap.NewProduction(
			zap.ErrorOutput(io),
		)
		zapcore.AddSync(io)
		zapcore.NewJSONEncoder(cfg)
		sugar := logger.Sugar()
		sugar.Infow("Failed to fetch URL.",
			// Structured context as loosely-typed key-value pairs.
			"url", "url-string",
			"attempt", "retryNum",
			"backoff", "time.Second",
		)
		sugar.Infof("Failed to fetch URL: %s", "the url")*/
	default:
		fmt.Printf("dispatch default:%s.", name)
	}
}

/**
 * 使用corba自動解析的[]string做參數使用
 *
 * command format: ./go_bananas task --name={task_name} argu1 argu2 ..
 *
 */
func IncognitoDispatch(name string, argu []string) {
	switch name {
	case "corbaArgu":
		fmt.Printf("name=%v argu=%v", name, argu)
		fmt.Println("corbaArgu")
	default:
		fmt.Printf("dispatch default:%s.", name)
	}
}
