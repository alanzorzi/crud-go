package main

import (
	"database/sql"

	"github.com/alanzorzi/crud-go/app/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Importação anônima para registrar o driver
)

var (
	db *sql.DB
)

func main() {

	var err error

	db, err = sql.Open("mysql", "root:teste123@tcp(localhost:3306)/users_db")
	if err != nil {

		panic(err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r, db)

	r.Run(":8080")
}
