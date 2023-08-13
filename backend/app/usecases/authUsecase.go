package usecases

import (
	"todolist/app/models"
	"todolist/app/repositories"
	"todolist/app/utils"
	"golang.org/x/crypto/bcrypt"
	"errors"
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
	dbUser, err := u.AuthRepository.GetUserByEmail(user.Email)
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

func (u *AuthUsecase) Signup(user *models.Auth) (string, error) {
	jwtToken := ""

	if (user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || user.PasswordConfirmation == "") {
		return jwtToken, errors.New("Please fill out all fields!")
	}
	if (!utils.IsValidEmail(user.Email)) {
		return jwtToken, errors.New("Please enter a valid email address!")
	}
	if (len(user.Password) < 8) {
		return jwtToken, errors.New("Password must be at least 8 characters long!")
	}
	if (user.Password != user.PasswordConfirmation) {
		return jwtToken, errors.New("Passwords do not match!")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return jwtToken, err
	}

	newUser := models.Auth{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		PasswordHash: hashedPassword,
	}

	err = u.AuthRepository.CreateUser(&newUser)
	if (err != nil) {
		return jwtToken, err
	}

	jwtToken, err = utils.CreateJWTToken(&newUser)
	return jwtToken, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
