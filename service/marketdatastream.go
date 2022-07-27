package service

import (
	"github.com/ruslanec/tinkoffbroker/domain"
	"golang.org/x/net/context"
)

// Сервис получения биржевой информации
type MarketDataStreamService interface {
	// Подписка на свечи
	SubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error
	// Закрытие подписки на свечи
	UnsubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error
	// Подписка на стакан
	SubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error
	// Закрытие подписки на стакан
	UnsubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error
	// Подписка на ленту сделок
	SubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error
	// Закрытие подписки на ленту сделок
	UnsubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error
	// Подписка на торговые статусы
	SubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error
	// Закрытие подписки на торговые статусы
	UnsubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error
	// Подписка на последнюю цену инструмента
	SubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error
	// Закрытие подписки на последнюю цену инструмента
	UnsubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error
	// Запрос активных подписок
	MySubscriptions(ctx context.Context) error
	// Получение данных по подпискам
	Recv(ctx context.Context) (interface{}, error)
}
