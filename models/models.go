package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
}

type UserType struct {
	ID          uuid.UUID
	Username    string
	Password    string
	CardNumbers *[]uint
	TotalDebts  uint
}

type CreateUserReq struct {
	Username    string
	Password    string
	CardNumbers *[]uint
}

type CreateUserRes struct {
	Token     string
	ExpiresAt string
}

type LoginUserReq struct {
	Username string
	Password string
}

type LoginUserRes struct {
	Token     string
	ExpiresAt string
}
