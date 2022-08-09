package service

import (
	"context"
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
)

// Справочная информация о ценных бумагах
type Instruments interface {
	InstrumentsService
}

// Сервис предоставления справочной информации о ценных бумагах
type InstrumentsService interface {
	// Метод получения списка акций
	Shares(ctx context.Context) ([]*domain.Share, error)
	// Метод получения фьючерса по FIGI
	ShareByFigi(ctx context.Context, figi string) (*domain.Share, error)
	// Метод получения списка облигаций
	Bonds(ctx context.Context) (shares []*domain.Bond, err error)
	// Метод получения облигации по FIGI
	BondByFigi(ctx context.Context, figi string) (*domain.Bond, error)
	// Запрос купонов по облигации
	BondCoupons(ctx context.Context, figi string, from, to time.Time) ([]*domain.Coupon, error)
	// Метод получения накопленного купонного дохода по облигации
	AccruedInterests(ctx context.Context, figi string, from, to time.Time) ([]*domain.AccruedInterest, error)
	// Метод получения списка валют
	Currencies(ctx context.Context) (shares []*domain.Currency, err error)
	// Метод получения валюты по FIGI
	CurrencyByFigi(ctx context.Context, figi string) (*domain.Currency, error)
	// Метод получения списка инвестиционных фондов
	Etfs(ctx context.Context) ([]*domain.Etf, error)
	// Метод получения инвестиционного фонда по FIGI
	EtfByFigi(ctx context.Context, figi string) (*domain.Etf, error)
	// Метод получения списка фьючерсов
	Future(ctx context.Context) ([]*domain.Future, error)
	// Метод получения фьючерса по FIGI
	FutureByFigi(ctx context.Context, figi string) (*domain.Future, error)
	// Метод получения расписания торгов торговых площадок
	TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*domain.TradingSchedule, error)
	// Метод получения размера гарантийного обеспечения по фьючерсам
	FuturesMargin(ctx context.Context, figi string) (*domain.FuturesMargin, error)
	// Метод получения основной информации об инструменте
	InstrumentByFigi(ctx context.Context, figi string) (*domain.Instrument, error)
	// Метод для получения событий выплаты дивидендов по инструменту
	Dividends(ctx context.Context, figi string, from, to time.Time) ([]*domain.Dividend, error)
	// Метод получения актива по его идентификатору
	AssetByID(ctx context.Context, id string) (*domain.AssetFull, error)
	// Метод получения списка активов
	Assets(ctx context.Context) ([]*domain.Asset, error)
	// Метод получения списка избранных инструментов
	Favorites(ctx context.Context) ([]*domain.FavoriteInstrument, error)
	// Метод редактирования списка избранных инструментов
	EditFavorites(ctx context.Context, figies []string, action domain.EditFavoritesActionType) ([]*domain.FavoriteInstrument, error)
	// Метод получения списка стран
	Countries(ctx context.Context) ([]*domain.Country, error)
	// Метод поиска инструмента
	FindInstrument(ctx context.Context, query string) ([]*domain.InstrumentShort, error)
	// Метод получения списка брендов
	Brands(ctx context.Context) ([]*domain.Brand, error)
	// Метод получения бренда по его идентификатору
	BrandByID(ctx context.Context, id string) (*domain.Brand, error)
}
