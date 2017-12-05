package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	xrCache "github.com/go-xorm/xorm-redis-cache"
	"gopkg.in/logger.v1"
)

var engine *xorm.Engine

type User2 struct {
	Id        int64
	Name      string    `xorm:"varchar(50) notnull unique 'user_name'";json:"name"`
	CreatedAt time.Time `xorm:"created"`
}

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}

func init() {
	//	gob.Register(User2{})
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		//todo something
	}

	//connection pool settings
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(10)
	log.Info("_")
	xormLogger := engine.Logger()
	//xormLogger.SetLevel(core.LOG_DEBUG)
	xormLogger.SetLevel(core.LOG_DEBUG)
	engine.ShowSQL(true)
	engine.ShowExecTime(true)

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "test_")
	//      cacheMapper := core.NewCacheMapper(core.SnakeMapper{})
	engine.SetTableMapper(tbMapper)
	cacher := xrCache.NewRedisCacher("127.0.0.1:6379", "", xrCache.DEFAULT_EXPIRATION, xormLogger)
	engine.SetDefaultCacher(cacher)
	//cacher := xorm.NewLRUCacher2(xorm.(), 1000)

	//engine.MapCacher(&User2{}, cacher)

}
func main() {
	user := &User2{}
	engine.Id(1).Get(user)
	log.Info(user)

	engine.Id(1).Get(user)
	fmt.Println("vim-go")
}
