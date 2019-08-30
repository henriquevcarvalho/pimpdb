package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"pimpdb"
)

var err error
var db *pimpdb.PimpDB

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

type SessionCache struct {
	User		string
	Id 			string
}

func rescue() {
	r := recover()
	if r != nil {
		fmt.Println("Panic has been recover.. why?", err)
		main()
	}
}

func main() {

	defer rescue()

	db = pimpdb.PimpDB{}.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.GET("/get/:id", get)
	e.POST("/save", save)

	e.Logger.Fatal(e.Start(":1328"))

}

func get(c echo.Context) error {

	id := c.Param("id")

	if x, found := db.Get(id); found {
		return c.JSON(http.StatusOK, x.(*SessionCache).Id)
	}

	return c.JSON(http.StatusOK, false)
}

func save(c echo.Context) error {

	x := new(SessionCache)
	if err = c.Bind(x); err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{
			http.StatusInternalServerError,
			err.Error(),
		})
	}

	err := db.Save( x.Id, x)

	if err != nil {
		return c.JSON(http.StatusOK, false)
	}

	return c.JSON(http.StatusOK, true)
}

