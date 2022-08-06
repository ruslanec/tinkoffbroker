package tinkoffbroker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

type Services struct {
	Users            service.UsersService
	Instruments      service.InstrumentsService
	OrdersStream     service.OrdersStreamService
	Orders           service.OrdersService
	MarketData       service.MarketDataService
	MarketDataStream service.MarketDataStreamService
	StopOrders       service.StopOrdersService
	Operations       service.OperationsService
}

type client struct {
	conn      *grpc.ClientConn
	accountId string
	services  Services
}

//Опциональные параметры
type Option func(Client)

func WithServices(s Services) Option {
	return func(c Client) {
		c.(*client).services = s
	}
}

func WithUsers(s service.UsersService) Option {
	return func(c Client) {
		c.(*client).services.Users = s
	}
}

func WithStopOrders(s service.StopOrdersService) Option {
	return func(c Client) {
		c.(*client).services.StopOrders = s
	}
}

func WithOrders(s service.OrdersService) Option {
	return func(c Client) {
		c.(*client).services.Orders = s
	}
}

func WithOrdersStream(s service.OrdersStreamService) Option {
	return func(c Client) {
		c.(*client).services.OrdersStream = s
	}
}

func WithOperations(s service.OperationsService) Option {
	return func(c Client) {
		c.(*client).services.Operations = s
	}
}

func WithMarketData(s service.MarketDataService) Option {
	return func(c Client) {
		c.(*client).services.MarketData = s
	}
}

func WithMarketDataStream(s service.MarketDataStreamService) Option {
	return func(c Client) {
		c.(*client).services.MarketDataStream = s
	}
}

func WithInstruments(s service.InstrumentsService) Option {
	return func(c Client) {
		c.(*client).services.Instruments = s
	}
}

//Конструктор
func NewClient(conn *grpc.ClientConn, accountId string, opts ...Option) Client {
	c := client{
		conn:      conn,
		accountId: accountId,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

func (c *client) Init(ctx context.Context) error { //TODO Избавиться от Init
	accounts, err := c.services.Users.Accounts(ctx)
	if err != nil {
		return err
	}
	var account *domain.Account
	for _, v := range accounts {
		if v.Id == c.accountId {
			account = v
		}
	}
	if account == nil {
		return fmt.Errorf("accountid %s not found", c.accountId)
	}

	return nil
}

//Получение данных по подпискам
func (c *client) Recv(ctx context.Context) (interface{}, error) {
	if c.services.MarketDataStream == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.Recv(ctx)
}

func (c *client) Run(ctx context.Context) (err error) {

	return nil
}

//Close
func (c *client) Close() {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}
}

//Метод получения списка акций
func (c *client) Shares(ctx context.Context) ([]*domain.Share, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Shares(ctx)
}

//Метод получения фьючерса по FIGI
func (c *client) ShareByFigi(ctx context.Context, figi string) (*domain.Share, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.ShareByFigi(ctx, figi)
}

//Метод получения списка облигаций
func (c *client) Bonds(ctx context.Context) (shares []*domain.Bond, err error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Bonds(ctx)
}

//Метод получения облигации по FIGI
func (c *client) BondByFigi(ctx context.Context, figi string) (*domain.Bond, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.BondByFigi(ctx, figi)
}

//Запрос купонов по облигации
func (c *client) BondCoupons(ctx context.Context, figi string, from, to time.Time) ([]*domain.Coupon, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.BondCoupons(ctx, figi, from, to)
}

//Метод получения накопленного купонного дохода по облигации
func (c *client) AccruedInterests(ctx context.Context, figi string, from, to time.Time) ([]*domain.AccruedInterest, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.AccruedInterests(ctx, figi, from, to)
}

//Метод получения списка валют
func (c *client) Currencies(ctx context.Context) (shares []*domain.Currency, err error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Currencies(ctx)
}

//Метод получения валюты по FIGI
func (c *client) CurrencyByFigi(ctx context.Context, figi string) (*domain.Currency, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.CurrencyByFigi(ctx, figi)
}

//Метод получения списка инвестиционных фондов
func (c *client) Etfs(ctx context.Context) ([]*domain.Etf, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Etfs(ctx)
}

//Метод получения инвестиционного фонда по его идентификатору
func (c *client) EtfByFigi(ctx context.Context, figi string) (*domain.Etf, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.EtfByFigi(ctx, figi)
}

//Метод получения списка фьючерсов
func (c *client) Future(ctx context.Context) ([]*domain.Future, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Future(ctx)
}

//Метод получения фьючерса по FIGI
func (c *client) FutureByFigi(ctx context.Context, figi string) (*domain.Future, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.FutureByFigi(ctx, figi)
}

//Метод получения расписания торгов торговых площадок
func (c *client) TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*domain.TradingSchedule, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.TradingSchedules(ctx, exchange, from, to)
}

//Метод получения размера гарантийного обеспечения по фьючерсам
func (c *client) FuturesMargin(ctx context.Context, figi string) (*domain.FuturesMargin, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.FuturesMargin(ctx, figi)
}

