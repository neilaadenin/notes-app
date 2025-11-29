package routes

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
	"backend/middleware"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)
}

func NotesRoutes(r *gin.Engine) {
	notes := r.Group("/api/notes")
	notes.Use(middleware.AuthMiddleware())
	{
		notes.GET("", controllers.GetNotes)
		notes.POST("", controllers.CreateNote)
	}
}

