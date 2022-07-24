package tinkoffbroker

import "time"

// Информация о цене
type LastPrice struct {
	Figi  string     `json:"figi,omitempty"`  // Идентификатор инструмента
	Price *Quotation `json:"price,omitempty"` // Последняя цена за один инструмент. Для получения стоимости лота требуется умножить на лотность инструмента
	Time  *time.Time `json:"time,omitempty"`  // Время получения последней цены в часовом поясе UTC по времени биржи
}

// Запрос подписки на последнюю цену
type LastPriceInstrument struct {
	Figi string `json:"figi,omitempty"` // Figi-идентификатор инструмента
}
