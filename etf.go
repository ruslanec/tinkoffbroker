package tinkoffbroker

import "time"

// Объект передачи информации об инвестиционном фонде
type Etf struct {
	Figi                  string                `json:"figi"`                     // Figi-идентификатор инструмента
	Ticker                string                `json:"ticker"`                   // Тикер инструмента
	ClassCode             string                `json:"class_code"`               // Класс-код (секция торгов)
	Isin                  string                `json:"isin"`                     // Isin-идентификатор инструмента
	Lot                   int32                 `json:"lot"`                      // Лотность инструмента Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*
	Currency              string                `json:"currency"`                 // Валюта расчётов
	Klong                 *Quotation            `json:"klong"`                    // Коэффициент ставки риска длинной позиции по инструменту
	Kshort                *Quotation            `json:"kshort"`                   // Коэффициент ставки риска короткой позиции по инструменту
	Dlong                 *Quotation            `json:"dlong"`                    // Ставка риска минимальной маржи в лонг
	Dshort                *Quotation            `json:"dshort"`                   // Ставка риска минимальной маржи в шорт
	DlongMin              *Quotation            `json:"dlong_min"`                // Ставка риска начальной маржи в лонг
	DshortMin             *Quotation            `json:"dshort_min"`               // Ставка риска начальной маржи в шорт
	ShortEnabledFlag      bool                  `json:"short_enabled_flag"`       // Признак доступности для операций в шорт
	Name                  string                `json:"name"`                     // Название инструмента
	Exchange              string                `json:"exchange"`                 // Торговая площадка
	FixedCommission       *Quotation            `json:"fixed_commission"`         // Размер фиксированной комиссии фонда
	FocusType             string                `json:"focus_type"`               // Возможные значения: </br>**equity** — акции;</br>**fixed_income** — облигации;</br>**mixed_allocation** — смешанный;</br>**money_market** — денежный рынок;</br>**real_estate** — недвижимость;</br>**commodity** — товары;</br>**specialty** — специальный;</br>**private_equity** — private equity;</br>**alternative_investment** — альтернативные инвестиции
	ReleasedDate          *time.Time            `json:"released_date"`            // Дата выпуска в часовом поясе UTC
	NumShares             *Quotation            `json:"num_shares"`               // Количество акций фонда в обращении
	CountryOfRisk         string                `json:"country_of_risk"`          // Код страны риска, те страны, в которой компания ведёт основной бизнес
	CountryOfRiskName     string                `json:"country_of_risk_name"`     // Наименование страны риска, те страны, в которой компания ведёт основной бизнес
	Sector                string                `json:"sector"`                   // Сектор экономики
	RebalancingFreq       string                `json:"rebalancing_freq"`         // Частота ребалансировки
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           // Текущий режим торгов инструмента
	OtcFlag               bool                  `json:"otc_flag"`                 // Признак внебиржевой ценной бумаги
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       // Признак доступности для покупки
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      // Признак доступности для продажи
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      // Шаг цены
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` // Признак доступности торгов через API
}
