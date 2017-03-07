package main

import (
	"fmt"
	"log"
	// "net/http"
	// "strconv"
	// "database/sql"
	// _ "github.com/lib/pq"

	"github.com/spf13/viper"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

//------------------------------------
// Model
//------------------------------------



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
	if ok != true {log.Fatal(err)}

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

	type Post struct {
		gorm.Model

		Title string
		Content string
		User User
		UserId uint
	}

	type User struct {
	    gorm.Model

	    Name string
	}

	db.Model(&post).Related(&user)

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




	//------------------------------------
	// Start Server
	//------------------------------------
	// e.Logger.Fatal(e.Start(":3000"))
}

