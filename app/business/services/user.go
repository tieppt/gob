package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/tieppt/gob/app/business/dto"
	"github.com/tieppt/gob/app/foundation/config"
	"github.com/tieppt/gob/app/foundation/database"
	"github.com/tieppt/gob/app/foundation/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *dto.UserRegisterDto) (*models.User, error) {
	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dbUser := &models.User{
		Username: user.Username,
		Password: string(hashPw),
		Email:    user.Email,
	}

	result := database.DBInst.DB.Create(dbUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return dbUser, nil
}

func Login(user *dto.UserLoginDto) (*dto.LoginResponseDto, error) {
	var dbUser models.User
	result := database.DBInst.DB.Where("username = ?", user.Username).First(&dbUser)
	if result.Error != nil {
		return nil, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return CreateResponseToken(dbUser)
}

func CreateResponseToken(dbUser models.User) (*dto.LoginResponseDto, error) {
	jwtToken, token, err := config.CreateJWTToken(dbUser.ID.String())
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		Token:    token,
		ExpireAt: jwtToken.Claims.(*jwt.StandardClaims).ExpiresAt,
	}, nil
}

func GetUserById(id string) (*models.User, error) {
	var dbUser models.User
	result := database.DBInst.DB.Where("id = ?", id).First(&dbUser)

	if result.Error != nil {
		return nil, result.Error
	}
	return &dbUser, nil
}
