package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

// data key must be small case and struct must be capital
// var data = `
// appver: 0.1.0
// appname: Golem CMS
// `

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
		log.Fatal(4, "Fail to load config file: %v", err)
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
