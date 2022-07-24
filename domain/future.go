package domain

import "time"

// Объект передачи информации о фьючерсе.
type Future struct {
	Figi              string                `json:"figi,omitempty"`                 // Figi-идентификатор инструмента.
	Ticker            string                `json:"ticker,omitempty"`               // Тикер инструмента.
	ClassCode         string                `json:"class_code,omitempty"`           // Класс-код (секция торгов).
	Lot               int32                 `json:"lot,omitempty"`                  // Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https:// tinkoff.github.io/investAPI/glossary#lot)
	Currency          string                `json:"currency,omitempty"`             // Валюта расчётов.
	Klong             *Quotation            `json:"klong,omitempty"`                // Коэффициент ставки риска длинной позиции по клиенту.
	Kshort            *Quotation            `json:"kshort,omitempty"`               // Коэффициент ставки риска короткой позиции по клиенту.
	Dlong             *Quotation            `json:"dlong,omitempty"`                // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort            *Quotation            `json:"dshort,omitempty"`               // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin          *Quotation            `json:"dlong_min,omitempty"`            // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin         *Quotation            `json:"dshort_min,omitempty"`           // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabled      bool                  `json:"short_enabled,omitempty"`        // Признак доступности для операций шорт.
	Name              string                `json:"name,omitempty"`                 // Название инструмента.
	Exchange          string                `json:"exchange,omitempty"`             // Торговая площадка.
	FirstTradeDate    *time.Time            `json:"first_trade_date,omitempty"`     // Дата начала обращения контракта в часовом поясе UTC.
	LastTradeDate     *time.Time            `json:"last_trade_date,omitempty"`      // Дата в часовом поясе UTC, до которой возможно проведение операций с фьючерсом.
	FuturesType       string                `json:"futures_type,omitempty"`         // Тип фьючерса. Возможные значения: </br>**physical_delivery** — физические поставки; </br>**cash_settlement** — денежный эквивалент.
	AssetType         string                `json:"asset_type,omitempty"`           // Тип актива. Возможные значения: </br>**commodity** — товар; </br>**currency** — валюта; </br>**security** — ценная бумага; </br>**index** — индекс.
	BasicAsset        string                `json:"basic_asset,omitempty"`          // Основной актив.
	BasicAssetSize    *Quotation            `json:"basic_asset_size,omitempty"`     // Размер основного актива.
	CountryOfRisk     string                `json:"country_of_risk,omitempty"`      // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName string                `json:"country_of_risk_name,omitempty"` // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector            string                `json:"sector,omitempty"`               // Сектор экономики.
	ExpirationDate    *time.Time            `json:"expiration_date,omitempty"`      // Дата истечения срока в часов поясе UTC.
	TradingStatus     SecurityTradingStatus `json:"trading_status,omitempty"`       // Текущий режим торгов инструмента.
	Otc               bool                  `json:"otc,omitempty"`                  // Признак внебиржевой ценной бумаги.
	BuyAvailable      bool                  `json:"buy_available,omitempty"`        // Признак доступности для покупки.
	SellAvailable     bool                  `json:"sell_available,omitempty"`       // Признак доступности для продажи.
	MinPriceIncrement *Quotation            `json:"min_price_increment,omitempty"`  // Шаг цены.
	ApiTradeAvailable bool                  `json:"api_trade_available,omitempty"`  // Признак доступности торгов через API.
}

// Данные по фьючерсу
type FuturesMargin struct {
	InitialMarginOnBuy      *MoneyValue `json:"initial_margin_on_buy,omitempty"`      //Гарантийное обеспечение при покупке.
	InitialMarginOnSell     *MoneyValue `json:"initial_margin_on_sell,omitempty"`     //Гарантийное обеспечение при продаже.
	MinPriceIncrement       *Quotation  `json:"min_price_increment,omitempty"`        //Шаг цены.
	MinPriceIncrementAmount *Quotation  `json:"min_price_increment_amount,omitempty"` //Стоимость шага цены.
}
