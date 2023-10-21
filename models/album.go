package models

type Album struct {
	ID     string  `json:"id,omitempty" bson:"id,omitempty"`
	Title  string  `json:"title,omitempty" bson:"title,omitempty"`
	Artist string  `json:"artist,omitempty" bson:"artist,omitempty"`
	Price  float64 `json:"price,omitempty" bson:"price,omitempty"`
}
