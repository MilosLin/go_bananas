package behavior

import (
	"strings"
	"time"

	"github.com/MilosLin/go_bananas/core/env"
	"github.com/MilosLin/go_bananas/core/logger"
	"github.com/MilosLin/go_bananas/core/model/database"
	dto "github.com/MilosLin/go_bananas/core/model/database/dto/testdb"
	simplejs "github.com/bitly/go-simplejson"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type assets struct{}

var assetsObj *assets

func Assets() *assets {
	if assetsObj == nil {
		assetsObj = new(assets)
	}
	return assetsObj
}

func (instance assets) InsertData(itemName string, dynamicCol map[string]interface{}) (int64, int64) {
	testM := database.GetConn("testDBM")

	var params []interface{}
	params = append(params, itemName)

	placeholder := []string{}

	for k, v := range dynamicCol {
		params = append(params, k)
		params = append(params, v)
		placeholder = append(placeholder, "?")
		placeholder = append(placeholder, "?")
	}

	stmt, err := testM.Prepare("INSERT INTO assets VALUES (?, COLUMN_CREATE(" + strings.Join(placeholder, ", ") + "));")

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	res, err := stmt.Exec(params...)
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

func (instance assets) GetAllData() (ret []dto.Asset) {
	testS := database.GetConn("testDBS")

	stmt, err := testS.Prepare(
		`SELECT
          item_name,
          COLUMN_JSON(dynamic_cols)
        FROM assets;
        `,
	)

	if err != nil {
		logger.Fatal("Prepare SQL Failed", zap.Error(err), zap.Stack(""))
	}
	rows, err := stmt.Query()
	if err != nil {
		logger.Fatal("Exec Stmt Failed", zap.Error(err))
	}
	defer rows.Close()
	for rows.Next() {
		data := dto.Asset{}
		var jsonString string
		if err := rows.Scan(
			&data.ItemName,
			&jsonString,
		); err != nil {
			logger.Fatal("Fatch Data Error", zap.Error(err))
		}
		if js, err := simplejs.NewJson([]byte(jsonString)); err == nil {
			data.Color, _ = js.Get("color").String()
			data.Size, _ = js.Get("size").String()

			timeString, _ := js.Get("expiry_date").String()
			data.ExpiryDate, _ = time.Parse(env.DateTime, timeString)

			priceFloat, _ := js.Get("price").Float64()
			data.Price = decimal.NewFromFloat(priceFloat)
		} else {
			logger.Fatal("Parse JSON Error", zap.Error(err))
		}
		ret = append(ret, data)
	}

	return
}

func TruncateAsserts() int64 {
	testM := database.GetConn("testDBM")

	stmt, err := testM.Prepare("TRUNCATE TABLE `assets`")

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
