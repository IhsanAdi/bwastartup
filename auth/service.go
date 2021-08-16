package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

// should be placed in .env file
var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// claims is same as payload
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
