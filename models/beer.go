package models

type BeerDTO struct {
	ID               int     `json:"id"`
	Beername         string  `db:"beer_name" json:"beer_name"`
	Creator          string  `db:"creator" json:"creator"`
	OriginCountryId  int     `db:"origin_country_id" json:"origin_country_id"`
	CurrentCountryId int     `db:"current_country_id" json:"current_country_id"`
	Alcohol          float32 `db:"alcohol" json:"alcohol"`
}

type Beer struct {
	ID             int     `json:"id"`
	Beername       string  `db:"beer_name" json:"beer_name"`
	Creator        string  `db:"creator" json:"creator"`
	OriginCountry  string  `json:"origin_country"`
	CurrentCountry string  `json:"current_country"`
	Alcohol        float32 `db:"alcohol" json:"alcohol"`
}
