package models

type Country struct {
	ID   int    `json:"id"`
	Name string `db:"country_name" json:"name"`
}
