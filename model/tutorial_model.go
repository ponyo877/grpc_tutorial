package model

import "time"

type Tutorial struct {
	PlaceID   string    `json:"place_id" gorm:"primary_key"`
	LogoUrlID string    `json:"logo_url_id"`
	ColorCode string    `json:"color_code"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func (Tutorial) TableName() string {
	return "tutorial_tbl"
}

type GCSCredential struct {
	PrivateKeyStr string `json:"private_key"`
	ClientEmail   string `json:"client_email"`
}
