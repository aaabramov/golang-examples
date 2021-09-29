package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting")
	db, _ := sql.Open("sqlite3", "./data.sqlite3")

	orm := ORM{db}

	newTask := InsertTask{text: "Task #2"}
	affected, _ := orm.Insert(newTask)
	fmt.Println("Inserted ", affected, " rows")

	all, _ := orm.SelectAll()
	for _, t := range all {
		fmt.Println("Loaded task: id=", t.id, ",text=", t.text)
	}

}
