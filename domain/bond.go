package domain

import (
	"time"
)

// Тип купонов.
type CouponType int32

const (
	CouponType_COUPON_TYPE_UNSPECIFIED CouponType = 0 // Неопределенное значение
	CouponType_COUPON_TYPE_CONSTANT    CouponType = 1 // Постоянный
	CouponType_COUPON_TYPE_FLOATING    CouponType = 2 // Плавающий
	CouponType_COUPON_TYPE_DISCOUNT    CouponType = 3 // Дисконт
	CouponType_COUPON_TYPE_MORTGAGE    CouponType = 4 // Ипотечный
	CouponType_COUPON_TYPE_FIX         CouponType = 5 // Фиксированный
	CouponType_COUPON_TYPE_VARIABLE    CouponType = 6 // Переменный
	CouponType_COUPON_TYPE_OTHER       CouponType = 7 // Прочее
)

// Объект передачи информации об облигации.
type Bond struct {
	Figi                  string                `json:"figi,omitempty"`                     // Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker,omitempty"`                   // Тикер инструмента.
	ClassCode             string                `json:"class_code,omitempty"`               // Класс-код (секция торгов).
	Isin                  string                `json:"isin,omitempty"`                     // Isin-идентификатор инструмента.
	Lot                   int32                 `json:"lot,omitempty"`                      // Лотность инструмента. Возможно совершение операций только на количества ценной бумаги, кратные параметру *lot*. Подробнее: [лот](https:// tinkoff.github.io/investAPI/glossary#lot)
	Currency              string                `json:"currency,omitempty"`                 // Валюта расчётов.
	Klong                 *Quotation            `json:"klong,omitempty"`                    // Коэффициент ставки риска длинной позиции по инструменту.
	Kshort                *Quotation            `json:"kshort,omitempty"`                   // Коэффициент ставки риска короткой позиции по инструменту.
	Dlong                 *Quotation            `json:"dlong,omitempty"`                    // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort,omitempty"`                   // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min,omitempty"`                // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https:// help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min,omitempty"`               // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https:// help.tinkoff.ru/margin-trade/short/risk-rate/)
	ShortEnabled          bool                  `json:"short_enabled,omitempty"`            // Признак доступности для операций в шорт.
	Name                  string                `json:"name,omitempty"`                     // Название инструмента.
	Exchange              string                `json:"exchange,omitempty"`                 // Торговая площадка.
	CouponQuantityPerYear int32                 `json:"coupon_quantity_per_year,omitempty"` // Количество выплат по купонам в год.
	MaturityDate          *time.Time            `json:"maturity_date,omitempty"`            // Дата погашения облигации в часовом поясе UTC.
	Nominal               *MoneyValue           `json:"nominal,omitempty"`                  // Номинал облигации.
	StateRegDate          *time.Time            `json:"state_reg_date,omitempty"`           // Дата выпуска облигации в часовом поясе UTC.
	PlacementDate         *time.Time            `json:"placement_date,omitempty"`           // Дата размещения в часовом поясе UTC.
	PlacementPrice        *MoneyValue           `json:"placement_price,omitempty"`          // Цена размещения.
	AciValue              *MoneyValue           `json:"aci_value,omitempty"`                // Значение НКД (накопленного купонного дохода) на дату.
	CountryOfRisk         string                `json:"country_of_risk,omitempty"`          // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name,omitempty"`     // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector                string                `json:"sector,omitempty"`                   // Сектор экономики.
	IssueKind             string                `json:"issue_kind,omitempty"`               // Форма выпуска. Возможные значения: </br>**documentary** — документарная; </br>**non_documentary** — бездокументарная.
	IssueSize             int64                 `json:"issue_size,omitempty"`               // Размер выпуска.
	IssueSizePlan         int64                 `json:"issue_size_plan,omitempty"`          // Плановый размер выпуска.
	TradingStatus         SecurityTradingStatus `json:"trading_status,omitempty"`           // Текущий режим торгов инструмента.
	Otc                   bool                  `json:"otc,omitempty"`                      // Признак внебиржевой ценной бумаги.
	BuyAvailable          bool                  `json:"buy_available,omitempty"`            // Признак доступности для покупки.
	SellAvailable         bool                  `json:"sell_available,omitempty"`           // Признак доступности для продажи.
	FloatingCoupon        bool                  `json:"floating_coupon,omitempty"`          // Признак облигации с плавающим купоном.
	Perpetual             bool                  `json:"perpetual,omitempty"`                // Признак бессрочной облигации.
	Amortization          bool                  `json:"amortization,omitempty"`             // Признак облигации с амортизацией долга.
	MinPriceIncrement     *Quotation            `json:"min_price_increment,omitempty"`      // Шаг цены.
	ApiTradeAvailable     bool                  `json:"api_trade_available,omitempty"`      // Признак доступности торгов через API.
}

// Объект передачи информации о купоне облигации
type Coupon struct {
	Figi            string      `json:"figi,omitempty"`              // Figi-идентификатор инструмента
	CouponDate      *time.Time  `json:"coupon_date,omitempty"`       // Дата события
	CouponNumber    int64       `json:"coupon_number,omitempty"`     // Номер купона
	FixDate         *time.Time  `json:"fix_date,omitempty"`          // (Опционально) Дата фиксации реестра для выплаты купона
	PayOneBond      *MoneyValue `json:"pay_one_bond,omitempty"`      // Выплата на одну облигацию
	CouponType      CouponType  `json:"coupon_type,omitempty"`       // Тип купона
	CouponStartDate *time.Time  `json:"coupon_start_date,omitempty"` // Начало купонного периода
	CouponEndDate   *time.Time  `json:"coupon_end_date,omitempty"`   // Окончание купонного периода
	CouponPeriod    int32       `json:"coupon_period,omitempty"`     // Купонный период в днях
}

// Операция начисления купонов.
type AccruedInterest struct {
	Date         *time.Time `json:"date,omitempty"`          // Дата и время выплаты в часовом поясе UTC
	Value        *Quotation `json:"value,omitempty"`         // Величина выплаты
	ValuePercent *Quotation `json:"value_percent,omitempty"` // Величина выплаты в процентах от номинала
	Nominal      *Quotation `json:"nominal,omitempty"`       // Номинал облигации
}
