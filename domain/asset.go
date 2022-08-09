package domain

import (
	"time"
)

// Тип структурной ноты
type StructuredProductType int32

const (
	StructuredProductTypeUnspecified    StructuredProductType = 0 // Тип не определён.
	StructuredProductTypeDeliverable    StructuredProductType = 1 // Поставочный.
	StructuredProductTypeNonDeliverable StructuredProductType = 2 // Беспоставочный.
)

// Тип актива
type AssetType int32

// Тип актива.
const (
	AssetTypeUnspecified AssetType = 0 // Тип не определён.
	AssetTypeCurrency    AssetType = 1 // Валюта.
	AssetTypeCommodity   AssetType = 2 // Товар.
	AssetTypeIndex       AssetType = 3 // Индекс.
	AssetTypeSecurity    AssetType = 4 // Ценная бумага.
)

// Информация об активе
type Asset struct {
	UID         string             `json:"uid,omitempty"`         // Уникальный идентификатор актива.
	Type        AssetType          `json:"type,omitempty"`        // Тип актива
	Name        string             `json:"name,omitempty"`        // Наименование актива
	Instruments []*AssetInstrument `json:"instruments,omitempty"` // Массив идентификаторов инструментов
}

// Идентификаторы инструмента
type AssetInstrument struct {
	UID            string            `json:"uid,omitempty"`             // uid идентификатор инструмента
	Figi           string            `json:"figi,omitempty"`            // figi идентификатор инструмента
	InstrumentType string            `json:"instrument_type,omitempty"` // Тип инструмента
	Ticker         string            `json:"ticker,omitempty"`          // Тикер инструмента
	ClassCode      string            `json:"class_code,omitempty"`      // Класс-код (секция торгов)
	Links          []*InstrumentLink `json:"links,omitempty"`           // Массив связанных инструментов
}

// Связь с другим инструментом
type InstrumentLink struct {
	Type          string `json:"type,omitempty"`           // Тип связи
	InstrumentUID string `json:"instrument_uid,omitempty"` // uid идентификатор связанного инструмента
}

// Данные по активу
type AssetFull struct {
	UID           string             `json:"uid,omitempty"`            // Уникальный идентификатор актива
	Type          AssetType          `json:"type,omitempty"`           // Тип актива
	Name          string             `json:"name,omitempty"`           // Наименование актива
	NameBrief     string             `json:"name_brief,omitempty"`     // Короткое наименование актива
	Description   string             `json:"description,omitempty"`    // Описание актива
	DeletedAt     *time.Time         `json:"deleted_at,omitempty"`     // Дата и время удаления актива
	RequiredTests []string           `json:"required_tests,omitempty"` // Тестирование клиентов
	Currency      *AssetCurrency     `json:"currency,omitempty"`       // Валюта. Обязательно и заполняется только для type = "ASSET_TYPE_CURRENCY".
	Security      *AssetSecurity     `json:"security,omitempty"`       // Ценная бумага. Обязательно и заполняется только для type = "ASSET_TYPE_SECURITY".
	GosRegCode    string             `json:"gos_reg_code,omitempty"`   // Номер государственной регистрации
	Cfi           string             `json:"cfi,omitempty"`            // Код CFI
	CodeNsd       string             `json:"code_nsd,omitempty"`       // Код НРД инструмента
	Status        string             `json:"status,omitempty"`         // Статус актива
	Brand         *Brand             `json:"brand,omitempty"`          // Бренд
	UpdatedAt     *time.Time         `json:"updated_at,omitempty"`     // Дата и время последнего обновления записи
	BrCode        string             `json:"br_code,omitempty"`        // Код типа ц.б. по классификации Банка России
	BrCodeName    string             `json:"br_code_name,omitempty"`   // Наименование кода типа ц.б. по классификации Банка России
	Instruments   []*AssetInstrument `json:"instruments,omitempty"`    // Массив идентификаторов инструментов
}

// Валюта
type AssetCurrency struct {
	BaseCurrency string // ISO-код валюты
}

// Ценная бумага
type AssetSecurity struct {
	Isin                string                    `json:"isin,omitempty"`                 // ISIN-идентификатор ценной бумаги
	Type                string                    `json:"type,omitempty"`                 // Тип ценной бумаги
	Share               *AssetShare               `json:"share,omitempty"`                // Акция. Заполняется только для акций (тип актива asset.type = "ASSET_TYPE_SECURITY" и security.type = share).
	Bond                *AssetBond                `json:"bond,omitempty"`                 // Облигация. Заполняется только для облигаций (тип актива asset.type = "ASSET_TYPE_SECURITY" и security.type = bond).
	Sp                  *AssetStructuredProduct   `json:"sp,omitempty"`                   // Структурная нота. Заполняется только для структурных продуктов (тип актива asset.type = "ASSET_TYPE_SECURITY" и security.type = sp).
	Etf                 *AssetEtf                 `json:"etf,omitempty"`                  // Фонд. Заполняется только для фондов (тип актива asset.type = "ASSET_TYPE_SECURITY" и security.type = etf).
	ClearingCertificate *AssetClearingCertificate `json:"clearing_certificate,omitempty"` // Клиринговый сертификат участия. Заполняется только для клиринговых сертификатов (тип актива asset.type = "ASSET_TYPE_SECURITY" и security.type = clearing_certificate).
}

