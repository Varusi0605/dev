package validation

// import (
// 	"os"
// 	"shopping-site/api/models"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func GenerateToken(user *models.Users) (string, error) {
// 	claims := &helpers.JWTClaims{
// 		Userid:   user.UserID,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Role:     user.Role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenstr, err := token.SignedString([]byte(os.Getenv("secret_key")))
// 	if err != nil {
// 		loggers.WarningLog.Println(err)
// 		return "", err
// 	}

// 	return tokenstr, nil
// }
