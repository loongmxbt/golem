package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

var (
	AppVer  string
	AppName string

	DbHost   string
	DbDriver string
	DbUname  string
	DbPasswd string
	DbDbname string
)

func LoadConfig() {
	config, err := toml.LoadFile("config/config.toml")

	if err != nil {
		log.Fatalf("Fail to load config file: %v", err)
	} else {
		// retrieve data directly
		AppVer = config.Get("app.ver").(string)
		AppName = config.Get("app.name").(string)
		log.Print(AppName, " Ver: ", AppVer)

		DbHost = config.Get("db.host").(string)
		DbDriver = config.Get("db.driver").(string)
		DbUname = config.Get("db.uname").(string)
		DbPasswd = config.Get("db.passwd").(string)
		DbDbname = config.Get("db.dbname").(string)

	}
}
