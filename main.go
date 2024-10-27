package main

import (
	"backend/databases"
	"backend/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Postgres")
	}
	defer db.Close()

	// Inisialisasi routes API
	routes.InitRoutes(router, db)

	// Menjalankan server
	router.Run(":5000")
}
