package tinkoffbroker

import (
	"context"
	"errors"
	"time"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	INSTRUMENT_STATUS = tkf.InstrumentStatus_INSTRUMENT_STATUS_BASE // Базовый список инструментов (по умолчанию). Инструменты доступные для торговли через TINKOFF INVEST API.
)

var (
	ErrWrongArg   = errors.New("wrong argument")
	ErrEmptyField = errors.New("returned empty field")
)

// Сервис предоставления справочной информации о ценных бумагах
type instrumentsService struct {
	conn   *grpc.ClientConn
	client tkf.InstrumentsServiceClient
}

func NewInstrumentsService(conn *grpc.ClientConn) InstrumentsService {
	instrumentsServiceClient := tkf.NewInstrumentsServiceClient(conn)

	return &instrumentsService{
		conn:   conn,
		client: instrumentsServiceClient,
	}
}

// Метод получения расписания торгов торговых площадок
func (s *instrumentsService) TradingSchedules(ctx context.Context, exchange string, from, to time.Time) ([]*TradingSchedule, error) {
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

	var schedules []*TradingSchedule
	for _, v := range resp.GetExchanges() {
		var days []*TradingDay
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
			days = append(days, &TradingDay{
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
		schedules = append(schedules, &TradingSchedule{
			Exchange: v.GetExchange(),
			Days:     days,
		})
	}
	return schedules, nil
}

// Метод получения облигации по FIGI
func (s *instrumentsService) BondByFigi(ctx context.Context, figi string) (*Bond, error) {
	resp, err := s.client.BondBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return convBond(resp.GetInstrument()), nil
}

// Метод получения списка облигаций
func (s *instrumentsService) Bonds(ctx context.Context) ([]*Bond, error) {
	resp, err := s.client.Bonds(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var bonds []*Bond
	for _, v := range resp.GetInstruments() {
		bonds = append(bonds, convBond(v))
	}
	return bonds, nil
}

// Запрос купонов по облигации
func (s *instrumentsService) BondCoupons(ctx context.Context, figi string, from, to *time.Time) ([]*Coupon, error) {
	resp, err := s.client.GetBondCoupons(ctx, &tkf.GetBondCouponsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var coupons []*Coupon
	for _, v := range resp.GetEvents() {
		couponDate := v.GetCouponDate().AsTime()
		fixDate := v.GetFixDate().AsTime()
		couponStartDate := v.GetCouponStartDate().AsTime()
		couponEndDate := v.GetCouponEndDate().AsTime()
		coupons = append(coupons, &Coupon{
			Figi:            v.GetFigi(),
			CouponDate:      &couponDate,
			CouponNumber:    v.GetCouponNumber(),
			FixDate:         &fixDate,
			PayOneBond:      convMoneyValue(v.GetPayOneBond()),
			CouponType:      CouponType(v.GetCouponType()),
			CouponStartDate: &couponStartDate,
			CouponEndDate:   &couponEndDate,
			CouponPeriod:    v.GetCouponPeriod(),
		})
	}
	return coupons, nil
}

// Метод получения валюты по FIGI
func (s *instrumentsService) CurrencyByFigi(ctx context.Context, figi string) (*Currency, error) {
	resp, err := s.client.CurrencyBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return convCurrency(resp.GetInstrument()), nil
}

// Метод получения списка валют
func (s *instrumentsService) Currencies(ctx context.Context) ([]*Currency, error) {
	resp, err := s.client.Currencies(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var currencies []*Currency
	for _, v := range resp.GetInstruments() {
		currencies = append(currencies, convCurrency(v))
	}
	return currencies, nil
}

// Метод получения инвестиционного фонда по его идентификатору
func (s *instrumentsService) EtfByFigi(ctx context.Context, figi string) (*Etf, error) {
	resp, err := s.client.EtfBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return convEtf(resp.GetInstrument()), nil
}

// Метод получения списка инвестиционных фондов
func (s *instrumentsService) Etfs(ctx context.Context) ([]*Etf, error) {
	resp, err := s.client.Etfs(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var etfs []*Etf
	for _, v := range resp.GetInstruments() {
		etfs = append(etfs, convEtf(v))
	}
	return etfs, nil
}

// Метод получения фьючерса по FIGI
func (s *instrumentsService) ShareByFigi(ctx context.Context, figi string) (*Share, error) {
	resp, err := s.client.ShareBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return convShare(resp.GetInstrument()), nil
}

// Метод получения списка акций
func (s *instrumentsService) Shares(ctx context.Context) ([]*Share, error) {
	resp, err := s.client.Shares(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var shares []*Share
	for _, v := range resp.GetInstruments() {
		shares = append(shares, convShare(v))
	}

	return shares, nil
}

// Метод получения фьючерса по FIGI
func (s *instrumentsService) FutureByFigi(ctx context.Context, figi string) (*Future, error) {
	resp, err := s.client.FutureBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	return convFuture(resp.GetInstrument()), nil
}

// Метод получения списка фьючерсов
func (s *instrumentsService) Future(ctx context.Context) ([]*Future, error) {
	resp, err := s.client.Futures(ctx, &tkf.InstrumentsRequest{
		InstrumentStatus: INSTRUMENT_STATUS,
	})
	if err != nil {
		return nil, err
	}

	var futures []*Future
	for _, v := range resp.GetInstruments() {
		futures = append(futures, convFuture(v))
	}
	return futures, nil
}

// Метод получения накопленного купонного дохода по облигации
func (s *instrumentsService) AccruedInterests(ctx context.Context, figi string, from, to *time.Time) ([]*AccruedInterest, error) {
	resp, err := s.client.GetAccruedInterests(ctx, &tkf.GetAccruedInterestsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var interests []*AccruedInterest
	for _, v := range resp.GetAccruedInterests() {
		date := v.GetDate().AsTime()
		interests = append(interests, &AccruedInterest{
			Date:         &date,
			Value:        convQuotation(v.GetValue()),
			ValuePercent: convQuotation(v.GetValuePercent()),
			Nominal:      convQuotation(v.GetNominal()),
		})
	}

	return interests, nil
}

// Метод получения размера гарантийного обеспечения по фьючерсам
func (s *instrumentsService) FuturesMargin(ctx context.Context, figi string) (*FuturesMargin, error) {
	resp, err := s.client.GetFuturesMargin(ctx, &tkf.GetFuturesMarginRequest{
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}
	return &FuturesMargin{
		InitialMarginOnBuy:      convMoneyValue(resp.InitialMarginOnBuy),
		InitialMarginOnSell:     convMoneyValue(resp.InitialMarginOnSell),
		MinPriceIncrement:       convQuotation(resp.MinPriceIncrement),
		MinPriceIncrementAmount: convQuotation(resp.MinPriceIncrementAmount),
	}, nil

}

// Метод получения основной информации об инструменте
func (s *instrumentsService) InstrumentByFigi(ctx context.Context, figi string) (*Instrument, error) {
	resp, err := s.client.GetInstrumentBy(ctx, &tkf.InstrumentRequest{
		IdType: tkf.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		Id:     figi,
	})
	if err != nil {
		return nil, err
	}

	var instrument *Instrument
	if tkfInstrument := resp.GetInstrument(); tkfInstrument != nil {
		instrument = &Instrument{
			Figi:              tkfInstrument.GetFigi(),
			Ticker:            tkfInstrument.GetTicker(),
			ClassCode:         tkfInstrument.GetClassCode(),
			Isin:              tkfInstrument.GetIsin(),
			Lot:               tkfInstrument.GetLot(),
			Currency:          tkfInstrument.GetCurrency(),
			Klong:             convQuotation(tkfInstrument.GetKlong()),
			Kshort:            convQuotation(tkfInstrument.GetKshort()),
			Dlong:             convQuotation(tkfInstrument.GetDlong()),
			Dshort:            convQuotation(tkfInstrument.GetDshort()),
			DlongMin:          convQuotation(tkfInstrument.GetDlongMin()),
			DshortMin:         convQuotation(tkfInstrument.GetDlongMin()),
			ShortEnabled:      tkfInstrument.GetShortEnabledFlag(),
			Name:              tkfInstrument.GetName(),
			Exchange:          tkfInstrument.GetExchange(),
			CountryOfRisk:     tkfInstrument.GetCountryOfRisk(),
			CountryOfRiskName: tkfInstrument.GetCountryOfRiskName(),
			InstrumentType:    tkfInstrument.GetInstrumentType(),
			TradingStatus:     SecurityTradingStatus(tkfInstrument.GetTradingStatus()),
			Otc:               tkfInstrument.GetOtcFlag(),
			BuyAvailable:      tkfInstrument.GetBuyAvailableFlag(),
			SellAvailable:     tkfInstrument.GetSellAvailableFlag(),
			MinPriceIncrement: convQuotation(tkfInstrument.GetMinPriceIncrement()),
			ApiTradeAvailable: tkfInstrument.GetApiTradeAvailableFlag(),
		}
	}

	return instrument, nil
}

// Метод для получения событий выплаты дивидендов по инструменту
func (s *instrumentsService) Dividends(ctx context.Context, figi string, from, to *time.Time) ([]*Dividend, error) {
	if from.After(*to) {
		return nil, ErrInputArgument
	}
	resp, err := s.client.GetDividends(ctx, &tkf.GetDividendsRequest{
		Figi: figi,
		From: timestamppb.New(*from),
		To:   timestamppb.New(*to),
	})
	if err != nil {
		return nil, err
	}

	var dividends []*Dividend
	for _, v := range resp.GetDividends() {
		paymentDate := v.GetPaymentDate().AsTime()
		declaredDate := v.GetDeclaredDate().AsTime()
		lastBuyDate := v.GetLastBuyDate().AsTime()
		recordDate := v.GetRecordDate().AsTime()
		createdAt := v.GetCreatedAt().AsTime()

		dividends = append(dividends, &Dividend{
			DividendNet:  convMoneyValue(v.GetDividendNet()),
			PaymentDate:  &paymentDate,
			DeclaredDate: &declaredDate,
			LastBuyDate:  &lastBuyDate,
			DividendType: v.GetDividendType(),
			RecordDate:   &recordDate,
			Regularity:   v.GetRegularity(),
			ClosePrice:   convMoneyValue(v.GetClosePrice()),
			YieldValue:   convQuotation(v.GetYieldValue()),
			CreatedAt:    &createdAt,
		})
	}

	return dividends, nil
}

// Метод получения актива по его идентификатору
func (s *instrumentsService) AssetById(ctx context.Context, id string) (*AssetFull, error) {
	if id == "" {
		return nil, ErrWrongArg
	}

	resp, err := s.client.GetAssetBy(ctx, &tkf.AssetRequest{Id: id})
	if err != nil {
		return nil, err
	}

	tkfAF := resp.GetAsset()
	if tkfAF == nil {
		return nil, ErrEmptyField
	}

	deletedAt := tkfAF.GetDeletedAt().AsTime()

	// Convert tkf.AssetCurrency to AssetCurrency
	var currency *AssetCurrency
	if tkfAF.GetType() == tkf.AssetType_ASSET_TYPE_CURRENCY {
		currency = &AssetCurrency{
			BaseCurrency: tkfAF.GetCurrency().GetBaseCurrency(),
		}
	}

	// Convert tkf.AssetSecurity to AssetSecurity
	var security *AssetSecurity
	tkfAS := tkfAF.GetSecurity()
	if tkfAS != nil && tkfAF.GetType() == tkf.AssetType_ASSET_TYPE_SECURITY {
		security = convAssetSecurity(tkfAS)
	}

	// Convert []*tkf.AssetInstrument to []*AssetInstrument
	var instruments []*AssetInstrument
	for _, tkfInstrument := range tkfAF.GetInstruments() {
		instruments = append(instruments, convAssetInstrument(tkfInstrument))
	}

	return &AssetFull{
		Uid:           tkfAF.GetUid(),
		Type:          AssetType(tkfAF.GetType()),
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
		Brand:         convBrand(tkfAF.GetBrand()),
		BrCode:        tkfAF.GetBrCode(),
		BrCodeName:    tkfAF.GetBrCodeName(),
		Instruments:   instruments,
	}, nil
}

// Метод получения списка активов
func (s *instrumentsService) Assets(ctx context.Context) ([]*Asset, error) {
	resp, err := s.client.GetAssets(ctx, &tkf.AssetsRequest{})
	if err != nil {
		return nil, err
	}

	tkfAssets := resp.GetAssets()
	if tkfAssets == nil {
		return nil, ErrEmptyField
	}

	// Convert tkf.Asset to Asset
	var assets []*Asset
	for _, tkfAsset := range tkfAssets {
		// Convert []*tkf.AssetInstrument to []*AssetInstrument
		var instruments []*AssetInstrument
		for _, tkfInstrument := range tkfAsset.GetInstruments() {
			instruments = append(instruments, convAssetInstrument(tkfInstrument))
		}

		assets = append(assets, &Asset{
			Uid:         tkfAsset.GetUid(),
			Type:        AssetType(tkfAsset.GetType()),
			Name:        tkfAsset.GetName(),
			Instruments: instruments,
		})
	}

	return assets, nil
}

// Метод получения списка избранных инструментов
func (s *instrumentsService) Favorites(ctx context.Context) ([]*FavoriteInstrument, error) {
	resp, err := s.client.GetFavorites(ctx, &tkf.GetFavoritesRequest{})
	if err != nil {
		return nil, err
	}

	var favoriteInstruments []*FavoriteInstrument
	for _, tkfFInstrument := range resp.GetFavoriteInstruments() {
		favoriteInstruments = append(favoriteInstruments, convFavoriteInstrument(tkfFInstrument))
	}

	return favoriteInstruments, nil
}

// Метод редактирования списка избранных инструментов
func (s *instrumentsService) EditFavorites(ctx context.Context, figies []string, action EditFavoritesActionType) ([]*FavoriteInstrument, error) {
	if len(figies) == 0 || action == EDIT_FAVORITES_ACTION_TYPE_UNSPECIFIED {
		return nil, ErrWrongArg
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

	var favoriteInstruments []*FavoriteInstrument
	for _, tkfFInstrument := range resp.GetFavoriteInstruments() {
		favoriteInstruments = append(favoriteInstruments, convFavoriteInstrument(tkfFInstrument))
	}

	return favoriteInstruments, nil
}

// Метод получения списка стран
func (s *instrumentsService) Countries(ctx context.Context) ([]*Country, error) {
	resp, err := s.client.GetCountries(ctx, &tkf.GetCountriesRequest{})
	if err != nil {
		return nil, err
	}

	var countries []*Country
	for _, tkfCountry := range resp.GetCountries() {
		countries = append(countries, &Country{
			AlfaTwo:   tkfCountry.GetAlfaTwo(),
			AlfaThree: tkfCountry.GetAlfaThree(),
			Name:      tkfCountry.GetName(),
			NameBrief: tkfCountry.GetNameBrief(),
		})
	}
	return countries, nil
}

// Метод поиска инструмента
func (s *instrumentsService) FindInstrument(ctx context.Context, query string) ([]*InstrumentShort, error) {
	if len(query) == 0 {
		return nil, ErrInputArgument
	}

	resp, err := s.client.FindInstrument(ctx, &tkf.FindInstrumentRequest{
		Query: query,
	})
	if err != nil {
		return nil, err
	}

	var instruments []*InstrumentShort
	for _, tkfInstrument := range resp.GetInstruments() {
		instruments = append(instruments, convInstrumentShort(tkfInstrument))
	}
	return instruments, nil
}

// Метод получения списка брендов
func (s *instrumentsService) Brands(ctx context.Context) ([]*Brand, error) {
	resp, err := s.client.GetBrands(ctx, &tkf.GetBrandsRequest{})
	if err != nil {
		return nil, err
	}

	var brands []*Brand
	for _, tkfBrand := range resp.GetBrands() {
		brands = append(brands, convBrand(tkfBrand))
	}

	return brands, nil
}

// Метод получения бренда по его идентификатору
func (s *instrumentsService) BrandById(ctx context.Context, id string) (*Brand, error) {
	if len(id) == 0 {
		return nil, ErrInputArgument
	}

	resp, err := s.client.GetBrandBy(ctx, &tkf.GetBrandRequest{})
	if err != nil {
		return nil, err
	}

	return convBrand(resp), nil
}
