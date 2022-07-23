package service

import (
	"context"
	"fmt"

	domain "github.com/ruslanec/tinkoffbroker"
	tkf "github.com/ruslanec/tinkoffbroker/service/proto"
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

	var tkf_candles []*tkf.CandleInstrument
	for _, candle := range candles {
		tkf_candles = append(tkf_candles, &tkf.CandleInstrument{
			Figi:     candle.Figi,
			Interval: tkf.SubscriptionInterval(candle.Interval),
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &tkf.SubscribeCandlesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkf_candles,
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

	var tkf_candles []*tkf.CandleInstrument
	for _, candle := range candles {
		tkf_candles = append(tkf_candles, &tkf.CandleInstrument{
			Figi:     candle.Figi,
			Interval: tkf.SubscriptionInterval(candle.Interval),
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeCandlesRequest{
			SubscribeCandlesRequest: &tkf.SubscribeCandlesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkf_candles,
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

	var tkf_orderbooks []*tkf.OrderBookInstrument
	for _, orderbook := range orderbooks {
		tkf_orderbooks = append(tkf_orderbooks, &tkf.OrderBookInstrument{
			Figi:  orderbook.Figi,
			Depth: orderbook.Depth,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeOrderBookRequest{
			SubscribeOrderBookRequest: &tkf.SubscribeOrderBookRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkf_orderbooks,
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

	var tkf_orderbooks []*tkf.OrderBookInstrument
	for _, orderbook := range orderbooks {
		tkf_orderbooks = append(tkf_orderbooks, &tkf.OrderBookInstrument{
			Figi:  orderbook.Figi,
			Depth: orderbook.Depth,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeOrderBookRequest{
			SubscribeOrderBookRequest: &tkf.SubscribeOrderBookRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkf_orderbooks,
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

	var tkf_trades []*tkf.TradeInstrument
	for _, trade := range trades {
		tkf_trades = append(tkf_trades, &tkf.TradeInstrument{
			Figi: trade.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeTradesRequest{
			SubscribeTradesRequest: &tkf.SubscribeTradesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkf_trades,
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

	var tkf_trades []*tkf.TradeInstrument
	for _, trade := range trades {
		tkf_trades = append(tkf_trades, &tkf.TradeInstrument{
			Figi: trade.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeTradesRequest{
			SubscribeTradesRequest: &tkf.SubscribeTradesRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkf_trades,
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

	var tkf_instruments []*tkf.InfoInstrument
	for _, instrument := range instruments {
		tkf_instruments = append(tkf_instruments, &tkf.InfoInstrument{
			Figi: instrument.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &tkf.SubscribeInfoRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkf_instruments,
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

	var tkf_instruments []*tkf.InfoInstrument
	for _, instrument := range instruments {
		tkf_instruments = append(tkf_instruments, &tkf.InfoInstrument{
			Figi: instrument.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeInfoRequest{
			SubscribeInfoRequest: &tkf.SubscribeInfoRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkf_instruments,
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

	var tkf_lastprices []*tkf.LastPriceInstrument
	for _, lastprice := range lastprices {
		tkf_lastprices = append(tkf_lastprices, &tkf.LastPriceInstrument{
			Figi: lastprice.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &tkf.SubscribeLastPriceRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_SUBSCRIBE,
				Instruments:        tkf_lastprices,
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

	var tkf_lastprices []*tkf.LastPriceInstrument
	for _, lastprice := range lastprices {
		tkf_lastprices = append(tkf_lastprices, &tkf.LastPriceInstrument{
			Figi: lastprice.Figi,
		})
	}

	return s.stream.Send(&tkf.MarketDataRequest{
		Payload: &tkf.MarketDataRequest_SubscribeLastPriceRequest{
			SubscribeLastPriceRequest: &tkf.SubscribeLastPriceRequest{
				SubscriptionAction: tkf.SubscriptionAction_SUBSCRIPTION_ACTION_UNSUBSCRIBE,
				Instruments:        tkf_lastprices,
			},
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
