package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/loongmxbt/golem/modules/config"
)

var (
	tables    []interface{}
	HasEngine bool
)

// Where call init() ?
func init() {
	tables = append(tables, new(User), new(Node))
}

func NewEngine(engine *xorm.Engine) (err error) {

	connstr := ""
	switch config.DbDriver {
	case "mysql":
		connstr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			config.DbUname, config.DbPasswd, config.DbHost, config.DbDbname)

	case "postgres":
		// TODO
	case "sqlite3":
		// TODO
	default:
		return fmt.Errorf("Unknown database type: %s", config.DbDriver)
	}

	engine, err = xorm.NewEngine(config.DbDriver, connstr)
	if err != nil {
		return fmt.Errorf("models.init(fail to connect to database): %v", err)
	}

	return engine.Sync2(tables...)
}

func InitORM() {
	var engine *xorm.Engine
	if err := NewEngine(engine); err != nil {
		log.Fatalf("Fail to initialize ORM engine: %v", err)
	}
	HasEngine = true
	log.Println("Init ORM succeed!")
}