// Акция
type AssetShare struct {
	Type               ShareType  `json:"type,omitempty"`                 // Тип акции
	IssueSize          Quotation  `json:"issue_size,omitempty"`           // Объем выпуска (шт.)
	Nominal            Quotation  `json:"nominal,omitempty"`              // Номинал
	NominalCurrency    string     `json:"nominal_currency,omitempty"`     // Валюта номинала
	PrimaryIndex       string     `json:"primary_index,omitempty"`        // Индекс (Bloomberg)
	DividendRate       Quotation  `json:"dividend_rate,omitempty"`        // Ставка дивиденда (для привилегированных акций)
	PreferredShareType string     `json:"preferred_share_type,omitempty"` // Тип привилегированных акций
	IpoDate            *time.Time `json:"ipo_date,omitempty"`             // Дата IPO
	RegistryDate       *time.Time `json:"registry_date,omitempty"`        // Дата регистрации
	DivYield           bool       `json:"div_yield,omitempty"`            // Признак наличия дивидендной доходности.
	IssueKind          string     `json:"issue_kind,omitempty"`           // Форма выпуска ФИ
	PlacementDate      *time.Time `json:"placement_date,omitempty"`       // Дата размещения акции
	RepresIsin         string     `json:"repres_isin,omitempty"`          // ISIN базового актива
	IssueSizePlan      Quotation  `json:"issue_size_plan,omitempty"`      // Объявленное количество шт.
	TotalFloat         Quotation  `json:"total_float,omitempty"`          // Количество акций в свободном обращении
}

// Облигация
type AssetBond struct {
	CurrentNominal        Quotation  `json:"current_nominal,omitempty"`          // Текущий номинал
	BorrowName            string     `json:"borrow_name,omitempty"`              // Наименование заемщика
	IssueSize             Quotation  `json:"issue_size,omitempty"`               // Объем эмиссии облигации (стоимость)
	Nominal               Quotation  `json:"nominal,omitempty"`                  // Номинал облигации
	NominalCurrency       string     `json:"nominal_currency,omitempty"`         // Валюта номинала
	IssueKind             string     `json:"issue_kind,omitempty"`               // Форма выпуска облигации
	InterestKind          string     `json:"interest_kind,omitempty"`            // Форма дохода облигации
	CouponQuantityPerYear int32      `json:"coupon_quantity_per_year,omitempty"` // Количество выплат в год
	IndexedNominal        bool       `json:"indexed_nominal,omitempty"`          // Признак облигации с индексируемым номиналом
	Subordinated          bool       `json:"subordinated,omitempty"`             // Признак субординированной облигации
	Collateral            bool       `json:"collateral,omitempty"`               // Признак обеспеченной облигации
	TaxFree               bool       `json:"tax_free,omitempty"`                 // Признак показывает, что купоны облигации не облагаются налогом (для mass market)
	Amortization          bool       `json:"amortization,omitempty"`             // Признак облигации с амортизацией долга
	FloatingCoupon        bool       `json:"floating_coupon,omitempty"`          // Признак облигации с плавающим купоном
	Perpetual             bool       `json:"perpetual,omitempty"`                // Признак бессрочной облигации
	MaturityDate          *time.Time `json:"maturity_date,omitempty"`            // Дата погашения облигации
	ReturnCondition       string     `json:"return_condition,omitempty"`         // Описание и условия получения дополнительного дохода
	StateRegDate          *time.Time `json:"state_reg_date,omitempty"`           // Дата выпуска облигации
	PlacementDate         *time.Time `json:"placement_date,omitempty"`           // Дата размещения облигации
	PlacementPrice        Quotation  `json:"placement_price,omitempty"`          // Цена размещения облигации
	IssueSizePlan         Quotation  `json:"issue_size_plan,omitempty"`          // Объявленное количество шт.
}

// Структурная нота
type AssetStructuredProduct struct {
	BorrowName      string                `json:"borrow_name,omitempty"`      // Наименование заемщика
	Nominal         Quotation             `json:"nominal,omitempty"`          // Номинал
	NominalCurrency string                `json:"nominal_currency,omitempty"` // Валюта номинала
	Type            StructuredProductType `json:"type,omitempty"`             // Тип структурной ноты
	LogicPortfolio  string                `json:"logic_portfolio,omitempty"`  // Стратегия портфеля
	AssetType       AssetType             `json:"asset_type,omitempty"`       // Тип базового актива
	BasicAsset      string                `json:"basic_asset,omitempty"`      // Вид базового актива в зависимости от типа базового актива
	SafetyBarrier   Quotation             `json:"safety_barrier,omitempty"`   // Барьер сохранности (в процентах)
	MaturityDate    *time.Time            `json:"maturity_date,omitempty"`    // Дата погашения
	IssueSizePlan   Quotation             `json:"issue_size_plan,omitempty"`  // Объявленное количество шт.
	IssueSize       Quotation             `json:"issue_size,omitempty"`       // Объем размещения
	PlacementDate   *time.Time            `json:"placement_date,omitempty"`   // Дата размещения ноты
	IssueKind       string                `json:"issue_kind,omitempty"`       // Форма выпуска
}

