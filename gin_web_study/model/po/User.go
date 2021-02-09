package po

import (
	"gin_web_study/util/selfdefinedtype/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"comment:'名称'"`
	json json.JSON
}
