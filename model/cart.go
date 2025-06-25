package model

import (
	"deleteProduct-cart/config"
	"fmt"
)

func DeleteItem(userID, productID string) error {
	key := fmt.Sprintf("cart:%s", userID)
	return config.RedisClient.HDel(config.Ctx, key, productID).Err()
}