// Фонд.
type AssetEtf struct {
	TotalExpense              Quotation    `json:"total_expense,omitempty"`                // Суммарные расходы фонда (в %).
	HurdleRate                Quotation    `json:"hurdle_rate,omitempty"`                  // Барьерная ставка доходности после которой фонд имеет право на perfomance fee (в процентах).
	PerformanceFee            Quotation    `json:"performance_fee,omitempty"`              // Комиссия за успешные результаты фонда (в процентах).
	FixedCommission           Quotation    `json:"fixed_commission,omitempty"`             // Фиксированная комиссия за управление (в процентах).
	PaymentType               string       `json:"payment_type,omitempty"`                 // Тип распределения доходов от выплат по бумагам.
	Watermark                 bool         `json:"watermark,omitempty"`                    // Признак необходимости выхода фонда в плюс для получения комиссии.
	BuyPremium                Quotation    `json:"buy_premium,omitempty"`                  // Премия (надбавка к цене) при покупке доли в фонде (в процентах).
	SellDiscount              Quotation    `json:"sell_discount,omitempty"`                // Ставка дисконта (вычет из цены) при продаже доли в фонде (в процентах).
	Rebalancing               bool         `json:"rebalancing,omitempty"`                  // Признак ребалансируемости портфеля фонда.
	RebalancingFreq           string       `json:"rebalancing_freq,omitempty"`             // Периодичность ребалансировки.
	ManagementType            string       `json:"management_type,omitempty"`              // Тип управления.
	PrimaryIndex              string       `json:"primary_index,omitempty"`                // Индекс, который реплицирует (старается копировать) фонд.
	FocusType                 string       `json:"focus_type,omitempty"`                   // База ETF.
	Leveraged                 bool         `json:"leveraged,omitempty"`                    // Признак использования заемных активов (плечо).
	NumShare                  Quotation    `json:"num_share,omitempty"`                    // Количество акций в обращении.
	Ucits                     bool         `json:"ucits,omitempty"`                        // Признак обязательства по отчетности перед регулятором.
	ReleasedDate              *time.Time   `json:"released_date,omitempty"`                // Дата выпуска.
	Description               string       `json:"description,omitempty"`                  // Описание фонда.
	PrimaryIndexDescription   string       `json:"primary_index_description,omitempty"`    // Описание индекса, за которым следует фонд.
	PrimaryIndexCompany       string       `json:"primary_index_company,omitempty"`        // Основные компании, в которые вкладывается фонд.
	IndexRecoveryPeriod       Quotation    `json:"index_recovery_period,omitempty"`        // Срок восстановления индекса (после просадки).
	InavCode                  string       `json:"inav_code,omitempty"`                    // IVAV-код.
	DivYield                  bool         `json:"div_yield,omitempty"`                    // Признак наличия дивидендной доходности.
	ExpenseCommission         Quotation    `json:"expense_commission,omitempty"`           // Комиссия на покрытие расходов фонда (в процентах).
	PrimaryIndexTrackingError Quotation    `json:"primary_index_tracking_error,omitempty"` // Ошибка следования за индексом (в процентах).
	RebalancingPlan           string       `json:"rebalancing_plan,omitempty"`             // Плановая ребалансировка портфеля.
	TaxRate                   string       `json:"tax_rate,omitempty"`                     // Ставки налогообложения дивидендов и купонов.
	RebalancingDates          []*time.Time `json:"rebalancing_dates,omitempty"`            // Даты ребалансировок.
	IssueKind                 string       `json:"issue_kind,omitempty"`                   // Форма выпуска.
	Nominal                   Quotation    `json:"nominal,omitempty"`                      // Номинал.
	NominalCurrency           string       `json:"nominal_currency,omitempty"`             // Валюта номинала.
}

// Клиринговый сертификат участия.
type AssetClearingCertificate struct {
	Nominal         Quotation `json:"nominal,omitempty"`          // Номинал.
	NominalCurrency string    `json:"nominal_currency,omitempty"` // Валюта номинала.
}

// Бренд.
type Brand struct {
	UID               string `json:"uid,omitempty"`                  // uid идентификатор бренда.
	Name              string `json:"name,omitempty"`                 // Наименование бренда.
	Description       string `json:"description,omitempty"`          // Описание.
	Info              string `json:"info,omitempty"`                 // Информация о бренде.
	Company           string `json:"company,omitempty"`              // Компания.
	Sector            string `json:"sector,omitempty"`               // Сектор.
	CountryOfRisk     string `json:"country_of_risk,omitempty"`      // Код страны риска.
	CountryOfRiskName string `json:"country_of_risk_name,omitempty"` // Наименование страны риска.
}
