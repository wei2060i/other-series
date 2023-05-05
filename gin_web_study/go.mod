module gin_web_study

go 1.15

//go get -u go.uber.org/zap
//go get -u gopkg.in/natefinch/lumberjack.v2
//go get -u github.com/gin-gonic/gin
//go get -u github.com/dgrijalva/jwt-go
//jorm 不要 go get 下载

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.11
)
