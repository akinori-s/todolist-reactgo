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

func (u *AuthUsecase) Login(user *models.Auth) (string, *models.Auth, error) {
	dbUser, err := u.AuthRepository.GetUserByEmail(user.Email)
	if err != nil {
		return "", nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
	if err != nil {
		return "", nil, errors.New("Invalid email or password.")
	}

	jwtToken, err := utils.CreateJWTToken(dbUser)
	dbUser.ID = 0
	dbUser.PasswordHash = ""
	return jwtToken, dbUser, err
}

func (u *AuthUsecase) Signup(user *models.Auth) (string, *models.Auth, error) {
	jwtToken := ""

	if (user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" || user.PasswordConfirmation == "") {
		return jwtToken, nil, errors.New("Please fill out all fields!")
	}
	if (!utils.IsValidEmail(user.Email)) {
		return jwtToken, nil, errors.New("Please enter a valid email address!")
	}
	if (len(user.Password) < 8) {
		return jwtToken, nil, errors.New("Password must be at least 8 characters long!")
	}
	if (user.Password != user.PasswordConfirmation) {
		return jwtToken, nil, errors.New("Passwords do not match!")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return jwtToken, nil, err
	}

	newUser := &models.Auth{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		PasswordHash: hashedPassword,
	}

	err = u.AuthRepository.CreateUser(newUser)
	if (err != nil) {
		return jwtToken, nil, err
	}

	jwtToken, err = utils.CreateJWTToken(newUser)
	newUser.ID = 0
	newUser.PasswordHash = ""
	return jwtToken, newUser, err
}

func (u *AuthUsecase) CheckLogin(tokenString string) (*models.Auth, error) {
	user, err := utils.ParseJWTTokenToUser(tokenString)
	user.ID = 0
	return user, err
}
