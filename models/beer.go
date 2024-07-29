package models

import "time"

type Beer struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Beername  string     `json:"beer_name"`
	Company   string     `json:"creator"`
	CountryID uint       `json:"country_id"`
	Country   Country    `json:"country" gorm:"references:ID"`
	Alcohol   float32    `json:"alcohol"`
}
