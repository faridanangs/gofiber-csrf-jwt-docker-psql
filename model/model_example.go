package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}
type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type UserResponse struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Post struct {
	ID       string `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
}
