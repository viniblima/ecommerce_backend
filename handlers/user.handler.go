package handlers

import (
	"errors"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string) bool
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHash(hashedPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

type User models.User

func (u *User) FindUserByEmail(email string) (*User, error) {
	var err error
	err = database.DB.Db.Debug().Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func GetByEmail(email string, admin bool) (models.User, error) {
	item := models.User{}
	var err error
	dbResult := database.DB.Db.Where("email = ?", email).First(&item)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("user not found")
	}

	return item, err
}

func GetUserByID(id string) (models.User, error) {
	item := models.User{}
	var err error
	dbResult := database.DB.Db.Where("id = ?", id).Take(&item)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("user not found")
	}
	return item, err
}
