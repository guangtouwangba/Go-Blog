package conf

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strconv"
)

type HttpServer struct {
	Addr string `yaml:"addr,omitempty" default:"0.0.0.0"`
	Port int64  `yaml:"port"`
}

func (h *HttpServer) GetServerConfig() string {
	if h.Addr == "" {
		return "0.0.0.0" + ":" + strconv.FormatInt(h.Port, 10)
	}
	return h.Addr + ":" + strconv.FormatInt(h.Port, 10)
}

type Db interface {
	GetDb() string
}

type Mysql struct {
	URL      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (m *Mysql) GetDb() string {
	return m.URL
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
	log.Println(string(yamlFile))
	err = yaml.NewDecoder(bytes.NewReader(yamlFile)).Decode(conf)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("load http server address:", conf.HttpServer.GetServerConfig())
	log.Println("load mysql url:", conf.Mysql.URL)
	return conf
}
