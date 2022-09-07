package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", nil)
}
