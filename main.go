package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "goapp:password123@tcp(127.0.0.1:3306)/subscribers"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("БД подключена")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Сервер работает, брат 🚀")
	})

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", nil)
}
