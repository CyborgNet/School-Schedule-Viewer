package controllers

import (
	"../config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	password := c.PostForm("password")
	if password == config.GetSturentPassword() {
		session := sessions.Default(c)
		session.Set("PASSWORD", password)
		session.Save()
		c.Redirect(302, "/")
	}
}
