package model

type CityCreate struct {
	ID        uint   `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	CityImage string `json:"cityimage" binding:"required"`
}

type CityParam struct {
	ID   uint   `json:"-"`
	Name string `json:"name"`
}

type SearchCity struct {
	Name string `json:"name"`
}
