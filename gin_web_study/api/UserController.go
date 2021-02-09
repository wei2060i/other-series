package api

import (
	"fmt"
	"gin_web_study/config/global"
	"gin_web_study/model/po"
	"github.com/gin-gonic/gin"
)

//后台登录
type AdminLoginVo struct {
	Name string `json:"name" binding:"required,min=1"`
	Pwd  string `json:"pwd" binding:"required,min=10"`
}

func GetLeaves(c *gin.Context) {
	//var adminLogin AdminLoginVo
	//a := AdminLogin()
	//fmt.Println("aaaaaaa",a)
	//_ = c.BindJSON(&adminLogin)
	fmt.Println(55555)

	a, b := AdminLogin()

	fmt.Println(a)
	fmt.Println(b)

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i int     /* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}
	for i = 0; i < 11; i++ {
		fmt.Println(n[i])
	}
}

func AdminLogin() (int, int) {
	panic("阿斯顿撒旦")
	return 1, 3
}

func SaveTest(c *gin.Context) {
	var u = po.User{Name: "小舞"}
	global.Mysql.Create(&u)
	fmt.Println(u)
	fmt.Println(u.ID)
}
