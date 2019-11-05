package routes

import (
	"bubble/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Index)
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodoList)
		v1.POST("todo", controllers.CreateTodo)
		v1.GET("todo/:id", controllers.GetTodo)
		v1.PUT("todo/:id", controllers.UpdateTodo)
		v1.DELETE("todo/:id", controllers.DeleteTodo)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})
	return r
}
