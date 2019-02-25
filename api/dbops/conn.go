package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConnent *sql.DB
	err       error
)

func init() {
	dbConnent, err = sql.Open("mysql", "root:freeman123@/video_service?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
