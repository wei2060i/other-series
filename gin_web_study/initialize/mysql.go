package initialize

import (
	"context"
	"fmt"
	"gin_web_study/config/global"
	"gin_web_study/model/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// 初始化mysql数据库
func Mysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1",
		"3306",
		"chat",
	)
	init := false
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("初始化mysql异常:连接超时(%ds)", 10))
				}
				// 此处需return避免协程空跑
				return
			}
		}
	}()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		// 指定表前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,//取消复数形式
		},
	})

	if err != nil {
		panic(fmt.Sprintf("初始化mysql异常: %v", err))
	}

	init = true
	// 开启mysql日志
	if true {
		db = db.Debug()
	}

	global.Mysql = db
	// 表结构
	autoMigrate()
	global.Log.Info("初始化mysql完成")
	//初始化数据库日志监听器
	//binlog()
}

// 自动迁移表结构
func autoMigrate() {
	global.Mysql.AutoMigrate(
		new(po.User),

	)
}
