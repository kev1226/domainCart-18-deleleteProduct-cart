package view

import (
	"deleteProduct-cart/presenter"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFromCart(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener el usuario"})
		return
	}
	userID := fmt.Sprintf("%v", userIDInterface)

	productID := c.Param("productId")

	resp, err := presenter.DeleteProduct(
		userID,
		productID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el producto"})
		return
	}
	c.JSON(http.StatusOK, resp)
}
