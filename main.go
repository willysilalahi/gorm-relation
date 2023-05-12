package main

import (
	"fmt"
	"gorm-relation/database"
	"gorm-relation/model"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Clear")
	msg_migrate := database.Migrate()
	msg_seed := database.Seed()

	fmt.Println(msg_migrate)
	fmt.Println(msg_seed)

	var tags []model.Tag
	db, _ := database.GetDB()

	data := db.Preload("Posts").Find(&tags).RowsAffected
	if data == 0 {
		fmt.Println("Gak ada data nya bang")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": tags,
		})
	})
	r.Run()
}
