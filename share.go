package tinkoffbroker

import (
	"time"
)

//Тип акций.
type ShareType int32

const (
	ShareType_SHARE_TYPE_UNSPECIFIED     ShareType = 0 //Значение не определено.
	ShareType_SHARE_TYPE_COMMON          ShareType = 1 //Обыкновенная
	ShareType_SHARE_TYPE_PREFERRED       ShareType = 2 //Привилегированная
	ShareType_SHARE_TYPE_ADR             ShareType = 3 //Американские депозитарные расписки
	ShareType_SHARE_TYPE_GDR             ShareType = 4 //Глобальные депозитарные расписки
	ShareType_SHARE_TYPE_MLP             ShareType = 5 //Товарищество с ограниченной ответственностью
	ShareType_SHARE_TYPE_NY_REG_SHRS     ShareType = 6 //Акции из реестра Нью-Йорка
	ShareType_SHARE_TYPE_CLOSED_END_FUND ShareType = 7 //Закрытый инвестиционный фонд
	ShareType_SHARE_TYPE_REIT            ShareType = 8 //Траст недвижимости
)

//Объект передачи информации об акции.
type Share struct {
	Figi                  string                `json:"figi,omitempty"`                     //Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker,omitempty"`                   //Тикер инструмента.
	ClassCode             string                `json:"class_code,omitempty"`               //Класс-код (секция торгов).
	Isin                  string                `json:"isin,omitempty"`                     //Isin-идентификатор инструмента.
	Lot                   int32                 `json:"lot,omitempty"`                      //Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https://tinkoff.github.io/investAPI/glossary#lot)
	Currency              string                `json:"currency,omitempty"`                 //Валюта расчётов.
	Klong                 *Quotation            `json:"klong,omitempty"`                    //Коэффициент ставки риска длинной позиции по инструменту.
	Kshort                *Quotation            `json:"kshort,omitempty"`                   //Коэффициент ставки риска короткой позиции по инструменту.
	Dlong                 *Quotation            `json:"dlong,omitempty"`                    //Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort,omitempty"`                   //Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min,omitempty"`                //Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min,omitempty"`               //Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabledFlag      bool                  `json:"short_enabled_flag,omitempty"`       //Признак доступности для операций в шорт.
	Name                  string                `json:"name,omitempty"`                     //Название инструмента.
	Exchange              string                `json:"exchange,omitempty"`                 //Торговая площадка.
	IpoDate               *time.Time            `json:"ipo_date,omitempty"`                 //Дата IPO акции в часовом поясе UTC.
	IssueSize             int64                 `json:"issue_size,omitempty"`               //Размер выпуска.
	CountryOfRisk         string                `json:"country_of_risk,omitempty"`          //Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name,omitempty"`     //Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector                string                `json:"sector,omitempty"`                   //Сектор экономики.
	IssueSizePlan         int64                 `json:"issue_size_plan,omitempty"`          //Плановый размер выпуска.
	Nominal               *MoneyValue           `json:"nominal,omitempty"`                  //Номинал.
	TradingStatus         SecurityTradingStatus `json:"trading_status,omitempty"`           //Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag,omitempty"`                 //Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag,omitempty"`       //Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag,omitempty"`      //Признак доступности для продажи.
	DivYieldFlag          bool                  `json:"div_yield_flag,omitempty"`           //Признак наличия дивидендной доходности.
	ShareType             ShareType             `json:"share_type,omitempty"`               //Тип акции. Возможные значения: [ShareType](https://tinkoff.github.io/investAPI/instruments#sharetype)
	MinPriceIncrement     *Quotation            `json:"min_price_increment,omitempty"`      //Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag,omitempty"` //Признак доступности торгов через API.
}

// Информация о выплате дивидентов
type Dividend struct {
	DividendNet  *MoneyValue `json:"dividend_net,omitempty"`  // Величина дивиденда на 1 ценную бумагу (включая валюту).
	PaymentDate  *time.Time  `json:"payment_date,omitempty"`  // Дата фактических выплат в часовом поясе UTC.
	DeclaredDate *time.Time  `json:"declared_date,omitempty"` // Дата объявления дивидендов в часовом поясе UTC.
	LastBuyDate  *time.Time  `json:"last_buy_date,omitempty"` // Последний день (включительно) покупки для получения выплаты в часовом поясе UTC.
	DividendType string      `json:"dividend_type,omitempty"` // Тип выплаты. Возможные значения: Regular Cash – регулярные выплаты, Cancelled – выплата отменена, Daily Accrual – ежедневное начисление, Return of Capital – возврат капитала, прочие типы выплат.
	RecordDate   *time.Time  `json:"record_date,omitempty"`   // Дата фиксации реестра в часовом поясе UTC.
	Regularity   string      `json:"regularity,omitempty"`    // Регулярность выплаты. Возможные значения: Annual – ежегодная, Semi-Anl – каждые полгода, прочие типы выплат.
	ClosePrice   *MoneyValue `json:"close_price,omitempty"`   // Цена закрытия инструмента на момент ex_dividend_date.
	YieldValue   *Quotation  `json:"yield_value,omitempty"`   // Величина доходности.
	CreatedAt    *time.Time  `json:"created_at,omitempty"`    // Дата и время создания записи в часовом поясе UTC.
}
