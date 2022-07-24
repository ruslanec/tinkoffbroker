package service

import (
	"context"

	"github.com/ruslanec/tinkoffbroker/domain"
)

// Управление стоп-заявками
type StopOrders interface {
	StopOrdersService
}

// Сервис стоп-заявок
type StopOrdersService interface {
	StopOrderTakeProfit(ctx context.Context, figi string, quantity int64, price, stopprice domain.Quotation)
	StopOrderStopLoss(ctx context.Context, figi string, quantity int64, price, stopprice domain.Quotation)
	StopOrderSellTakeProfit(ctx context.Context, figi string, quantity int64, price, stopprice domain.Quotation)
	StopOrderSellTillDate(ctx context.Context, figi string, quantity int64, price, stopprice domain.Quotation)
	StopOrders(ctx context.Context)
	CancelStopOrder(ctx context.Context)
}
