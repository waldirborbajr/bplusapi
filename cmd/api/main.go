package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

var router *chi.Mux

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	err := http.ListenAndServe(":8005", router)
	catch(err)
}
