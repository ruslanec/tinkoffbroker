package marketdatastream

import (
	"context"
	"reflect"
	"testing"

	"github.com/ruslanec/tinkoffbroker/domain"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

func TestNewMarketDataStreamService(t *testing.T) {
	type args struct {
		conn *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want service.MarketDataStreamService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarketDataStreamService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarketDataStreamService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketDataStreamService_SubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*domain.CandleInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.SubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_UnsubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*domain.CandleInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.UnsubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_SubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*domain.OrderBookInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.SubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_UnsubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*domain.OrderBookInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.UnsubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_SubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*domain.TradeInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.SubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_UnsubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*domain.TradeInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.UnsubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_SubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*domain.InfoInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.SubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_UnsubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*domain.InfoInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.UnsubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_SubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*domain.LastPriceInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.SubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.SubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_UnsubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*domain.LastPriceInstrument
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnsubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("marketDataStreamService.UnsubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marketDataStreamService_Recv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *marketDataStreamService
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
				t.Errorf("marketDataStreamService.Recv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketDataStreamService.Recv() = %v, want %v", got, tt.want)
			}
		})
	}
}
