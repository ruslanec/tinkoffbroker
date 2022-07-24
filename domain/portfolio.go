package domain

import (
	"errors"
	"fmt"
)

// Текущий портфель по счёту
type Portfolio struct {
	TotalAmountShares     *MoneyValue          `json:"total_amount_shares,omitempty"`     // Общая стоимость акций в портфеле в рублях
	TotalAmountBonds      *MoneyValue          `json:"total_amount_bonds,omitempty"`      // Общая стоимость облигаций в портфеле в рублях
	TotalAmountEtf        *MoneyValue          `json:"total_amount_etf,omitempty"`        // Общая стоимость фондов в портфеле в рублях
	TotalAmountCurrencies *MoneyValue          `json:"total_amount_currencies,omitempty"` // Общая стоимость валют в портфеле в рублях
	TotalAmountFutures    *MoneyValue          `json:"total_amount_futures,omitempty"`    // Общая стоимость фьючерсов в портфеле в рублях
	ExpectedYield         *Quotation           `json:"expected_yield,omitempty"`          // Текущая относительная доходность портфеля, в %
	Positions             []*PortfolioPosition `json:"positions,omitempty"`               // Список позиций портфеля
	AcceptableLevelOfRisk map[string]float64   `json:"risk,omitempty"`                    // Допустимый уровень риска для инструмента
}

// Допустимый уровень риска для инструмента
func (p *Portfolio) RiskLevel(figi string) (float64, error) {
	if p.AcceptableLevelOfRisk != nil {
		return 0, errors.New("field Risk not initialized")
	}
	risk, ok := p.AcceptableLevelOfRisk[figi]
	if !ok {
		return 0, errors.New("risk value not setted")
	}
	return risk, nil
}

// Позиции портфеля
type PortfolioPosition struct {
	Figi                     string      `json:"figi,omitempty"`                        // Figi-идентификатора инструмента
	InstrumentType           string      `json:"instrument_type,omitempty"`             // Тип инструмента
	Quantity                 *Quotation  `json:"quantity,omitempty"`                    // Количество инструмента в портфеле в штуках
	AveragePositionPrice     *MoneyValue `json:"average_position_price,omitempty"`      // Средневзвешенная цена позиции. **Возможна задержка до секунды для пересчёта**.
	ExpectedYield            *Quotation  `json:"expected_yield,omitempty"`              // Текущая рассчитанная относительная доходность позиции, в %.
	CurrentNkd               *MoneyValue `json:"current_nkd,omitempty"`                 // Текущий НКД
	AveragePositionPricePt   *Quotation  `json:"average_position_price_pt,omitempty"`   // Средняя цена лота в позиции в пунктах (для фьючерсов). **Возможна задержка до секунды для пересчёта**.
	CurrentPrice             *MoneyValue `json:"current_price,omitempty"`               // Текущая цена инструмента
	AveragePositionPriceFifo *MoneyValue `json:"average_position_price_fifo,omitempty"` // Средняя цена лота в позиции по методу FIFO. **Возможна задержка до секунды для пересчёта**.
	QuantityLots             *Quotation  `json:"quantity_lots,omitempty"`               // Количество лотов в портфеле
}

func (s *PortfolioPosition) String() string {
	return fmt.Sprintf("%s", s.Figi)
}

// Список позиций по счёту.
type Positions struct {
	Money                   []*MoneyValue         `json:"money,omitempty"`                      // Массив валютных позиций портфеля
	Blocked                 []*MoneyValue         `json:"blocked,omitempty"`                    // Массив заблокированных валютных позиций портфеля
	Securities              []*PositionInstrument `json:"securities,omitempty"`                 // Список ценно-бумажных позиций портфеля
	LimitsLoadingInProgress bool                  `json:"limits_loading_in_progress,omitempty"` // Признак идущей в данный момент выгрузки лимитов
	Futures                 []*PositionInstrument `json:"futures,omitempty"`                    // Список фьючерсов портфеля
}

//Баланс позиции инструмента
type PositionInstrument struct {
	Figi    string `json:"figi,omitempty"`    // Figi-идентификатор бумаги
	Blocked int64  `json:"blocked,omitempty"` // Заблокировано
	Balance int64  `json:"balance,omitempty"` // Текущий незаблокированный баланс
}
