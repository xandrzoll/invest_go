package main

import (
	"encoding/csv"
	"fmt"
	moex "invest/src/moex_service"
	pg "invest/src/storage"
	"log"
	"os"
)

func main() {

	investService := moex.NewMoexApiService()
	securities := investService.GetSecuritiesList()
	data, err := investService.FetchCandlesAsync(
		securities,
		"2025-02-10",
		"2025-02-10",
		1,
	)
	if err != nil {
		fmt.Println("Error fetching candles:", err)
		return
	}

	file, err := os.Create("candles_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Ticker", "Open", "Close", "High", "Low", "Volume", "Begin"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	for _, candle := range data {
		row := []string{
			fmt.Sprintf("%s", candle.Security),
			fmt.Sprintf("%f", candle.Open),
			fmt.Sprintf("%f", candle.Close),
			fmt.Sprintf("%f", candle.High),
			fmt.Sprintf("%f", candle.Low),
			//fmt.Sprintf("%f", candle.Value),
			fmt.Sprintf("%f", candle.Volume),
			fmt.Sprintf("%s", candle.Timestamp),
		}
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing CSV row:", err)
			return
		}
	}

	fmt.Println("Data successfully saved to candles_data.csv")

	connStr := "postgres://invest_user:invest_password@localhost:5432/invest"
	pgService, err := pg.NewPostgreService(connStr)
	if err != nil {
		log.Fatalf("Failed to create PostgreService: %v\n", err)
	}
	defer pgService.Close()

	batchSize := 500
	err = pgService.InsertCandles(data, batchSize)
	if err != nil {
		log.Fatalf("Failed to insert candles: %v\n", err)
	}

	log.Println("All candles inserted successfully")
}
