/* Сервис получения информации о портфеле по конкретному счёту*/
package operations

import (
	"context"
	"time"

	"github.com/ruslanec/tinkoffbroker"
	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type operationsService struct {
	conn      *grpc.ClientConn
	accountId string
	client    tkf.OperationsServiceClient
}

// Конструктор сервиса
func NewOperationsService(conn *grpc.ClientConn, accountId string) service.OperationsService {
	client := tkf.NewOperationsServiceClient(conn)

	return &operationsService{
		conn:      conn,
		accountId: accountId,
		client:    client,
	}
}

// Метод получения портфеля по счёту
func (s *operationsService) Portfolio(ctx context.Context) (*domain.Portfolio, error) { // TOD remove connections
	if s.accountId == "" {
		return nil, tinkoffbroker.ErrArgEmptyAccounID
	}

	resp, err := s.client.GetPortfolio(ctx, &tkf.PortfolioRequest{
		AccountId: s.accountId,
	})
	if err != nil {
		return nil, err
	}

	var positions []*domain.PortfolioPosition
	for _, v := range resp.GetPositions() {
		positions = append(positions, service.ConvPortfolioPosition(v))
	}

	return &domain.Portfolio{
		TotalAmountShares:     service.ConvMoneyValueFromTkf(resp.TotalAmountShares),
		TotalAmountBonds:      service.ConvMoneyValueFromTkf(resp.TotalAmountBonds),
		TotalAmountEtf:        service.ConvMoneyValueFromTkf(resp.TotalAmountEtf),
		TotalAmountCurrencies: service.ConvMoneyValueFromTkf(resp.TotalAmountCurrencies),
		TotalAmountFutures:    service.ConvMoneyValueFromTkf(resp.TotalAmountFutures),
		ExpectedYield:         service.ConvQuotationFromTkf(resp.ExpectedYield),
		Positions:             positions,
	}, nil
}

// Метод получения списка операций по счёту
func (s *operationsService) Operations(ctx context.Context, from, to *time.Time, state domain.OperationState, figi string) ([]*domain.Operation, error) {
	if s.accountId == "" {
		return nil, tinkoffbroker.ErrArgCandleUnspecified
	}

	resp, err := s.client.GetOperations(ctx, &tkf.OperationsRequest{
		AccountId: s.accountId,
		From:      timestamppb.New(*from),
		To:        timestamppb.New(*to),
		State:     tkf.OperationState(state),
		Figi:      figi,
	})
	if err != nil {
		return nil, err
	}

	var operations []*domain.Operation
	for _, v := range resp.GetOperations() {
		operations = append(operations, service.ConvOperation(v))
	}
	return operations, nil
}

// Метод получения списка позиций по счёту
func (s *operationsService) Positions(ctx context.Context) (*domain.Positions, error) {
	if s.accountId == "" {
		return nil, tinkoffbroker.ErrArgCandleUnspecified
	}

	resp, err := s.client.GetPositions(ctx, &tkf.PositionsRequest{
		AccountId: s.accountId,
	})
	if err != nil {
		return nil, err
	}

	var money []*domain.MoneyValue
	for _, v := range resp.GetMoney() {
		money = append(money, service.ConvMoneyValueFromTkf(v))
	}

	var blocked []*domain.MoneyValue
	for _, v := range resp.GetBlocked() {
		blocked = append(blocked, service.ConvMoneyValueFromTkf(v))
	}

	var securities []*domain.PositionInstrument
	for _, v := range resp.GetSecurities() {
		securities = append(securities, &domain.PositionInstrument{
			Figi:    v.GetFigi(),
			Blocked: v.GetBlocked(),
			Balance: v.GetBalance(),
		})
	}

	var futures []*domain.PositionInstrument
	for _, v := range resp.GetFutures() {
		futures = append(futures, &domain.PositionInstrument{
			Figi:    v.GetFigi(),
			Blocked: v.GetBlocked(),
			Balance: v.GetBalance(),
		})
	}
	return &domain.Positions{
		Money:                   money,
		Blocked:                 blocked,
		Securities:              securities,
		LimitsLoadingInProgress: resp.GetLimitsLoadingInProgress(),
		Futures:                 futures,
	}, nil
}
