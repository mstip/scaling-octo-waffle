package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Todo struct {
	ID   int64  `db:"id"`
	Text string `db:"text"`
	Done bool   `db:"done"`
}

func prepareDB(db *sqlx.DB) error {
	createTodoTable := `
	CREATE TABLE IF NOT EXISTS todos
	(
		id bigserial NOT NULL,
		text text,
		done boolean,
		PRIMARY KEY (id)
	);
	`

	_, err := db.Exec(createTodoTable)
	return err
}

func AddTodo(db *sqlx.DB, todo *Todo) (*Todo, error) {
	result, err := db.NamedQuery(
		"INSERT INTO todos(text, done) VALUES(:text, :done) RETURNING id, text, done",
		&todo,
	)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		err = result.StructScan(todo)
		if err != nil {
			return nil, err
		}
	}
	return todo, nil
}

func GetAllTodos(db *sqlx.DB) ([]Todo, error) {
	var todos []Todo
	err := db.Select(&todos, "SELECT id, text, done FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	return todos, err
}

func GetTodoById(db *sqlx.DB, id int64) (*Todo, error) {
	var todo Todo
	err := db.Get(&todo, "SELECT id, text, done FROM todos WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func DeleteTodo(db *sqlx.DB, todo *Todo) error {
	_, err := db.NamedExec("DELETE FROM todos where id=:id", todo)
	return err
}

func UpdateTodo(db *sqlx.DB, todo *Todo) error {
	_, err := db.NamedExec("UPDATE todos SET text=:text, done=:done where id=:id", todo)
	return err
}

func main() {
	db, err := sqlx.Connect(
		"postgres",
		"user=postgres dbname=postgres password=qqqq sslmode=disable",
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err := prepareDB(db); err != nil {
		log.Fatal(err)
	}

	todoItem := &Todo{
		Text: "Woop",
		Done: false,
	}
	todoItem, err = AddTodo(db, todoItem)
	if err != nil {
		log.Fatal(err)
	}

	todos, err := GetAllTodos(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(todos)

	aTodo, err := GetTodoById(db, 3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(aTodo)

	err = DeleteTodo(db, todoItem)
	if err != nil {
		log.Fatal(err)
	}
	aTodo.Text = "changed!"
	aTodo.Done = true
	err = UpdateTodo(db, aTodo)
	if err != nil {
		log.Fatal(err)
	}
}
