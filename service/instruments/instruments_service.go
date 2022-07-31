package instruments

import (
	"context"
	"time"

	"github.com/ruslanec/tinkoffbroker"
	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/ruslanec/tinkoffbroker/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	INSTRUMENT_STATUS = tkf.InstrumentStatus_INSTRUMENT_STATUS_BASE // Базовый список инструментов (по умолчанию). Инструменты доступные для торговли через TINKOFF INVEST API.
)

// Сервис предоставления справочной информации о ценных бумагах
type instrumentsService struct {
	conn   *grpc.ClientConn
	client tkf.InstrumentsServiceClient
}

func NewInstrumentsService(conn *grpc.ClientConn) service.InstrumentsService {
	instrumentsServiceClient := tkf.NewInstrumentsServiceClient(conn)

	return &instrumentsService{
		conn:   conn,
		client: instrumentsServiceClient,
	}
}

// Метод получения расписания торгов торговых площадок
func (s *instrumentsService) TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*domain.TradingSchedule, error) {
	var resp *tkf.TradingSchedulesResponse
	var err error

	if exchange == "" {
		resp, err = s.client.TradingSchedules(ctx, &tkf.TradingSchedulesRequest{
			From: timestamppb.New(from),
			To:   timestamppb.New(to),
		})
	} else {
		resp, err = s.client.TradingSchedules(ctx, &tkf.TradingSchedulesRequest{
			Exchange: exchange,
			From:     timestamppb.New(from),
			To:       timestamppb.New(to),
		})
	}
	if err != nil {
		return nil, err
	}

	var schedules []*domain.TradingSchedule
	for _, v := range resp.GetExchanges() {
		var days []*domain.TradingDay
		for _, day := range v.GetDays() {
			date := day.GetDate().AsTime()
			startTime := day.GetStartTime().AsTime()
			endTime := day.GetEndTime().AsTime()
			openingAuctionStartTime := day.GetOpeningAuctionStartTime().AsTime()
			closingAuctionEndTime := day.GetClosingAuctionEndTime().AsTime()
			eveningOpeningAuctionStartTime := day.GetEveningOpeningAuctionStartTime().AsTime()
			eveningStartTime := day.GetEveningStartTime().AsTime()
			eveningEndTime := day.GetEveningEndTime().AsTime()
			clearingStartTime := day.GetClearingStartTime().AsTime()
			clearingEndTime := day.GetClearingEndTime().AsTime()
			premarketStartTime := day.GetPremarketStartTime().AsTime()
			premarketEndTime := day.GetPremarketEndTime().AsTime()
			days = append(days, &domain.TradingDay{
				Date:                           &date,
				IsTradingDay:                   day.GetIsTradingDay(),
				StartTime:                      &startTime,
				EndTime:                        &endTime,
				OpeningAuctionStartTime:        &openingAuctionStartTime,
				ClosingAuctionEndTime:          &closingAuctionEndTime,
				EveningOpeningAuctionStartTime: &eveningOpeningAuctionStartTime,
				EveningStartTime:               &eveningStartTime,
				EveningEndTime:                 &eveningEndTime,
				ClearingStartTime:              &clearingStartTime,
				ClearingEndTime:                &clearingEndTime,
				PremarketStartTime:             &premarketStartTime,
				PremarketEndTime:               &premarketEndTime,
			})
		}
		schedules = append(schedules, &domain.TradingSchedule{
			Exchange: v.GetExchange(),
			Days:     days,
		})
	}
	return schedules, nil
}

// Метод получения облигации по FIGI
func (s *instrumentsService) BondByFigi(ctx context.Context, figi string) (*domain.Bond, error) {
	resp, err := s.client.BondBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return service.ConvBond(resp.GetInstrument()), nil
}

