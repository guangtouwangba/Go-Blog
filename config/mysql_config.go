package config

type Db interface {
	GetDb() string
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
