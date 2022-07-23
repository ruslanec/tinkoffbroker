package tinkoffbroker

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
	CouponQuantityPerYear int32                 `json:"coupon_quantity_per_year"` // Количество выплат по купонам в год.
	MaturityDate          *time.Time            `json:"maturity_date"`            // Дата погашения облигации в часовом поясе UTC.
	Nominal               *MoneyValue           `json:"nominal"`                  // Номинал облигации.
	StateRegDate          *time.Time            `json:"state_reg_date"`           // Дата выпуска облигации в часовом поясе UTC.
	PlacementDate         *time.Time            `json:"placement_date"`           // Дата размещения в часовом поясе UTC.
	PlacementPrice        *MoneyValue           `json:"placement_price"`          // Цена размещения.
	AciValue              *MoneyValue           `json:"aci_value"`                // Значение НКД (накопленного купонного дохода) на дату.
	CountryOfRisk         string                `json:"country_of_risk"`          // Код страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	CountryOfRiskName     string                `json:"country_of_risk_name"`     // Наименование страны риска, т.е. страны, в которой компания ведёт основной бизнес.
	Sector                string                `json:"sector"`                   // Сектор экономики.
	IssueKind             string                `json:"issue_kind"`               // Форма выпуска. Возможные значения: </br>**documentary** — документарная; </br>**non_documentary** — бездокументарная.
	IssueSize             int64                 `json:"issue_size"`               // Размер выпуска.
	IssueSizePlan         int64                 `json:"issue_size_plan"`          // Плановый размер выпуска.
	TradingStatus         SecurityTradingStatus `json:"trading_status"`           // Текущий режим торгов инструмента.
	OtcFlag               bool                  `json:"otc_flag"`                 // Признак внебиржевой ценной бумаги.
	BuyAvailableFlag      bool                  `json:"buy_available_flag"`       // Признак доступности для покупки.
	SellAvailableFlag     bool                  `json:"sell_available_flag"`      // Признак доступности для продажи.
	FloatingCouponFlag    bool                  `json:"floating_coupon_flag"`     // Признак облигации с плавающим купоном.
	PerpetualFlag         bool                  `json:"perpetual_flag"`           // Признак бессрочной облигации.
	AmortizationFlag      bool                  `json:"amortization_flag"`        // Признак облигации с амортизацией долга.
	MinPriceIncrement     *Quotation            `json:"min_price_increment"`      // Шаг цены.
	ApiTradeAvailableFlag bool                  `json:"api_trade_available_flag"` // Признак доступности торгов через API.
}

// Объект передачи информации о купоне облигации
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
