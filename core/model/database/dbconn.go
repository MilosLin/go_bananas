/**
 * 資料庫連線池，用於控管各db的連線
 */
package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MilosLin/go_bananas/core/config"
	"github.com/MilosLin/go_bananas/core/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

//資料庫連線物件
var pool map[string]*sql.DB

//查詢條件影響筆數上限
var QueryLimit int

//批次Insert/Update SQL筆數上限
var MaxBatchSQL int

func init() {
	pool = make(map[string]*sql.DB)
	c := config.Instance()
	if limit := c.GetInt("database.queryLimit"); limit > 0 {
		QueryLimit = limit
	} else {
		QueryLimit = 500000
	}

	if maxBatchSQL := c.GetInt("database.maxBatchSQL"); maxBatchSQL > 0 {
		MaxBatchSQL = maxBatchSQL
	} else {
		MaxBatchSQL = 500
	}
}

/**
 * 依照資料庫名稱取得DB連線
 */
func GetConn(db_name string) *sql.DB {
	if value, ok := pool[db_name]; ok {
		return value
	} else {
		c := config.Instance()
		group := "database." + db_name

		conn_str := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=UTC",
			c.GetString(group+".account"),
			c.GetString(group+".password"),
			c.GetString(group+".ip"),
			c.GetString(group+".port"),
			c.GetString(group+".db"),
		)
		db, err := sql.Open("mysql", conn_str)

		if err != nil {
			logger.Fatal("Connect Database Failed", zap.Error(err), zap.String("connStr", conn_str))
		}

		db.SetMaxOpenConns(c.GetInt("database.maxOpenConns"))
		db.SetMaxIdleConns(c.GetInt("database.maxIdleConns"))
		db.SetConnMaxLifetime(time.Second * c.GetDuration("database.connMaxLifeTime"))
		pool[db_name] = db
		return db
	}
}

/**
 * 關閉所有連線
 */
func CloseConn() {
	for k, v := range pool {
		logger.Debug("close db conn", zap.String("db", k))
		v.Close()
	}
}
