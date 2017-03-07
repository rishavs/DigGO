package main

import (
	"fmt"
	"log"
	// "net/http"
	// "strconv"
	// "database/sql"
	// _ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"./controllers"
	"./models"
	// "./views"
)

//------------------------------------
// Main
//------------------------------------

func main() {

	//------------------------------------
	// Read Config
	//------------------------------------

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db_uri, ok := viper.Get("PG_URL").(string)
	if ok != true {
		log.Fatal(err)
	}

	fmt.Println(db_uri)

	//------------------------------------
	// Connect to Db
	//------------------------------------

	db, err := gorm.Open("postgres", db_uri)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected successfully to DB")
	}

	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Post{})

	//------------------------------------
	// Create Server
	//------------------------------------
	e := echo.New()

	//------------------------------------
	// Register Middlewares
	//------------------------------------
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//------------------------------------
	// Routes
	//------------------------------------
	e.GET("/p/", controllers.IndexPosts)
	e.GET("/p/:id", controllers.GetPost)
	e.POST("/p/:id", controllers.CreatePost)

	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	//------------------------------------
	// Start Server
	//------------------------------------
	e.Logger.Fatal(e.Start(":3000"))
}
