package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang/VideoServer/api/conf"
)


var (
	db *sql.DB
	err error
)


func init() {
	db, err = sql.Open("mysql", conf.DBSC)
	if err != nil {
		panic(err.Error())
	}
}

func ClearTables()  {
	db.Exec("TRUNCATE users")
	db.Exec("TRUNCATE video_info")
	db.Exec("TRUNCATE comments")
	db.Exec("TRUNCATE session")
}