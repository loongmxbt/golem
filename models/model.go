package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"fmt"
)

var (
	tables    []interface{}
	HasEngine bool
)

func init() {
	tables = append(tables, new(User), new(Node))
}

func NewEngine(engine *xorm.Engine) (err error) {
	engine, err = xorm.NewEngine("mysql", "root:mxbtsql1014@/GolemCMS?charset=utf8")

	if err != nil {
		return fmt.Errorf("models.init(fail to connect to database): %v", err)
	}

	return engine.Sync2(tables...)
}
