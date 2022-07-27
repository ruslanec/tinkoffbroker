package tinkoffbroker

import (
	"context"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
)

type Client interface {
	Run(ctx context.Context) error
	Close()
	Init(ctx context.Context) error
	// Получение данных по подпискам
	Recv(ctx context.Context) (interface{}, error)

	// Сервис предоставления справочной информации о ценных бумагах
	// Метод получения списка акций
	Shares(ctx context.Context) ([]*domain.Share, error)
	// Метод получения фьючерса по FIGI
	ShareByFigi(ctx context.Context, figi string) (*domain.Share, error)
	// Метод получения списка облигаций
	Bonds(ctx context.Context) (shares []*domain.Bond, err error)
	// Метод получения облигации по FIGI
	BondByFigi(ctx context.Context, figi string) (*domain.Bond, error)
	// Запрос купонов по облигации
	BondCoupons(ctx context.Context, figi string, from, to *time.Time) ([]*domain.Coupon, error)
	// Метод получения накопленного купонного дохода по облигации
	AccruedInterests(ctx context.Context, figi string, from, to *time.Time) ([]*domain.AccruedInterest, error)
	// Метод получения списка валют
	Currencies(ctx context.Context) (shares []*domain.Currency, err error)
	// Метод получения валюты по FIGI
	CurrencyByFigi(ctx context.Context, figi string) (*domain.Currency, error)
	// Метод получения списка инвестиционных фондов
	Etfs(ctx context.Context) ([]*domain.Etf, error)
	// Метод получения инвестиционного фонда по его идентификатору
	EtfByFigi(ctx context.Context, figi string) (*domain.Etf, error)
	// Метод получения списка фьючерсов
	Future(ctx context.Context) ([]*domain.Future, error)
	// Метод получения фьючерса по FIGI
	FutureByFigi(ctx context.Context, figi string) (*domain.Future, error)
	// Метод получения расписания торгов торговых площадок
	TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*domain.TradingSchedule, error)
	// Метод получения размера гарантийного обеспечения по фьючерсам
	FuturesMargin(ctx context.Context, figi string) (*domain.FuturesMargin, error)
	// Метод получения основной информации об инструменте
	InstrumentByFigi(ctx context.Context, figi string) (*domain.Instrument, error)
	// Метод для получения событий выплаты дивидендов по инструменту
	Dividends(ctx context.Context, figi string, from, to *time.Time) ([]*domain.Dividend, error)
	// Метод получения актива по его идентификатору
	AssetById(ctx context.Context, id string) (*domain.AssetFull, error)
	// Метод получения списка активов
	Assets(ctx context.Context) ([]*domain.Asset, error)
	// Метод получения списка избранных инструментов
	Favorites(ctx context.Context) ([]*domain.FavoriteInstrument, error)
	// Метод редактирования списка избранных инструментов
	EditFavorites(ctx context.Context, figies []string, action domain.EditFavoritesActionType) ([]*domain.FavoriteInstrument, error)
	// Метод получения списка стран
	Countries(ctx context.Context) ([]*domain.Country, error)
	// Метод поиска инструмента
	FindInstrument(ctx context.Context, query string) ([]*domain.InstrumentShort, error)
	// Метод получения списка брендов
	Brands(ctx context.Context) ([]*domain.Brand, error)
	// Метод получения бренда по его идентификатору
	BrandById(ctx context.Context, id string) (*domain.Brand, error)

	// Сервис получения биржевой информации
	// Метод запроса статуса торгов по инструментам
	TradingStatus(ctx context.Context, figi string) (*domain.InstrumentTradingStatus, error)
	// Метод запроса последних цен по инструментам
	LastPrices(ctx context.Context, figi []string) ([]*domain.LastPrice, error)
	// Метод запроса исторических свечей по инструменту
	Candles(ctx context.Context, figi string, from, to time.Time, interval domain.CandleInterval) ([]*domain.Candle, error)
	// Метод получения стакана по инструменту
	OrderBook(ctx context.Context, figi string, depth int32) (*domain.OrderBook, error)

	// Сервис получения биржевой информации
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

	// Сервис работы с торговыми поручениями
	// Метод выставления рыночной заявки на покупку
	OrderBuyLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления лимитной заявки на продажу
	OrderSellLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления рыночной заявки на покупку
	OrderBuyMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления рыночной заявки на продажу
	OrderSellMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод отмены биржевой заявки
	CancelOrder(ctx context.Context, orderId string) (*time.Time, error)
	// Метод получения статуса торгового поручения
	OrderState(ctx context.Context, orderId string) (*domain.OrderState, error)
	// Метод получения списка активных заявок по счёту
	Orders(ctx context.Context) ([]*domain.OrderState, error)

	// Потоковый сервис получения информации о сделках пользователя
	// Создать подписку на поток сделок пользователя
	SubscribeOrderTrades(ctx context.Context) error
	// Отписаться от потока сделок пользователя
	UnsubscribeOrderTrades(ctx context.Context) error

	// Сервис получения информации о портфеле по конкретному счёту
	// Метод получения портфеля по счёту
	Portfolio(ctx context.Context) (*domain.Portfolio, error)
	// Метод получения списка операций по счёту
	Operations(ctx context.Context, from, to *time.Time, state domain.OperationState, figi string) ([]*domain.Operation, error)
	// Метод получения списка позиций по счёту
	Positions(ctx context.Context) (*domain.Positions, error)

	// Сервис стоп-заявок
	// ------------------------------------------------------

	// Сервис предоставления информации об аккаунтах
	// Метод получения открытых и активных счетов пользователя
	Accounts(ctx context.Context) ([]*domain.Account, error)
	// Запрос тарифных лимитов пользователя
	UserTariff(ctx context.Context) (*domain.UserTariff, error)
	// Расчёт маржинальных показателей по счёту
	MarginAttributes(ctx context.Context) (*domain.MarginAttributes, error)
}
