package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 调用init函数
	"log"
)


const (
	getAllByProductIDAndCustomerID = `select * from p_orders where product_id in (:product_id) and customer_id=:customer_id`
)

// GetAllByProductIDAndCustomerID
// @param driver_id
// @param rate_date
// @return []Order, error
func GetAllByProductIDAndCustomerID(ctx context.Context, productIDs []uint64, customerID uint64) ([]Order, error) {
	var orderList []Order

	params := map[string]interface{}{
		"product_id" : productIDs,
		"customer_id": customerID,
	}

	// getAllByProductIDAndCustomerID 是 const 类型的 sql 字符串
	sql, args, err := sqlutil.Named(getAllByProductIDAndCustomerID, params)
	if err != nil {
		return nil, err
	}

	err = dao.QueryList(ctx, sqldbInstance, sql, args, &orderList)
	if err != nil {
		return nil, err
	}

	return orderList, err
}

func main() {
	// db 是一个 sql.DB 类型的对象
	// 该对象线程安全，且内部已包含了一个连接池
	// 连接池的选项可以在 sql.Open 中设置，这里为了简单省略了
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id int
		name string
	)
	rows, err := db.Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// 必须要把 rows 里的内容读完，否则连接永远不会释放, 迭代器
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}