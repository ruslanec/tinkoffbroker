package domain

import "time"

//Текущий статус заявки (поручения)
type OrderExecutionReportStatus int32

const (
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_UNSPECIFIED   OrderExecutionReportStatus = 0
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_FILL          OrderExecutionReportStatus = 1 //Исполнена
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_REJECTED      OrderExecutionReportStatus = 2 //Отклонена
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_CANCELLED     OrderExecutionReportStatus = 3 //Отменена пользователем
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_NEW           OrderExecutionReportStatus = 4 //Новая
	OrderExecutionReportStatus_EXECUTION_REPORT_STATUS_PARTIALLYFILL OrderExecutionReportStatus = 5 //Частично исполнена
)

//Направление операции
type OrderDirection int32

const (
	OrderDirection_ORDER_DIRECTION_UNSPECIFIED OrderDirection = 0 //Значение не указано
	OrderDirection_ORDER_DIRECTION_BUY         OrderDirection = 1 //Покупка
	OrderDirection_ORDER_DIRECTION_SELL        OrderDirection = 2 //Продажа
)

//Тип заявки
type OrderType int32

const (
	OrderType_ORDER_TYPE_UNSPECIFIED OrderType = 0 //Значение не указано
	OrderType_ORDER_TYPE_LIMIT       OrderType = 1 //Лимитная
	OrderType_ORDER_TYPE_MARKET      OrderType = 2 //Рыночная
)

//Сделки в рамках торгового поручения.
type OrderStage struct {
	Price    *MoneyValue `json:"price,omitempty"`    //Цена.
	Quantity int64       `json:"quantity,omitempty"` //Количество лотов.
	TradeId  string      `json:"trade_id,omitempty"` //Идентификатор торговой операции.
}

//Информация о торговом поручении.
type OrderState struct {
	OrderId               string                     `json:"order_id,omitempty"`                //Идентификатор заявки.
	ExecutionReportStatus OrderExecutionReportStatus `json:"execution_report_status,omitempty"` //Текущий статус заявки.
	LotsRequested         int64                      `json:"lots_requested,omitempty"`          //Запрошено лотов.
	LotsExecuted          int64                      `json:"lots_executed,omitempty"`           //Исполнено лотов.
	InitialOrderPrice     *MoneyValue                `json:"initial_order_price,omitempty"`     //Начальная цена заявки. Произведение количества запрошенных лотов на цену.
	ExecutedOrderPrice    *MoneyValue                `json:"executed_order_price,omitempty"`    //Исполненная цена заявки. Произведение средней цены покупки на количество лотов.
	TotalOrderAmount      *MoneyValue                `json:"total_order_amount,omitempty"`      //Итоговая стоимость заявки, включающая все комиссии.
	AveragePositionPrice  *MoneyValue                `json:"average_position_price,omitempty"`  //Средняя цена позиции по сделке.
	InitialCommission     *MoneyValue                `json:"initial_commission,omitempty"`      //Начальная комиссия. Комиссия, рассчитанная на момент подачи заявки.
	ExecutedCommission    *MoneyValue                `json:"executed_commission,omitempty"`     //Фактическая комиссия по итогам исполнения заявки.
	Figi                  string                     `json:"figi,omitempty"`                    //Figi-идентификатор инструмента.
	Direction             OrderDirection             `json:"direction,omitempty"`               //Направление заявки.
	InitialSecurityPrice  *MoneyValue                `json:"initial_security_price,omitempty"`  //Начальная цена инструмента. Цена инструмента на момент выставления заявки.
	Stages                []*OrderStage              `json:"stages,omitempty"`                  //Стадии выполнения заявки.
	ServiceCommission     *MoneyValue                `json:"service_commission,omitempty"`      //Сервисная комиссия.
	Currency              string                     `json:"currency,omitempty"`                //Валюта заявки.
	OrderType             OrderType                  `json:"order_type,omitempty"`              //Тип заявки.
	OrderDate             *time.Time                 `json:"order_date,omitempty"`              //Дата и время выставления заявки в часовом поясе UTC.
}

//Информация о выставлении поручения.
type PostOrderResponse struct {
	OrderId               string                     `json:"order_id,omitempty"`                //Идентификатор заявки.
	ExecutionReportStatus OrderExecutionReportStatus `json:"execution_report_status,omitempty"` //Текущий статус заявки.
	LotsRequested         int64                      `json:"lots_requested,omitempty"`          //Запрошено лотов.
	LotsExecuted          int64                      `json:"lots_executed,omitempty"`           //Исполнено лотов.
	InitialOrderPrice     *MoneyValue                `json:"initial_order_price,omitempty"`     //Начальная цена заявки. Произведение количества запрошенных лотов на цену.
	ExecutedOrderPrice    *MoneyValue                `json:"executed_order_price,omitempty"`    //Исполненная цена заявки. Произведение средней цены покупки на количество лотов.
	TotalOrderAmount      *MoneyValue                `json:"total_order_amount,omitempty"`      //Итоговая стоимость заявки, включающая все комиссии.
	InitialCommission     *MoneyValue                `json:"initial_commission,omitempty"`      //Начальная комиссия. Комиссия рассчитанная при выставлении заявки.
	ExecutedCommission    *MoneyValue                `json:"executed_commission,omitempty"`     //Фактическая комиссия по итогам исполнения заявки.
	AciValue              *MoneyValue                `json:"aci_value,omitempty"`               //Значение НКД (накопленного купонного дохода) на дату. Подробнее: [НКД при выставлении торговых поручений](https://tinkoff.github.io/investAPI/head-orders#coupon)
	Figi                  string                     `json:"figi,omitempty"`                    // Figi-идентификатор инструмента.
	Direction             OrderDirection             `json:"direction,omitempty"`               //Направление сделки.
	InitialSecurityPrice  *MoneyValue                `json:"initial_security_price,omitempty"`  //Начальная цена инструмента заявки.
	OrderType             OrderType                  `json:"order_type,omitempty"`              //Тип заявки.
	Message               string                     `json:"message,omitempty"`                 //Дополнительные данные об исполнении заявки.
	InitialOrderPricePt   *Quotation                 `json:"initial_order_price_pt,omitempty"`  //Начальная цена заявки в пунктах (для фьючерсов).
}
