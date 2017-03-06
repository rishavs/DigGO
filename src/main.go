package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


//------------------------------------
// Model
//------------------------------------

type (
	post struct {
		ID   int    `json:"id"`
		Title string `json:"title"`
		// Content string `json:"content"`
		// Link string `json:"link"`
		// Is_deleted bool `json:"is_deleted"`
		// Created_at string `json:"created_at"`
		// Updated_at string `json:"updated_at"`
		// Viewed_at string `json:"viewed_at"`
	}
)

var (
	posts = map[int]*post{}
	seq   = 1
)




//------------------------------------
// Handlers
//------------------------------------

func createPost(c echo.Context) error {
	p := &post{
		ID: seq,
	}
	if err := c.Bind(p); err != nil {
		return err
	}
	posts[p.ID] = p
	seq++
	return c.JSON(http.StatusCreated, p)
}

func getPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, posts[id])
}

func updatePost(c echo.Context) error {
	p := new(post)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	posts[id].Title = p.Title
	return c.JSON(http.StatusOK, posts[id])
}

func deletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(posts, id)
	return c.NoContent(http.StatusNoContent)
}



func request(c echo.Context) error {
	req := c.Request()
	format := "<pre><code>Protocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n</code></pre>"
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
}


//------------------------------------
// Main
//------------------------------------

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", createPost)
	e.GET("/p/:id", getPost)
	e.PUT("/p/:id", updatePost)
	e.DELETE("/p/:id", deletePost)


	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}