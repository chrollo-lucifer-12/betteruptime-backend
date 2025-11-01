package main

import (
	"github.com/chrollo-lucifer-12/betteruptime/db"
	"github.com/gin-gonic/gin"
)

type AddWebsiteRequest struct {
	Url string `json:"url"`
}

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	api := r.Group("api/v1")
	s.RegisterWebsiteRoutes(api)
}

func (s *Server) RegisterWebsiteRoutes(r *gin.RouterGroup) {
	r.POST("/user", s.addUser)
	r.POST("/website", s.addWebsiteHandler)
	r.GET("/status/:websiteId", func(c *gin.Context) {
		//	userId := c.Param("websiteId")

	})
}

func (s *Server) addUser(c *gin.Context) {
	r := AddUserRequest{}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		c.JSON(411, gin.H{"message": err.Error()})
		return
	}
	newUser := db.User{
		Username: r.Username,
		Password: r.Password,
	}
	if err := s.db.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"message": err.Error})
		return
	}
	c.JSON(201, gin.H{"id": newUser.ID})
}

func (s *Server) addWebsiteHandler(c *gin.Context) {
	r := AddWebsiteRequest{}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		c.JSON(411, gin.H{"message": err.Error()})
		return
	}
	newWebsite := db.Website{
		Url:    r.Url,
		UserID: 1,
	}
	if err := s.db.Create(&newWebsite).Error; err != nil {
		c.JSON(500, gin.H{"message": err.Error})
	}
	c.JSON(201, gin.H{"id": newWebsite.ID})
}
