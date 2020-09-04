package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"wumiao/config"
)

var Db *xorm.EngineGroup

func init() {

	m := config.MasterDbConfig
	s := config.SlaveDbConfig
	conn := []string{
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", m.User, m.Pwd, m.Host, m.Port, m.DbName), // 第一个默认是master
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", s.User, s.Pwd, s.Host, s.Port, s.DbName), // 第二个开始都是slave
	}

	engine, err := xorm.NewEngineGroup(config.DriverName, conn)

	if err != nil {
		log.Fatal("dbhelper.DbInstanceMaster,", err)
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	//engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cache)
	Db = engine
}

func GetMysqlGroup() *xorm.EngineGroup {
	return Db
}
