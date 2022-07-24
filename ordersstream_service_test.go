package tinkoffbroker

import (
	"context"
	"reflect"
	"testing"

	"google.golang.org/grpc"
)

func TestNewOrdersStreamService(t *testing.T) {
	type args struct {
		conn *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want OrdersStreamService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrdersStreamService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrdersStreamService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ordersStreamService_SubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ordersStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ordersStreamService.SubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ordersStreamService_UnsubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ordersStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ordersStreamService.UnsubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ordersStreamService_Recv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *ordersStreamService
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Recv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ordersStreamService.Recv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ordersStreamService.Recv() = %v, want %v", got, tt.want)
			}
		})
	}
}
