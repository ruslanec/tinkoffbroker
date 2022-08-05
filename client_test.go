package tinkoffbroker

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"google.golang.org/grpc"
)

func TestNewClient(t *testing.T) {
	type args struct {
		conn      *grpc.ClientConn
		accountId string
		opts      []Option
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.conn, tt.args.accountId, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Init(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("client.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Recv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Recv(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Recv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Recv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Run(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("client.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Close(t *testing.T) {
	tests := []struct {
		name string
		c    *client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Close()
		})
	}
}

func Test_client_Shares(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Shares(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Shares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Shares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_ShareByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ShareByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.ShareByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.ShareByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Bonds(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantShares []*domain.Bond
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShares, err := tt.c.Bonds(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Bonds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShares, tt.wantShares) {
				t.Errorf("client.Bonds() = %v, want %v", gotShares, tt.wantShares)
			}
		})
	}
}

func Test_client_BondByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Bond
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BondByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.BondByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.BondByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_BondCoupons(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Coupon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BondCoupons(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.BondCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.BondCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_AccruedInterests(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.AccruedInterest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.AccruedInterests(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AccruedInterests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.AccruedInterests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Currencies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantShares []*domain.Currency
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShares, err := tt.c.Currencies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Currencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShares, tt.wantShares) {
				t.Errorf("client.Currencies() = %v, want %v", gotShares, tt.wantShares)
			}
		})
	}
}

func Test_client_CurrencyByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CurrencyByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CurrencyByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CurrencyByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Etfs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Etfs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Etfs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Etfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_EtfByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EtfByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.EtfByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.EtfByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Future(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Future(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Future() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Future() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_FutureByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FutureByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.FutureByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.FutureByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_TradingSchedules(t *testing.T) {
	type args struct {
		ctx      context.Context
		exchange string
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.TradingSchedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.TradingSchedules(tt.args.ctx, tt.args.exchange, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.TradingSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.TradingSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_FuturesMargin(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.FuturesMargin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FuturesMargin(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.FuturesMargin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.FuturesMargin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_InstrumentByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Instrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.InstrumentByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.InstrumentByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.InstrumentByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Dividends(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Dividend
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Dividends(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Dividends() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Dividends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_AssetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.AssetFull
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.AssetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AssetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.AssetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Assets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Assets(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Assets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Assets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Favorites(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Favorites(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Favorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Favorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_EditFavorites(t *testing.T) {
	type args struct {
		ctx    context.Context
		figies []string
		action domain.EditFavoritesActionType
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditFavorites(tt.args.ctx, tt.args.figies, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.EditFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.EditFavorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Countries(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Country
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Countries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Countries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Countries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_FindInstrument(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.InstrumentShort
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FindInstrument(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.FindInstrument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.FindInstrument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Brands(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Brands(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Brands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Brands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_BrandById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BrandById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.BrandById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.BrandById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_TradingStatus(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.InstrumentTradingStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.TradingStatus(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.TradingStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.TradingStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_LastPrices(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi []string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.LastPrice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.LastPrices(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.LastPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.LastPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Candles(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		from     time.Time
		to       time.Time
		interval domain.CandleInterval
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Candle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Candles(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Candles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Candles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderBook(t *testing.T) {
	type args struct {
		ctx   context.Context
		figi  string
		depth int32
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.OrderBook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBook(tt.args.ctx, tt.args.figi, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderBuyLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBuyLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderBuyLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderBuyLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderSellLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderSellLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderSellLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderSellLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderBuyMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBuyMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderBuyMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderBuyMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderSellMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *domain.Quotation
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderSellMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderSellMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderSellMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CancelOrder(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CancelOrder(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_OrderState(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderState(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.OrderState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.OrderState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Orders(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Orders(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Orders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Orders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Portfolio(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Portfolio
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Portfolio(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Portfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Portfolio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Operations(t *testing.T) {
	type args struct {
		ctx   context.Context
		from  *time.Time
		to    *time.Time
		state domain.OperationState
		figi  string
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Operation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Operations(tt.args.ctx, tt.args.from, tt.args.to, tt.args.state, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Operations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Operations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Positions(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.Positions
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Positions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Positions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Positions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Accounts(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    []*domain.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Accounts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.Accounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UserTariff(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.UserTariff
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UserTariff(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.UserTariff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.UserTariff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_MarginAttributes(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		want    *domain.MarginAttributes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarginAttributes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.MarginAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.MarginAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*domain.CandleInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*domain.CandleInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_SubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*domain.OrderBookInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*domain.OrderBookInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_SubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*domain.TradeInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*domain.TradeInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_SubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*domain.InfoInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*domain.InfoInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_SubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*domain.LastPriceInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("client.SubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UnsubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*domain.LastPriceInstrument
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("client.UnsubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
