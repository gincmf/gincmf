/**
** @创建时间: 2020/10/29 4:51 下午
** @作者　　: return
** @描述　　:
 */
package migrate
import "gincmf/plugins/demoPlugin/model"

type Demo struct {}

func (migrate *Demo) AutoMigrate () {
	new(model.Demo).AutoMigrate()
}
