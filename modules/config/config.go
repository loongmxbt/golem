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

	// i18n
	Langs, Names []string
)

func LoadConfig() {
	conf, err := toml.LoadFile("config/app.toml")

	if err != nil {
		log.Fatalf("Fail to load config file: %v", err)
	} else {
		// retrieve data directly
		AppVer = conf.Get("app.ver").(string)
		AppName = conf.Get("app.name").(string)
		log.Print(AppName, " Ver: ", AppVer)

		DbHost = conf.Get("db.host").(string)
		DbDriver = conf.Get("db.driver").(string)
		DbUname = conf.Get("db.uname").(string)
		DbPasswd = conf.Get("db.passwd").(string)
		DbDbname = conf.Get("db.dbname").(string)

		LangsInter := conf.Get("i18n.langs").([]interface{})
		NamesInter := conf.Get("i18n.names").([]interface{})

		if len(LangsInter) != len(NamesInter) {
			log.Println("Langs and Names mismatch!")
		}

		Langs = make([]string, len(LangsInter))
		Names = make([]string, len(NamesInter))

		for i := range LangsInter {
			Langs[i] = LangsInter[i].(string)
			Names[i] = NamesInter[i].(string)
		}
	}
}
