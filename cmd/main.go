package main

import (
	"Go-Blog/config"
	"Go-Blog/internal/adapter/inbound/rest/routing"
	"Go-Blog/internal/adapter/outbound/db"
	"log"
)

func init() {
	dbTool := db.MysqlTool{}
	dbTool.CreateDatabaseConnection()
}

func main() {
	r := routing.InitRouter()
	log.Println("Server is running on  " + config.GetYamlConfig().HttpServer.GetServerConfig())
	err := r.Run(config.GetYamlConfig().HttpServer.GetServerConfig())
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
