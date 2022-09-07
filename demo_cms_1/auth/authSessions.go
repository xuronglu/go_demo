package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func AuthSession() gin.HandlerFunc {
	var store = sessions.NewCookieStore([]byte("mayanan.cn"))
	return func(context *gin.Context) {
		session, err := store.Get(context.Request, "session-name")
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"err": err})
			context.Abort() // 结束后续的请求处理
			return
		}
		fooInfo := session.Values["admin"]
		if fooInfo == "adminList" {
			context.Next()
			return
		}
		context.JSON(http.StatusUnauthorized, gin.H{
			"err": "身份认证失败",
		})
		context.Abort()
		return
	}
}
