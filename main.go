package main

import (
	"ginTest/controller/controller"
	"ginTest/dao/db"
	"github.com/gin-gonic/gin"
)
func main() {
	dns :="root:admin@tcp(localhost:3306)/nvz?parseTime=true"
	db.Init(dns)

	engine := gin.New()
	//加载html文件
	engine.LoadHTMLGlob("views/*")
	//加载静态资源文件   /leaGin 是访问路径url的前缀  后者是静态资源地址
	// 例如：
	//   /leaGin/images/1.jpg 这个url文件，存储在/static/1.jpg
	engine.Static("/leaGin","/static/*")
	engine.GET("/index",controller.ToIndex)
}
