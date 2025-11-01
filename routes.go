package main

import (
	"github.com/chrollo-lucifer-12/betteruptime/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	r.POST("/signup", s.addUser)
	r.POST("/login", s.loginUser)
	r.POST("/website", s.addWebsiteHandler)
	r.GET("/status/:websiteId", func(c *gin.Context) {
		//	userId := c.Param("websiteId")

	})
}

func (s *Server) addUser(c *gin.Context) {
	var r AddUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	var existingUser db.User
	if err := s.db.Where("username = ?", r.Username).First(&existingUser).Error; err == nil {
		c.JSON(409, gin.H{"message": "username already taken"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to hash password"})
		return
	}
	newUser := db.User{
		Username: r.Username,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"id": newUser.ID})
}

func (s *Server) loginUser(c *gin.Context) {
	var r AddUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var user db.User
	if err := s.db.Where("username = ?", r.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"message": "invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		c.JSON(401, gin.H{"message": "invalid username or password"})
		return
	}

	newSession := s.createSession(user.ID)
	c.JSON(201, gin.H{"token": newSession.Token})
}

func (s *Server) addWebsiteHandler(c *gin.Context) {
	r := AddWebsiteRequest{}
	if err := c.ShouldBindJSON(&r); err != nil {
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
