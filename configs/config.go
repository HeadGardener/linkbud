package configs

import "github.com/spf13/viper"

type DBConfig struct {
	Host    string
	DBName  string
	SSLMode string
}

type ServerConfig struct {
	Port string
}

func (c *DBConfig) Init() {
	c.Host = viper.GetString("db.host")
	c.DBName = viper.GetString("db.name")
	c.SSLMode = viper.GetString("db.sslmode")
}

func (c *ServerConfig) Init() {
	c.Port = viper.GetString("port")
}
