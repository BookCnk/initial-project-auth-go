package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

type UserToken struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primarykey"`
	UserID    string         `json:"user_id" gorm:"type:varchar(255)"`
	User      User           `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Iat       time.Time      `json:"iat" gorm:"type:timestamp;not null"`
	Exp       time.Time      `json:"exp" gorm:"type:timestamp;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type JwtUserToken struct {
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	Sub string `json:"sub"`
}

func (u UserToken) ToToken(secret []byte) (string, error) {
	claims := JwtUserToken{
		Sub: u.UserID,
		Iat: u.Iat.Unix(),
		Exp: u.Exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": claims.Sub,
		"iat": claims.Iat,
		"exp": claims.Exp,
	})

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
