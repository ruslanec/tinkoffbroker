package marketdata

import (
	"context"
	"strconv"
	"time"

	"github.com/ruslanec/timerange-go"
	"github.com/ruslanec/tinkoffbroker"
	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/* Сервис получения биржевой информации:
1. свечи;
2. стаканы;
3. торговые статусы;
4. лента сделок. */
type marketDataService struct {
	conn   *grpc.ClientConn
	client tkf.MarketDataServiceClient
}

/*
Глубина запроса свечей в зависимости от интервала:
Интервал свечи		Допустимый период
1 минута			от 1 минут до 1 дня
5 минут				от 5 минут до 1 дня
15 минут			от 15 минут до 1 дня
1 час				от 1 часа до 1 недели
1 день				от 1 дня до 1 года
*/
const (
	maxPeriodIntervalForMinuteCandles = 24 * time.Hour
	maxPeriodIntervalForHourCandles   = 7 * 24 * time.Hour
	maxPeriodIntervalForDayCandles    = 365 * 24 * time.Hour
)

// Доступные значения глубины стакана: 1, 10, 20, 30, 40, 50.
var allowableOrderBookDepth = []int32{1, 10, 20, 30, 40, 50}

// Конструктор сервиса
func NewMarketDataService(conn *grpc.ClientConn) service.MarketDataService {
	return &marketDataService{
		conn:   conn,
		client: tkf.NewMarketDataServiceClient(conn),
	}
}

// Метод запроса последних цен по инструментам
func (s *marketDataService) LastPrices(ctx context.Context, figi []string) ([]*domain.LastPrice, error) {
	resp, err := s.client.GetLastPrices(ctx, &tkf.GetLastPricesRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	var prices []*domain.LastPrice
	for _, v := range resp.GetLastPrices() {
		t := v.GetTime().AsTime()
		prices = append(prices, &domain.LastPrice{
			Figi:  v.GetFigi(),
			Price: service.ConvQuotationFromTkf(v.GetPrice()),
			Time:  &t,
		})
	}
	return prices, nil
}

// Метод запроса исторических свечей по инструменту
func (s *marketDataService) Candles(ctx context.Context, figi string, from, to time.Time, interval domain.CandleInterval) ([]*domain.Candle, error) {
	ctx, cancel := context.WithTimeout(ctx, tinkoffbroker.DefaultRequestTimeout) // TODO
	defer cancel()

	var timerangeInterval time.Duration
	switch tkf.CandleInterval(interval) {
	case tkf.CandleInterval_CANDLE_INTERVAL_UNSPECIFIED:
		return nil, tinkoffbroker.ErrArgCandleUnspecified
	case tkf.CandleInterval_CANDLE_INTERVAL_1_MIN, tkf.CandleInterval_CANDLE_INTERVAL_5_MIN, tkf.CandleInterval_CANDLE_INTERVAL_15_MIN:
		timerangeInterval = maxPeriodIntervalForMinuteCandles
	case tkf.CandleInterval_CANDLE_INTERVAL_HOUR:
		timerangeInterval = maxPeriodIntervalForHourCandles
	case tkf.CandleInterval_CANDLE_INTERVAL_DAY:
		timerangeInterval = maxPeriodIntervalForDayCandles
	default:
		return nil, tinkoffbroker.ErrArgCandleUnspecified
	}

	var candles []*domain.Candle
	timerangeIter := timerange.New(from, to, timerangeInterval)
	for timerangeIter.Next() {
		start := timerangeIter.Current()
		if start == to {
			continue
		}
		end := start.Add(timerangeInterval)
		if end.After(to) {
			end = to
		}

		var header metadata.MD
		resp, err := s.client.GetCandles(ctx, &tkf.GetCandlesRequest{
			Figi:     figi,
			From:     timestamppb.New(start),
			To:       timestamppb.New(end),
			Interval: tkf.CandleInterval(interval),
		}, grpc.Header(&header))
		if err != nil {
			return nil, err
		}

		// Контроль ограничений по тарифу. При исчерпании кол-ва запросов делаем паузу до сброса счетчика запросов
		if limitRemaining, ok := header["x-ratelimit-remaining"]; ok {
			if len(limitRemaining) > 0 && limitRemaining[0] == "0" {
				if limitReset, ok := header["x-ratelimit-reset"]; ok {
					if len(limitReset) > 0 {
						count, err := strconv.Atoi(limitReset[0])
						if err == nil {
							time.Sleep(time.Second * time.Duration(count+5))
						}
					}
				}
			}
		}

		for _, v := range resp.GetCandles() {
			dt := v.Time.AsTime()
			candles = append(candles, &domain.Candle{
				Figi:     figi,
				DateTime: &dt,
				Interval: interval,
				Open:     service.ConvQuotationFromTkf(v.GetOpen()),
				High:     service.ConvQuotationFromTkf(v.GetHigh()),
				Low:      service.ConvQuotationFromTkf(v.GetLow()),
				Close:    service.ConvQuotationFromTkf(v.GetClose()),
				Volume:   v.GetVolume(),
			})
		}
	}
	return candles, nil
}

// Метод получения стакана по инструменту
func (s *marketDataService) OrderBook(ctx context.Context, figi string, depth int32) (*domain.OrderBook, error) {
	if !contains(allowableOrderBookDepth, depth) {
		return nil, tinkoffbroker.ErrArgCandleUnspecified
	}

	resp, err := s.client.GetOrderBook(ctx, &tkf.GetOrderBookRequest{
		Figi:  figi,
		Depth: depth,
	})
	if err != nil {
		return nil, err
	}

	var bids, asks []*domain.Order
	for _, v := range resp.GetBids() {
		bids = append(bids, &domain.Order{
			Price:    service.ConvQuotationFromTkf(v.GetPrice()),
			Quantity: v.GetQuantity(),
		})
	}
	for _, v := range resp.GetAsks() {
		asks = append(asks, &domain.Order{
			Price:    service.ConvQuotationFromTkf(v.GetPrice()),
			Quantity: v.GetQuantity(),
		})
	}

	return &domain.OrderBook{
		Figi:       resp.GetFigi(),
		Depth:      resp.GetDepth(),
		Bids:       bids,
		Asks:       asks,
		LastPrice:  service.ConvQuotationFromTkf(resp.GetLastPrice()),
		ClosePrice: service.ConvQuotationFromTkf(resp.GetClosePrice()),
		LimitUp:    service.ConvQuotationFromTkf(resp.GetLimitUp()),
		LimitDown:  service.ConvQuotationFromTkf(resp.GetLimitDown()),
	}, nil
}

// Метод запроса статуса торгов по инструментам
func (s *marketDataService) TradingStatus(ctx context.Context, figi string) (*domain.InstrumentTradingStatus, error) {
	resp, err := s.client.GetTradingStatus(ctx, &tkf.GetTradingStatusRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	return &domain.InstrumentTradingStatus{
		Figi:          resp.GetFigi(),
		TradingStatus: domain.SecurityTradingStatus(resp.GetTradingStatus()),
	}, nil
}

func contains(array []int32, value int32) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

// Метод запроса обезличенных сделок за последний час.
func (s *marketDataService) LastTrades(ctx context.Context, figi string, from, to time.Time) ([]*domain.Trade, error) {
	if figi == "" {
		return nil, tinkoffbroker.ErrArgEmptyFigi
	}

	if from.After(to) {
		return nil, tinkoffbroker.ErrArgTimeInterval
	}

	resp, err := s.client.GetLastTrades(ctx, &tkf.GetLastTradesRequest{
		Figi: figi,
		From: timestamppb.New(from),
		To:   timestamppb.New(to),
	})
	if err != nil {
		return nil, err
	}

	tkfTrades := resp.GetTrades()
	trades := make([]*domain.Trade, 0, len(tkfTrades))
	for _, tkfTrade := range tkfTrades {
		trades = append(trades, service.ConvTrade(tkfTrade))
	}

	return trades, nil
}
