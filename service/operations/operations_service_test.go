/* Сервис получения информации о портфеле по конкретному счёту*/

package operations

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
)

func TestNewOperationsService(t *testing.T) {
	type args struct {
		conn      *grpc.ClientConn
		accountId string
	}
	tests := []struct {
		name string
		args args
		want service.OperationsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperationsService(tt.args.conn, tt.args.accountId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperationsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operationsService_Portfolio(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *operationsService
		args    args
		want    *domain.Portfolio
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Portfolio(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("operationsService.Portfolio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operationsService.Portfolio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operationsService_Operations(t *testing.T) {
	type args struct {
		ctx   context.Context
		from  *time.Time
		to    *time.Time
		state domain.OperationState
		figi  string
	}
	tests := []struct {
		name    string
		s       *operationsService
		args    args
		want    []*domain.Operation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Operations(tt.args.ctx, tt.args.from, tt.args.to, tt.args.state, tt.args.figi)
			if (err != nil) != tt.wantErr {
				t.Errorf("operationsService.Operations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operationsService.Operations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operationsService_Positions(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *operationsService
		args    args
		want    *domain.Positions
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Positions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("operationsService.Positions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("operationsService.Positions() = %v, want %v", got, tt.want)
			}
		})
	}
}
