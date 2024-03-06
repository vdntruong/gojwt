package tjwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Generator struct {
	Issuer        string
	SecretKey     string
	SigningMethod jwt.SigningMethod
}

func NewGenerator(issuer, secret string, signingMethod jwt.SigningMethod) *Generator {
	return &Generator{
		Issuer:        issuer,
		SecretKey:     secret,
		SigningMethod: signingMethod,
	}
}

func (g *Generator) Sign(claims TClaims) (string, error) {
	token := jwt.NewWithClaims(g.SigningMethod, claims)
	tokenString, err := token.SignedString([]byte(g.SecretKey))
	return tokenString, err
}

func (g *Generator) Verify(tokenString string) (*TClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(g.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
