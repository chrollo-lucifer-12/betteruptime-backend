package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("api/v1")
	RegisterWebsiteRoutes(api)
}

func RegisterWebsiteRoutes(r *gin.RouterGroup) {
	r.POST("/website", func(c *gin.Context) {

	})
	r.GET("/status/:websiteId", func(c *gin.Context) {
		//	userId := c.Param("websiteId")

	})
}