//Метод получения основной информации об инструменте
func (c *client) InstrumentByFigi(ctx context.Context, figi string) (*domain.Instrument, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.InstrumentByFigi(ctx, figi)
}

//Метод для получения событий выплаты дивидендов по инструменту
func (c *client) Dividends(ctx context.Context, figi string, from, to time.Time) ([]*domain.Dividend, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Dividends(ctx, figi, from, to)
}

//Метод получения актива по его идентификатору
func (c *client) AssetById(ctx context.Context, id string) (*domain.AssetFull, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.AssetById(ctx, id)
}

//Метод получения списка активов
func (c *client) Assets(ctx context.Context) ([]*domain.Asset, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Assets(ctx)
}

//Метод получения списка избранных инструментов
func (c *client) Favorites(ctx context.Context) ([]*domain.FavoriteInstrument, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Favorites(ctx)
}

//Метод редактирования списка избранных инструментов
func (c *client) EditFavorites(ctx context.Context, figies []string, action domain.EditFavoritesActionType) ([]*domain.FavoriteInstrument, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.EditFavorites(ctx, figies, action)
}

//Метод получения списка стран
func (c *client) Countries(ctx context.Context) ([]*domain.Country, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Countries(ctx)
}

//Метод поиска инструмента
func (c *client) FindInstrument(ctx context.Context, query string) ([]*domain.InstrumentShort, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.FindInstrument(ctx, query)
}

//Метод получения списка брендов
func (c *client) Brands(ctx context.Context) ([]*domain.Brand, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.Brands(ctx)
}

//Метод получения бренда по его идентификатору
func (c *client) BrandById(ctx context.Context, id string) (*domain.Brand, error) {
	if c.services.Instruments == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Instruments.BrandById(ctx, id)
}

//Метод запроса статуса торгов по инструментам
func (c *client) TradingStatus(ctx context.Context, figi string) (*domain.InstrumentTradingStatus, error) {
	if c.services.MarketData == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketData.TradingStatus(ctx, figi)
}

//Метод запроса последних цен по инструментам
func (c *client) LastPrices(ctx context.Context, figi []string) ([]*domain.LastPrice, error) {
	if c.services.MarketData == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketData.LastPrices(ctx, figi)
}

//Метод запроса исторических свечей по инструменту
func (c *client) Candles(ctx context.Context, figi string, from, to time.Time, interval domain.CandleInterval) ([]*domain.Candle, error) {
	if c.services.MarketData == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketData.Candles(ctx, figi, from, to, interval)
}

//Метод получения стакана по инструменту
func (c *client) OrderBook(ctx context.Context, figi string, depth int32) (*domain.OrderBook, error) {
	if c.services.MarketData == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketData.OrderBook(ctx, figi, depth)
}

//Метод запроса обезличенных сделок за последний час.
func (c *client) LastTrades(ctx context.Context, figi string, from, to time.Time) ([]*domain.Trade, error) {
	if c.services.MarketData == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.MarketData.LastTrades(ctx, figi, from, to)
}

//Метод выставления рыночной заявки на покупку
func (c *client) OrderBuyLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	//проверить уровень доступа к текущему счёту (нужен доступ к usersService - Accounts)
	//проверить будет ли шорт
	//проверить разрешен и шорт по счету (нужен доступ к operationsService - Portfolio)
	//проверить маржинальные показатели по счету (нужен доступ к usersService - MarginAttributes)
	//проверить ограничения по тарифу (нужен доступ к usersService - UserTariff)
	//проверить наличие стоп-заявок (нужен доступ к StopOrdersService - StopOrders)
	//отменить выставленные стоп-заявки (нужен доступ к StopOrdersService - CancelStopOrders)
	return c.services.Orders.OrderBuyLimit(ctx, figi, quantity, price)
	//выствить стоп-заявки (нужен доступ к StopOrdersService - PostStopOrders)
}

//Метод выставления лимитной заявки на продажу
func (c *client) OrderSellLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	//проверить уровень доступа к текущему счёту (нужен доступ к usersService - Accounts)
	//проверить наличие доступного количества (нужен доступ к operationsService - Portfolio)
	//проверить наличие стоп-заявок (нужен доступ к StopOrdersService - StopOrders)
	//отменить выставленные стоп-заявки (нужен доступ к StopOrdersService - CancelStopOrders)
	return c.services.Orders.OrderSellLimit(ctx, figi, quantity, price)
	//выствить стоп-заявки (нужен доступ к StopOrdersService - PostStopOrders)
}

//Метод выставления рыночной заявки на покупку
func (c *client) OrderBuyMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	//проверить уровень доступа к текущему счёту (нужен доступ к usersService - Accounts)
	//проверить будет ли шорт
	//проверить разрешен и шорт по счету (нужен доступ к operationsService - Portfolio)
	//проверить маржинальные показатели по счету (нужен доступ к usersService - MarginAttributes)
	//проверить ограничения по тарифу (нужен доступ к usersService - UserTariff)
	//проверить наличие стоп-заявок (нужен доступ к StopOrdersService - StopOrders)
	//отменить выставленные стоп-заявки (нужен доступ к StopOrdersService - CancelStopOrders)
	return c.services.Orders.OrderBuyMarket(ctx, figi, quantity, price)
	//выствить стоп-заявки (нужен доступ к StopOrdersService - PostStopOrders)
}

