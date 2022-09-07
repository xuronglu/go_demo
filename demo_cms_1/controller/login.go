package controller

import (
	"demo_cms_1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func LoginGetController(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPostController(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	status := model.LoginModel(username, password)
	if status == 1 {
		c.JSON(http.StatusOK, gin.H{"status": 400, "msg": "请输入正确的用户名"})
		return
	}
	if status == 2 {
		c.JSON(http.StatusOK, gin.H{"status": 400, "msg": "密码错误!"})
		return
	}
	var store = sessions.NewCookieStore([]byte("mayanan.cn"))
	// 获取一个session对象，session-name是session的名字
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 在session中存储值
	session.Values["admin"] = "adminList"
	// 保存更改  session默认过期时间：一个月
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "登录成功!"})
}

func AdminPutAdminPutAdminLogout(c *gin.Context) {
	story := sessions.NewCookieStore([]byte("mayanan.cn"))
	session, err := story.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	//c.String(200, "session删除成功")
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "退出成功"})
}
