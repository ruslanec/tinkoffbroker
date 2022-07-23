package service

import (
	"reflect"
	"testing"

	tkf "github.com/ruslanec/tinkoffbroker/service/proto"

	"github.com/stretchr/testify/assert"
)

func Test_ordersService_addOrder(t *testing.T) {
	service := NewOrdersService(nil, "111111")
	type args struct {
		order interface{}
	}
	tests := []struct {
		name  string
		s     *ordersService
		args  args
		id    string
		order interface{}
		want  interface{}
	}{
		{
			name: "add PostOrder",
			s:    service.(*ordersService),
			id:   "id1",
			args: args{
				order: &tkf.PostOrderResponse{OrderId: "id1"},
			},
			want: &tkf.PostOrderResponse{},
		},
		{
			name: "add OrderState",
			s:    service.(*ordersService),
			id:   "id2",
			args: args{
				order: &tkf.OrderState{OrderId: "id2"},
			},
			want: &tkf.OrderState{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.addOrder(tt.args.order)
		})
		order := tt.s.orders[tt.id]
		assert.NotNil(t, order)
		assert.IsType(t, tt.want, order)
	}
}

func Test_ordersService_order(t *testing.T) {
	service := NewOrdersService(nil, "111111")
	type args struct {
		orderId string
	}
	tests := []struct {
		name string
		s    *ordersService
		args args
		want interface{}
	}{
		{
			name: "get by not exist orderId",
			s:    service.(*ordersService),
			args: args{
				orderId: "unknown",
			},
			want: nil,
		},
		{
			name: "get by exist orderId",
			s:    service.(*ordersService),
			args: args{
				orderId: "id1",
			},
			want: &tkf.PostOrderResponse{OrderId: "id1"},
		},
	}
	for _, tt := range tests {
		tt.s.orders["id1"] = &tkf.PostOrderResponse{OrderId: "id1"}
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.order(tt.args.orderId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersService_removeOrder(t *testing.T) {
	service := NewOrdersService(nil, "111111")
	type args struct {
		orderId string
	}
	tests := []struct {
		name string
		s    *ordersService
		args args
		want int
	}{
		{
			name: "remove exist orderId",
			s:    service.(*ordersService),
			args: args{
				orderId: "id1",
			},
			want: 0,
		},
		{
			name: "remove not exist orderId",
			s:    service.(*ordersService),
			args: args{
				orderId: "id2",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt.s.orders["id1"] = &tkf.PostOrderResponse{OrderId: "id1"}
		t.Run(tt.name, func(t *testing.T) {
			tt.s.removeOrder(tt.args.orderId)
		})
		assert.Equal(t, tt.want, len(tt.s.orders))
	}
}

func Test_ordersService_orderIdList(t *testing.T) {
	service := NewOrdersService(nil, "111111")
	tests := []struct {
		name string
		s    *ordersService
		want []string
	}{
		{
			name: "remove not exist orderId",
			s:    service.(*ordersService),
			want: []string{"id1", "id2"},
		},
	}
	for _, tt := range tests {
		tt.s.orders["id1"] = &tkf.PostOrderResponse{OrderId: "id1"}
		tt.s.orders["id2"] = &tkf.PostOrderResponse{OrderId: "id2"}
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.orderIdList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersService.orderIdList() = %v, want %v", got, tt.want)
			}
		})
	}
}
