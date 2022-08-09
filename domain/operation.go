package domain

import "time"

// Статус запрашиваемых операций
type OperationState int32

const (
	OperationStateUnspecidied OperationState = 0 // Статус операции не определён
	OperationStateExecuted    OperationState = 1 // Исполнена
	OperationStateCanceled    OperationState = 2 // Отменена
)

// Тип операции
type OperationType int32

const (
	OperationTypeUnspecidied              OperationType = 0  // Тип операции не определён
	OperationTypeInput                    OperationType = 1  // Завод денежных средств
	OperationTypeBondTax                  OperationType = 2  // Удержание налога по купонам
	OperationTypeOutputSecurities         OperationType = 3  // Вывод ЦБ
	OperationTypeOvernight                OperationType = 4  // Доход по сделке РЕПО овернайт
	OperationTypeTax                      OperationType = 5  // Удержание налога
	OperationTypeBondRepaymentFull        OperationType = 6  // Полное погашение облигаций
	OperationTypeSellCard                 OperationType = 7  // Продажа ЦБ с карты
	OperationTypeDividendTax              OperationType = 8  // Удержание налога по дивидендам
	OperationTypeOutput                   OperationType = 9  // Вывод денежных средств
	OperationTypeBondRepayment            OperationType = 10 // Частичное погашение облигаций
	OperationTypeTaxCorrection            OperationType = 11 // Корректировка налога
	OperationTypeErviceFee                OperationType = 12 // Удержание комиссии за обслуживание брокерского счёта
	OperationTypeBenefitTax               OperationType = 13 // Удержание налога за материальную выгоду
	OperationTypeMarginFee                OperationType = 14 // Удержание комиссии за непокрытую позицию
	OperationTypeBuy                      OperationType = 15 // Покупка ЦБ
	OperationTypeBuyCard                  OperationType = 16 // Покупка ЦБ с карты
	OperationTypeInputSecurities          OperationType = 17 // Завод ЦБ
	OperationTypeSellMargin               OperationType = 18 // Продажа в результате Margin-call
	OperationTypeBrokerFee                OperationType = 19 // Удержание комиссии за операцию
	OperationTypeBuyMargin                OperationType = 20 // Покупка в результате Margin-call
	OperationTypeDividend                 OperationType = 21 // Выплата дивидендов
	OperationTypeSell                     OperationType = 22 // Продажа ЦБ
	OperationTypeCoupon                   OperationType = 23 // Выплата купонов
	OperationTypeSuccessFee               OperationType = 24 // Удержание комиссии SuccessFee
	OperationTypeDividendTransfer         OperationType = 25 // Передача дивидендного дохода
	OperationTypeAccruingVarmargin        OperationType = 26 // Зачисление вариационной маржи
	OperationTypeWritingOffVarmargin      OperationType = 27 // Списание вариационной маржи
	OperationTypeDeliveryBuy              OperationType = 28 // Покупка в рамках экспирации фьючерсного контракта
	OperationTypeDeliverySell             OperationType = 29 // Продажа в рамках экспирации фьючерсного контракта
	OperationTypeTrackMfee                OperationType = 30 // Комиссия за управление по счёту автоследования
	OperationTypeTrackPfee                OperationType = 31 // Комиссия за результат по счёту автоследования
	OperationTypeTaxProgressive           OperationType = 32 // Удержание налога по ставке 15%
	OperationTypeBondTaxProgressive       OperationType = 33 // Удержание налога по купонам по ставке 15%
	OperationTypeDividendTaxProgressive   OperationType = 34 // Удержание налога по дивидендам по ставке 15%
	OperationTypeBenefitTaxProgressive    OperationType = 35 // Удержание налога за материальную выгоду по ставке 15%
	OperationTypeTaxCorrectionProgressive OperationType = 36 // Корректировка налога по ставке 15%
	OperationTypeRepoProgressive          OperationType = 37 // Удержание налога за возмещение по сделкам РЕПО по ставке 15%
	OperationTypeTaxRepo                  OperationType = 38 // Удержание налога за возмещение по сделкам РЕПО
	OperationTypeTaxRepoHold              OperationType = 39 // Удержание налога по сделкам РЕПО
	OperationTypeTaxRepoRefund            OperationType = 40 // Возврат налога по сделкам РЕПО
	OperationTypeTaxRepoHoldProgressive   OperationType = 41 // Удержание налога по сделкам РЕПО по ставке 15%
	OperationTypeTaxRepoRefundProgressive OperationType = 42 // Возврат налога по сделкам РЕПО по ставке 15%
	OperationTypeDivExt                   OperationType = 43 // Выплата дивидендов на карту
	OperationTypeTaxCorrectionCoupon      OperationType = 44 // Корректировка налога по купонам
)

// Сделка по операции.
type OperationTrade struct {
	TradeID  string      `json:"trade_id,omitempty"`  // Идентификатор сделки
	DateTime *time.Time  `json:"date_time,omitempty"` // Дата и время сделки в часовом поясе UTC
	Quantity int64       `json:"quantity,omitempty"`  // Количество инструментов
	Price    *MoneyValue `json:"price,omitempty"`     // Цена
}

// Данные по операции.
type Operation struct {
	ID                string            `json:"id,omitempty"`                  // Идентификатор операции
	ParentOperationID string            `json:"parent_operation_id,omitempty"` // Идентификатор родительской операции
	Currency          string            `json:"currency,omitempty"`            // Валюта операции
	Payment           *MoneyValue       `json:"payment,omitempty"`             // Сумма операции
	Price             *MoneyValue       `json:"price,omitempty"`               // Цена операции
	State             OperationState    `json:"state,omitempty"`               // Статус операции
	Quantity          int64             `json:"quantity,omitempty"`            // Количество лотов инструмента
	QuantityRest      int64             `json:"quantity_rest,omitempty"`       // Неисполненный остаток по сделке
	Figi              string            `json:"figi,omitempty"`                // Figi-идентификатор инструмента, связанного с операцией
	InstrumentType    string            `json:"instrument_type,omitempty"`     // Тип инструмента. Возможные значения: </br>**bond** — облигация; </br>**share** — акция; </br>**currency** — валюта; </br>**etf** — фонд; </br>**futures** — фьючерс.
	Date              *time.Time        `json:"date,omitempty"`                // Дата и время операции в формате часовом поясе UTC
	Type              string            `json:"type,omitempty"`                // Текстовое описание типа операции
	OperationType     OperationType     `json:"operation_type,omitempty"`      // Тип операции
	Trades            []*OperationTrade `json:"trades,omitempty"`              // Массив сделок
}
