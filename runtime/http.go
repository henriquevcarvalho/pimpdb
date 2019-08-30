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
var Log *log.Logger

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

type SessionCache struct {
	User		string
	Id 			string
	Sid 		string
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

	Log = log.New(f, "prefix", log.LstdFlags)

	db = pimpdb.PimpDB{}.Init()

	e = echo.New()
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
	if x, found := db.Get(id); found {
		Log.Println("[x] Getting Hoe nr: " + id, x)
		return c.JSON(http.StatusOK, x.(*SessionCache).Id)
	}

	Log.Println("[x] Failed to pimp: " + id)
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

	err := db.Save(x.Sid, x)
	if err != nil {
		Log.Println("[x] Failed to pimp: ", err)
		return c.JSON(http.StatusOK, false)
	}

	Log.Println("[x] Saving Hoe nr: " + x.Sid, x)
	return c.JSON(http.StatusOK, true)
}

