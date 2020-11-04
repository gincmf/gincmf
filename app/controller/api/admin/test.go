/**
** @创建时间: 2020/11/4 7:56 下午
** @作者　　: return
** @描述　　:
 */
package admin

import (
	"gincmf/app/util"
	"github.com/gin-gonic/gin"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/controller"
)

type TestController struct {
	rc controller.RestController
}

func (rest *TestController) Get(c *gin.Context) {
	qn := cmf.QiNiu{Bucket: "vlog"}
	key,_ := qn.UploadFile("test/1111.md",util.CurrentPath() + "/todo.md")
	rest.rc.Success(c, "操作成功Get："+key, nil)

}
