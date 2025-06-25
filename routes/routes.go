package routes

import (
	"deleteProduct-cart/view"

	"github.com/gin-gonic/gin"
	commonjwt "github.com/kev1226/auth-common-go/jwt"
)

func SetupRoutes(r *gin.Engine) {
	group := r.Group("/cart")
	group.Use(commonjwt.AuthGuard("user")) // Admin también puede eliminar
	group.DELETE("/:productId", view.DeleteFromCart)
}
