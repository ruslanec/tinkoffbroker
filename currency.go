package tinkoffbroker

// Объект передачи информации о валюте.
type Currency struct {
	Figi                  string                `json:"figi"`                     // Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker"`                   // Тикер инструмента.
	ClassCode             string                `json:"class_code"`               // Класс-код (секция торгов).
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
	Nominal               *MoneyValue           `json:"nominal"`                  // Номинал.
	CountryOfRisk         string                `json:"country_of_risk"`          // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name"`     // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           // Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag"`                 // Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       // Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      // Признак доступности для продажи.
	IsoCurrencyName       string                `json:"iso_currency_name"`        // Строковый ISO-код валюты.
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      // Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` // Признак доступности торгов через API.
}
