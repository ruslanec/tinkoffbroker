package tinkoffbroker

import (
	"fmt"
	"strconv"
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
	Units int64 `json:"units"` // целая часть суммы, может быть отрицательным числом
	Nano  int32 `json:"nano"`  // дробная часть суммы, может быть отрицательным числом
}

func (s *Quotation) String() string {
	units := strconv.FormatInt(s.Units, 10)
	nano := strconv.FormatInt(int64(s.Nano), 10)
	return fmt.Sprintf("%s.%s", units, nano)
}

func (s *Quotation) Float64() float64 {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		fmt.Printf("Quotation Float64(): %v", err)
	}
	return f
}

// Денежная сумма в определенной валюте
type MoneyValue struct {
	Currency string `json:"currency"` // строковый ISO-код валюты
	Units    int64  `json:"units"`    // целая часть суммы, может быть отрицательным числом
	Nano     int32  `json:"nano"`     // дробная часть суммы, может быть отрицательным числом
}

func (s *MoneyValue) String() string {
	units := strconv.FormatInt(s.Units, 10)
	nano := strconv.FormatInt(int64(s.Nano), 10)
	return fmt.Sprintf("%s.%s %s", units, nano, s.Currency)
}

func (s *MoneyValue) Float64() float64 {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		fmt.Printf("MoneyValue Float64(): %v", err)
	}
	return f
}