// Метод получения списка облигаций
func (s *instrumentsService) Bonds(ctx context.Context) ([]*domain.Bond, error) {
	resp, err := s.client.Bonds(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var bonds []*domain.Bond
	for _, v := range resp.GetInstruments() {
		bonds = append(bonds, service.ConvBond(v))
	}
	return bonds, nil
}

// Запрос купонов по облигации
func (s *instrumentsService) BondCoupons(ctx context.Context, figi string, from, to *time.Time) ([]*domain.Coupon, error) {
	resp, err := s.client.GetBondCoupons(ctx, &tkf.GetBondCouponsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var coupons []*domain.Coupon
	for _, v := range resp.GetEvents() {
		couponDate := v.GetCouponDate().AsTime()
		fixDate := v.GetFixDate().AsTime()
		couponStartDate := v.GetCouponStartDate().AsTime()
		couponEndDate := v.GetCouponEndDate().AsTime()
		coupons = append(coupons, &domain.Coupon{
			Figi:            v.GetFigi(),
			CouponDate:      &couponDate,
			CouponNumber:    v.GetCouponNumber(),
			FixDate:         &fixDate,
			PayOneBond:      service.ConvMoneyValueFromTkf(v.GetPayOneBond()),
			CouponType:      domain.CouponType(v.GetCouponType()),
			CouponStartDate: &couponStartDate,
			CouponEndDate:   &couponEndDate,
			CouponPeriod:    v.GetCouponPeriod(),
		})
	}
	return coupons, nil
}

// Метод получения валюты по FIGI
func (s *instrumentsService) CurrencyByFigi(ctx context.Context, figi string) (*domain.Currency, error) {
	resp, err := s.client.CurrencyBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return service.ConvCurrency(resp.GetInstrument()), nil
}

// Метод получения списка валют
func (s *instrumentsService) Currencies(ctx context.Context) ([]*domain.Currency, error) {
	resp, err := s.client.Currencies(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var currencies []*domain.Currency
	for _, v := range resp.GetInstruments() {
		currencies = append(currencies, service.ConvCurrency(v))
	}
	return currencies, nil
}

// Метод получения инвестиционного фонда по его идентификатору
func (s *instrumentsService) EtfByFigi(ctx context.Context, figi string) (*domain.Etf, error) {
	resp, err := s.client.EtfBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return service.ConvEtf(resp.GetInstrument()), nil
}

// Метод получения списка инвестиционных фондов
func (s *instrumentsService) Etfs(ctx context.Context) ([]*domain.Etf, error) {
	resp, err := s.client.Etfs(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var etfs []*domain.Etf
	for _, v := range resp.GetInstruments() {
		etfs = append(etfs, service.ConvEtf(v))
	}
	return etfs, nil
}

// Метод получения фьючерса по FIGI
func (s *instrumentsService) ShareByFigi(ctx context.Context, figi string) (*domain.Share, error) {
	resp, err := s.client.ShareBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return service.ConvShare(resp.GetInstrument()), nil
}

// Метод получения списка акций
func (s *instrumentsService) Shares(ctx context.Context) ([]*domain.Share, error) {
	resp, err := s.client.Shares(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var shares []*domain.Share
	for _, v := range resp.GetInstruments() {
		shares = append(shares, service.ConvShare(v))
	}

	return shares, nil
}

// Метод получения фьючерса по FIGI
func (s *instrumentsService) FutureByFigi(ctx context.Context, figi string) (*domain.Future, error) {
	resp, err := s.client.FutureBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return service.ConvFuture(resp.GetInstrument()), nil
}

// Метод получения списка фьючерсов
func (s *instrumentsService) Future(ctx context.Context) ([]*domain.Future, error) {
	resp, err := s.client.Futures(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var futures []*domain.Future
	for _, v := range resp.GetInstruments() {
		futures = append(futures, service.ConvFuture(v))
	}
	return futures, nil
}

// Метод получения накопленного купонного дохода по облигации
func (s *instrumentsService) AccruedInterests(ctx context.Context, figi string, from, to *time.Time) ([]*domain.AccruedInterest, error) {
	resp, err := s.client.GetAccruedInterests(ctx, &tkf.GetAccruedInterestsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var interests []*domain.AccruedInterest
	for _, v := range resp.GetAccruedInterests() {
		date := v.GetDate().AsTime()
		interests = append(interests, &domain.AccruedInterest{
			Date:         &date,
			Value:        service.ConvQuotationFromTkf(v.GetValue()),
			ValuePercent: service.ConvQuotationFromTkf(v.GetValuePercent()),
			Nominal:      service.ConvQuotationFromTkf(v.GetNominal()),
		})
	}

	return interests, nil
}

// Метод получения размера гарантийного обеспечения по фьючерсам
func (s *instrumentsService) FuturesMargin(ctx context.Context, figi string) (*domain.FuturesMargin, error) {
	resp, err := s.client.GetFuturesMargin(ctx, &tkf.GetFuturesMarginRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}
	return &domain.FuturesMargin{
		InitialMarginOnBuy:      service.ConvMoneyValueFromTkf(resp.InitialMarginOnBuy),
		InitialMarginOnSell:     service.ConvMoneyValueFromTkf(resp.InitialMarginOnSell),
		MinPriceIncrement:       service.ConvQuotationFromTkf(resp.MinPriceIncrement),
		MinPriceIncrementAmount: service.ConvQuotationFromTkf(resp.MinPriceIncrementAmount),
	}, nil

}

// Метод получения основной информации об инструменте
func (s *instrumentsService) InstrumentByFigi(ctx context.Context, figi string) (*domain.Instrument, error) {
	resp, err := s.client.GetInstrumentBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	var instrument *domain.Instrument
	if tkfInstrument := resp.GetInstrument(); tkfInstrument != nil {
		instrument = &domain.Instrument{
			Figi:              tkfInstrument.GetFigi(),
			Ticker:            tkfInstrument.GetTicker(),
			ClassCode:         tkfInstrument.GetClassCode(),
			Isin:              tkfInstrument.GetIsin(),
			Lot:               tkfInstrument.GetLot(),
			Currency:          tkfInstrument.GetCurrency(),
			Klong:             service.ConvQuotationFromTkf(tkfInstrument.GetKlong()),
			Kshort:            service.ConvQuotationFromTkf(tkfInstrument.GetKshort()),
			Dlong:             service.ConvQuotationFromTkf(tkfInstrument.GetDlong()),
			Dshort:            service.ConvQuotationFromTkf(tkfInstrument.GetDshort()),
			DlongMin:          service.ConvQuotationFromTkf(tkfInstrument.GetDlongMin()),
			DshortMin:         service.ConvQuotationFromTkf(tkfInstrument.GetDlongMin()),
			ShortEnabled:      tkfInstrument.GetShortEnabledFlag(),
			Name:              tkfInstrument.GetName(),
			Exchange:          tkfInstrument.GetExchange(),
			CountryOfRisk:     tkfInstrument.GetCountryOfRisk(),
			CountryOfRiskName: tkfInstrument.GetCountryOfRiskName(),
			InstrumentType:    tkfInstrument.GetInstrumentType(),
			TradingStatus:     domain.SecurityTradingStatus(tkfInstrument.GetTradingStatus()),
			Otc:               tkfInstrument.GetOtcFlag(),
			BuyAvailable:      tkfInstrument.GetBuyAvailableFlag(),
			SellAvailable:     tkfInstrument.GetSellAvailableFlag(),
			MinPriceIncrement: service.ConvQuotationFromTkf(tkfInstrument.GetMinPriceIncrement()),
			ApiTradeAvailable: tkfInstrument.GetApiTradeAvailableFlag(),
		}
	}

	return instrument, nil
}

// Метод для получения событий выплаты дивидендов по инструменту
func (s *instrumentsService) Dividends(ctx context.Context, figi string, from, to *time.Time) ([]*domain.Dividend, error) {
	if from.After(*to) {
		return nil, tinkoffbroker.ErrArgTimeInterval
	}
	resp, err := s.client.GetDividends(ctx, &tkf.GetDividendsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var dividends []*domain.Dividend
	for _, v := range resp.GetDividends() {
		paymentDate := v.GetPaymentDate().AsTime()
		declaredDate := v.GetDeclaredDate().AsTime()
		lastBuyDate := v.GetLastBuyDate().AsTime()
		recordDate := v.GetRecordDate().AsTime()
		createdAt := v.GetCreatedAt().AsTime()

		dividends = append(dividends, &domain.Dividend{
			DividendNet:  service.ConvMoneyValueFromTkf(v.GetDividendNet()),
			PaymentDate:  &paymentDate,
			DeclaredDate: &declaredDate,
			LastBuyDate:  &lastBuyDate,
			DividendType: v.GetDividendType(),
			RecordDate:   &recordDate,
			Regularity:   v.GetRegularity(),
			ClosePrice:   service.ConvMoneyValueFromTkf(v.GetClosePrice()),
			YieldValue:   service.ConvQuotationFromTkf(v.GetYieldValue()),
			CreatedAt:    &createdAt,
		})
	}

	return dividends, nil
}

// Метод получения актива по его идентификатору
func (s *instrumentsService) AssetById(ctx context.Context, id string) (*domain.AssetFull, error) {
	if id == "" {
		return nil, tinkoffbroker.ErrArgEmptyID
	}

	resp, err := s.client.GetAssetBy(ctx, &tkf.AssetRequest{Id: id})
	if err != nil {
		return nil, err
	}

	tkfAF := resp.GetAsset()
	if tkfAF == nil {
		return nil, tinkoffbroker.ErrRetEmptyField
	}
	deletedAt := tkfAF.GetDeletedAt().AsTime()

	// service.Convert tkf.AssetCurrency to AssetCurrency
	var currency *domain.AssetCurrency
	if tkfAF.GetType() == tkf.AssetType_ASSET_TYPE_CURRENCY {
		currency = &domain.AssetCurrency{
			BaseCurrency: tkfAF.GetCurrency().GetBaseCurrency(),
		}
	}

	// service.Convert tkf.AssetSecurity to AssetSecurity
	var security *domain.AssetSecurity
	tkfAS := tkfAF.GetSecurity()
	if tkfAS != nil && tkfAF.GetType() == tkf.AssetType_ASSET_TYPE_SECURITY {
		security = service.ConvAssetSecurity(tkfAS)
	}

	// service.Convert []*tkf.AssetInstrument to []*AssetInstrument
	var instruments []*domain.AssetInstrument
	for _, tkfInstrument := range tkfAF.GetInstruments() {
		instruments = append(instruments, service.ConvAssetInstrument(tkfInstrument))
	}

	return &domain.AssetFull{
		Uid:           tkfAF.GetUid(),
		Type:          domain.AssetType(tkfAF.GetType()),
		Name:          tkfAF.GetName(),
		NameBrief:     tkfAF.GetNameBrief(),
		Description:   tkfAF.GetDescription(),
		DeletedAt:     &deletedAt,
		RequiredTests: tkfAF.GetRequiredTests(),
		Currency:      currency,
		Security:      security,
		GosRegCode:    tkfAF.GetGosRegCode(),
		Cfi:           tkfAF.GetCfi(),
		CodeNsd:       tkfAF.GetCodeNsd(),
		Status:        tkfAF.GetStatus(),
		Brand:         service.ConvBrand(tkfAF.GetBrand()),
		BrCode:        tkfAF.GetBrCode(),
		BrCodeName:    tkfAF.GetBrCodeName(),
		Instruments:   instruments,
	}, nil
}

// Метод получения списка активов
func (s *instrumentsService) Assets(ctx context.Context) ([]*domain.Asset, error) {
	resp, err := s.client.GetAssets(ctx, &tkf.AssetsRequest{})
	if err != nil {
		return nil, err
	}

	tkfAssets := resp.GetAssets()
	if tkfAssets == nil {
		return nil, tinkoffbroker.ErrRetEmptyField
	}

	// service.Convert tkf.Asset to Asset
	var assets []*domain.Asset
	for _, tkfAsset := range tkfAssets {
		// service.Convert []*tkf.AssetInstrument to []*AssetInstrument
		var instruments []*domain.AssetInstrument
		for _, tkfInstrument := range tkfAsset.GetInstruments() {
			instruments = append(instruments, service.ConvAssetInstrument(tkfInstrument))
		}

		assets = append(assets, &domain.Asset{
			Uid:         tkfAsset.GetUid(),
			Type:        domain.AssetType(tkfAsset.GetType()),
			Name:        tkfAsset.GetName(),
			Instruments: instruments,
		})
	}

	return assets, nil
}

// Метод получения списка избранных инструментов
func (s *instrumentsService) Favorites(ctx context.Context) ([]*domain.FavoriteInstrument, error) {
	resp, err := s.client.GetFavorites(ctx, &tkf.GetFavoritesRequest{})
	if err != nil {
		return nil, err
	}

	var favoriteInstruments []*domain.FavoriteInstrument
	for _, tkfFInstrument := range resp.GetFavoriteInstruments() {
		favoriteInstruments = append(favoriteInstruments, service.ConvFavoriteInstrument(tkfFInstrument))
	}

	return favoriteInstruments, nil
}

// Метод редактирования списка избранных инструментов
func (s *instrumentsService) EditFavorites(ctx context.Context, figies []string, action domain.EditFavoritesActionType) ([]*domain.FavoriteInstrument, error) {
	if len(figies) == 0 || action == domain.EDIT_FAVORITES_ACTION_TYPE_UNSPECIFIED {
		return nil, tinkoffbroker.ErrArgEmptyFigi
	}

	var tkfInstruments []*tkf.EditFavoritesRequestInstrument
	for _, figi := range figies {
		tkfInstruments = append(tkfInstruments, &tkf.EditFavoritesRequestInstrument{
			Figi: figi,
		})
	}

	resp, err := s.client.EditFavorites(ctx, &tkf.EditFavoritesRequest{
		Instruments: tkfInstruments,
		ActionType:  tkf.EditFavoritesActionType(action),
	})
	if err != nil {
		return nil, err
	}

	var favoriteInstruments []*domain.FavoriteInstrument
	for _, tkfFInstrument := range resp.GetFavoriteInstruments() {
		favoriteInstruments = append(favoriteInstruments, service.ConvFavoriteInstrument(tkfFInstrument))
	}

	return favoriteInstruments, nil
}

// Метод получения списка стран
func (s *instrumentsService) Countries(ctx context.Context) ([]*domain.Country, error) {
	resp, err := s.client.GetCountries(ctx, &tkf.GetCountriesRequest{})
	if err != nil {
		return nil, err
	}

	var countries []*domain.Country
	for _, tkfCountry := range resp.GetCountries() {
		countries = append(countries, &domain.Country{
			AlfaTwo:   tkfCountry.GetAlfaTwo(),
			AlfaThree: tkfCountry.GetAlfaThree(),
			Name:      tkfCountry.GetName(),
			NameBrief: tkfCountry.GetNameBrief(),
		})
	}
	return countries, nil
}

// Метод поиска инструмента
func (s *instrumentsService) FindInstrument(ctx context.Context, query string) ([]*domain.InstrumentShort, error) {
	if len(query) == 0 {
		return nil, tinkoffbroker.ErrArgEmptyQuery
	}

	resp, err := s.client.FindInstrument(ctx, &tkf.FindInstrumentRequest{
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	var instruments []*domain.InstrumentShort
	for _, tkfInstrument := range resp.GetInstruments() {
		instruments = append(instruments, service.ConvInstrumentShort(tkfInstrument))
	}
	return instruments, nil
}

// Метод получения списка брендов
func (s *instrumentsService) Brands(ctx context.Context) ([]*domain.Brand, error) {
	resp, err := s.client.GetBrands(ctx, &tkf.GetBrandsRequest{})
	if err != nil {
		return nil, err
	}

	var brands []*domain.Brand
	for _, tkfBrand := range resp.GetBrands() {
		brands = append(brands, service.ConvBrand(tkfBrand))
	}

	return brands, nil
}

// Метод получения бренда по его идентификатору
func (s *instrumentsService) BrandById(ctx context.Context, id string) (*domain.Brand, error) {
	if len(id) == 0 {
		return nil, tinkoffbroker.ErrArgEmptyID
	}

	resp, err := s.client.GetBrandBy(ctx, &tkf.GetBrandRequest{})
	if err != nil {
		return nil, err
	}

	return service.ConvBrand(resp), nil
}
