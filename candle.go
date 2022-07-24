package tinkoffbroker

import "time"

//Интервал свечей.
type CandleInterval int32

const (
	CandleInterval_CANDLE_INTERVAL_UNSPECIFIED CandleInterval = 0 //Интервал не определён.
	CandleInterval_CANDLE_INTERVAL_1_MIN       CandleInterval = 1 //1 минута.
	CandleInterval_CANDLE_INTERVAL_5_MIN       CandleInterval = 2 //5 минут.
	CandleInterval_CANDLE_INTERVAL_15_MIN      CandleInterval = 3 //15 минут.
	CandleInterval_CANDLE_INTERVAL_HOUR        CandleInterval = 4 //1 час.
	CandleInterval_CANDLE_INTERVAL_DAY         CandleInterval = 5 //1 день.
)

var (
	CandleIntervalToDuration = map[CandleInterval]time.Duration{
		CandleInterval_CANDLE_INTERVAL_1_MIN:  time.Minute,
		CandleInterval_CANDLE_INTERVAL_5_MIN:  time.Minute * 5,
		CandleInterval_CANDLE_INTERVAL_15_MIN: time.Minute * 15,
		CandleInterval_CANDLE_INTERVAL_HOUR:   time.Hour,
		CandleInterval_CANDLE_INTERVAL_DAY:    time.Hour * 24,
	}

	CandleDurationToInterval = map[time.Duration]CandleInterval{
		time.Minute:      CandleInterval_CANDLE_INTERVAL_1_MIN,
		time.Minute * 5:  CandleInterval_CANDLE_INTERVAL_5_MIN,
		time.Minute * 15: CandleInterval_CANDLE_INTERVAL_15_MIN,
		time.Hour:        CandleInterval_CANDLE_INTERVAL_HOUR,
		time.Hour * 24:   CandleInterval_CANDLE_INTERVAL_DAY,
	}
)

// Интервал свечи
type SubscriptionInterval int32 // TODO: заменить на CandleInterval

const (
	SubscriptionInterval_SUBSCRIPTION_INTERVAL_UNSPECIFIED  SubscriptionInterval = 0 // Интервал свечи не определён
	SubscriptionInterval_SUBSCRIPTION_INTERVAL_ONE_MINUTE   SubscriptionInterval = 1 // Минутные свечи
	SubscriptionInterval_SUBSCRIPTION_INTERVAL_FIVE_MINUTES SubscriptionInterval = 2 // Пятиминутные свечи
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

//Запрос изменения статус подписки на свечи
type CandleInstrument struct {
	Figi     string               `json:"figi,omitempty"`     // Figi-идентификатор инструмента
	Interval SubscriptionInterval `json:"interval,omitempty"` // Интервал свечей
}
