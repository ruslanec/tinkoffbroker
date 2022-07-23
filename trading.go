package tinkoffbroker

import "time"

// Направление сделки
type TradeDirection int32

const (
	TradeDirection_TRADE_DIRECTION_UNSPECIFIED TradeDirection = 0 // Направление сделки не определено
	TradeDirection_TRADE_DIRECTION_BUY         TradeDirection = 1 // Покупка
	TradeDirection_TRADE_DIRECTION_SELL        TradeDirection = 2 // Продажа
)

// Данные по торговой площадке
type TradingSchedule struct {
	Exchange string        `json:"exchange"` // Наименование торговой площадки.
	Days     []*TradingDay `json:"days"`     // Массив с торговыми и неторговыми днями.
}

// Информация о времени торгов
type TradingDay struct {
	Date                           *time.Time `json:"date"`                               // Дата.
	IsTradingDay                   bool       `json:"is_trading_day"`                     // Признак торгового дня на бирже.
	StartTime                      *time.Time `json:"start_time"`                         // Время начала торгов по часовому поясу UTC.
	EndTime                        *time.Time `json:"end_time"`                           // Время окончания торгов по часовому поясу UTC.
	OpeningAuctionStartTime        *time.Time `json:"opening_auction_start_time"`         // Время начала аукциона открытия в часовом поясе UTC.
	ClosingAuctionEndTime          *time.Time `json:"closing_auction_end_time"`           // Время окончания аукциона закрытия в часовом поясе UTC.
	EveningOpeningAuctionStartTime *time.Time `json:"evening_opening_auction_start_time"` // Время начала аукциона открытия вечерней сессии в часовом поясе UTC.
	EveningStartTime               *time.Time `json:"evening_start_time"`                 // Время начала вечерней сессии в часовом поясе UTC.
	EveningEndTime                 *time.Time `json:"evening_end_time"`                   // Время окончания вечерней сессии в часовом поясе UTC.
	ClearingStartTime              *time.Time `json:"clearing_start_time"`                // Время начала основного клиринга в часовом поясе UTC.
	ClearingEndTime                *time.Time `json:"clearing_end_time"`                  // Время окончания основного клиринга в часовом поясе UTC.
	PremarketStartTime             *time.Time `json:"premarket_start_time"`               // Время начала премаркета в часовом поясе UTC.
	PremarketEndTime               *time.Time `json:"premarket_end_time"`                 // Время окончания премаркета в часовом поясе UTC.
}

// Информация о торговом статусе
type InstrumentTradingStatus struct {
	Figi          string                `json:"figi"`           // Figi-идентификатор инструмента
	TradingStatus SecurityTradingStatus `json:"trading_status"` // Статус торговли инструментом
}

// Запрос подписки на поток обезличенных сделок
type TradeInstrument struct {
	Figi string `json:"figi"` // Figi-идентификатор инструмента.
}
