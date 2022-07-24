package service

import (
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"golang.org/x/net/context"
)

// Подача торговых поручений
type Orders interface {
	OrdersService
	OrdersStreamService
}

// Сервис работы с торговыми поручениями
type OrdersService interface {
	// Метод выставления рыночной заявки на покупку
	OrderBuyLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления лимитной заявки на продажу
	OrderSellLimit(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления рыночной заявки на покупку
	OrderBuyMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод выставления рыночной заявки на продажу
	OrderSellMarket(ctx context.Context, figi string, quantity int64, price *domain.Quotation) (*domain.PostOrderResponse, error)
	// Метод отмены биржевой заявки
	CancelOrder(ctx context.Context, orderId string) (*time.Time, error)
	// Метод получения статуса торгового поручения
	OrderState(ctx context.Context, orderId string) (*domain.OrderState, error)
	// Метод получения списка активных заявок по счёту
	Orders(ctx context.Context) ([]*domain.OrderState, error)
}

// Потоковый сервис получения информации о сделках пользователя
type OrdersStreamService interface {
	// Создать подписку на поток сделок пользователя
	SubscribeOrderTrades(ctx context.Context) error
	// Отписаться от потока сделок пользователя
	UnsubscribeOrderTrades(ctx context.Context) error
	// Получение сделок пользователя по подписке
	Recv(ctx context.Context) (interface{}, error)
}
