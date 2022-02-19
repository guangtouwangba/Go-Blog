package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type YmlConfig struct {
	HttpServer
	Mysql
	Redis
}

func GetYamlConfig() *YmlConfig {
	conf := new(YmlConfig)
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.NewDecoder(bytes.NewReader(yamlFile)).Decode(conf)
	if err != nil {
		log.Panicln(err)
	}
	//log.Println("config:", conf)
	//log.Println("load http server address:", conf.HttpServer.GetServerConfig())
	//log.Println("load mysql url:", conf.Mysql.URL)
	return conf
}
