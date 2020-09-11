package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(server *gin.Engine) {
	fmt.Println("Setting up routers ...")

	// Profile
	server.GET("/profile", getProfile())

	// Experience
	server.GET("/experiences", getAllExperiences())
	server.POST("/experience/:id", addExperience())
	server.DELETE("experience/:id", deleteExperience())

	// Project
	server.GET("/projects", getAllProjects())
	server.POST("/project/:id", addProject())
	server.DELETE("project/:id", deleteProject())

	// Blogs
	server.GET("/blogs", getAllBlogs())
	server.POST("/blog/:id", addBlog())
	server.DELETE("blog/:id", deleteBlog())
}
