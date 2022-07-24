package tinkoffbroker

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestWithRepository(t *testing.T) {
	type args struct {
		s repository.Repository
	}
	tests := []struct {
		name string
		args args
		want BackTestOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRepository(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithPeriod(t *testing.T) {
	type args struct {
		p TimePeriod
	}
	tests := []struct {
		name string
		args args
		want BackTestOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPeriod(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBackTest(t *testing.T) {
	type args struct {
		period      *TimePeriod
		instruments []string
		broker      Client
		opts        []BackTestOption
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
			if got := NewBackTest(tt.args.period, tt.args.instruments, tt.args.broker, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBackTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Init(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_Run(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_Close(t *testing.T) {
	tests := []struct {
		name string
		c    *backtestClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Close()
		})
	}
}

func Test_backtestClient_Shares(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Shares(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Shares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Shares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_ShareByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ShareByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.ShareByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.ShareByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Bonds(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		c          *backtestClient
		args       args
		wantShares []*Bond
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShares, err := tt.c.Bonds(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Bonds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShares, tt.wantShares) {
				t.Errorf("backtestClient.Bonds() = %v, want %v", gotShares, tt.wantShares)
			}
		})
	}
}

func Test_backtestClient_BondByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Bond
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BondByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.BondByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.BondByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_BondCoupons(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Coupon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BondCoupons(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.BondCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.BondCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_AccruedInterests(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*AccruedInterest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.AccruedInterests(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.AccruedInterests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.AccruedInterests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Currencies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		c          *backtestClient
		args       args
		wantShares []*Currency
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShares, err := tt.c.Currencies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Currencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShares, tt.wantShares) {
				t.Errorf("backtestClient.Currencies() = %v, want %v", gotShares, tt.wantShares)
			}
		})
	}
}

func Test_backtestClient_CurrencyByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CurrencyByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.CurrencyByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.CurrencyByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Etfs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Etfs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Etfs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Etfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_EtfByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EtfByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.EtfByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.EtfByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Future(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Future(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Future() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Future() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_FutureByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FutureByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.FutureByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.FutureByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_TradingSchedules(t *testing.T) {
	type args struct {
		ctx      context.Context
		exchange string
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*TradingSchedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.TradingSchedules(tt.args.ctx, tt.args.exchange, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.TradingSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.TradingSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_FuturesMargin(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *FuturesMargin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FuturesMargin(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.FuturesMargin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.FuturesMargin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_InstrumentByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Instrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.InstrumentByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.InstrumentByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.InstrumentByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Dividends(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Dividend
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Dividends(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Dividends() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Dividends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_AssetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *AssetFull
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.AssetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.AssetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.AssetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Assets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Assets(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Assets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Assets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Favorites(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Favorites(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Favorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Favorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_EditFavorites(t *testing.T) {
	type args struct {
		ctx    context.Context
		figies []string
		action EditFavoritesActionType
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditFavorites(tt.args.ctx, tt.args.figies, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.EditFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.EditFavorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Countries(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Country
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Countries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Countries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Countries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_FindInstrument(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*InstrumentShort
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FindInstrument(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.FindInstrument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.FindInstrument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Brands(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Brands(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Brands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Brands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_BrandById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.BrandById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.BrandById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.BrandById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_TradingStatus(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *InstrumentTradingStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.TradingStatus(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.TradingStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.TradingStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_LastPrices(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi []string
	}
	tests := []struct {
		name    string
		с       *backtestClient
		args    args
		want    []*LastPrice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.с.LastPrices(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.LastPrices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.LastPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Candles(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		from     time.Time
		to       time.Time
		interval CandleInterval
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Candle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Candles(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Candles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Candles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderBook(t *testing.T) {
	type args struct {
		ctx   context.Context
		figi  string
		depth int32
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *OrderBook
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBook(tt.args.ctx, tt.args.figi, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderBuyLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *Quotation
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBuyLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderBuyLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderBuyLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderSellLimit(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *Quotation
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderSellLimit(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderSellLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderSellLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderBuyMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *Quotation
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderBuyMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderBuyMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderBuyMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderSellMarket(t *testing.T) {
	type args struct {
		ctx      context.Context
		figi     string
		quantity int64
		price    *Quotation
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *PostOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderSellMarket(tt.args.ctx, tt.args.figi, tt.args.quantity, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderSellMarket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderSellMarket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_CancelOrder(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		c       *backtestClient
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
				t.Errorf("backtestClient.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_OrderState(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderId string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.OrderState(tt.args.ctx, tt.args.orderId)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.OrderState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.OrderState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Orders(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*OrderState
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Orders(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Orders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Orders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_SubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeOrderTrades(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeOrderTrades(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeOrderTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_Portfolio(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Portfolio
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Portfolio(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Portfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Portfolio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Operations(t *testing.T) {
	type args struct {
		ctx   context.Context
		from  *time.Time
		to    *time.Time
		state OperationState
		figi  string
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Operation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Operations(tt.args.ctx, tt.args.from, tt.args.to, tt.args.state, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Operations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Operations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Positions(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *Positions
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Positions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Positions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Positions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_Accounts(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    []*Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Accounts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Accounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_UserTariff(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *UserTariff
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UserTariff(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UserTariff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.UserTariff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_MarginAttributes(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		want    *MarginAttributes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarginAttributes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.MarginAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.MarginAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backtestClient_SubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*CandleInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeCandles(t *testing.T) {
	type args struct {
		ctx     context.Context
		candles []*CandleInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeCandles(tt.args.ctx, tt.args.candles); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeCandles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_SubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*OrderBookInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeOrderBook(t *testing.T) {
	type args struct {
		ctx        context.Context
		orderbooks []*OrderBookInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeOrderBook(tt.args.ctx, tt.args.orderbooks); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeOrderBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_SubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*TradeInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeTrades(t *testing.T) {
	type args struct {
		ctx    context.Context
		trades []*TradeInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeTrades(tt.args.ctx, tt.args.trades); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeTrades() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_SubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*InfoInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeInfo(t *testing.T) {
	type args struct {
		ctx         context.Context
		instruments []*InfoInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeInfo(tt.args.ctx, tt.args.instruments); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_SubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*LastPriceInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.SubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_UnsubscribeLastPrices(t *testing.T) {
	type args struct {
		ctx        context.Context
		lastprices []*LastPriceInstrument
	}
	tests := []struct {
		name    string
		c       *backtestClient
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnsubscribeLastPrices(tt.args.ctx, tt.args.lastprices); (err != nil) != tt.wantErr {
				t.Errorf("backtestClient.UnsubscribeLastPrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backtestClient_Recv(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		c       *backtestClient
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
				t.Errorf("backtestClient.Recv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backtestClient.Recv() = %v, want %v", got, tt.want)
			}
		})
	}
}
