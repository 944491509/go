package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 连接池对象

func insert() (err error) {
	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	if err != nil {
		return
	}
	_, err = stmt.Exec("wenhao", "研发部门", "2012-12-09")
	if err != nil {
		return
	}
	return

}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go"
	db, err = sql.Open("mysql", dsn) // open 不会去校验数据是否正确
	if err != nil {
		// dsn 格式不正确的时候会报错
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func query() (err error) {
	// sql语句
	sqlStr := `select * from userinfo;`
	// 执行
	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	// 查看结果
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			return
		}
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功!")

	err = insert()
	if err != nil {
		fmt.Printf("insert db failed, err:%v\n", err)
		return
	}
	fmt.Println("新增数据成功!")
	err = query()
	if err != nil {
		fmt.Printf("select db failed, err:%v\n", err)
		return
	}

}
