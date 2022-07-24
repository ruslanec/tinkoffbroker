package tinkoffbroker

// Информация о стакане
type OrderBook struct {
	Figi       string     `json:"figi,omitempty"`        // Figi-идентификатор инструмента
	Depth      int32      `json:"depth,omitempty"`       // Глубина стакана
	Bids       []*Order   `json:"bids,omitempty"`        // Множество пар значений на покупку
	Asks       []*Order   `json:"asks,omitempty"`        // Множество пар значений на продажу
	LastPrice  *Quotation `json:"last_price,omitempty"`  // Цена последней сделки
	ClosePrice *Quotation `json:"close_price,omitempty"` // Цена закрытия
	LimitUp    *Quotation `json:"limit_up,omitempty"`    // Верхний лимит цены
	LimitDown  *Quotation `json:"limit_down,omitempty"`  // Нижний лимит цены
}

// Массив предложений/спроса
type Order struct {
	Price    *Quotation `json:"price,omitempty"`    // Цена за 1 лот
	Quantity int64      `json:"quantity,omitempty"` // Количество в лотах
}

// Запрос подписки на стаканы
type OrderBookInstrument struct {
	Figi  string `pjson:"figi" json:"figi,omitempty"` // Figi-идентификатор инструмента
	Depth int32  `json:"depth,omitempty"`             // Глубина стакана
}
