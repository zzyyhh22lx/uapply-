package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	User uint8 = 1
)

// TokenExpireDuration Expiration time 7 days
const TokenExpireDuration = time.Hour * 24 * 7

// Set the signature key
var mySercet = []byte("bms")

type Claim struct {
	Bo UserClaims
}

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenToken Build Web JWT
// typ 看上面的 vars
func GenToken(ID uint, typ uint8) (string, error) {
	// Create our own statement
	var c Claim
	switch typ {
	case User:
		{
			c.Bo = UserClaims{
				UserID: ID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
					Issuer:    "uapply",                                   // 签发人
				},
			}
		}
	}
	// Create a signature object using the SigningMethodHS256 signature method
	var token *jwt.Token
	switch typ {
	case User:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, c.Bo)
	}
	// Use the specified seek signature and get the fully encoded token
	return token.SignedString(mySercet)
}

func ParseTokenUser(tokenString string) (*UserClaims, error) {
	// parse token
	var mc = new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
