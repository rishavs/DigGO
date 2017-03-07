package controllers

import (
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
	"net/http"
)

//create user
func IndexPosts(c echo.Context) error {
	return c.String(http.StatusOK, "Index")
}

//create user
func CreatePost(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}

func GetPost(c echo.Context) error {
	return c.String(http.StatusOK, "show")
}
