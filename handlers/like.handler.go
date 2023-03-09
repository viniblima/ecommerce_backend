package handlers

import (
	"errors"
	"fmt"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func IsProductLiked(userID string, productID string) (models.Like, error) {
	var like models.Like
	var err error
	dbResult := database.DB.Db.Where("user_id = ? AND product_id = ?", userID, productID).Take(&like)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		fmt.Println("caiu aqui")

		err = errors.New("user not found")

	}
	return like, err
}

func DeleteLike(id string) (models.Like, error) {
	var like models.Like
	var err error
	dbResult := database.DB.Db.Where("id = ?", id).Delete(&like)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		fmt.Println("caiu aqui")

		err = errors.New("Error on delete like")

	}
	return like, err
}

func CreateLike(userID string, productID string) (models.Like, error) {
	var like models.Like
	var err error
	like.UserID = userID
	like.ProductID = productID
	dbResult := database.DB.Db.Create(&like)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		fmt.Println("caiu aqui")

		err = errors.New("Error on create like")

	}
	return like, err
}
