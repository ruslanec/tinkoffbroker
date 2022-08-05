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
	Portfolio(ctx context.Context) (*domain.Portfolio, error)
	// Метод получения списка операций по счёту
	Operations(ctx context.Context, from, to *time.Time, state domain.OperationState, figi string) ([]*domain.Operation, error)
	// Метод получения списка позиций по счёту
	Positions(ctx context.Context) (*domain.Positions, error)
	//Метод получения доступного остатка для вывода средств
	//GetWithdrawLimits //TODO Not implemented
	//Метод получения брокерского отчёта.
	//GetBrokerReport //TODO Not implemented
	//Метод получения отчёта "Справка о доходах за пределами РФ".
	//GetDividendsForeignIssuer //TODO Not implemented
	//Метод получения списка операций по счёту с пагинацией.
	//GetOperationsByCursor //TODO Not implemented
}
