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
