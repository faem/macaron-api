package models

import (
	"log"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func init() {
	x, err := xorm.NewEngine("postgres", "host=localhost port=5432 user=fahim password=1234 dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	engine = x

	engine.SetLogger(nil)
	engine.SetMapper(core.GonicMapper{})

	err = engine.Sync(new(Profile))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
