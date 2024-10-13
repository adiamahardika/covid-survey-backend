package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Survey struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Affected bool   `json:"affected"`
}

func InitRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/surveys", func(c *gin.Context) {
		// Mendapatkan daftar data
		rows, err := db.Query("SELECT id, name, age, address, affected FROM surveys")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var surveys []Survey
		for rows.Next() {
			var id, age int
			var name, address string
			var affected bool

			rows.Scan(&id, &name, &age, &address, &affected)

			surveys = append(surveys, Survey{
				ID:       id,
				Name:     name,
				Age:      age,
				Address:  address,
				Affected: affected,
			})
		}

		c.JSON(200, surveys)
	})

	router.POST("/surveys", func(c *gin.Context) {
		// Menambahkan data baru
		var survey Survey

		if err := c.BindJSON(&survey); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("INSERT INTO surveys (name, age, address, affected) VALUES (?, ?, ?, ?)",
			survey.Name, survey.Age, survey.Address, survey.Affected)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Data added successfully"})
	})

	router.GET("/surveys/:id", func(c *gin.Context) {
		id := c.Param("id")
		var survey Survey

		err := db.QueryRow("SELECT id, name, age, address, affected FROM surveys WHERE id = ?", id).
			Scan(&survey.ID, &survey.Name, &survey.Age, &survey.Address, &survey.Affected)
		if err != nil {
			c.JSON(404, gin.H{"error": "Data not found"})
			return
		}

		c.JSON(200, survey)
	})

	router.PUT("/surveys/:id", func(c *gin.Context) {
		id := c.Param("id")
		var survey Survey

		if err := c.BindJSON(&survey); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE surveys SET name = ?, age = ?, address = ?, affected = ? WHERE id = ?",
			survey.Name, survey.Age, survey.Address, survey.Affected, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Data updated successfully"})
	})

	router.DELETE("/surveys/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM surveys WHERE id = ?", id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Data deleted successfully"})
	})
}
