package setting

import (
	"log"

	"gopkg.in/yaml.v2"
)

// data key must be small case and struct must be capital
var data = `
appver: 0.1.0
appname: Golem CMS
`

type ConfigFile struct {
	AppVer  string
	AppName string
}

func LoadConfig() {
	Cfg := ConfigFile{}

	err := yaml.Unmarshal([]byte(data), &Cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("%v", Cfg.AppName)
}
