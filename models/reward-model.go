package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	Brand        *string  `json:"brand,omitempty"`
	Currency     *string  `json:"currency,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
}
