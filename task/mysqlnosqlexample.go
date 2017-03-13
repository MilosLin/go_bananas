package task

import (
	"fmt"

	"github.com/MilosLin/go_bananas/core/model/database/behavior"
)

type MysqlNosqlExample struct{}

func NewMysqlNosqlExample() *MysqlNosqlExample {
	return &MysqlNosqlExample{}
}

// run task MysqlNosqlExample
//
// task --name=mysqlnosqlexample
func (instance *MysqlNosqlExample) Run(argu *string) error {
	behavior.TruncateAsserts()

	dy := make(map[string]interface{})
	dy["color"] = "blue"
	dy["size"] = "XL"
	dy["expiry_date"] = "2017-08-05 00:00:00"
	behavior.Assets().InsertData(
		"MariaDB T-shirt",
		dy,
	)

	dy = make(map[string]interface{})
	dy["color"] = "black"
	dy["price"] = 500
	dy["expiry_date"] = "2017-08-05 23:59:59"
	behavior.Assets().InsertData(
		"Thinkpad Laptop",
		dy,
	)
	fmt.Printf("%+v\n", behavior.Assets().GetAllData())
	behavior.TruncateAsserts()
	return nil
}
