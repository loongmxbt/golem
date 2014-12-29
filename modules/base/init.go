package base

import (
	"github.com/go-xorm/xorm"
	"github.com/loongmxbt/golem/models"
	"github.com/loongmxbt/golem/modules/config"
)

var engine *xorm.Engine

func GlobalInit() {
	// Load Config
	config.LoadConfig()
	// InitORM
	models.InitORM()

}
