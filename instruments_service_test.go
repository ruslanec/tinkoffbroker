package tinkoffbroker

import (
	"context"
	"reflect"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestNewInstrumentsService(t *testing.T) {
	type args struct {
		conn *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want InstrumentsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInstrumentsService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInstrumentsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_TradingSchedules(t *testing.T) {
	type args struct {
		ctx      context.Context
		exchange string
		from     time.Time
		to       time.Time
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*TradingSchedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.TradingSchedules(tt.args.ctx, tt.args.exchange, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.TradingSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.TradingSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_BondByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Bond
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.BondByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.BondByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.BondByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Bonds(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Bond
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Bonds(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Bonds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Bonds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_BondCoupons(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Coupon
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.BondCoupons(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.BondCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.BondCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_CurrencyByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CurrencyByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.CurrencyByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.CurrencyByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Currencies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Currency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Currencies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Currencies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Currencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_EtfByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.EtfByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.EtfByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.EtfByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Etfs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Etf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Etfs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Etfs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Etfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_ShareByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ShareByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.ShareByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.ShareByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Shares(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Share
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Shares(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Shares() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Shares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_FutureByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FutureByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.FutureByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.FutureByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Future(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Future
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Future(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Future() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Future() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_AccruedInterests(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*AccruedInterest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AccruedInterests(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.AccruedInterests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.AccruedInterests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_FuturesMargin(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *FuturesMargin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FuturesMargin(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.FuturesMargin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.FuturesMargin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_InstrumentByFigi(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Instrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.InstrumentByFigi(tt.args.ctx, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.InstrumentByFigi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.InstrumentByFigi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Dividends(t *testing.T) {
	type args struct {
		ctx  context.Context
		figi string
		from *time.Time
		to   *time.Time
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Dividend
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Dividends(tt.args.ctx, tt.args.figi, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Dividends() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Dividends() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_AssetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *AssetFull
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AssetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.AssetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.AssetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Assets(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Asset
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Assets(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Assets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Assets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Favorites(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Favorites(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Favorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Favorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_EditFavorites(t *testing.T) {
	type args struct {
		ctx    context.Context
		figies []string
		action EditFavoritesActionType
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*FavoriteInstrument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.EditFavorites(tt.args.ctx, tt.args.figies, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.EditFavorites() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.EditFavorites() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Countries(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Country
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Countries(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Countries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Countries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_FindInstrument(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*InstrumentShort
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindInstrument(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.FindInstrument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.FindInstrument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_Brands(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    []*Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Brands(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.Brands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.Brands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instrumentsService_BrandById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		s       *instrumentsService
		args    args
		want    *Brand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.BrandById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("instrumentsService.BrandById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instrumentsService.BrandById() = %v, want %v", got, tt.want)
			}
		})
	}
}
