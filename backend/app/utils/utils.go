package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"todolist/app/models"
	"time"
	"regexp"
	"errors"
)

var mySigningKey = []byte("very secret")

type AuthClaim struct {
	ID				int
	FirstName		string
	LastName		string
	Email			string
	jwt.RegisteredClaims
}

func IsValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CreateJWTToken(user *models.Auth) (string, error) {
	claims := AuthClaim{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt:	jwt.NewNumericDate(time.Now().Add((24 * 7) * time.Hour)),
			IssuedAt:	jwt.NewNumericDate(time.Now()),
			NotBefore:	jwt.NewNumericDate(time.Now()),
			Issuer:		"todolist",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseJWTTokenToUser(tokenString string) (*models.Auth, error) {
	authClaim := &AuthClaim{}
	token, err := jwt.ParseWithClaims(tokenString, authClaim, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		user := new(models.Auth)
		user.ID = claims.ID
		user.FirstName = claims.FirstName
		user.LastName = claims.LastName	
		user.Email = claims.Email
		return user, nil
	}
	return nil, errors.New("Invalid token")
	
}
