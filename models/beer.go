package models

type Beer struct {
	ID             int     `json:"id"`
	Beername       string  `db:"beer_name" json:"beer_name"`
	Creator        string  `db:"creator" json:"creator"`
	OriginCountry  string  `db:"origin_country" json:"origin_country"`
	CurrentCountry string  `db:"current_country" json:"current_country"`
	Alcohol        float32 `db:"alcohol" json:"alcohol"`
}
