package marketdatastream

import (
	"context"
	"fmt"

	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

type marketDataStreamService struct {
	conn   *grpc.ClientConn
	client tkf.MarketDataStreamServiceClient
	stream tkf.MarketDataStreamService_MarketDataStreamClient
}

// Конструктор сервиса
func NewMarketDataStreamService(conn *grpc.ClientConn) service.MarketDataStreamService {
	return &marketDataStreamService{
		conn:   conn,
		client: tkf.NewMarketDataStreamServiceClient(conn),
	}
}

// Подписка на свечи
func (s *marketDataStreamService) SubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("SubscribeCandles: %v", err)
		}
	}

	var tkfCandles []*tkf.CandleInstrument
	for _, candle := range candles {
		tkfCandles = append(tkfCandles, &tkf.CandleInstrument{
			Figi:     candle.Figi,
			Interval: tkf.SubscriptionInterval(candle.Interval),
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &tkf.SubscribeCandlesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkfCandles,
			},
		},
	})
}

// Закрытие подписки на свечи
func (s *marketDataStreamService) UnsubscribeCandles(ctx context.Context, candles []*domain.CandleInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("UnsubscribeCandles: %v", err)
		}
	}

	tkfCandles := make([]*tkf.CandleInstrument, 0, len(candles))
	for _, candle := range candles {
		tkfCandles = append(tkfCandles, &tkf.CandleInstrument{
			Figi:     candle.Figi,
			Interval: tkf.SubscriptionInterval(candle.Interval),
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &tkf.SubscribeCandlesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkfCandles,
			},
		},
	})
}

// Подписка на стакан
func (s *marketDataStreamService) SubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("SubscribeOrderBook: %v", err)
		}
	}

	tkfOrderbooks := make([]*tkf.OrderBookInstrument, 0, len(orderbooks))
	for _, orderbook := range orderbooks {
		tkfOrderbooks = append(tkfOrderbooks, &tkf.OrderBookInstrument{
			Figi:  orderbook.Figi,
			Depth: orderbook.Depth,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeOrderBookRequest{
			SubscribeOrderBookRequest: &tkf.SubscribeOrderBookRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkfOrderbooks,
			},
		},
	})
}

// Закрытие подписки на стакан
func (s *marketDataStreamService) UnsubscribeOrderBook(ctx context.Context, orderbooks []*domain.OrderBookInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("UnsubscribeOrderBook: %v", err)
		}
	}

	tkfOrderbooks := make([]*tkf.OrderBookInstrument, 0, len(orderbooks))
	for _, orderbook := range orderbooks {
		tkfOrderbooks = append(tkfOrderbooks, &tkf.OrderBookInstrument{
			Figi:  orderbook.Figi,
			Depth: orderbook.Depth,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeOrderBookRequest{
			SubscribeOrderBookRequest: &tkf.SubscribeOrderBookRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkfOrderbooks,
			},
		},
	})
}

// Подписка на ленту сделок
func (s *marketDataStreamService) SubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("SubscribeTrades: %v", err)
		}
	}

	tkfTrades := make([]*tkf.TradeInstrument, 0, len(trades))
	for _, trade := range trades {
		tkfTrades = append(tkfTrades, &tkf.TradeInstrument{
			Figi: trade.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeTradesRequest{
			SubscribeTradesRequest: &tkf.SubscribeTradesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkfTrades,
			},
		},
	})
}

// Закрытие подписки на ленту сделок
func (s *marketDataStreamService) UnsubscribeTrades(ctx context.Context, trades []*domain.TradeInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("UnsubscribeTrades: %v", err)
		}
	}

	tkfTrades := make([]*tkf.TradeInstrument, 0, len(trades))
	for _, trade := range trades {
		tkfTrades = append(tkfTrades, &tkf.TradeInstrument{
			Figi: trade.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeTradesRequest{
			SubscribeTradesRequest: &tkf.SubscribeTradesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkfTrades,
			},
		},
	})
}

