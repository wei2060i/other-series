package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// zap日志
	Log *zap.SugaredLogger
	// mysql实例
	Mysql *gorm.DB

)

