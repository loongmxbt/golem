package setting

import (
	"log"

	"github.com/go-xorm/xorm"
	"github.com/loongmxbt/golem/models"
)

var engine *xorm.Engine

func GlobalInit() {
	// Load Config
	LoadConfig()
	// InitORM
	InitORM()

}

func InitORM() {
	var engine *xorm.Engine
	if err := models.NewEngine(engine); err != nil {
		log.Fatal(4, "Fail to initialize ORM engine: %v", err)
	}
}

func LoadConfig() {

}
