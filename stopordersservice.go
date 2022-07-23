package tinkoffbroker

import (
	"context"
)

// Управление стоп-заявками
type StopOrders interface {
	StopOrdersService
}

// Сервис стоп-заявок
type StopOrdersService interface {
	StopOrderTakeProfit(ctx context.Context, figi string, quantity int64, price, stopprice Quotation)
	StopOrderStopLoss(ctx context.Context, figi string, quantity int64, price, stopprice Quotation)
	StopOrderSellTakeProfit(ctx context.Context, figi string, quantity int64, price, stopprice Quotation)
	StopOrderSellTillDate(ctx context.Context, figi string, quantity int64, price, stopprice Quotation)
	StopOrders(ctx context.Context)
	CancelStopOrder(ctx context.Context)
}
