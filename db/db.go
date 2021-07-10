package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chiwon99881/todolist/env"
	"github.com/chiwon99881/todolist/types"
)

// DB is connection func for database
func DB() *sql.DB {
	env.Start()
	pqConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", pqConn)

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}

	return db
}

// Close is func that terminate database
func Close() {
	DB().Close()
}

// SelectAllToDo is excute select SQL
func SelectAllToDo() []*types.ToDo {
	var toDos []*types.ToDo
	stmt := `select * from "todo"`
	rows, err := DB().Query(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		toDo := &types.ToDo{}
		rows.Scan(&toDo.ID, &toDo.Caption, &toDo.Excute)
		toDos = append(toDos, toDo)
	}
	return toDos
}

// SelectToDo is find To Do
func SelectToDo(ID int) *types.ToDo {
	toDo := &types.ToDo{}
	stmt := `select * from "todo" where "id"=$1`
	row := DB().QueryRow(stmt, ID)

	row.Scan(&toDo.ID, &toDo.Caption, &toDo.Excute)
	return toDo
}

// InsertToDo is excute insert SQL
func InsertToDo(caption string, check bool) {
	stmt := `insert into "todo"("caption", "excute") values($1, $2)`
	_, err := DB().Exec(stmt, caption, check)
	if err != nil {
		panic(err.Error())
	}
}

// UpdateToDo is true -> false or false -> true your ToDo
func UpdateToDo(toDoID int, check bool) {
	toggleExcute := !check
	stmt := `update "todo" set "excute"=$2 where "id"=$1`
	_, err := DB().Exec(stmt, toDoID, toggleExcute)
	if err != nil {
		panic(err.Error())
	}
}