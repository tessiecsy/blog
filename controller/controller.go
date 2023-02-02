package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "用户不存在")
		fmt.Println("用户不存在")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

//博客

// 博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

// 添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(http.StatusMovedPermanently, "/post_index")
}

// 跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

//跳转到博客详细
func PostDetail(c *gin.Context){
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))  //把markdown转化为html格式

	c.HTML(200, "detail.html", gin.H{
		"Title":p.Title,
		"Content":template.HTML(content),
	})


}