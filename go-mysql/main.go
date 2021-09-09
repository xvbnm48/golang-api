package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("GO Mysql Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testDb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// fmt.Println("success connected to database")
	// insert, err := db.Query("INSERT INTO user VALUES('SAKURA ENDO')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()
	// fmt.Println("success insert data")

	results, err := db.Query("SELECT name FROM user")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
