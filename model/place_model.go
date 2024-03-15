package model


type PlaceCreate struct {
	ID 				uint	    `json:"id" binding:"required"`
	Name 			string 		`json:"name" binding:"required"`
	CityID 			uint    	`json:"cityid" binding:"required"`
	PlaceDesc 		string 		`json:"place_desc" binding:"required"`
	PlaceAddress	string 		`json:"place_address" binding:"required"`
	PlaceHistory 	string 		`json:"place_history" binding:"required"`
	PlaceFasil 		string 		`json:"place_fasil" binding:"required"`
	PlaceActivity 	string 		`json:"place_activity" binding:"required"`
	PlaceBestTime 	string 		`json:"place_besttime" binding:"required"`
	PlaceOpen 		string 		`json:"place_open" binding:"required"`
	PlacePrice 		string    	`json:"place_price" binding:"required"`
	PlaceRules 		string 		`json:"place_rules" binding:"required"`
	PlaceEvent 		string 		`json:"place_event" binding:"required"`
	PlaceAward 		string 		`json:"place_award" binding:"required"`
	PlaceImage 		string 		`json:"place_image" binding:"required"`
}

type PlaceParam struct{
	ID 		uint 		`json:"-"`
	Name 	string  	`json:"-"`
	CityID 	uint  		`json:"-"`
}