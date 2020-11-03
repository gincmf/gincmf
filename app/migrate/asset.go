package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type asset struct {
	Migrate
}

func (_ *asset) AutoMigrate() {
	_ = cmf.NewDb().AutoMigrate(&model.Asset{})
}

