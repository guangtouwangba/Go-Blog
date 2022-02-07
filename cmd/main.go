package main

import (
	"Go-Blog/internal/adapter/inbound/rest/routing"
)

func main() {
	r := routing.InitRouter()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
