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

package users

import (
	"context"
	"strings"

	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

type usersService struct {
	conn      *grpc.ClientConn
	accountId string
	client    tkf.UsersServiceClient
}

func NewUsersService(conn *grpc.ClientConn, accountId string) service.UsersService {
	return &usersService{
		conn:      conn,
		accountId: accountId,
		client:    tkf.NewUsersServiceClient(conn),
	}
}

// Метод получения открытых и активных счетов пользователя
func (s *usersService) Accounts(ctx context.Context) ([]*domain.Account, error) {
	resp, err := s.client.GetAccounts(ctx, &tkf.GetAccountsRequest{})
	if err != nil {
		return nil, err
	}

	tkfAccounts := resp.GetAccounts()
	accounts := make([]*domain.Account, 0, len(tkfAccounts))
	for _, tkfAccount := range tkfAccounts {
		openedDate := service.ConvTimestamp(tkfAccount.OpenedDate)
		closedDate := service.ConvTimestamp(tkfAccount.ClosedDate)

		accounts = append(accounts, &domain.Account{
			Id:          tkfAccount.GetId(),
			Type:        domain.AccountType(tkfAccount.GetType()),
			Name:        tkfAccount.GetName(),
			Status:      domain.AccountStatus(tkfAccount.GetStatus()),
			OpenedDate:  openedDate,
			ClosedDate:  closedDate,
			AccessLevel: domain.AccessLevel(tkfAccount.GetAccessLevel()),
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
func (s *usersService) MarginAttributes(ctx context.Context) (*domain.MarginAttributes, error) {
	resp, err := s.client.GetMarginAttributes(ctx, &tkf.GetMarginAttributesRequest{
		AccountId: s.accountId,
	})
	if err != nil {
		return nil, err
	}

	return &domain.MarginAttributes{
		LiquidPortfolio:       service.ConvMoneyValueFromTkf(resp.GetLiquidPortfolio()),
		StartingMargin:        service.ConvMoneyValueFromTkf(resp.GetStartingMargin()),
		MinimalMargin:         service.ConvMoneyValueFromTkf(resp.GetMinimalMargin()),
		FundsSufficiencyLevel: service.ConvQuotationFromTkf(resp.GetFundsSufficiencyLevel()),
		AmountOfMissingFunds:  service.ConvMoneyValueFromTkf(resp.GetAmountOfMissingFunds()),
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
