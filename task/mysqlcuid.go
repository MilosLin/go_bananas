package task

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MilosLin/go_bananas/core/env"
	"github.com/MilosLin/go_bananas/core/model/database/behavior"
	dto "github.com/MilosLin/go_bananas/core/model/database/dto/testdb"
	"github.com/olekukonko/tablewriter"
	"github.com/shopspring/decimal"
)

type MysqlCUID struct{}

func NewMysqlCUID() *MysqlCUID {
	return &MysqlCUID{}
}

// run task MysqlCUID
//
// task --name=mysqlcuid
func (instance *MysqlCUID) Run(argu *string) error {
	// insert 4 data
	for i := 0; i < 4; i++ {
		o := dto.Order{
			UserID:    "a" + strconv.Itoa(i),
			OrderTime: time.Now(),
			Money:     decimal.NewFromFloat(10),
			Remark:    "Remark String",
		}
		behavior.InsertOrder(o)
	}

	fmt.Printf("Insert 4 data into database\n")
	// read data
	instance.showOrder(behavior.GetOrderByMaxID("4"))

	fmt.Printf("Update SET money=50 Where o_id=3 \n")
	// update data by id
	behavior.UpdateMoneyByOID("3", "50")

	fmt.Printf("Delete Where o_id=1 \n")
	// delete data
	behavior.DeleteOrderByOID("1")

	// read data again
	instance.showOrder(behavior.GetOrderByMaxID("4"))

	// truncate table
	behavior.TruncateOrder()
	return nil
}

// showOrder
//
// print data on screen, format by table
func (instance *MysqlCUID) showOrder(orders []dto.Order) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetHeader([]string{
		"o_id",
		"user_id",
		"order_time",
		"money",
		"remark",
		"update_time",
		"create_time",
	})

	for _, o := range orders {
		table.Append([]string{
			strconv.Itoa(o.OID),
			o.UserID,
			o.OrderTime.Format(env.DateTimeWithTimeZone),
			o.Money.String(),
			o.Remark,
			o.UpdateTime.Format(env.DateTimeWithTimeZone),
			o.CreateTime.Format(env.DateTimeWithTimeZone),
		})
	}
	table.Render()
}
