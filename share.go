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
	Figi                  string                `json:"figi"`                     //Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker"`                   //Тикер инструмента.
	ClassCode             string                `json:"class_code"`               //Класс-код (секция торгов).
	Isin                  string                `json:"isin"`                     //Isin-идентификатор инструмента.
	Lot                   int32                 `json:"lot"`                      //Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https://tinkoff.github.io/investAPI/glossary#lot)
	Currency              string                `json:"currency"`                 //Валюта расчётов.
	Klong                 *Quotation            `json:"klong"`                    //Коэффициент ставки риска длинной позиции по инструменту.
	Kshort                *Quotation            `json:"kshort"`                   //Коэффициент ставки риска короткой позиции по инструменту.
	Dlong                 *Quotation            `json:"dlong"`                    //Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort"`                   //Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min"`                //Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min"`               //Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabledFlag      bool                  `json:"short_enabled_flag"`       //Признак доступности для операций в шорт.
	Name                  string                `json:"name"`                     //Название инструмента.
	Exchange              string                `json:"exchange"`                 //Торговая площадка.
	IpoDate               *time.Time            `json:"ipo_date"`                 //Дата IPO акции в часовом поясе UTC.
	IssueSize             int64                 `json:"issue_size"`               //Размер выпуска.
	CountryOfRisk         string                `json:"country_of_risk"`          //Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name"`     //Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector                string                `json:"sector"`                   //Сектор экономики.
	IssueSizePlan         int64                 `json:"issue_size_plan"`          //Плановый размер выпуска.
	Nominal               *MoneyValue           `json:"nominal"`                  //Номинал.
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           //Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag"`                 //Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       //Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      //Признак доступности для продажи.
	DivYieldFlag          bool                  `json:"div_yield_flag"`           //Признак наличия дивидендной доходности.
	ShareType             ShareType             `json:"share_type"`               //Тип акции. Возможные значения: [ShareType](https://tinkoff.github.io/investAPI/instruments#sharetype)
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      //Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` //Признак доступности торгов через API.
}

// Информация о выплате дивидентов
type Dividend struct {
	DividendNet  *MoneyValue `json:"dividend_net"`  // Величина дивиденда на 1 ценную бумагу (включая валюту).
	PaymentDate  *time.Time  `json:"payment_date"`  // Дата фактических выплат в часовом поясе UTC.
	DeclaredDate *time.Time  `json:"declared_date"` // Дата объявления дивидендов в часовом поясе UTC.
	LastBuyDate  *time.Time  `json:"last_buy_date"` // Последний день (включительно) покупки для получения выплаты в часовом поясе UTC.
	DividendType string      `json:"dividend_type"` // Тип выплаты. Возможные значения: Regular Cash – регулярные выплаты, Cancelled – выплата отменена, Daily Accrual – ежедневное начисление, Return of Capital – возврат капитала, прочие типы выплат.
	RecordDate   *time.Time  `json:"record_date"`   // Дата фиксации реестра в часовом поясе UTC.
	Regularity   string      `json:"regularity"`    // Регулярность выплаты. Возможные значения: Annual – ежегодная, Semi-Anl – каждые полгода, прочие типы выплат.
	ClosePrice   *MoneyValue `json:"close_price"`   // Цена закрытия инструмента на момент ex_dividend_date.
	YieldValue   *Quotation  `json:"yield_value"`   // Величина доходности.
	CreatedAt    *time.Time  `json:"created_at"`    // Дата и время создания записи в часовом поясе UTC.
}
