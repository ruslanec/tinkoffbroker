package tinkoffbroker

import "time"

// Тип счёта
type AccountType int32

const (
	AccountType_ACCOUNT_TYPE_UNSPECIFIED AccountType = 0 // Тип аккаунта не определён
	AccountType_ACCOUNT_TYPE_TINKOFF     AccountType = 1 // Брокерский счёт Тинькофф
	AccountType_ACCOUNT_TYPE_TINKOFF_IIS AccountType = 2 // ИИС счёт
	AccountType_ACCOUNT_TYPE_INVEST_BOX  AccountType = 3 // Инвесткопилка
)

// Статус счёта
type AccountStatus int32

const (
	AccountStatus_ACCOUNT_STATUS_UNSPECIFIED AccountStatus = 0 //Статус счёта не определён
	AccountStatus_ACCOUNT_STATUS_NEW         AccountStatus = 1 //Новый, в процессе открытия
	AccountStatus_ACCOUNT_STATUS_OPEN        AccountStatus = 2 //Открытый и активный счёт
	AccountStatus_ACCOUNT_STATUS_CLOSED      AccountStatus = 3 //Закрытый счёт
)

// Уровень доступа к счёту
type AccessLevel int32

const (
	AccessLevel_ACCOUNT_ACCESS_LEVEL_UNSPECIFIED AccessLevel = 0 // Уровень доступа не определён
	AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS AccessLevel = 1 // Полный доступ к счёту
	AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY   AccessLevel = 2 // Доступ с уровнем прав "только чтение"
	AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS   AccessLevel = 3 // Доступ отсутствует
)

// Информация о счёте
type Account struct {
	Id          string        `json:"id"`           // Идентификатор счёта
	Type        AccountType   `json:"type"`         // Тип счёта
	Name        string        `json:"name"`         // Название счёта
	Status      AccountStatus `json:"status"`       // Статус счёта
	OpenedDate  *time.Time    `json:"opened_date"`  // Дата открытия счёта в часовом поясе UTC
	ClosedDate  *time.Time    `json:"closed_date"`  // Дата закрытия счёта в часовом поясе UTC
	AccessLevel AccessLevel   `json:"access_level"` // Уровень доступа к текущему счёту (определяется токеном)
}

// Маржинальные показатели по счёту
type MarginAttributes struct {
	LiquidPortfolio       *MoneyValue `json:"liquid_portfolio"`        // Ликвидная стоимость портфеля
	StartingMargin        *MoneyValue `json:"starting_margin"`         // Начальная маржа — начальное обеспечение для совершения новой сделки
	MinimalMargin         *MoneyValue `json:"minimal_margin"`          // Минимальная маржа — это минимальное обеспечение для поддержания позиции
	FundsSufficiencyLevel *Quotation  `json:"funds_sufficiency_level"` // Уровень достаточности средств. Соотношение стоимости ликвидного портфеля к начальной марже.
	AmountOfMissingFunds  *MoneyValue `json:"amount_of_missing_funds"` // Объем недостающих средств. Разница между стартовой маржой и ликвидной стоимости портфеля.
}

//Текущие лимиты пользователя.
type UserTariff struct {
	UnaryMethodLimitsPerMinute map[string]UnaryLimit `json:"unary_limits"`  //Массив лимитов пользователя по unary-запросам
	StreamLimits               []*StreamLimit        `json:"stream_limits"` //Массив лимитов пользователей для stream-соединений
}

// Лимит unary-методов
type UnaryLimit struct {
	MaxValue     int32 `json:"max_value"`     //Максимальное количество unary-запросов в минуту
	CurrentValue int32 `json:"current_value"` //Текущее количество unary-запросов в минуту
}

// Лимит stream-соединений
type StreamLimit struct {
	Limit   int32    `json:"limit"`   //Максимальное количество stream-соединений
	Streams []string `json:"streams"` //Названия stream-методов
}
