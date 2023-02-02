package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/login", controller.GoLogin)
	r.POST("/login", controller.Login)

	r.GET("/", controller.Index) //首页 
	r.GET("/register", controller.GoRegister)
	r.POST("/register", controller.Register)
	
	//博客
	r.GET("/post_index", controller.GetPostIndex)  //跳到博客首页
	r.GET("/post_add", controller.GoAddPost)  //跳到添加博客页面
	r.POST("/post", controller.AddPost)  //添加博客
	r.GET("/post_detail", controller.PostDetail) //跳到博客详细

	

	r.Run()

}
