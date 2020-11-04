/**
** @创建时间: 2020/10/29 4:33 下午
** @作者　　: return
** @描述　　:
 */
package router

import (
	"gincmf/app/middleware"
	"gincmf/plugins/demoPlugin/controller"
	cmf "github.com/gincmf/cmf/bootstrap"

)

func ApiListenRouter() {
	adminGroup := cmf.Group("api/admin", middleware.ApiBaseController)
	{
		adminGroup.Get("/demo",new(controller.DemoController).Get)
	}
}
