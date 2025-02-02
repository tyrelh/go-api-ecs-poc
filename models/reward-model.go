package models

type Reward struct {
	Brand        *string  `json:"brand,omitempty"`
	Currency     *string  `json:"currency,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
	Id           *int     `json:"id,omitempty"`
}
