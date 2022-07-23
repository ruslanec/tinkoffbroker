package tinkoffbroker

// Информация о стакане
type OrderBook struct {
	Figi       string     `json:"figi"`        // Figi-идентификатор инструмента
	Depth      int32      `json:"depth"`       // Глубина стакана
	Bids       []*Order   `json:"bids"`        // Множество пар значений на покупку
	Asks       []*Order   `json:"asks"`        // Множество пар значений на продажу
	LastPrice  *Quotation `json:"last_price"`  // Цена последней сделки
	ClosePrice *Quotation `json:"close_price"` // Цена закрытия
	LimitUp    *Quotation `json:"limit_up"`    // Верхний лимит цены
	LimitDown  *Quotation `json:"limit_down"`  // Нижний лимит цены
}

// Массив предложений/спроса
type Order struct {
	Price    *Quotation `json:"price"`    // Цена за 1 лот
	Quantity int64      `json:"quantity"` // Количество в лотах
}

// Запрос подписки на стаканы
type OrderBookInstrument struct {
	Figi  string `pjson:"figi"` // Figi-идентификатор инструмента
	Depth int32  `json:"depth"` // Глубина стакана
}
