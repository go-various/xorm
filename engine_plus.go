package xorm

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-various/xorm/log"
	"io"
)

var ErrInitEngine = errors.New("init xorm engine: config or config.Url can not be null")


func NewEnginePlus(c *Config, w ...io.Writer) (EngineInterface, error) {
	if nil == c || "" == c.Master {
		return nil, ErrInitEngine
	}
	var err error
	var engine EngineInterface
	if nil == c.Slaves || len(c.Slaves) == 0{
		engine, err = NewEngine(c.Driver, c.Master)
	}else {
		conns := []string{c.Master}
		conns = append(conns, c.Slaves...)
		engine, err = NewEngineGroup(c.Driver,  conns)
	}
	if err != nil {
		return nil, err
	}

	if len(w) > 0{
		engine.SetLogger(log.NewSimpleLogger(io.MultiWriter(w...)))
	}

	engine.ShowSQL(c.ShowSql)
	engine.SetMaxIdleConns(c.MaxIdle)
	engine.SetMaxOpenConns(c.MaxConn)

	return engine, nil
}