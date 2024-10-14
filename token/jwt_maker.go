package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	minSecrectKeySize = 32
)

type JWTMaker struct {
	secrectKey string
}

type MyCustomClaims struct {
	Payload *Payload
	jwt.RegisteredClaims
}

func NewJWTMaker(secrect string) (Maker, error) {
	if len(secrect) < minSecrectKeySize {
		return nil, fmt.Errorf("Invalid key size: must be at leaset %d", minSecrectKeySize)
	}
	return JWTMaker{
		secrect,
	}, nil
}

func (maker JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayLoad(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssuedAt),
			NotBefore: jwt.NewNumericDate(payload.IssuedAt),
		},
	})

	token, err := jwtToken.SignedString([]byte(maker.secrectKey))

	return token, payload, nil
}

func (maker JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secrectKey), nil
	}
	parsed, err := jwt.ParseWithClaims(token, MyCustomClaims{}, keyFunc)
	if err != nil {

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	myCustomClaims, ok := parsed.Claims.(*MyCustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	payload := myCustomClaims.Payload

	return payload, nil
}
