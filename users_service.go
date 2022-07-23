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

package tinkoffbroker

import (
	"context"
	"strings"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"google.golang.org/grpc"
)

type usersService struct {
	conn   *grpc.ClientConn
	client tkf.UsersServiceClient
}

func NewUsersService(conn *grpc.ClientConn) UsersService {
	return &usersService{
		conn:   conn,
		client: tkf.NewUsersServiceClient(conn),
	}
}

// Метод получения открытых и активных счетов пользователя
func (s *usersService) Accounts(ctx context.Context) ([]*Account, error) {
	resp, err := s.client.GetAccounts(context.Background(), &tkf.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	for _, v := range resp.GetAccounts() {
		openedDate := v.GetOpenedDate().AsTime()
		closedDate := v.GetClosedDate().AsTime()
		accounts = append(accounts, &Account{
			Id:          v.GetId(),
			Type:        AccountType(v.GetType()),
			Name:        v.GetName(),
			Status:      AccountStatus(v.GetStatus()),
			OpenedDate:  &openedDate,
			ClosedDate:  &closedDate,
			AccessLevel: AccessLevel(v.GetAccessLevel()),
		})
	}
	return accounts, nil
}

// Запрос тарифных лимитов пользователя
func (s *usersService) UserTariff(ctx context.Context) (*UserTariff, error) {
	resp, err := s.client.GetUserTariff(ctx, &tkf.GetUserTariffRequest{})
	if err != nil {
		return nil, err
	}

	unaryLimits := make(map[string]UnaryLimit)
	for _, limit := range resp.GetUnaryLimits() {
		for _, method := range limit.GetMethods() {
			name := strings.Split(method, "/")[1] // TODO: add check
			unaryLimits[name] = UnaryLimit{
				MaxValue: limit.GetLimitPerMinute(),
			}
		}
	}

	var streamLimits []*StreamLimit
	for _, v := range resp.GetStreamLimits() {
		streamLimits = append(streamLimits, &StreamLimit{
			Limit:   v.GetLimit(),
			Streams: v.GetStreams(),
		})
	}
	return &UserTariff{
		UnaryMethodLimitsPerMinute: unaryLimits,
		StreamLimits:               streamLimits,
	}, nil
}

// Расчёт маржинальных показателей по счёту пользователя
func (s *usersService) MarginAttributes(ctx context.Context, accountID string) (*MarginAttributes, error) {
	resp, err := s.client.GetMarginAttributes(ctx, &tkf.GetMarginAttributesRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}
	return &MarginAttributes{
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
