package tinkoffbroker

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"google.golang.org/grpc"
)

var _ OrdersService = (*ordersService)(nil)

// Сервис предназначен для работы с торговыми поручениями
type ordersService struct {
	conn      *grpc.ClientConn
	accountID string
	client    tkf.OrdersServiceClient
	orders    map[string]interface{}
	mu        sync.Mutex
}

func NewOrdersService(conn *grpc.ClientConn, accountID string) OrdersService {
	ordersServiceClient := tkf.NewOrdersServiceClient(conn)

	return &ordersService{
		conn:      conn,
		accountID: accountID,
		client:    ordersServiceClient,
		orders:    make(map[string]interface{}),
	}
}

// Добавление заявки в список заявок сервиса
func (s *ordersService) addOrder(order interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch order := order.(type) {
	case *tkf.PostOrderResponse:
		s.orders[order.OrderId] = order
	case *tkf.OrderState:
		s.orders[order.OrderId] = order
	}
}

// Получение заявки из списка заявок сервиса
func (s *ordersService) order(orderId string) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, found := s.orders[orderId]
	if !found {
		return nil
	}
	return order
}

// Удаление заявки из списка заявок сервиса
func (s *ordersService) removeOrder(orderId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.orders, orderId)
}

// Получение списка orderId заявок из списка сервиса
func (s *ordersService) orderIdList() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var orderIdList []string
	for k := range s.orders {
		orderIdList = append(orderIdList, k)
	}

	return orderIdList
}

// Метод выставления рыночной заявки на покупку
func (s *ordersService) OrderBuyLimit(ctx context.Context, figi string, quantity int64, price *Quotation) (*PostOrderResponse, error) {
	const (
		ORDER_DIRECTION = tkf.OrderDirection_ORDER_DIRECTION_BUY
		ORDER_TYPE      = tkf.OrderType_ORDER_TYPE_LIMIT
	)

	return s.postOrder(ctx, figi, quantity, price, ORDER_DIRECTION, ORDER_TYPE)
}

// Метод выставления лимитной заявки на продажу
func (s *ordersService) OrderSellLimit(ctx context.Context, figi string, quantity int64, price *Quotation) (*PostOrderResponse, error) {
	const (
		ORDER_DIRECTION = tkf.OrderDirection_ORDER_DIRECTION_SELL
		ORDER_TYPE      = tkf.OrderType_ORDER_TYPE_LIMIT
	)

	return s.postOrder(ctx, figi, quantity, price, ORDER_DIRECTION, ORDER_TYPE)
}

// Метод выставления рыночной заявки на покупку
func (s *ordersService) OrderBuyMarket(ctx context.Context, figi string, quantity int64, price *Quotation) (*PostOrderResponse, error) {
	const (
		ORDER_DIRECTION = tkf.OrderDirection_ORDER_DIRECTION_BUY
		ORDER_TYPE      = tkf.OrderType_ORDER_TYPE_MARKET
	)

	return s.postOrder(ctx, figi, quantity, price, ORDER_DIRECTION, ORDER_TYPE)
}

// Метод выставления рыночной заявки на продажу
func (s *ordersService) OrderSellMarket(ctx context.Context, figi string, quantity int64, price *Quotation) (*PostOrderResponse, error) {
	const (
		ORDER_DIRECTION = tkf.OrderDirection_ORDER_DIRECTION_SELL
		ORDER_TYPE      = tkf.OrderType_ORDER_TYPE_MARKET
	)

	return s.postOrder(ctx, figi, quantity, price, ORDER_DIRECTION, ORDER_TYPE)
}

// Метод отмены биржевой заявки
func (s *ordersService) CancelOrder(ctx context.Context, orderId string) (*time.Time, error) {
	resp, err := s.client.CancelOrder(ctx, &tkf.CancelOrderRequest{
		AccountId: s.accountID,
		OrderId:   orderId,
	})
	if err != nil {
		return nil, err
	}
	s.removeOrder(orderId)
	t := resp.GetTime().AsTime()
	return &t, nil
}

// Метод получения статуса торгового поручения
func (s *ordersService) OrderState(ctx context.Context, orderId string) (*OrderState, error) {
	resp, err := s.client.GetOrderState(ctx, &tkf.GetOrderStateRequest{
		AccountId: s.accountID,
		OrderId:   orderId,
	})
	if err != nil {
		return nil, err
	}

	var stages []*OrderStage
	for _, v := range resp.GetStages() {
		stages = append(stages, &OrderStage{
			Price:    convMoneyValue(v.GetPrice()),
			Quantity: v.GetQuantity(),
			TradeId:  v.GetTradeId(),
		})
	}

	orderState := convOrderState(resp)
	orderState.Stages = stages
	return orderState, nil
}

// Метод получения списка активных заявок по счёту
func (s *ordersService) Orders(ctx context.Context) ([]*OrderState, error) {
	resp, err := s.client.GetOrders(ctx, &tkf.GetOrdersRequest{
		AccountId: s.accountID,
	})
	if err != nil {
		return nil, err
	}

	var orders []*OrderState
	for _, v := range resp.GetOrders() {
		orders = append(orders, convOrderState(v))
	}
	// var orders []*stockbroker.OrderState
	// var order *stockbroker.OrderState
	// orderIdList := s.orderIdList() // Получаем перечень orderId из списка заявок сервиса
	// for _, v := range resp.GetOrders() {
	// 	order = &stockbroker.OrderState(*v)
	// 	orders = append(orders, order)

	// 	// Синхронизация списка заявок сервиса с сервисом Тинькофф (активные заявки)
	// 	s.addOrder(order)             // Обновляем записи в списке заявок сервиса
	// 	func(s []string, id string) { // Удаляем orderId из перечня
	// 		for i, v := range s {
	// 			if v == id {
	// 				s = append(s[:i], s[i+1:]...)
	// 			}
	// 		}
	// 	}(orderIdList, v.OrderId)
	// }

	// for _, v := range orderIdList { // Удаляем заявки из списка заявок сервиса, отсутствующие в списке активных заявок сервиса Тинькофф
	// 	delete(s.orders, v)
	// }

	return orders, nil
}

// Метод выставления заявки
func (s *ordersService) postOrder(ctx context.Context,
	figi string,
	quantity int64,
	price *Quotation,
	orderdirection tkf.OrderDirection,
	ordertype tkf.OrderType) (*PostOrderResponse, error) {
	id := uuid.New()
	p := tkf.Quotation{
		Units: price.Units,
		Nano:  price.Nano,
	}
	resp, err := s.client.PostOrder(ctx, &tkf.PostOrderRequest{
		Figi:      figi,
		Quantity:  quantity,
		Price:     &p,
		Direction: orderdirection,
		AccountId: s.accountID,
		OrderType: ordertype,
		OrderId:   id.String(),
	})
	if err != nil {
		return nil, err
	}

	return convPostOrderResponse(resp), nil
}
