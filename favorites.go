package tinkoffbroker

// Тип действия со списком избранных инструментов.
type EditFavoritesActionType int32

const (
	EDIT_FAVORITES_ACTION_TYPE_UNSPECIFIED EditFavoritesActionType = 0 // Тип не определён.
	EDIT_FAVORITES_ACTION_TYPE_ADD         EditFavoritesActionType = 1 // Добавить в список.
	EDIT_FAVORITES_ACTION_TYPE_DEL         EditFavoritesActionType = 2 // Удалить из списка.
)

// Массив избранных инструментов.
type FavoriteInstrument struct {
	Figi              string `json:"figi,omitempty"`                // Figi-идентификатор инструмента.
	Ticker            string `json:"ticker,omitempty"`              // Тикер инструмента.
	ClassCode         string `json:"class_code,omitempty"`          // Класс-код инструмента.
	Isin              string `json:"isin,omitempty"`                // Isin-идентификатор инструмента.
	InstrumentType    string `json:"instrument_type,omitempty"`     // Тип инструмента.
	Otc               bool   `json:"otc,omitempty"`                 // Признак внебиржевой ценной бумаги.
	ApiTradeAvailable bool   `json:"api_trade_available,omitempty"` // Признак доступности торгов через API.
}
