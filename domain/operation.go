package domain

import "time"

//Статус запрашиваемых операций
type OperationState int32

const (
	OperationState_OPERATION_STATE_UNSPECIFIED OperationState = 0 //Статус операции не определён
	OperationState_OPERATION_STATE_EXECUTED    OperationState = 1 //Исполнена
	OperationState_OPERATION_STATE_CANCELED    OperationState = 2 //Отменена
)

//Тип операции
type OperationType int32

const (
	OperationType_OPERATION_TYPE_UNSPECIFIED                 OperationType = 0  //Тип операции не определён
	OperationType_OPERATION_TYPE_INPUT                       OperationType = 1  //Завод денежных средств
	OperationType_OPERATION_TYPE_BOND_TAX                    OperationType = 2  //Удержание налога по купонам
	OperationType_OPERATION_TYPE_OUTPUT_SECURITIES           OperationType = 3  //Вывод ЦБ
	OperationType_OPERATION_TYPE_OVERNIGHT                   OperationType = 4  //Доход по сделке РЕПО овернайт
	OperationType_OPERATION_TYPE_TAX                         OperationType = 5  //Удержание налога
	OperationType_OPERATION_TYPE_BOND_REPAYMENT_FULL         OperationType = 6  //Полное погашение облигаций
	OperationType_OPERATION_TYPE_SELL_CARD                   OperationType = 7  //Продажа ЦБ с карты
	OperationType_OPERATION_TYPE_DIVIDEND_TAX                OperationType = 8  //Удержание налога по дивидендам
	OperationType_OPERATION_TYPE_OUTPUT                      OperationType = 9  //Вывод денежных средств
	OperationType_OPERATION_TYPE_BOND_REPAYMENT              OperationType = 10 //Частичное погашение облигаций
	OperationType_OPERATION_TYPE_TAX_CORRECTION              OperationType = 11 //Корректировка налога
	OperationType_OPERATION_TYPE_SERVICE_FEE                 OperationType = 12 //Удержание комиссии за обслуживание брокерского счёта
	OperationType_OPERATION_TYPE_BENEFIT_TAX                 OperationType = 13 //Удержание налога за материальную выгоду
	OperationType_OPERATION_TYPE_MARGIN_FEE                  OperationType = 14 //Удержание комиссии за непокрытую позицию
	OperationType_OPERATION_TYPE_BUY                         OperationType = 15 //Покупка ЦБ
	OperationType_OPERATION_TYPE_BUY_CARD                    OperationType = 16 //Покупка ЦБ с карты
	OperationType_OPERATION_TYPE_INPUT_SECURITIES            OperationType = 17 //Завод ЦБ
	OperationType_OPERATION_TYPE_SELL_MARGIN                 OperationType = 18 //Продажа в результате Margin-call
	OperationType_OPERATION_TYPE_BROKER_FEE                  OperationType = 19 //Удержание комиссии за операцию
	OperationType_OPERATION_TYPE_BUY_MARGIN                  OperationType = 20 //Покупка в результате Margin-call
	OperationType_OPERATION_TYPE_DIVIDEND                    OperationType = 21 //Выплата дивидендов
	OperationType_OPERATION_TYPE_SELL                        OperationType = 22 //Продажа ЦБ
	OperationType_OPERATION_TYPE_COUPON                      OperationType = 23 //Выплата купонов
	OperationType_OPERATION_TYPE_SUCCESS_FEE                 OperationType = 24 //Удержание комиссии SuccessFee
	OperationType_OPERATION_TYPE_DIVIDEND_TRANSFER           OperationType = 25 //Передача дивидендного дохода
	OperationType_OPERATION_TYPE_ACCRUING_VARMARGIN          OperationType = 26 //Зачисление вариационной маржи
	OperationType_OPERATION_TYPE_WRITING_OFF_VARMARGIN       OperationType = 27 //Списание вариационной маржи
	OperationType_OPERATION_TYPE_DELIVERY_BUY                OperationType = 28 //Покупка в рамках экспирации фьючерсного контракта
	OperationType_OPERATION_TYPE_DELIVERY_SELL               OperationType = 29 //Продажа в рамках экспирации фьючерсного контракта
	OperationType_OPERATION_TYPE_TRACK_MFEE                  OperationType = 30 //Комиссия за управление по счёту автоследования
	OperationType_OPERATION_TYPE_TRACK_PFEE                  OperationType = 31 //Комиссия за результат по счёту автоследования
	OperationType_OPERATION_TYPE_TAX_PROGRESSIVE             OperationType = 32 //Удержание налога по ставке 15%
	OperationType_OPERATION_TYPE_BOND_TAX_PROGRESSIVE        OperationType = 33 //Удержание налога по купонам по ставке 15%
	OperationType_OPERATION_TYPE_DIVIDEND_TAX_PROGRESSIVE    OperationType = 34 //Удержание налога по дивидендам по ставке 15%
	OperationType_OPERATION_TYPE_BENEFIT_TAX_PROGRESSIVE     OperationType = 35 //Удержание налога за материальную выгоду по ставке 15%
	OperationType_OPERATION_TYPE_TAX_CORRECTION_PROGRESSIVE  OperationType = 36 //Корректировка налога по ставке 15%
	OperationType_OPERATION_TYPE_TAX_REPO_PROGRESSIVE        OperationType = 37 //Удержание налога за возмещение по сделкам РЕПО по ставке 15%
	OperationType_OPERATION_TYPE_TAX_REPO                    OperationType = 38 //Удержание налога за возмещение по сделкам РЕПО
	OperationType_OPERATION_TYPE_TAX_REPO_HOLD               OperationType = 39 //Удержание налога по сделкам РЕПО
	OperationType_OPERATION_TYPE_TAX_REPO_REFUND             OperationType = 40 //Возврат налога по сделкам РЕПО
	OperationType_OPERATION_TYPE_TAX_REPO_HOLD_PROGRESSIVE   OperationType = 41 //Удержание налога по сделкам РЕПО по ставке 15%
	OperationType_OPERATION_TYPE_TAX_REPO_REFUND_PROGRESSIVE OperationType = 42 //Возврат налога по сделкам РЕПО по ставке 15%
	OperationType_OPERATION_TYPE_DIV_EXT                     OperationType = 43 //Выплата дивидендов на карту
	OperationType_OPERATION_TYPE_TAX_CORRECTION_COUPON       OperationType = 44 //Корректировка налога по купонам
)