// Подписка на торговые статусы
func (s *marketDataStreamService) SubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("SubscribeInfo: %v", err)
		}
	}

	tkfInstruments := make([]*tkf.InfoInstrument, 0, len(instruments))
	for _, instrument := range instruments {
		tkfInstruments = append(tkfInstruments, &tkf.InfoInstrument{
			Figi: instrument.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &tkf.SubscribeInfoRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkfInstruments,
			},
		},
	})
}

// Закрытие подписки на торговые статусы
func (s *marketDataStreamService) UnsubscribeInfo(ctx context.Context, instruments []*domain.InfoInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("UnsubscribeInfo: %v", err)
		}
	}

	tkfInstruments := make([]*tkf.InfoInstrument, 0, len(instruments))
	for _, instrument := range instruments {
		tkfInstruments = append(tkfInstruments, &tkf.InfoInstrument{
			Figi: instrument.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &tkf.SubscribeInfoRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkfInstruments,
			},
		},
	})
}

// Подписка на последнюю цену инструмента
func (s *marketDataStreamService) SubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("SubscribeLastPrices: %v", err)
		}
	}

	tkfLastprices := make([]*tkf.LastPriceInstrument, 0, len(lastprices))
	for _, lastprice := range lastprices {
		tkfLastprices = append(tkfLastprices, &tkf.LastPriceInstrument{
			Figi: lastprice.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &tkf.SubscribeLastPriceRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkfLastprices,
			},
		},
	})
}

// Закрытие подписки на последнюю цену инструмента
func (s *marketDataStreamService) UnsubscribeLastPrices(ctx context.Context, lastprices []*domain.LastPriceInstrument) error {
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("UnsubscribeLastPrices: %v", err)
		}
	}

	tkfLastprices := make([]*tkf.LastPriceInstrument, 0, len(lastprices))
	for _, lastprice := range lastprices {
		tkfLastprices = append(tkfLastprices, &tkf.LastPriceInstrument{
			Figi: lastprice.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &tkf.SubscribeLastPriceRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkfLastprices,
			},
		},
	})
}

// Запрос активных подписок
func (s *marketDataStreamService) MySubscriptions(ctx context.Context) error { // TODO Not working
	var err error
	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return fmt.Errorf("MySubscriptions: %v", err)
		}
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_GetMySubscriptions{
			GetMySubscriptions: &tkf.GetMySubscriptions{},
		},
	})
}

// Получение данных по подпискам
func (s *marketDataStreamService) Recv(ctx context.Context) (interface{}, error) {
	var err error

	if s.stream == nil {
		s.stream, err = s.client.MarketDataStream(ctx)
		if err != nil {
			return nil, fmt.Errorf("Recv: %v", err)
		}
	}

	resp, err := s.stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("Recv: %v", err)
	}
	switch v := resp.GetPayload().(type) {
	case *tkf.MarketDataResponse_SubscribeCandlesResponse:
		return v.SubscribeCandlesResponse.GetCandlesSubscriptions(), nil
	case *tkf.MarketDataResponse_SubscribeOrderBookResponse:
		return v.SubscribeOrderBookResponse.GetOrderBookSubscriptions(), nil
	case *tkf.MarketDataResponse_SubscribeTradesResponse:
		return v.SubscribeTradesResponse.GetTradeSubscriptions(), nil
	case *tkf.MarketDataResponse_SubscribeInfoResponse:
		return v.SubscribeInfoResponse.GetInfoSubscriptions(), nil
	case *tkf.MarketDataResponse_SubscribeLastPriceResponse:
		return v.SubscribeLastPriceResponse.GetLastPriceSubscriptions(), nil
	case *tkf.MarketDataResponse_Candle:
		return v.Candle, nil
	case *tkf.MarketDataResponse_Orderbook:
		return v.Orderbook, nil
	case *tkf.MarketDataResponse_Trade:
		return v.Trade, nil
	case *tkf.MarketDataResponse_LastPrice:
		return v.LastPrice, nil
	case *tkf.MarketDataResponse_TradingStatus:
		return v.TradingStatus, nil
	case *tkf.MarketDataResponse_Ping:
		return v.Ping.Time.AsTime(), nil
	default:
		return nil, fmt.Errorf("received unknown response from stream: %T", v)
	}
}
