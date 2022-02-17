package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strconv"
)

type HttpServer struct {
	Host string `yaml:"host,omitempty"`
	Port int64  `yaml:"port"`
}

func (h *HttpServer) GetServerConfig() string {
	if h.Host == "" {
		return "0.0.0.0" + ":" + strconv.FormatInt(h.Port, 10)
	}
	return h.Host + ":" + strconv.FormatInt(h.Port, 10)
}

type Db interface {
	GetDb() string
}

type YmlConfig struct {
	HttpServer
	Mysql
	Redis
}

type Mysql struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
	Params   string `yaml:"params,omitempty"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Update   bool   `yaml:"update"`
}

func (m *Mysql) GetDb() string {
	//log.Println("mysql url:", m.URL)
	//log.Println("mysql user:", m.User)
	//log.Println("mysql password:", m.Password)
	dbAddr := m.User + ":" + m.Password + "@tcp(" + m.URL + ")" + "/"
	return dbAddr
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
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
