package routers

import (
	"bubble/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.Index)
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controller.GetTodoList)
		v1.POST("todo", controller.CreateTodo)
		v1.GET("todo/:id", controller.GetTodo)
		v1.PUT("todo/:id", controller.UpdateTodo)
		v1.DELETE("todo/:id", controller.DeleteTodo)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})
	return r
}
