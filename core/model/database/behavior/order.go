package behavior

import (
	"github.com/MilosLin/go_bananas/core/model/database"

	"github.com/MilosLin/go_bananas/core/logger"

	dto "github.com/MilosLin/go_bananas/core/model/database/dto/testdb"
	"go.uber.org/zap"
)

func InsertOrder(o dto.Order) (int64, int64) {
	testM := database.GetConn("testDBM")

	stmt, err := testM.Prepare("INSERT INTO `order`(`user_id`, `order_time`, `money`, `remark`) VALUES(?,?,?,?)")

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	res, err := stmt.Exec(o.UserID, o.OrderTime, o.Money.String(), o.Remark)
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		logger.Fatal("Get Last Insert Id Failed", zap.Error(err))
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		logger.Fatal("Get Affected Rows Failed", zap.Error(err))
	}
	return lastId, rowCnt
}

func GetOrderByMaxID(maxOID string) (ret []dto.Order) {
	testS := database.GetConn("testDBS")

	stmt, err := testS.Prepare(
		"SELECT `o_id`, `user_id`, `order_time`, `money`, `remark`, `update_time`, `create_time` FROM `order` WHERE `o_id` <= ?",
	)

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	rows, err := stmt.Query(maxOID)
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}
	defer rows.Close()
	for rows.Next() {
		orderData := dto.Order{}
		if err := rows.Scan(
			&orderData.OID,
			&orderData.UserID,
			&orderData.OrderTime,
			&orderData.Money,
			&orderData.Remark,
			&orderData.UpdateTime,
			&orderData.CreateTime,
		); err != nil {
			logger.Fatal("Fatch Data Error", zap.Error(err))
		}
		ret = append(ret, orderData)
	}

	return
}

func UpdateMoneyByOID(OID, money string) (int64, int64) {
	testM := database.GetConn("testDBM")

	stmt, err := testM.Prepare("UPDATE `order` SET `money` = ? WHERE `o_id` = ?")

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	res, err := stmt.Exec(money, OID)
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		logger.Fatal("Get Last Insert Id Failed", zap.Error(err))
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		logger.Fatal("Get Affected Rows Failed", zap.Error(err))
	}
	return lastId, rowCnt
}

func DeleteOrderByOID(OID string) int64 {
	testM := database.GetConn("testDBM")

	stmt, err := testM.Prepare("DELETE FROM `order` WHERE `o_id` = ?")

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	res, err := stmt.Exec(OID)
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		logger.Fatal("Get Affected Rows Failed", zap.Error(err))
	}
	return rowCnt
}

func TruncateOrder() int64 {
	testM := database.GetConn("testDBM")

	stmt, err := testM.Prepare("TRUNCATE TABLE `order`")

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	res, err := stmt.Exec()
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		logger.Fatal("Get Affected Rows Failed", zap.Error(err))
	}
	return rowCnt
}

// MultiInsertOrder
//
// 批次insert範例
func MultiInsertOrder(orders []dto.Order) (rowCnt int64) {
	if len(orders) <= 0 {
		return 0
	}

	testM := database.GetConn("testDBM")

	for cut := 0; cut < len(orders); cut += database.MaxBatchSQL {
		startIndex := cut
		endIndex := startIndex + database.MaxBatchSQL
		if endIndex > len(orders) {
			endIndex = len(orders)
		}
		sql := "INSERT INTO `order`(`user_id`, `order_time`, `money`, `remark`) VALUES(?,?,?,?)"

		for i := 0; i < (len(orders[startIndex:endIndex]) - 1); i++ {
			sql += ",(?,?,?,?)"
		}
		sql += ";"
		var params []interface{}
		for _, o := range orders[startIndex:endIndex] {
			params = append(params, o.UserID)
			params = append(params, o.OrderTime)
			params = append(params, o.Money.String())
			params = append(params, o.Remark)
		}

		stmt, err := testM.Prepare(sql)

		if err != nil {
			logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
		}
		res, err := stmt.Exec(params...)
		if err != nil {
			logger.Fatal("Exec Stmt Failed", zap.Error(err))
		}

		count, err := res.RowsAffected()
		if err != nil {
			logger.Fatal("Get Affected Rows Failed", zap.Error(err))
		}
		rowCnt += count
	}
	return
}
