package model

import "time"

type RefreshTokenId uint

type RefreshToken struct {
	ModelAt

	ID        RefreshTokenId `json:"ID" binding:"-"`
	UserID    uint           `json:"userID"`
	Token     string         `json:"token"`
	ExpiresAt time.Time      `json:"expiresAt"`
}
