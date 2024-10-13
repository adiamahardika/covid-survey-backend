package main

import (
	"backend/routes"
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// Koneksi ke SQLite
	db, err := sql.Open("sqlite", "./databases/survey.db")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to SQLite")
	}
	defer db.Close()

	// Inisialisasi routes API
	routes.InitRoutes(router, db)

	// Menjalankan server
	router.Run(":5000")
}
