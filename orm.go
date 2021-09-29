package main

import (
	"database/sql"
	"fmt"
)

type ORM struct {
	db *sql.DB
}

func (orm *ORM) CreateDDL() (int64, error) {
	prepare, err := orm.db.Prepare("CREATE TABLE IF NOT EXISTS tasks(id INTEGER PRIMARY KEY, text VARCHAR(64))")
	if err != nil {
		return 0, err
	}
	exec, err := prepare.Exec()
	if err != nil {
		return 0, err
	}
	fmt.Println("DDL created")
	return exec.RowsAffected()
}

func (orm *ORM) Insert(task InsertTask) (int64, error) {
	statement, _ := orm.db.Prepare("INSERT INTO tasks (text) VALUES (?)")
	exec, err := statement.Exec(task.text)
	if err != nil {
		return 0, err
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func (orm *ORM) SelectAll() ([]Task, error) {
	rows, err := orm.db.Query("SELECT id, text FROM tasks")
	if err != nil {
		panic("failed to rows select")
	}
	var task Task
	var tasks []Task
	for rows.Next() {
		rows.Scan(&task.id, &task.text)
		tasks = append(tasks, task)
	}
	return tasks, nil
}
