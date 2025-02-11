package models

type Candle struct {
	Security string  // Тикер акции
	Open     float64 // Цена открытия
	Close    float64 // Цена закрытия
	High     float64 // Максимальная цена
	Low      float64 // Минимальная цена
	Value    float64 //
	Volume   float64 // Объем торгов
	//Timestamp time.Time // Время свечи
	Timestamp string //
}

type CandleResponse struct {
	Candles struct {
		Data [][]interface{} `json:"data"`
	} `json:"candles"`
}
