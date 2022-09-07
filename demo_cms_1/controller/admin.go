package controller

import (
	"demo_cms_1/dao"
	"demo_cms_1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminGetAdmin(c *gin.Context) {
	var adminList []model.DemoAdmin
	start := c.Query("start")
	end := c.Query("end")
	username := c.Query("username")
	if start == "" && end == "" && username == "" {
		//.Limit(2 * id)
		dao.DB.Order("ID desc").Find(&adminList)
		c.HTML(http.StatusOK, "adminList.html", adminList)
		return
	}
	if username == "" {
		dao.DB.Debug().Order("ID desc").Where("created_at BETWEEN ? AND ?", start, end).Find(&adminList)
		c.HTML(http.StatusOK, "adminList.html", adminList)
		return
	}
	if start != "" && end != "" && username != "" {
		dao.DB.Debug().Order("ID desc").Where("created_at BETWEEN ? AND ?", start, end).Where("name", username).Find(&adminList)
		c.HTML(http.StatusOK, "adminList.html", adminList)
		return
	}
	dao.DB.Debug().Order("ID desc").Where("name = ?", username).Find(&adminList)
	c.HTML(http.StatusOK, "adminList.html", adminList)
}

func AdminGetAdminAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "adminAdd.html", nil)
}

func AdminGetAdminEdit(c *gin.Context) {
	id := c.Param("id")
	var adminEdit model.DemoAdmin
	dao.DB.First(&adminEdit, id)
	c.HTML(http.StatusOK, "adminEdit.html", adminEdit)
}

func AdminPostAdmin(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	rePassword := c.Request.FormValue("rePassword")
	phone := c.Request.FormValue("phone")
	email := c.Request.FormValue("email")
	if password != rePassword {
		fmt.Println("新密码跟确认密码不一致")
		return
	}
	model.AdminPostMode(username, password, phone, email)
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "成功"})
}

func AdminDeleteAdmin(c *gin.Context) {
	var delete model.DemoAdmin
	id := c.Param("id")
	dao.DB.Unscoped().Where("ID = ?", id).Delete(&delete)
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "成功"})
}

func AdminPutAdmin(c *gin.Context) {
	username := c.Request.FormValue("name")
	id := c.Request.FormValue("id")
	password := c.Request.FormValue("password")
	rePassword := c.Request.FormValue("rePassword")
	phone := c.Request.FormValue("phone")
	email := c.Request.FormValue("email")
	if password != rePassword {
		fmt.Println("新密码跟确认密码不一致")
		return
	}
	model.AdminPutMode(username, password, phone, email, id)
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "成功"})
}

func AdminPutAdminPutAdminToDown(c *gin.Context) {
	id := c.Request.FormValue("id")
	status := c.Request.FormValue("status")
	orBool, err := strconv.ParseBool(status)
	if err != nil {
		panic(err)
	}
	dao.DB.Debug().Model(&model.DemoAdmin{}).Where("ID = ?", id).Update("Status", orBool)
	c.JSON(http.StatusOK, gin.H{"status": 200, "msg": "成功"})
}
