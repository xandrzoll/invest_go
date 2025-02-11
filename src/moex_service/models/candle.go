package models

type Candle struct {
	Security string
	Open     float64
	Close    float64
	High     float64
	Low      float64
	//Value     float64
	Volume    float64
	Timestamp string
}

type CandleResponse struct {
	Candles struct {
		Data [][]interface{} `json:"data"`
	} `json:"candles"`
}
