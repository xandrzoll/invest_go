package moex_service

import (
	"encoding/json"
	"fmt"
	"invest/src/moex_service/models"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL    = "https://iss.moex.com/iss/"
	candlesURL = baseURL +
		"engines/%s/markets/%s/boards/%s/securities/%s/candles.json?from=%s&till=%s&interval=%d&start=%d"
)

type MoexApiService struct {
	client      *http.Client
	rateLimiter *time.Ticker
}

func NewMoexApiService() *MoexApiService {
	return &MoexApiService{
		client:      &http.Client{},
		rateLimiter: time.NewTicker(time.Second / 10)}
}

func (s *MoexApiService) GetSecuritiesList() []models.Security {
	var securities []models.Security

	securities = append(securities,
		models.Security{Ticker: "SBER", Name: "Сбербанк", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "ABIO", Name: "Артген", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "AFLT", Name: "Аэрофлот", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "ALRS", Name: "Алроса", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "DATA", Name: "Аренадата", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "DIAS", Name: "Диасофт", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "ETLN", Name: "Эталон", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "FESH", Name: "ДВМП", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "GAZP", Name: "Газпром", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "GMKN", Name: "НорНикель", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "HNFG", Name: "Хэндерсон", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "HYDR", Name: "РусГидро", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "IRAO", Name: "ИнтерРао", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "LKOH", Name: "Лукойл", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "MBNK", Name: "МТС Банк", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "MTSS", Name: "МТС", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "NLMK", Name: "НЛМК", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "NVTK", Name: "Новатэк", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "PHOR", Name: "Фосагро", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "POSI", Name: "Позитив", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "ROSN", Name: "Роснефть", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "RTKM", Name: "Ростелеком", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "RUAL", Name: "Русал", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "SFIN", Name: "ЭсЭфАй", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "SNGS", Name: "Сургут", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	securities = append(securities,
		models.Security{Ticker: "SNGSP", Name: "Сургут Преф", Engine: "stock", Market: "shares", Board: "TQBR"},
	)
	return securities
}

func (s *MoexApiService) fetchCandles(engine string, market string, board string, security string, dttmFrom string,
	dttmTill string, interval int, start int) (models.CandleResponse, error) {

	candlesURL := fmt.Sprintf(candlesURL, engine, market, board, security, dttmFrom, dttmTill, interval, start)

	resp, err := s.client.Get(candlesURL)
	if err != nil {
		return models.CandleResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.CandleResponse{}, err
	}

	var candleResponse models.CandleResponse
	if err := json.Unmarshal(body, &candleResponse); err != nil {
		return models.CandleResponse{}, err
	}

	return candleResponse, nil
}

func (s *MoexApiService) parseCandlesResponse(security string, candlesResp models.CandleResponse) ([]models.Candle, error) {
	var candles []models.Candle

	for _, data := range candlesResp.Candles.Data {
		candle := models.Candle{
			Security:  security,
			Open:      data[0].(float64),
			Close:     data[1].(float64),
			High:      data[2].(float64),
			Low:       data[3].(float64),
			Value:     data[4].(float64),
			Volume:    data[5].(float64),
			Timestamp: data[6].(string),
		}
		candles = append(candles, candle)
	}
	return candles, nil
}

func (s *MoexApiService) FetchCandles(
	securities []models.Security, dttmFrom string, dttmTill string, interval int) ([]models.Candle, error) {

	var candleResp []models.Candle

	for _, security := range securities {
		start := 0
		for {
			resp, err := s.fetchCandles(
				security.Engine,
				security.Market,
				security.Board,
				security.Ticker,
				dttmFrom,
				dttmTill,
				interval,
				start)

			if err != nil {
				fmt.Println("Error fetching candles:", err)
				return []models.Candle{}, err
			}
			len_data := len(resp.Candles.Data)
			if len_data > 0 {
				start += len_data
				candles, err := s.parseCandlesResponse(security.Ticker, resp)
				if err != nil {
					fmt.Println("Error parsing candles:", err)
					break
				}
				candleResp = append(candleResp, candles...)
			} else {
				break
			}
		}
	}
	return candleResp, nil
}

func (s *MoexApiService) FetchCandlesAsync(
	securities []models.Security, dttmFrom string, dttmTill string, interval int) ([]models.Candle, error) {

	var candleResp []models.Candle

	var wg sync.WaitGroup

	for _, security := range securities {
		wg.Add(1)
		go func(security models.Security) {
			defer wg.Done()
			start := 0
			fmt.Println("Start requests for:", security.Ticker)
			for {
				<-s.rateLimiter.C
				resp, err := s.fetchCandles(
					security.Engine,
					security.Market,
					security.Board,
					security.Ticker,
					dttmFrom,
					dttmTill,
					interval,
					start)

				if err != nil {
					fmt.Printf("Error fetching candles for security %s: %v\n", security.Ticker, err)
					return
				}
				lenData := len(resp.Candles.Data)
				if lenData > 0 {
					start += lenData
					candles, err := s.parseCandlesResponse(security.Ticker, resp)
					if err != nil {
						fmt.Println("Error parsing candles:", err)
						break
					}
					candleResp = append(candleResp, candles...)
				} else {
					break
				}
			}
		}(security)
	}
	wg.Wait()
	return candleResp, nil
}
