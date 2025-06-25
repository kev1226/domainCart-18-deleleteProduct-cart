package presenter

import (
	"deleteProduct-cart/dto"
	"deleteProduct-cart/model"
)

func DeleteProduct(userID, productID string) (*dto.DeleteItemResponse, error) {
	err := model.DeleteItem(userID, productID)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteItemResponse{
		Message: "Producto eliminado correctamente",
	}, nil
}
