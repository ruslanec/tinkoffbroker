package tinkoffbroker

import (
	"time"

	"golang.org/x/net/context"
)

// Получение биржевой информации
type MarketData interface {
	MarketDataService
	MarketDataStreamService
}

// Сервис получения биржевой информации
type MarketDataService interface {
	// Метод запроса последних цен по инструментам
	LastPrices(ctx context.Context, figi []string) ([]*LastPrice, error)
	// Метод запроса исторических свечей по инструменту
	Candles(ctx context.Context, figi string, from, to time.Time, interval CandleInterval) ([]*Candle, error)
	// Метод получения стакана по инструменту
	OrderBook(ctx context.Context, figi string, depth int32) (*OrderBook, error)
	// Метод запроса статуса торгов по инструментам
	TradingStatus(ctx context.Context, figi string) (*InstrumentTradingStatus, error)
}

// Сервис получения биржевой информации
type MarketDataStreamService interface {
	// Подписка на свечи
	SubscribeCandles(ctx context.Context, candles []*CandleInstrument) error
	// Закрытие подписки на свечи
	UnsubscribeCandles(ctx context.Context, candles []*CandleInstrument) error
	// Подписка на стакан
	SubscribeOrderBook(ctx context.Context, orderbooks []*OrderBookInstrument) error
	// Закрытие подписки на стакан
	UnsubscribeOrderBook(ctx context.Context, orderbooks []*OrderBookInstrument) error
	// Подписка на ленту сделок
	SubscribeTrades(ctx context.Context, trades []*TradeInstrument) error
	// Закрытие подписки на ленту сделок
	UnsubscribeTrades(ctx context.Context, trades []*TradeInstrument) error
	// Подписка на торговые статусы
	SubscribeInfo(ctx context.Context, instruments []*InfoInstrument) error
	// Закрытие подписки на торговые статусы
	UnsubscribeInfo(ctx context.Context, instruments []*InfoInstrument) error
	// Подписка на последнюю цену инструмента
	SubscribeLastPrices(ctx context.Context, lastprices []*LastPriceInstrument) error
	// Закрытие подписки на последнюю цену инструмента
	UnsubscribeLastPrices(ctx context.Context, lastprices []*LastPriceInstrument) error
	// Получение данных по подпискам
	Recv(ctx context.Context) (interface{}, error)
}
