package tinkoffbroker

import (
	"golang.org/x/net/context"
)

// Информация об аккаунтах
type Users interface {
	UsersService
}

// Сервис предоставления информации об аккаунтах
type UsersService interface {
	// Метод получения открытых и активных счетов пользователя
	Accounts(ctx context.Context) ([]*Account, error)
	// Запрос тарифных лимитов пользователя
	UserTariff(ctx context.Context) (*UserTariff, error)
	// Расчёт маржинальных показателей по счёту
	MarginAttributes(ctx context.Context, accountID string) (*MarginAttributes, error)
}
