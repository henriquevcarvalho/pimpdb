package main

import (
	"encoding/json"
	"fmt"
	"github.com/badtheory/pimpdb"
	"github.com/badtheory/worst"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
	"os"
)

var err error
var db *pimpdb.PimpDB
var ws *worst.Worst

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


	f, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println("[x] PimpDB logs started")

	db = pimpdb.New()
	db.SetCacheOptions()

	ws = worst.New(worst.Options{
		Server: &http.Server{
			Addr: "localhost:1339",
		},
		Static: worst.Static{
			Url: "/public/*",
			Path: "/home/paulo/Desktop/public",
		},
	})

	ws.Router.Get("/get/{id}", get)
	ws.Router.Post("/save", save)
	ws.Run()

}

func get(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	if x, found := db.Cache.Get(id); found {
		log.Println("[x] Getting Hoe nr: " + id, x)
		ws.Router.Render.JSON(w, http.StatusOK, x.(*SessionCache).Id)
		return
	}

	log.Println("[x] Failed to pimp: " + id)
	ws.Router.Render.JSON(w, http.StatusOK, false)
	return

}

func save(w http.ResponseWriter, r *http.Request) {

	var x SessionCache
	if err = json.NewDecoder(r.Body).Decode(&x); err != nil {
		ws.Router.Render.JSON(w, http.StatusOK, ResponseError{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	err := db.Cache.Save(x.Sid, x)
	if err != nil {
		log.Println("[x] Failed to pimp: ", err)
		ws.Router.Render.JSON(w, http.StatusOK, false)
		return
	}

	log.Println("[x] Saving Hoe nr: " + x.Sid, x)
	ws.Router.Render.JSON(w, http.StatusOK, true)
	return

}

