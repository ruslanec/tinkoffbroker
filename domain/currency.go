package domain

import (
	"time"
)

// Объект передачи информации о валюте.
type Currency struct {
	Figi                string                `json:"figi,omitempty"`                    // Figi-идентификатор инструмента.
	Ticker              string                `json:"ticker,omitempty"`                  // Тикер инструмента.
	ClassCode           string                `json:"class_code,omitempty"`              // Класс-код (секция торгов).
	Isin                string                `json:"isin,omitempty"`                    // Isin-идентификатор инструмента.
	Lot                 int32                 `json:"lot,omitempty"`                     // Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https://tinkoff.github.io/investAPI/glossary#lot)
	Currency            string                `json:"currency,omitempty"`                // Валюта расчётов.
	Klong               *Quotation            `json:"klong,omitempty"`                   // Коэффициент ставки риска длинной позиции по инструменту.
	Kshort              *Quotation            `json:"kshort,omitempty"`                  // Коэффициент ставки риска короткой позиции по инструменту.
	Dlong               *Quotation            `json:"dlong,omitempty"`                   // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort              *Quotation            `json:"dshort,omitempty"`                  // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin            *Quotation            `json:"dlong_min,omitempty"`               // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin           *Quotation            `json:"dshort_min,omitempty"`              // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabled        bool                  `json:"short_enabled,omitempty"`           // Признак доступности для операций в шорт.
	Name                string                `json:"name,omitempty"`                    // Название инструмента.
	Exchange            string                `json:"exchange,omitempty"`                // Торговая площадка.
	Nominal             *MoneyValue           `json:"nominal,omitempty"`                 // Номинал.
	CountryOfRisk       string                `json:"country_of_risk,omitempty"`         // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName   string                `json:"country_of_risk_name,omitempty"`    // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	TradingStatus       SecurityTradingStatus `json:"trading_status,omitempty"`          // Текущий режим торгов инструмента.
	Otc                 bool                  `json:"otc,omitempty"`                     // Признак внебиржевой ценной бумаги.
	BuyAvailable        bool                  `json:"buy_available,omitempty"`           // Признак доступности для покупки.
	SellAvailable       bool                  `json:"sell_available,omitempty"`          // Признак доступности для продажи.
	IsoCurrencyName     string                `json:"iso_currency_name,omitempty"`       // Строковый ISO-код валюты.
	MinPriceIncrement   *Quotation            `json:"min_price_increment,omitempty"`     // Шаг цены.
	APITradeAvailable   bool                  `json:"api_trade_available,omitempty"`     // Признак доступности торгов через API.
	UID                 string                `json:"uid,omitempty"`                     // Уникальный идентификатор инструмента.
	RealExchange        RealExchange          `json:"real_exchange,omitempty"`           // Реальная площадка исполнения расчётов.
	PositionUID         string                `json:"position_uid,omitempty"`            // Уникальный идентификатор позиции инструмента.
	ForIIS              bool                  `json:"for_iis,omitempty"`                 // Признак доступности для ИИС.
	First1MinCandleDate *time.Time            `json:"first_1_min_candle_date,omitempty"` // Дата первой минутной свечи.
	First1DayCandleDate *time.Time            `json:"first_1_day_candle_date,omitempty"` // Дата первой дневной свечи.
}
