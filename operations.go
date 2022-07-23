package tinkoffbroker

import (
	"context"
	"time"
)

// Информация о портфеле по конкретному счёту
type Operations interface {
	OperationsService
}

// Сервис получения информации о портфеле по конкретному счёту
type OperationsService interface {
	// Метод получения портфеля по счёту
	Portfolio(ctx context.Context, accountId string) (*Portfolio, error)
	// Метод получения списка операций по счёту
	Operations(ctx context.Context, accountId string, from, to *time.Time, state OperationState, figi string) ([]*Operation, error)
	// Метод получения списка позиций по счёту
	Positions(ctx context.Context, accountId string) (*Positions, error)
}
