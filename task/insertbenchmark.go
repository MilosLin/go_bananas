package task

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MilosLin/go_bananas/core/model/database/behavior"
	dto "github.com/MilosLin/go_bananas/core/model/database/dto/testdb"
	"github.com/MilosLin/go_bananas/core/module/validator"
	"github.com/shopspring/decimal"
)

type MysqlInsertBenchmark struct {
	Num int `validate:"gt=0"`
}

func NewMysqlInsertBenchmark() *MysqlInsertBenchmark {
	return &MysqlInsertBenchmark{}
}

// run task MysqlInsertBenchmark
//
// task --name=mysqlinsertbenchmark --argu="-num=150"
func (instance *MysqlInsertBenchmark) Run(argu *string) error {
	argus := flag.NewFlagSet("argu", flag.ExitOnError)
	argus.IntVar(&instance.Num, "num", 1000, "insert how many rows in once")
	argus.Parse(strings.Fields(*argu))

	if err := validator.Struct(instance); err != nil {
		return err
	}
	fmt.Printf("prepare testing. num=%d \n", instance.Num)
	// Prepare Testing. truncate table
	behavior.TruncateOrder()

	orders := []dto.Order{}
	for i := 0; i < instance.Num; i++ {
		o := dto.Order{
			UserID:    "a" + strconv.Itoa(i),
			OrderTime: time.Now(),
			Money:     decimal.NewFromFloat(10),
			Remark:    "Remark String",
		}
		orders = append(orders, o)
	}

	fmt.Printf("start test insert data one by one... \n")
	start := time.Now()
	// insert data one by one
	for _, o := range orders {
		behavior.InsertOrder(o)
	}
	stop := time.Now()
	fmt.Printf("end test insert data one by one. latency:%v \n", stop.Sub(start))
	// truncate table
	behavior.TruncateOrder()

	fmt.Printf("start test multi insert... \n")
	start = time.Now()
	behavior.MultiInsertOrder(orders)
	stop = time.Now()
	fmt.Printf("end test multi insert. latency:%v \n", stop.Sub(start))

	// truncate table
	behavior.TruncateOrder()
	return nil
}
