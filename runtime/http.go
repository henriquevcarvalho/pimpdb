package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"pimpdb"
)

var err error
var db *pimpdb.PimpDB
var e *echo.Echo

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

type SessionCache struct {
	User string
	Id   string
	Sid  string
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

	f, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	db = pimpdb.New()
	db.SetCacheOptions()
	db.SetLoggerOptions()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.GET("/get/:id", get)
	e.POST("/save", save)

	e.Logger.SetOutput(f)
	e.Logger.Fatal(e.Start(":1328"))

}

func get(c echo.Context) error {

	id := c.Param("id")
	if x, found := db.Cache.Get(id); found {
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

	 if exists := db.Cache.Set(x.Sid, x); exists {
		 return c.JSON(http.StatusOK, true)
	 }

	return c.JSON(http.StatusOK, false)
}
