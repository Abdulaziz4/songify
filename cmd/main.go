package main

import (
	"github.com/Abdulaziz4/songify/config"
	"github.com/Abdulaziz4/songify/model"
	"github.com/Abdulaziz4/songify/router"
	"github.com/Abdulaziz4/songify/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	print(c.DbHost)

	dsn := "host=" + c.DbHost + " user=" + c.DbUser + " password=" + c.DbPassword + " dbname=" + c.DbName + " port=" + c.DbPort + " sslmode=disable TimeZone=Asia/Shanghai"

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

	r.Run(":" + c.Port)
}
