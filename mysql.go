package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type DbWorker struct {
	//mysql data source name
	Dsn string
}

func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}




func main() {
	dbw := DbWorker{
		Dsn: "root:root@tcp(127.0.0.1:3306)/test",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	check(err)

	rows, err := db.Query("SELECT * FROM test.announcement")
	check(err)

	for rows.Next() {

		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	rows.Close()




	defer db.Close()


}

