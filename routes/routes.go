package routes

import (
	"deleteProduct-cart/view"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	commonjwt "github.com/kev1226/auth-common-go/jwt"
)

func SetupRoutes(r *gin.Engine) {
	// Configuración global de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	group := r.Group("/cart")
	group.Use(commonjwt.AuthGuard("user"))
	group.DELETE("/:productId", view.DeleteFromCart)
}
