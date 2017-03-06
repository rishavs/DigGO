package main

import (
	"net/http"
	// "strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//------------------------------------
// Model
//------------------------------------

//------------------------------------
// Main
//------------------------------------

func main() {

	data, err := ioutil.ReadFile("config.yml")
	checkErr(err)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes

	e.Logger.Fatal(e.Start(":3000"))
}
