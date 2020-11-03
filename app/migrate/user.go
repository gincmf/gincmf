package migrate

import (
	"fmt"
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
	"github.com/gincmf/cmf/util"
	"time"
)

type user struct {
	Migrate
}

func (u *user) AutoMigrate() {
	_ = cmf.NewDb().AutoMigrate(&model.User{})

	userResult := cmf.NewDb().First(&model.User{}, "user_login = ?", "admin") // 查询

	if userResult.RowsAffected == 0 {
		fmt.Println("user", userResult)
		//新增一条管理员数据
		cmf.NewDb().Create(&model.User{UserType: 1, UserLogin: "admin", UserPass: util.GetMd5("123456"),LastLoginAt: time.Now().Unix(), CreateAt: time.Now().Unix()})
	}
}
