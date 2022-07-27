package service

import (
	"github.com/ruslanec/tinkoffbroker/domain"
	"golang.org/x/net/context"
)

// Информация об аккаунтах
type Users interface {
	UsersService
}

// Сервис предоставления информации об аккаунтах
type UsersService interface {
	// Метод получения открытых и активных счетов пользователя
	Accounts(ctx context.Context) ([]*domain.Account, error)
	// Запрос тарифных лимитов пользователя
	UserTariff(ctx context.Context) (*domain.UserTariff, error)
	// Расчёт маржинальных показателей по счёту
	MarginAttributes(ctx context.Context) (*domain.MarginAttributes, error)
}
