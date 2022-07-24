package marketdata

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

func TestNewMarketDataService(t *testing.T) {
	type args struct {
		conn *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want service.MarketDataService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMarketDataService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarketDataService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketDataService_LastPrices(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi []string
	}
	tests := []struct {
		name    string
		s       *marketDataService
		args    args
		want    []*domain.LastPrice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.LastPrices(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("marketDataService.LastPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketDataService.LastPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketDataService_Candles(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		from     time.Time
		to       time.Time
		interval domain.CandleInterval
	}
	tests := []struct {
		name    string
		s       *marketDataService
		args    args
		want    []*domain.Candle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Candles(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("marketDataService.Candles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketDataService.Candles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketDataService_OrderBook(t *testing.T) {
	type args struct {
		ctx   context.Context
		figi  string
		depth int32
	}
	tests := []struct {
		name    string
		s       *marketDataService
		args    args
		want    *domain.OrderBook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OrderBook(tt.args.ctx, tt.args.figi, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("marketDataService.OrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketDataService.OrderBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marketDataService_TradingStatus(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *marketDataService
		args    args
		want    *domain.InstrumentTradingStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.TradingStatus(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("marketDataService.TradingStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marketDataService.TradingStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		array []int32
		value int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.array, tt.args.value); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