//Сделка по операции.
type OperationTrade struct {
	TradeId  string      `json:"trade_id,omitempty"`  //Идентификатор сделки
	DateTime *time.Time  `json:"date_time,omitempty"` //Дата и время сделки в часовом поясе UTC
	Quantity int64       `json:"quantity,omitempty"`  //Количество инструментов
	Price    *MoneyValue `json:"price,omitempty"`     //Цена
}

//Данные по операции.
type Operation struct {
	Id                string            `json:"id,omitempty"`                  //Идентификатор операции
	ParentOperationId string            `json:"parent_operation_id,omitempty"` //Идентификатор родительской операции
	Currency          string            `json:"currency,omitempty"`            //Валюта операции
	Payment           *MoneyValue       `json:"payment,omitempty"`             //Сумма операции
	Price             *MoneyValue       `json:"price,omitempty"`               //Цена операции
	State             OperationState    `json:"state,omitempty"`               //Статус операции
	Quantity          int64             `json:"quantity,omitempty"`            //Количество лотов инструмента
	QuantityRest      int64             `json:"quantity_rest,omitempty"`       //Неисполненный остаток по сделке
	Figi              string            `json:"figi,omitempty"`                //Figi-идентификатор инструмента, связанного с операцией
	InstrumentType    string            `json:"instrument_type,omitempty"`     //Тип инструмента. Возможные значения: </br>**bond** — облигация; </br>**share** — акция; </br>**currency** — валюта; </br>**etf** — фонд; </br>**futures** — фьючерс.
	Date              *time.Time        `json:"date,omitempty"`                //Дата и время операции в формате часовом поясе UTC
	Type              string            `json:"type,omitempty"`                //Текстовое описание типа операции
	OperationType     OperationType     `json:"operation_type,omitempty"`      //Тип операции
	Trades            []*OperationTrade `json:"trades,omitempty"`              //Массив сделок
}
