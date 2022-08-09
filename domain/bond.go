package domain

import (
	"time"
)

// Тип купонов.
type CouponType int32

const (
	CouponTypeUnspecified CouponType = 0 // Неопределенное значение
	CouponTypeConstant    CouponType = 1 // Постоянный
	CouponTypeFloating    CouponType = 2 // Плавающий
	CouponTypeDiscount    CouponType = 3 // Дисконт
	CouponTypeMortgage    CouponType = 4 // Ипотечный
	CouponTypeFix         CouponType = 5 // Фиксированный
	CouponTypeVariable    CouponType = 6 // Переменный
	CouponTypeOther       CouponType = 7 // Прочее
)

// Реальная площадка исполнения расчётов.
type RealExchange int32

const (
	RealExchangeUnspecified RealExchange = 0 // Тип не определён.
	RealExchangeMOEX        RealExchange = 1 // Московская биржа.
	RealExchangeRTS         RealExchange = 2 // Санкт-Петербургская биржа.
	RealExchangeOTC         RealExchange = 3 // Внебиржевой инструмент.
)

// Объект передачи информации об облигации.
type Bond struct {
	Figi                  string                `json:"figi,omitempty"`                     // Figi-идентификатор инструмента.
	Ticker                string                `json:"ticker,omitempty"`                   // Тикер инструмента.
	ClassCode             string                `json:"class_code,omitempty"`               // Класс-код (секция торгов).
	Isin                  string                `json:"isin,omitempty"`                     // Isin-идентификатор инструмента.
	Lot                   int32                 `json:"lot,omitempty"`                      // Лотность инструмента.
	Currency              string                `json:"currency,omitempty"`                 // Валюта расчётов.
	Klong                 *Quotation            `json:"klong,omitempty"`                    // Коэффициент ставки риска длинной позиции по инструменту.
	Kshort                *Quotation            `json:"kshort,omitempty"`                   // Коэффициент ставки риска короткой позиции по инструменту.
	Dlong                 *Quotation            `json:"dlong,omitempty"`                    // Ставка риска минимальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	Dshort                *Quotation            `json:"dshort,omitempty"`                   // Ставка риска минимальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
	DlongMin              *Quotation            `json:"dlong_min,omitempty"`                // Ставка риска начальной маржи в лонг. Подробнее: [ставка риска в лонг](https://help.tinkoff.ru/margin-trade/long/risk-rate/)
	DshortMin             *Quotation            `json:"dshort_min,omitempty"`               // Ставка риска начальной маржи в шорт. Подробнее: [ставка риска в шорт](https://help.tinkoff.ru/margin-trade/short/risk-rate/)
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
	APITradeAvailable     bool                  `json:"api_trade_available,omitempty"`      // Признак доступности торгов через API.
	UID                   string                `json:"uid,omitempty"`                      // Уникальный идентификатор инструмента.
	RealExchange          RealExchange          `json:"real_exchange,omitempty"`            // Реальная площадка исполнения расчётов.
	PositionUID           string                `json:"position_uid,omitempty"`             // Уникальный идентификатор позиции инструмента.
	ForIIS                bool                  `json:"for_iis,omitempty"`                  // Признак доступности для ИИС.
	First1MinCandleDate   *time.Time            `json:"first_1_min_candle_date,omitempty"`  // Дата первой минутной свечи.
	First1DayCandleDate   *time.Time            `json:"first_1_day_candle_date,omitempty"`  // Дата первой дневной свечи.
}

// Объект передачи информации о купоне облигации.
type Coupon struct {
	Figi            string      `json:"figi"`              // Figi-идентификатор инструмента
	CouponDate      *time.Time  `json:"coupon_date"`       // Дата события
	CouponNumber    int64       `json:"coupon_number"`     // Номер купона
	FixDate         *time.Time  `json:"fix_date"`          // (Опционально) Дата фиксации реестра для выплаты купона
	PayOneBond      *MoneyValue `json:"pay_one_bond"`      // Выплата на одну облигацию
	CouponType      CouponType  `json:"coupon_type"`       // Тип купона
	CouponStartDate *time.Time  `json:"coupon_start_date"` // Начало купонного периода
	CouponEndDate   *time.Time  `json:"coupon_end_date"`   // Окончание купонного периода
	CouponPeriod    int32       `json:"coupon_period"`     // Купонный период в днях
}

// Операция начисления купонов.
type AccruedInterest struct {
	Date         *time.Time `json:"date"`          // Дата и время выплаты в часовом поясе UTC
	Value        *Quotation `json:"value"`         // Величина выплаты
	ValuePercent *Quotation `json:"value_percent"` // Величина выплаты в процентах от номинала
	Nominal      *Quotation `json:"nominal"`       // Номинал облигации
}
