package models

type Security struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
	Engine string `json:"engine"`
	Market string `json:"market"`
	Board  string `json:"board"`
}
