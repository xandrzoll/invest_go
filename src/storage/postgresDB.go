package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"invest/src/moex_service/models"
	"log"
	"time"
)

type PostgreService struct {
	pool *pgxpool.Pool
}

func NewPostgreService(connStr string) (*PostgreService, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &PostgreService{pool: pool}, nil
}

func (s *PostgreService) Close() {
	s.pool.Close()
}

func (s *PostgreService) InsertCandlesBatch(candles []models.Candle) error {
	query := `
		INSERT INTO data.candles (ticker, open, close, high, low, volume, dttm)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	batch := &pgx.Batch{}
	for _, candle := range candles {
		batch.Queue(query, candle.Security, candle.Open, candle.Close, candle.High, candle.Low, candle.Volume, candle.Timestamp)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	br := s.pool.SendBatch(ctx, batch)
	defer br.Close()

	_, err := br.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute batch: %w", err)
	}

	return nil
}

func (s *PostgreService) InsertCandles(candles []models.Candle, batchSize int) error {
	for i := 0; i < len(candles); i += batchSize {
		end := i + batchSize
		if end > len(candles) {
			end = len(candles)
		}

		batch := candles[i:end]
		err := s.InsertCandlesBatch(batch)
		if err != nil {
			return fmt.Errorf("failed to insert batch: %w", err)
		}

		log.Printf("Inserted batch from %d to %d\n", i, end)
	}

	return nil
}
