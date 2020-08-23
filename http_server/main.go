package main

import (
	"arc/src/env"
	"arc/src/handlers"
	"arc/src/managers"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	confPath = "conf/conf.json"
)

func init() {
	InitConf()
	InitLogger()
	InitTableManager()
}

func InitLogger() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}

func InitConf() {
	if err := env.Conf.Load(confPath); err != nil {
		logs.Error(err)
		logs.Error("Load conf(%s) failed.", confPath)
		return
	}
}

func getDB() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Conf.Mysql.User,
			env.Conf.Mysql.Password,
			env.Conf.Mysql.Host,
			env.Conf.Mysql.DBName,
		),
	)
	if err != nil {
		panic(err)
	}
	return db
}

func InitTableManager() {
	handlers.GlobalTableManager = managers.NewTableManager(getDB())
}

func main() {
	r := gin.Default()
	r.POST("/solved_problem", handlers.InsertSolvedProblems)
	r.GET("/solved_problem", handlers.GetSolvedProblems)
	if err := r.Run(fmt.Sprintf(":%d", env.Conf.Http.Port)); err != nil {
		logs.Error("Running go http server failed. :|")
		return
	}
}
