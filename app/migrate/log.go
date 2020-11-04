package migrate

import (
	"gincmf/app/model"
	cmf "github.com/gincmf/cmf/bootstrap"
)

type log struct {
	Migrate
}

func (_ *log) AutoMigrate() {
	_ = cmf.NewDb().AutoMigrate(&model.Log{})
}
