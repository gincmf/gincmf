package main

import (
	"gincmf/app/migrate"
	"gincmf/app/util"
	"gincmf/plugins"
	"gincmf/router"
	cmf "github.com/gincmf/cmf/bootstrap"
)



// test commit
func main() {

	//初始化配置设置
	cmf.Initialize(util.CurrentPath() + "/conf/config.json")

	//初始化路由设置
	router.ApiListenRouter()

	//初始化网页路由设置
	router.WebListenRouter()

	// 数据库迁移
	migrate.AutoMigrate()

	// 注册插件
	plugins.AutoRegister()

	//启动服务
	cmf.Start()
	//执行数据库迁移
}
