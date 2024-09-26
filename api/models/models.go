package models

type City struct {
	ID        int    `json:"id"`
	CityName  string `json:"city_name"`
	CountryID int    `json:"country_id"`
}

var Cities = map[int]City{
	1: {ID: 1, CityName: "Sample City", CountryID: 1},
}
var LastCityID int = 0
