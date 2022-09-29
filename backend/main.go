package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) handler(c *gin.Context) {
	var todos []Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func (r *repository) postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := Todos{
		Task: data.Task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": todo})
}

type Todos struct {
	gorm.Model
	Task string `json:"task"`
	Done bool   `json:"done"`
}

type DataRequest struct {
	Task string `json:"task" binding:"required"`
}

func main() {

	// db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	if err := sqlDB.Ping(); err != nil {
		panic("failed to ping database")
	}

	if err := db.AutoMigrate(&Todos{}); err != nil {
		panic("failed to migrate database")
	}

	if err := db.AutoMigrate(&Todos{}); err != nil {
		panic("failed to migrate database")
	}

	// db = make([]string, 0)
	repo := repository{
		db: db,
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/", repo.handler)
	r.POST("/send", repo.postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
