/**
 * 資料庫連線池，用於控管各db的連線
 */
package database

import (
	"fmt"
	"time"

	"github.com/MilosLin/go_bananas/core/config"
	"github.com/MilosLin/go_bananas/core/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

func init() {
	pool = make(map[string]*gorm.DB)
	if limit := config.Instance().GetInt("database.query_limit"); limit > 0 {
		Query_limit = limit
	} else {
		Query_limit = 500000
	}
}

var pool map[string]*gorm.DB
var Query_limit int

type DBConnPool struct {
}

/**
 * 依照資料庫名稱取得DB連線
 */
func GetConn(db_name string) *gorm.DB {
	if value, ok := pool[db_name]; ok {
		return value
	} else {
		c := config.Instance()
		group := fmt.Sprintf("database.%s", db_name)

		conn_str := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=UTC",
			c.GetString(group+".account"),
			c.GetString(group+".password"),
			c.GetString(group+".ip"),
			c.GetString(group+".port"),
			c.GetString(group+".db"),
		)
		db, _ := gorm.Open("mysql", conn_str)

		db.DB().SetMaxOpenConns(c.GetInt("database.max_open_conns"))
		db.DB().SetMaxIdleConns(c.GetInt("database.max_idle_conns"))
		db.DB().SetConnMaxLifetime(time.Second * c.GetDuration("database.conn_max_life_time"))
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
