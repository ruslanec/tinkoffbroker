package orders

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

func TestNewOrdersService(t *testing.T) {
	type args struct {
		conn      *grpc.ClientConn
		accountID string
	}
	tests := []struct {
		name string
		args args
		want service.OrdersService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrdersService(tt.args.conn, tt.args.accountID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrdersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_addOrder(t *testing.T) {
	type args struct {
		order interface{}
	}
	tests := []struct {
		name string
		s    *ordersService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.addOrder(tt.args.order)
		})
	}
}

func Test_ordersService_order(t *testing.T) {
	type args struct {
		orderId string
	}
	tests := []struct {
		name string
		s    *ordersService
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.order(tt.args.orderId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_removeOrder(t *testing.T) {
	type args struct {
		orderId string
	}
	tests := []struct {
		name string
		s    *ordersService
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.removeOrder(tt.args.orderId)
		})
	}
}

func Test_ordersService_orderIdList(t *testing.T) {
	tests := []struct {
		name string
		s    *ordersService
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.orderIdList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.orderIdList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_OrderBuyLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderBuyLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.OrderBuyLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.OrderBuyLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_OrderSellLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderSellLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.OrderSellLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.OrderSellLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_OrderBuyMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderBuyMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.OrderBuyMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.OrderBuyMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_OrderSellMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderSellMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.OrderSellMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.OrderSellMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_CancelOrder(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CancelOrder(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_OrderState(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderState(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.OrderState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.OrderState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_Orders(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    []*domain.OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Orders(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.Orders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.Orders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_postOrder(t *testing.T) {
	type args struct {
		ctx            context.Context
		figi           string
		quantity       int64
		price          *domain.Quotation
		orderdirection tkf.OrderDirection
		ordertype      tkf.OrderType
	}
	tests := []struct {
		name    string
		s       *ordersService
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.postOrder(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price, tt.args.orderdirection, tt.args.ordertype)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersService.postOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.postOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
