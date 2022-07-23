package tinkoffbroker

import "time"

// Объект передачи информации о фьючерсе.
type Future struct {
	Figi                  string                `json:"figi"`                     // Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker"`                   // Тикер инструмента.
	ClassCode             string                `json:"class_code"`               // Класс-код (секция торгов).
	Lot                   int32                 `json:"lot"`                      // Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https:// tinkoff.github.io/investAPI/glossary#lot)
	Currency              string                `json:"currency"`                 // Валюта расчётов.
	Klong                 *Quotation            `json:"klong"`                    // Коэффициент ставки риска длинной позиции по клиенту.
	Kshort                *Quotation            `json:"kshort"`                   // Коэффициент ставки риска короткой позиции по клиенту.
	Dlong                 *Quotation            `json:"dlong"`                    // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort"`                   // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min"`                // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min"`               // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabledFlag      bool                  `json:"short_enabled_flag"`       // Признак доступности для операций шорт.
	Name                  string                `json:"name"`                     // Название инструмента.
	Exchange              string                `json:"exchange"`                 // Торговая площадка.
	FirstTradeDate        *time.Time            `json:"first_trade_date"`         // Дата начала обращения контракта в часовом поясе UTC.
	LastTradeDate         *time.Time            `json:"last_trade_date"`          // Дата в часовом поясе UTC, до которой возможно проведение операций с фьючерсом.
	FuturesType           string                `json:"futures_type"`             // Тип фьючерса. Возможные значения: </br>**physical_delivery** — физические поставки; </br>**cash_settlement** — денежный эквивалент.
	AssetType             string                `json:"asset_type"`               // Тип актива. Возможные значения: </br>**commodity** — товар; </br>**currency** — валюта; </br>**security** — ценная бумага; </br>**index** — индекс.
	BasicAsset            string                `json:"basic_asset"`              // Основной актив.
	BasicAssetSize        *Quotation            `json:"basic_asset_size"`         // Размер основного актива.
	CountryOfRisk         string                `json:"country_of_risk"`          // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name"`     // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector                string                `json:"sector"`                   // Сектор экономики.
	ExpirationDate        *time.Time            `json:"expiration_date"`          // Дата истечения срока в часов поясе UTC.
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           // Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag"`                 // Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       // Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      // Признак доступности для продажи.
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      // Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` // Признак доступности торгов через API.
}

// Данные по фьючерсу
type FuturesMargin struct {
	InitialMarginOnBuy      *MoneyValue `json:"initial_margin_on_buy"`      //Гарантийное обеспечение при покупке.
	InitialMarginOnSell     *MoneyValue `json:"initial_margin_on_sell"`     //Гарантийное обеспечение при продаже.
	MinPriceIncrement       *Quotation  `json:"min_price_increment"`        //Шаг цены.
	MinPriceIncrementAmount *Quotation  `json:"min_price_increment_amount"` //Стоимость шага цены.

}
