package ordersstream

import (
	"context"
	"errors"
	"fmt"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

type ordersStreamService struct {
	conn         *grpc.ClientConn
	client       tkf.OrdersStreamServiceClient
	streamClient tkf.OrdersStreamService_TradesStreamClient
}

func NewOrdersStreamService(conn *grpc.ClientConn) service.OrdersStreamService {
	return &ordersStreamService{
		conn:   conn,
		client: tkf.NewOrdersStreamServiceClient(conn),
	}
}

// Создать подписку на поток сделок пользователя
func (s *ordersStreamService) SubscribeOrderTrades(ctx context.Context) error {
	var err error
	if s.streamClient == nil {
		s.streamClient, err = s.client.TradesStream(ctx, &tkf.TradesStreamRequest{})
		if err != nil {
			return err
		}
	}
	return nil
}

// Отписаться от потока сделок пользователя
func (s *ordersStreamService) UnsubscribeOrderTrades(ctx context.Context) error {
	if s.streamClient == nil {
		return errors.New("OrdersStreamService_TradesStreamClient is nil")
	}

	err := s.streamClient.CloseSend()
	if err != nil {
		return err
	}
	return nil
}

// Получение сделок пользователя по подписке
func (s *ordersStreamService) Recv(ctx context.Context) (interface{}, error) {
	if s.streamClient == nil {
		return nil, errors.New("OrdersStreamService_TradesStreamClient is nil")
	}

	resp, err := s.streamClient.Recv()
	if err != nil {
		return nil, err
	}

	switch v := resp.Payload.(type) {
	case *tkf.TradesStreamResponse_OrderTrades:
		return v.OrderTrades, nil
	case *tkf.TradesStreamResponse_Ping: // TODO: обработать отстутсвие ping
		return v.Ping.GetTime(), nil
	default:
		return nil, fmt.Errorf("received unknown response from stream: %v", v)
	}
}
