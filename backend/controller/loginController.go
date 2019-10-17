package controller

import (
	"github.com/gin-gonic/gin"
	. "blog/service"
)

func Login(c *gin.Context)  {
	cs := CourseService{}
	cc := cs.GetCourse()	
	c.JSON(200, cc)
}

func Register(c *gin.Context)  {
	c.String(200, "this is register page")
}
