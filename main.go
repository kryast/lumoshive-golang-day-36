package main

import (
	"day-36/router"
	"net/http"
)

func main() {
	r, log, err := router.InitRouter()
	if err != nil {
		log.Panic(err.Error())

	}

	log.Info("Server run on post 8080")
	http.ListenAndServe(":8080", r)
}
