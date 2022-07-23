/*
UsersService
Сервис предназначен для получения:
1. списка счетов пользователя;
2. маржинальных показателей по счёту.

Методы сервиса:
1. GetAccounts - Метод получения счетов пользователя.
2. GetMarginAttributes - Расчёт маржинальных показателей по счёту.
3. GetUserTariff - Запрос тарифа пользователя.
4. GetInfo - Метод получения информации о пользователе.
*/

package service

import (
	"context"
	"strings"

	domain "github.com/ruslanec/tinkoffbroker"
	tkf "github.com/ruslanec/tinkoffbroker/service/proto"
	"google.golang.org/grpc"
)

type usersService struct {
	conn   *grpc.ClientConn
	client tkf.UsersServiceClient
}

func NewUsersService(conn *grpc.ClientConn) service.UsersService {
	return &usersService{
		conn:   conn,
		client: tkf.NewUsersServiceClient(conn),
	}
}

// Метод получения открытых и активных счетов пользователя
func (s *usersService) Accounts(ctx context.Context) ([]*domain.Account, error) {
	resp, err := s.client.GetAccounts(context.Background(), &tkf.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}

	var accounts []*domain.Account
	for _, v := range resp.GetAccounts() {
		openedDate := v.GetOpenedDate().AsTime()
		closedDate := v.GetClosedDate().AsTime()
		accounts = append(accounts, &domain.Account{
			Id:          v.GetId(),
			Type:        domain.AccountType(v.GetType()),
			Name:        v.GetName(),
			Status:      domain.AccountStatus(v.GetStatus()),
			OpenedDate:  &openedDate,
			ClosedDate:  &closedDate,
			AccessLevel: domain.AccessLevel(v.GetAccessLevel()),
		})
	}
	return accounts, nil
}

// Запрос тарифных лимитов пользователя
func (s *usersService) UserTariff(ctx context.Context) (*domain.UserTariff, error) {
	resp, err := s.client.GetUserTariff(ctx, &tkf.GetUserTariffRequest{})
	if err != nil {
		return nil, err
	}

	unaryLimits := make(map[string]domain.UnaryLimit)
	for _, limit := range resp.GetUnaryLimits() {
		for _, method := range limit.GetMethods() {
			name := strings.Split(method, "/")[1] // TODO: add check
			unaryLimits[name] = domain.UnaryLimit{
				MaxValue: limit.GetLimitPerMinute(),
			}
		}
	}

	var streamLimits []*domain.StreamLimit
	for _, v := range resp.GetStreamLimits() {
		streamLimits = append(streamLimits, &domain.StreamLimit{
			Limit:   v.GetLimit(),
			Streams: v.GetStreams(),
		})
	}
	return &domain.UserTariff{
		UnaryMethodLimitsPerMinute: unaryLimits,
		StreamLimits:               streamLimits,
	}, nil
}

// Расчёт маржинальных показателей по счёту пользователя
func (s *usersService) MarginAttributes(ctx context.Context, accountID string) (*domain.MarginAttributes, error) {
	resp, err := s.client.GetMarginAttributes(ctx, &tkf.GetMarginAttributesRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}
	return &domain.MarginAttributes{
		LiquidPortfolio:       convMoneyValue(resp.GetLiquidPortfolio()),
		StartingMargin:        convMoneyValue(resp.GetStartingMargin()),
		MinimalMargin:         convMoneyValue(resp.GetMinimalMargin()),
		FundsSufficiencyLevel: convQuotation(resp.GetFundsSufficiencyLevel()),
		AmountOfMissingFunds:  convMoneyValue(resp.GetAmountOfMissingFunds()),
	}, nil
}

// Метод получения информации о пользователе
func (s *usersService) Info(ctx context.Context) (*tkf.GetInfoResponse, error) {
	resp, err := s.client.GetInfo(ctx, &tkf.GetInfoRequest{})
	if err != nil {
		return nil, err
	}
	return resp, err
}
