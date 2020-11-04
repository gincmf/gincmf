/**
** @创建时间: 2020/10/29 4:29 下午
** @作者　　: return
** @描述　　: 插件名采用大驼峰命名法，都带 Plugin类名后缀，如 DemoPlugin,CustomAdminLoginPlugin
 */
package controller

import (
"github.com/gin-gonic/gin"
"github.com/gincmf/cmf/controller"
)

type DemoController struct {
	rc controller.RestController
}

func (rest *DemoController) Get(c *gin.Context) {
	rest.rc.Success(c, "操作成功Get", nil)
}

func (rest *DemoController) Show(c *gin.Context) {
	var rewrite struct {
		Id int `uri:"id"`
	}
	if err := c.ShouldBindUri(&rewrite); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	rest.rc.Success(c, "操作成功show", nil)
}

func (rest *DemoController) Edit(c *gin.Context) {
	rest.rc.Success(c, "操作成功Edit", nil)
}

func (rest *DemoController) Store(c *gin.Context) {
	rest.rc.Success(c, "操作成功Store", nil)
}

func (rest *DemoController) Delete(c *gin.Context) {
	rest.rc.Success(c, "操作成功Delete", nil)
}