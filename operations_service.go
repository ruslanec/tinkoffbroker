/* Сервис получения информации о портфеле по конкретному счёту*/
package tinkoffbroker

import (
	"context"
	"time"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type operationsService struct {
	conn   *grpc.ClientConn
	client tkf.OperationsServiceClient
}

// Конструктор сервиса
func NewOperationsService(conn *grpc.ClientConn) OperationsService {
	client := tkf.NewOperationsServiceClient(conn)

	return &operationsService{
		conn:   conn,
		client: client,
	}
}

// Метод получения портфеля по счёту
func (s *operationsService) Portfolio(ctx context.Context, accountId string) (*Portfolio, error) {
	if accountId == "" {
		return nil, ErrCandleInterval
	}
	resp, err := s.client.GetPortfolio(ctx, &tkf.PortfolioRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}
	var positions []*PortfolioPosition
	for _, v := range resp.GetPositions() {
		positions = append(positions, convPortfolioPosition(v))
	}

	return &Portfolio{
		TotalAmountShares:     convMoneyValue(resp.TotalAmountShares),
		TotalAmountBonds:      convMoneyValue(resp.TotalAmountBonds),
		TotalAmountEtf:        convMoneyValue(resp.TotalAmountEtf),
		TotalAmountCurrencies: convMoneyValue(resp.TotalAmountCurrencies),
		TotalAmountFutures:    convMoneyValue(resp.TotalAmountFutures),
		ExpectedYield:         convQuotation(resp.ExpectedYield),
		Positions:             positions,
	}, nil
}

// Метод получения списка операций по счёту
func (s *operationsService) Operations(ctx context.Context, accountId string, from, to *time.Time, state OperationState, figi string) ([]*Operation, error) {
	if accountId == "" {
		return nil, ErrCandleInterval
	}
	resp, err := s.client.GetOperations(ctx, &tkf.OperationsRequest{
		AccountId: accountId,
		From:      timestamppb.New(*from),
		To:        timestamppb.New(*to),
		State:     tkf.OperationState(state),
		Figi:      figi,
	})
	if err != nil {
		return nil, err
	}

	var operations []*Operation
	for _, v := range resp.GetOperations() {
		operations = append(operations, convOperation(v))
	}
	return operations, nil
}

// Метод получения списка позиций по счёту
func (s *operationsService) Positions(ctx context.Context, accountId string) (*Positions, error) {
	resp, err := s.client.GetPositions(ctx, &tkf.PositionsRequest{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}

	var money []*MoneyValue
	for _, v := range resp.GetMoney() {
		money = append(money, &MoneyValue{
			Currency: v.GetCurrency(),
			Units:    v.GetUnits(),
			Nano:     v.GetNano(),
		})
	}

	var blocked []*MoneyValue
	for _, v := range resp.GetBlocked() {
		blocked = append(blocked, &MoneyValue{
			Currency: v.GetCurrency(),
			Units:    v.GetUnits(),
			Nano:     v.GetNano(),
		})
	}

	var securities []*PositionInstrument
	for _, v := range resp.GetSecurities() {
		securities = append(securities, &PositionInstrument{
			Figi:    v.GetFigi(),
			Blocked: v.GetBlocked(),
			Balance: v.GetBalance(),
		})
	}

	var futures []*PositionInstrument
	for _, v := range resp.GetFutures() {
		futures = append(futures, &PositionInstrument{
			Figi:    v.GetFigi(),
			Blocked: v.GetBlocked(),
			Balance: v.GetBalance(),
		})
	}
	return &Positions{
		Money:                   money,
		Blocked:                 blocked,
		Securities:              securities,
		LimitsLoadingInProgress: resp.GetLimitsLoadingInProgress(),
		Futures:                 futures,
	}, nil
}
