package model

import (
	"time"
)

type Car struct {
	Id        uint      `json:"id"`
	Brand     string    `json:"brand"`
	Series    string    `json:"series"`
	Year      time.Time `json:"year"`
	Fuel      string    `json:"fuel"`
	Gear      string    `json:"gear"`
	Situation string    `json:"situation"`
	Km        int       `json:"km"`
	Color     string    `json:"color"`
	Price     int       `json:"price"`
}
