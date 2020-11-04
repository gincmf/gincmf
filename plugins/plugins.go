/**
** @创建时间: 2020/10/29 4:53 下午
** @作者　　: return
** @描述　　:
 */
package plugins

import (
	"gincmf/plugins/demoPlugin/migrate"
	demoPlugin "gincmf/plugins/demoPlugin/router"
)

func AutoRegister()  {

	// 注册路由
	demoPlugin.ApiListenRouter()

	// 注册数据库迁移
	demoMigrate := migrate.Demo{}
	demoMigrate.AutoMigrate()
}
