package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Режим торгов инструмента
type SecurityTradingStatus int32

const (
	SecurityTradingStatusUnspecified                  SecurityTradingStatus = 0  // Торговый статус не определён
	SecurityTradingStatusNotAvailableForTrading       SecurityTradingStatus = 1  // Недоступен для торгов
	SecurityTradingStatusOpeningPeriod                SecurityTradingStatus = 2  // Период открытия торгов
	SecurityTradingStatusClosingPeriod                SecurityTradingStatus = 3  // Период закрытия торгов
	SecurityTradingStatusBreakInTrading               SecurityTradingStatus = 4  // Перерыв в торговле
	SecurityTradingStatusNormalTrading                SecurityTradingStatus = 5  // Нормальная торговля
	SecurityTradingStatusClosingAuction               SecurityTradingStatus = 6  // Аукцион закрытия
	SecurityTradingStatusDarkPoolAuction              SecurityTradingStatus = 7  // Аукцион крупных пакетов
	SecurityTradingStatusDiscreteAuction              SecurityTradingStatus = 8  // Дискретный аукцион
	SecurityTradingStatusOpeningAuctionPeriod         SecurityTradingStatus = 9  // Аукцион открытия
	SecurityTradingStatusTradingAtClosingAuctionPrice SecurityTradingStatus = 10 // Период торгов по цене аукциона закрытия
	SecurityTradingStatusSessionAssigned              SecurityTradingStatus = 11 // Сессия назначена
	SecurityTradingStatusSessionClose                 SecurityTradingStatus = 12 // Сессия закрыта
	SecurityTradingStatusSessionOpen                  SecurityTradingStatus = 13 // Сессия открыта
	SecurityTradingStatusDealerNormalTrading          SecurityTradingStatus = 14 // Доступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatusDealerBreakInTrading         SecurityTradingStatus = 15 // Перерыв торговли в режиме внутренней ликвидности брокера
	SecurityTradingStatusDealerNotAvailableForTrading SecurityTradingStatus = 16 // Недоступна торговля в режиме внутренней ликвидности брокера
)

// Котировка - денежная сумма без указания валюты
type Quotation struct {
	Value decimal.Decimal `json:"value,omitempty"`
}

func (s *Quotation) String() string {
	return s.Value.String()
}

func (s *Quotation) Display() string {
	return s.Value.StringFixed(4)
}

func (s *Quotation) Float64() float64 {
	return s.Value.InexactFloat64()
}

func (s *Quotation) Add(value *Quotation) *Quotation {
	return &Quotation{
		Value: s.Value.Add(value.Value),
	}
}

func (s *Quotation) Mul(value *Quotation) *Quotation {
	return &Quotation{
		Value: s.Value.Mul(value.Value),
	}
}

// Денежная сумма в определенной валюте
type MoneyValue struct {
	Currency string          `json:"currency,omitempty"` // строковый ISO-код валюты
	Value    decimal.Decimal `json:"value,omitempty"`
}

func (s *MoneyValue) GetCurrency() string {
	return s.Currency
}

func (s *MoneyValue) String() string {
	return s.Value.String()
}

func (s *MoneyValue) Display() string {
	return fmt.Sprintf("%s %s", s.Value.StringFixed(2), s.Currency)
}

func (s *MoneyValue) Float64() float64 {
	return s.Value.InexactFloat64()
}

func (s *MoneyValue) Add(value *MoneyValue) *MoneyValue {
	if s.Currency != value.Currency {
		return nil
	}

	return &MoneyValue{
		Currency: s.Currency,
		Value:    s.Value.Add(value.Value),
	}
}

func (s *MoneyValue) Mul(value *Quotation) *MoneyValue {
	return &MoneyValue{
		Currency: s.Currency,
		Value:    s.Value.Mul(value.Value),
	}
}

func (s *MoneyValue) Div(value *Quotation) *MoneyValue {
	if value.Value.IsZero() {
		return nil
	}

	return &MoneyValue{
		Currency: s.Currency,
		Value:    s.Value.Div(value.Value),
	}
}
