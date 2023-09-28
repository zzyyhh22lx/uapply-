package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	User uint8 = 1
	Orga uint8 = 2
	Depa uint8 = 3
	Inte uint8 = 4
)

// TokenExpireDuration Expiration time 7 days
const TokenExpireDuration = time.Hour * 24 * 7

// Set the signature key
var mySercet = []byte("bms")

type Claim struct {
	Uo UserClaims
	Oo OrgaClaims
	Do DepaClaims
	Io InteClaims
}

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type OrgaClaims struct {
	OrgaID uint `json:"orga_id"`
	jwt.StandardClaims
}

type DepaClaims struct {
	OrgaID uint `json:"orga_id"`
	DepaID uint `json:"depa_id"`
	jwt.StandardClaims
}

type InteClaims struct {
	InteID uint `json:"inte_id"`
	OrgaID uint `json:"orga_id"`
	DepaID uint `json:"depa_id"`
	jwt.StandardClaims
}

func GenTokenDepa(InteID uint, DepaID uint, OrgaID uint) (string, error) {
	var c Claim
	c.Io = InteClaims{
		DepaID: DepaID,
		OrgaID: OrgaID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "uapply",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.Io)
	return token.SignedString(mySercet)
}

func GenTokneDepa(DepaID uint, OrgaID uint) (string, error) {
	var c Claim
	c.Do = DepaClaims{
		DepaID: DepaID,
		OrgaID: OrgaID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "uapply",                                   // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.Do)
	return token.SignedString(mySercet)
}

// GenToken Build Web JWT
// typ 看上面的 vars
func GenToken(ID uint, typ uint8) (string, error) {
	// Create our own statement
	var c Claim
	switch typ {
	case User:
		{
			c.Uo = UserClaims{
				UserID: ID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
					Issuer:    "uapply",                                   // 签发人
				},
			}
		}
	case Orga:
		{
			c.Oo = OrgaClaims{
				OrgaID: ID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
					Issuer:    "uapply",                                   // 签发人
				},
			}
		}
	case Inte:
		{
			c.Io = InteClaims{
				InteID: ID,
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
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, c.Uo)
	case Orga:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, c.Oo)
	case Inte:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, c.Io)
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

func ParseTokenOrga(tokenString string) (*OrgaClaims, error) {
	// parse token
	var mc = new(OrgaClaims)
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

func ParseTokenDepa(tokenString string) (*DepaClaims, error) {
	// parse token
	var mc = new(DepaClaims)
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

func ParseTokenInte(tokenString string) (*InteClaims, error) {
	// parse token
	var mc = new(InteClaims)
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
