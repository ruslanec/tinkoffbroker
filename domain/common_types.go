package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

//Режим торгов инструмента
type SecurityTradingStatus int32

const (
	SecurityTradingStatus_SECURITY_TRADING_STATUS_UNSPECIFIED                      SecurityTradingStatus = 0  //Торговый статус не определён
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NOT_AVAILABLE_FOR_TRADING        SecurityTradingStatus = 1  //Недоступен для торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_PERIOD                   SecurityTradingStatus = 2  //Период открытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_PERIOD                   SecurityTradingStatus = 3  //Период закрытия торгов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_BREAK_IN_TRADING                 SecurityTradingStatus = 4  //Перерыв в торговле
	SecurityTradingStatus_SECURITY_TRADING_STATUS_NORMAL_TRADING                   SecurityTradingStatus = 5  //Нормальная торговля
	SecurityTradingStatus_SECURITY_TRADING_STATUS_CLOSING_AUCTION                  SecurityTradingStatus = 6  //Аукцион закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DARK_POOL_AUCTION                SecurityTradingStatus = 7  //Аукцион крупных пакетов
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DISCRETE_AUCTION                 SecurityTradingStatus = 8  //Дискретный аукцион
	SecurityTradingStatus_SECURITY_TRADING_STATUS_OPENING_AUCTION_PERIOD           SecurityTradingStatus = 9  //Аукцион открытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_TRADING_AT_CLOSING_AUCTION_PRICE SecurityTradingStatus = 10 //Период торгов по цене аукциона закрытия
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_ASSIGNED                 SecurityTradingStatus = 11 //Сессия назначена
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_CLOSE                    SecurityTradingStatus = 12 //Сессия закрыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_SESSION_OPEN                     SecurityTradingStatus = 13 //Сессия открыта
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NORMAL_TRADING            SecurityTradingStatus = 14 //Доступна торговля в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_BREAK_IN_TRADING          SecurityTradingStatus = 15 //Перерыв торговли в режиме внутренней ликвидности брокера
	SecurityTradingStatus_SECURITY_TRADING_STATUS_DEALER_NOT_AVAILABLE_FOR_TRADING SecurityTradingStatus = 16 //Недоступна торговля в режиме внутренней ликвидности брокера
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
