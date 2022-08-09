package service

import (
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"golang.org/x/net/context"
)

// Сервис получения биржевой информации
type MarketDataService interface {
	// Метод запроса последних цен по инструментам
	LastPrices(ctx context.Context, figi []string) ([]*domain.LastPrice, error)
	// Метод запроса исторических свечей по инструменту
	Candles(ctx context.Context, figi string, from, to time.Time, interval domain.CandleInterval) ([]*domain.Candle, error)
	// Метод получения стакана по инструменту
	OrderBook(ctx context.Context, figi string, depth int32) (*domain.OrderBook, error)
	// Метод запроса статуса торгов по инструментам
	TradingStatus(ctx context.Context, figi string) (*domain.InstrumentTradingStatus, error)
	// Метод запроса обезличенных сделок за последний час.
	LastTrades(ctx context.Context, figi string, from, to time.Time) ([]*domain.Trade, error)
}
