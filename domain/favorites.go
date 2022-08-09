package domain

// Тип действия со списком избранных инструментов.
type EditFavoritesActionType int32

const (
	EditFavoritesActionTypeUnspecified EditFavoritesActionType = 0 // Тип не определён.
	EditFavoritesActionTypeAdd         EditFavoritesActionType = 1 // Добавить в список.
	EditFavoritesActionTypeDel         EditFavoritesActionType = 2 // Удалить из списка.
)

// Массив избранных инструментов.
type FavoriteInstrument struct {
	Figi              string `json:"figi,omitempty"`                // Figi-идентификатор инструмента.
	Ticker            string `json:"ticker,omitempty"`              // Тикер инструмента.
	ClassCode         string `json:"class_code,omitempty"`          // Класс-код инструмента.
	Isin              string `json:"isin,omitempty"`                // Isin-идентификатор инструмента.
	InstrumentType    string `json:"instrument_type,omitempty"`     // Тип инструмента.
	Otc               bool   `json:"otc,omitempty"`                 // Признак внебиржевой ценной бумаги.
	APITradeAvailable bool   `json:"api_trade_available,omitempty"` // Признак доступности торгов через API.
}
