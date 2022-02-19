package config

import "strconv"

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
