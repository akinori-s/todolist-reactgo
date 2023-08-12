package usecases

import (
	"todolist/app/models"
	"todolist/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepository repositories.AuthRepository
}

func NewAuthUsecase(authRepository repositories.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: authRepository,
	}
}

func (u *AuthUsecase) Login(user *models.Auth) (string, error) {
	dbUser, err := u.AuthRepository.GetUserByUsername(user.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return "passwords dont match...", err
	}

	// wip: need to fix return vaue and manage JWT token.

	/*
	1. get user password with user.Username
	2. compare user.Password with user.Password
	3. if not match, return error
	4. if match, create token and return token
	*/

	return "SUCCESS", nil
}
