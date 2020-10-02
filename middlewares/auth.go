package middlewares

import (
	"../config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		passwordReceived := session.Get("PASSWORD")
		if c.FullPath() != "/login" {
			if passwordReceived == nil {
				c.Redirect(302, "login")
				return
			} else if passwordReceived.(string) != config.GetSturentPassword() {
				c.Redirect(302, "login")
				return
			}
		} else {
			if passwordReceived == nil {
				c.Next()
			} else if passwordReceived.(string) == config.GetSturentPassword() {
				c.Redirect(302, "/")
				return
			}
		}

		c.Next()
	}
}
