package conf

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strconv"
)

type HttpServer struct {
	Host string `yaml:"addr,omitempty" default:"0.0.0.0"`
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

type Mysql struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
	Params   string `yaml:"params，omitempty"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (m *Mysql) GetDb() string {
	log.Println("mysql url:", m.URL)
	log.Println("mysql user:", m.User)
	log.Println("mysql password:", m.Password)
	dbAddr := m.User + ":" + m.Password + "@tcp(" + m.URL + ")" + "/"
	return dbAddr
}

type YmlConfig struct {
	HttpServer
	Mysql
}

func GetYamlConfig() *YmlConfig {
	conf := new(YmlConfig)
	yamlFile, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.NewDecoder(bytes.NewReader(yamlFile)).Decode(conf)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("load http server address:", conf.HttpServer.GetServerConfig())
	log.Println("load mysql url:", conf.Mysql.URL)
	return conf
}
