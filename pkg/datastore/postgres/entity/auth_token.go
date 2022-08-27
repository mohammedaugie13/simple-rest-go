package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthToken struct {
	UserId          string         `gorm:"column:userId;not null;foreignKey:userId;references:userId;primaryKey"`
	RefreshToken    string         `gorm:"column:refreshToken;not null"`
	RefreshTokenExp time.Time      `gorm:"column:refreshTokenExp;not null"`
	CreatedAt       time.Time      `gorm:"column:createdAt;autoCreateTime:nano"`
	UpdatedAt       time.Time      `gorm:"column:updatedAt;autoUpdateTime:nano"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deletedAt"`
}

func (AuthToken) TableName() string {
	return "authToken"
}
