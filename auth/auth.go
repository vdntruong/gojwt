package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"gojwt/model"
	"gojwt/tjwt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthSvc struct {
	jwt *tjwt.Generator
}

func NewAuthSvc(jwt *tjwt.Generator) *AuthSvc {
	return &AuthSvc{
		jwt: jwt,
	}
}

func (a *AuthSvc) GetToken(ctx context.Context, username string, password string) (string, *model.User, error) {
	// call to services to validate user by username and password
	var u = model.User{
		EmployeeCode: fmt.Sprintf("%s-%d", username, rand.Uint32()),
		UserName:     username,
		UnitName:     "PhongNhanSu",
		Title:        "GiamDoc",
		PhoneNumber:  "123-456-789",
		Email:        fmt.Sprintf("%s@gmail.com", username),
	}

	var now = time.Now()
	var claims = tjwt.TClaims{
		EmployeeCode: u.EmployeeCode,
		UserName:     u.UserName,
		UnitName:     u.UnitName,
		Title:        u.Title,
		PhoneNumber:  u.PhoneNumber,
		Email:        u.Email,

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:  a.jwt.Issuer,
			Subject: u.EmployeeCode,

			ExpiresAt: jwt.NewNumericDate(now.Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        u.EmployeeCode,
		},
	}

	t, err := a.jwt.Sign(claims)
	if err != nil {
		return "", nil, err
	}

	return t, &u, nil
}

func (a *AuthSvc) VerifyToken(ctx context.Context, token string) (*model.User, error) {
	claims, err := a.jwt.Verify(token)
	if err != nil {
		return nil, err
	}

	return &model.User{
		EmployeeCode: claims.EmployeeCode,
		UserName:     claims.UserName,
		UnitName:     claims.UnitName,
		Title:        claims.Title,
		PhoneNumber:  claims.PhoneNumber,
		Email:        claims.Email,
	}, nil
}
