package main

import (
	"database/sql"
	"log"

	books "github.com/Mictrlan/library/books/controller/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	dbConn, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/library")
	if err != nil {
		log.Fatal(err)
	}

	booksCon := books.New(dbConn)
	booksCon.Register(router)

	router.Run(":8080")
}
