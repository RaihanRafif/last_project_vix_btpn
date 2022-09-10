package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/asaskevich/govalidator"

	"os"
	"tugas_akhir/controllers"
	"tugas_akhir/database"
	"tugas_akhir/middlewares"
	"tugas_akhir/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func drop(db *gorm.DB) {
	db.DropTableIfExists(
		&models.FileUpload{},
		&models.User{},
	)
}

func migrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.FileUpload{})
}

func create(database *gorm.DB) {
	drop(database)
	migrate(database)
}

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	println(os.Getenv("DB_DIALECT"))
	database := database.OpenDbConnection()
	defer database.Close()
	args := os.Args
	if len(args) > 1 {
		first := args[1]
		second := ""
		if len(args) > 2 {
			second = args[2]
		}

		if first == "create" {
			create(database)
		} else if first == "seed" {
			seeds.Seed()
			os.Exit(0)
		} else if first == "migrate" {
			migrate(database)
		}

		if second == "seed" {
			seeds.Seed()
			os.Exit(0)
		} else if first == "migrate" {
			migrate(database)
		}

		if first != "" && second == "" {
			os.Exit(0)
		}
	}
	migrate(database)
	goGonicEngine := gin.Default()
	goGonicEngine.Use(cors.Default())
	goGonicEngine.Use(middlewares.Benchmark())
	goGonicEngine.Use(middlewares.UserLoaderMiddleware())
	goGonicEngine.Static("/static", "./static")
	apiRouteGroup := goGonicEngine.Group("/api")

	controllers.RegisterUserRoutes(apiRouteGroup.Group("/users"))
	controllers.RegisterPhotoRoutes(apiRouteGroup.Group("/photos"))

	goGonicEngine.Run(":8080") // listen and serve on 0.0.0.0:8080
	fmt.Print("RAIHAN")
}
