package service

import (
	"context"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
)

// Информация о портфеле по конкретному счёту
type Operations interface {
	OperationsService
}

// Сервис получения информации о портфеле по конкретному счёту
type OperationsService interface {
	// Метод получения портфеля по счёту
	Portfolio(ctx context.Context, accountId string) (*domain.Portfolio, error)
	// Метод получения списка операций по счёту
	Operations(ctx context.Context, accountId string, from, to *time.Time, state domain.OperationState, figi string) ([]*domain.Operation, error)
	// Метод получения списка позиций по счёту
	Positions(ctx context.Context, accountId string) (*domain.Positions, error)
}
