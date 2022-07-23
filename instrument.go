package tinkoffbroker

import "time"

// Объект передачи основной информации об инструменте.
type Instrument struct {
	Figi                  string                `json:"figi"`                     // Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker"`                   // Тикер инструмента.
	ClassCode             string                `json:"class_code"`               // Класс-код инструмента.
	Isin                  string                `json:"isin"`                     // Isin-идентификатор инструмента.
	Lot                   int32                 `json:"lot"`                      // Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https:// tinkoff.github.io/investAPI/glossary#lot)
	Currency              string                `json:"currency"`                 // Валюта расчётов.
	Klong                 *Quotation            `json:"klong"`                    // Коэффициент ставки риска длинной позиции по инструменту.
	Kshort                *Quotation            `json:"kshort"`                   // Коэффициент ставки риска короткой позиции по инструменту.
	Dlong                 *Quotation            `json:"dlong"`                    // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort"`                   // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min"`                // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min"`               // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabledFlag      bool                  `json:"short_enabled_flag"`       // Признак доступности для операций в шорт.
	Name                  string                `json:"name"`                     // Название инструмента.
	Exchange              string                `json:"exchange"`                 // Торговая площадка.
	CountryOfRisk         string                `json:"country_of_risk"`          // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name"`     // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	InstrumentType        string                `json:"instrument_type"`          // Тип инструмента.
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           // Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag"`                 // Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       // Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      // Признак доступности для продажи.
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      // Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` // Признак доступности торгов через API.
}

// Запрос подписки на торговый статус
type InfoInstrument struct {
	Figi string `json:"figi"` // Figi-идентификатор инструмента
}

// Краткая информация об инструменте.
type InstrumentShort struct {
	Isin                string     `json:"isin,omitempty"`                    // Isin инструмента.
	Figi                string     `json:"figi,omitempty"`                    // Figi инструмента.
	Ticker              string     `json:"ticker,omitempty"`                  // Ticker инструмента.
	ClassCode           string     `json:"class_code,omitempty"`              // ClassCode инструмента.
	InstrumentType      string     `json:"instrument_type,omitempty"`         // Тип инструмента.
	Name                string     `json:"name,omitempty"`                    // Название инструмента.
	Uid                 string     `json:"uid,omitempty"`                     // Уникальный идентификатор инструмента.
	PositionUid         string     `json:"position_uid,omitempty"`            // Уникальный идентификатор позиции инструмента.
	ApiTradeAvailable   bool       `json:"api_trade_available,omitempty"`     // Признак доступности торгов через API.
	ForIis              bool       `json:"for_iis,omitempty"`                 // Признак доступности для ИИС.
	First1minCandleDate *time.Time `json:"first_1_min_candle_date,omitempty"` // Дата первой минутной свечи.
	First1dayCandleDate *time.Time `json:"first_1_day_candle_date,omitempty"` // Дата первой дневной свечи.
}
