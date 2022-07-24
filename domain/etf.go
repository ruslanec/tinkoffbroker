package domain

import "time"

// Объект передачи информации об инвестиционном фонде
type Etf struct {
	Figi              string                `json:"figi,omitempty"`                 // Figi-идентификатор инструмента
	Ticker            string                `json:"ticker,omitempty"`               // Тикер инструмента
	ClassCode         string                `json:"class_code,omitempty"`           // Класс-код (секция торгов)
	Isin              string                `json:"isin,omitempty"`                 // Isin-идентификатор инструмента
	Lot               int32                 `json:"lot,omitempty"`                  // Лотность инструмента Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*
	Currency          string                `json:"currency,omitempty"`             // Валюта расчётов
	Klong             *Quotation            `json:"klong,omitempty"`                // Коэффициент ставки риска длинной позиции по инструменту
	Kshort            *Quotation            `json:"kshort,omitempty"`               // Коэффициент ставки риска короткой позиции по инструменту
	Dlong             *Quotation            `json:"dlong,omitempty"`                // Ставка риска минимальной маржи в лонг
	Dshort            *Quotation            `json:"dshort,omitempty"`               // Ставка риска минимальной маржи в шорт
	DlongMin          *Quotation            `json:"dlong_min,omitempty"`            // Ставка риска начальной маржи в лонг
	DshortMin         *Quotation            `json:"dshort_min,omitempty"`           // Ставка риска начальной маржи в шорт
	ShortEnabled      bool                  `json:"short_enabled,omitempty"`        // Признак доступности для операций в шорт
	Name              string                `json:"name,omitempty"`                 // Название инструмента
	Exchange          string                `json:"exchange,omitempty"`             // Торговая площадка
	FixedCommission   *Quotation            `json:"fixed_commission,omitempty"`     // Размер фиксированной комиссии фонда
	FocusType         string                `json:"focus_type,omitempty"`           // Возможные значения: </br>**equity** — акции;</br>**fixed_income** — облигации;</br>**mixed_allocation** — смешанный;</br>**money_market** — денежный рынок;</br>**real_estate** — недвижимость;</br>**commodity** — товары;</br>**specialty** — специальный;</br>**private_equity** — private equity;</br>**alternative_investment** — альтернативные инвестиции
	ReleasedDate      *time.Time            `json:"released_date,omitempty"`        // Дата выпуска в часовом поясе UTC
	NumShares         *Quotation            `json:"num_shares,omitempty"`           // Количество акций фонда в обращении
	CountryOfRisk     string                `json:"country_of_risk,omitempty"`      // Код страны риска, те страны, в которой компания ведёт основной бизнес
	CountryOfRiskName string                `json:"country_of_risk_name,omitempty"` // Наименование страны риска, те страны, в которой компания ведёт основной бизнес
	Sector            string                `json:"sector,omitempty"`               // Сектор экономики
	RebalancingFreq   string                `json:"rebalancing_freq,omitempty"`     // Частота ребалансировки
	TradingStatus     SecurityTradingStatus `json:"trading_status,omitempty"`       // Текущий режим торгов инструмента
	Otc               bool                  `json:"otc,omitempty"`                  // Признак внебиржевой ценной бумаги
	BuyAvailable      bool                  `json:"buy_available,omitempty"`        // Признак доступности для покупки
	SellAvailable     bool                  `json:"sell_available,omitempty"`       // Признак доступности для продажи
	MinPriceIncrement *Quotation            `json:"min_price_increment,omitempty"`  // Шаг цены
	ApiTradeAvailable bool                  `json:"api_trade_available,omitempty"`  // Признак доступности торгов через API
}