//Метод выставления рыночной заявки на продажу
func (c *client) OrderSellMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	//проверить уровень доступа к текущему счёту (нужен доступ к usersService - Accounts)
	//проверить наличие доступного количества (нужен доступ к operationsService - Portfolio)
	//проверить наличие стоп-заявок (нужен доступ к StopOrdersService - StopOrders)
	//отменить выставленные стоп-заявки (нужен доступ к StopOrdersService - CancelStopOrders)
	return c.services.Orders.OrderSellMarket(ctx, figi, quantity, price)
	//выствить стоп-заявки (нужен доступ к StopOrdersService - PostStopOrders)
}

//Метод отмены биржевой заявки
func (c *client) CancelOrder(ctx context.Context, orderId string) (*time.Time, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Orders.CancelOrder(ctx, orderId)
}

//Метод получения статуса торгового поручения
func (c *client) OrderState(ctx context.Context, orderId string) (*domain.OrderState, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Orders.OrderState(ctx, orderId)
}

//Метод получения списка активных заявок по счёту
func (c *client) Orders(ctx context.Context) ([]*domain.OrderState, error) {
	if c.services.Orders == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Orders.Orders(ctx)
}

//Создать подписку на поток сделок пользователя
func (c *client) SubscribeOrderTrades(ctx context.Context) error {
	if c.services.OrdersStream == nil {
		return ErrSvcNotImplemented
	}

	return c.services.OrdersStream.SubscribeOrderTrades(ctx)
}

//Отписаться от потока сделок пользователя
func (c *client) UnsubscribeOrderTrades(ctx context.Context) error {
	if c.services.OrdersStream == nil {
		return ErrSvcNotImplemented
	}

	return c.services.OrdersStream.UnsubscribeOrderTrades(ctx)
}

//Метод получения портфеля по счёту
func (c *client) Portfolio(ctx context.Context) (*domain.Portfolio, error) {
	if c.services.Operations == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Operations.Portfolio(ctx)
}

//Метод получения списка операций по счёту
func (c *client) Operations(ctx context.Context, from, to *time.Time, state domain.OperationState, figi string) ([]*domain.Operation, error) {
	if c.services.Operations == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Operations.Operations(ctx, from, to, state, figi)
}

//Метод получения списка позиций по счёту
func (c *client) Positions(ctx context.Context) (*domain.Positions, error) {
	if c.services.Operations == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Operations.Positions(ctx)
}

//Метод получения открытых и активных счетов пользователя
func (c *client) Accounts(ctx context.Context) ([]*domain.Account, error) {
	if c.services.Users == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Users.Accounts(ctx)
}

//Запрос тарифных лимитов пользователя
func (c *client) UserTariff(ctx context.Context) (*domain.UserTariff, error) {
	if c.services.Users == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Users.UserTariff(ctx)
}

//Расчёт маржинальных показателей по счёту
func (c *client) MarginAttributes(ctx context.Context) (*domain.MarginAttributes, error) {
	if c.services.Users == nil {
		return nil, ErrSvcNotImplemented
	}
	return c.services.Users.MarginAttributes(ctx)
}

//Подписка на свечи
func (c *client) SubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.SubscribeCandles(ctx, candles)
}

//Закрытие подписки на свечи
func (c *client) UnsubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.UnsubscribeCandles(ctx, candles)
}

//Подписка на стакан
func (c *client) SubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.SubscribeOrderBook(ctx, orderbooks)
}

//Закрытие подписки на стакан
func (c *client) UnsubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.UnsubscribeOrderBook(ctx, orderbooks)
}

//Подписка на ленту сделок
func (c *client) SubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.SubscribeTrades(ctx, trades)
}

//Закрытие подписки на ленту сделок
func (c *client) UnsubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.UnsubscribeTrades(ctx, trades)
}

//Подписка на торговые статусы
func (c *client) SubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.SubscribeInfo(ctx, instruments)
}

//Закрытие подписки на торговые статусы
func (c *client) UnsubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.UnsubscribeInfo(ctx, instruments)
}

//Подписка на последнюю цену инструмента
func (c *client) SubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.SubscribeLastPrices(ctx, lastprices)
}

//Закрытие подписки на последнюю цену инструмента
func (c *client) UnsubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.UnsubscribeLastPrices(ctx, lastprices)
}

//Запрос активных подписок
func (c *client) MySubscriptions(ctx context.Context) error {
	if c.services.MarketDataStream == nil {
		return ErrSvcNotImplemented
	}
	return c.services.MarketDataStream.MySubscriptions(ctx)
}
