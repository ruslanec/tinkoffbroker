package domain

import "time"

// Интервал свечей.
type CandleInterval int32

const (
	CandleIntervalUnspecified CandleInterval = 0 // Интервал не определён.
	CandleInterval1Min        CandleInterval = 1 // 1 минута.
	CandleInterval5Min        CandleInterval = 2 // 5 минут.
	CandleInterval15Min       CandleInterval = 3 // 15 минут.
	CandleIntervalHour        CandleInterval = 4 // 1 час.
	CandleIntervalDay         CandleInterval = 5 // 1 день.
)

var (
	CandleIntervalToDuration = map[CandleInterval]time.Duration{
		CandleInterval1Min:  time.Minute,
		CandleInterval5Min:  time.Minute * 5,
		CandleInterval15Min: time.Minute * 15,
		CandleIntervalHour:  time.Hour,
		CandleIntervalDay:   time.Hour * 24,
	}

	CandleDurationToInterval = map[time.Duration]CandleInterval{
		time.Minute:      CandleInterval1Min,
		time.Minute * 5:  CandleInterval5Min,
		time.Minute * 15: CandleInterval15Min,
		time.Hour:        CandleIntervalHour,
		time.Hour * 24:   CandleIntervalDay,
	}
)

// Интервал свечи
type SubscriptionInterval int32 // TODO: заменить на CandleInterval

const (
	SubscriptionIntervalUnspecified SubscriptionInterval = 0 // Интервал свечи не определён
	SubscriptionIntervalOneMinute   SubscriptionInterval = 1 // Минутные свечи
	SubscriptionIntervalFiveMinutes SubscriptionInterval = 2 // Пятиминутные свечи
)

// Информация о свече
type Candle struct {
	Figi     string         `json:"figi,omitempty"`      // Figi-идентификатор инструмента
	DateTime *time.Time     `json:"date_time,omitempty"` // Время свечи в часовом поясе UTC
	Interval CandleInterval `json:"interval,omitempty"`  // Интервал свечи
	Open     *Quotation     `json:"open,omitempty"`      // Цена открытия за 1 лот
	High     *Quotation     `json:"high,omitempty"`      // Максимальная цена за 1 лот
	Low      *Quotation     `json:"low,omitempty"`       // Минимальная цена за 1 лот
	Close    *Quotation     `json:"close,omitempty"`     // Цена закрытия за 1 лот
	Volume   int64          `json:"volume,omitempty"`    // Объём сделок в лотах
}

// Запрос изменения статус подписки на свечи
type CandleInstrument struct {
	Figi     string               `json:"figi,omitempty"`     // Figi-идентификатор инструмента
	Interval SubscriptionInterval `json:"interval,omitempty"` // Интервал свечей
}
