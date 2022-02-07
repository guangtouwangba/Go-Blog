package main

import (
	"Go-Blog/conf"
	"Go-Blog/internal/adapter/inbound/rest/routing"
)

func main() {
	r := routing.InitRouter()
	err := r.Run(conf.GetYamlConfig().HttpServer.GetServerConfig())
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
