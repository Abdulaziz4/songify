package main

import (
	"os"

	"github.com/Abdulaziz4/songify/model"
	"github.com/Abdulaziz4/songify/router"
	"github.com/Abdulaziz4/songify/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// c, err := config.LoadConfig()
	// if err != nil {
	// 	panic(err)
	// }
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	// Establish connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Song{})

	// Bootstrap services
	songService := service.NewSongService(db)

	// Bootstrap routers
	songRouter := router.NewSongRouter(songService)

	// Bootstrap gin
	r := gin.Default()

	r.GET("/songs", songRouter.Get)
	r.POST("/songs", songRouter.Post)
	r.GET("/songs/:id", songRouter.GetById)
	r.DELETE("/songs/:id", songRouter.Delete)

	r.Run(":" + os.Getenv("PORT"))
}
