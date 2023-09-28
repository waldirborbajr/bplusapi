package main

import (
	"fmt"
	"net/http"

	"github.com/waldirborbajr/bplusapi/internal/handler"
)

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	addr := "3030"

	router := handler.NewHandler()

	err := http.ListenAndServe(fmt.Sprintf(":%s", addr), router)
	catch(err)
}
