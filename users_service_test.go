/*
UsersService
Сервис предназначен для получения:
1. списка счетов пользователя;
2. маржинальных показателей по счёту.

Методы сервиса:
1. GetAccounts - Метод получения счетов пользователя.
2. GetMarginAttributes - Расчёт маржинальных показателей по счёту.
3. GetUserTariff - Запрос тарифа пользователя.
4. GetInfo - Метод получения информации о пользователе.
*/

package tinkoffbroker

import (
	"context"
	"reflect"
	"testing"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"google.golang.org/grpc"
)

func TestNewUsersService(t *testing.T) {
	type args struct {
		conn *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want UsersService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersService(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Accounts(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *usersService
		args    args
		want    []*Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Accounts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.Accounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_UserTariff(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *usersService
		args    args
		want    *UserTariff
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UserTariff(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.UserTariff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.UserTariff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_MarginAttributes(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		s       *usersService
		args    args
		want    *MarginAttributes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarginAttributes(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.MarginAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.MarginAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersService_Info(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       *usersService
		args    args
		want    *tkf.GetInfoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Info(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersService.Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersService.Info() = %v, want %v", got, tt.want)
			}
		})
	}
}
