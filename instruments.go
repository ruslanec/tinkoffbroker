package tinkoffbroker

import (
	"time"

	"context"
)

// Справочная информация о ценных бумагах
type Instruments interface {
	InstrumentsService
}

// Сервис предоставления справочной информации о ценных бумагах
type InstrumentsService interface {
	// Метод получения списка акций
	Shares(ctx context.Context) ([]*Share, error)
	// Метод получения фьючерса по FIGI
	ShareByFigi(ctx context.Context, figi string) (*Share, error)
	// Метод получения списка облигаций
	Bonds(ctx context.Context) (shares []*Bond, err error)
	// Метод получения облигации по FIGI
	BondByFigi(ctx context.Context, figi string) (*Bond, error)
	// Запрос купонов по облигации
	BondCoupons(ctx context.Context, figi string, from, to *time.Time) ([]*Coupon, error)
	// Метод получения накопленного купонного дохода по облигации
	AccruedInterests(ctx context.Context, figi string, from, to *time.Time) ([]*AccruedInterest, error)
	// Метод получения списка валют
	Currencies(ctx context.Context) (shares []*Currency, err error)
	// Метод получения валюты по FIGI
	CurrencyByFigi(ctx context.Context, figi string) (*Currency, error)
	// Метод получения списка инвестиционных фондов
	Etfs(ctx context.Context) ([]*Etf, error)
	// Метод получения инвестиционного фонда по FIGI
	EtfByFigi(ctx context.Context, figi string) (*Etf, error)
	// Метод получения списка фьючерсов
	Future(ctx context.Context) ([]*Future, error)
	// Метод получения фьючерса по FIGI
	FutureByFigi(ctx context.Context, figi string) (*Future, error)
	// Метод получения расписания торгов торговых площадок
	TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*TradingSchedule, error)
	// Метод получения размера гарантийного обеспечения по фьючерсам
	FuturesMargin(ctx context.Context, figi string) (*FuturesMargin, error)
	// Метод получения основной информации об инструменте
	InstrumentByFigi(ctx context.Context, figi string) (*Instrument, error)
	// Метод для получения событий выплаты дивидендов по инструменту
	Dividends(ctx context.Context, figi string, from, to *time.Time) ([]*Dividend, error)
	// Метод получения актива по его идентификатору
	AssetById(ctx context.Context, id string) (*AssetFull, error)
	// Метод получения списка активов
	Assets(ctx context.Context) ([]*Asset, error)
	// Метод получения списка избранных инструментов
	Favorites(ctx context.Context) ([]*FavoriteInstrument, error)
	// Метод редактирования списка избранных инструментов
	EditFavorites(ctx context.Context, figies []string, action EditFavoritesActionType) ([]*FavoriteInstrument, error)
	// Метод получения списка стран
	Countries(ctx context.Context) ([]*Country, error)
	// Метод поиска инструмента
	FindInstrument(ctx context.Context, query string) ([]*InstrumentShort, error)
	// Метод получения списка брендов
	Brands(ctx context.Context) ([]*Brand, error)
	// Метод получения бренда по его идентификатору
	BrandById(ctx context.Context, id string) (*Brand, error)
}
