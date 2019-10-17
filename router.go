package main

import (
	"github.com/gin-gonic/gin"
	. "blog/frontend/controller"
)

func InitRouter() *gin.Engine {
	
	router := gin.Default()

	//auth group
	v := router.Group("/")
	{
		//多级分组
		v1 := v.Group("/v1")
		v1.GET("/login", Login)
		
		v2 := v.Group("/v2")
		v2.GET("/register", Register)
	}
	
	router.GET("/login", Login)
	
	return router
}