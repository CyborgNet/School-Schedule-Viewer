package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func ShowMainPage(c *gin.Context) {
	var classSubject [10]string
	for i := 0; i < 10; i++ {
		classSubject[i] = GiveSchedule(1, i + 1)
		if classSubject[i] == "" {
			classSubject[i] = "No data"
		}
	}
	render(c, gin.H{
		"title": "Daeshin Schedule",
		"class1": classSubject[0],
		"class2": classSubject[1],
		"class3": classSubject[2],
		"class4": classSubject[3],
		"class5": classSubject[4],
		"class6": classSubject[5],
		"class7": classSubject[6],
		"class8": classSubject[7],
		"class9": classSubject[8],
		"class10": classSubject[9],
	}, "main.html")
}

func ShowLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Daeshin Schedule",
	}, "login.html")
}
