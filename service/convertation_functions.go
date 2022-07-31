package service

import (
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Конвертация tkf.MoneyValue в domain,MoneyValue
func ConvMoneyValueFromTkf(moneyValue *tkf.MoneyValue) *domain.MoneyValue {
	if moneyValue == nil {
		return &domain.MoneyValue{
			Currency: moneyValue.GetCurrency(),
			Value:    decimal.Zero,
		}
	}

	return &domain.MoneyValue{
		Currency: moneyValue.GetCurrency(),
		Value:    toDecimal(moneyValue.GetUnits(), moneyValue.GetNano()),
	}
}

// Конвертация domain,MoneyValue в tkf.MoneyValue
func ConvMoneyValueToTkf(moneyValue *domain.MoneyValue) *tkf.MoneyValue {
	if moneyValue == nil {
		return &tkf.MoneyValue{
			Currency: moneyValue.GetCurrency(),
		}
	}

	units, nano := fromDecimal(moneyValue.Value)

	return &tkf.MoneyValue{
		Currency: moneyValue.GetCurrency(),
		Units:    units,
		Nano:     nano,
	}
}

// Конвертация tkf.Quotation в domain. Quotation
func ConvQuotationFromTkf(quotation *tkf.Quotation) *domain.Quotation {
	if quotation == nil {
		return &domain.Quotation{
			Value: decimal.Zero,
		}
	}

	return &domain.Quotation{
		Value: toDecimal(quotation.GetUnits(), quotation.GetNano()),
	}
}

// Конвертация domain.Quotation в tkf.Quotation
func ConvQuotationToTkf(quotation *domain.Quotation) *tkf.Quotation {
	if quotation == nil {
		return &tkf.Quotation{}
	}

	units, nano := fromDecimal(quotation.Value)

	return &tkf.Quotation{
		Units: units,
		Nano:  nano,
	}
}

func toDecimal(units int64, nano int32) decimal.Decimal {
	if units == 0 && nano == 0 {
		return decimal.Zero
	}

	if units < 0 {
		if nano > 0 {
			nano = -nano
		}
	}

	if units > 0 {
		if nano < 0 {
			nano = -nano
		}
	}

	value := decimal.NewFromInt(units)
	fractional := decimal.New(int64(nano), -9)
	return value.Add(fractional)
}

func fromDecimal(value decimal.Decimal) (units int64, nano int32) {
	fractional := value.Mod(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(1_000_000_000)).IntPart()
	return value.IntPart(), int32(fractional)
}

// Конвертация tkf.PortfolioPosition в domain.PortfolioPosition
func ConvPortfolioPosition(portfolioPosition *tkf.PortfolioPosition) *domain.PortfolioPosition {
	if portfolioPosition == nil {
		return nil
	}

	return &domain.PortfolioPosition{
		Figi:                     portfolioPosition.GetFigi(),
		InstrumentType:           portfolioPosition.GetInstrumentType(),
		Quantity:                 ConvQuotationFromTkf(portfolioPosition.GetQuantity()),
		AveragePositionPrice:     ConvMoneyValueFromTkf(portfolioPosition.GetAveragePositionPrice()),
		ExpectedYield:            ConvQuotationFromTkf(portfolioPosition.GetExpectedYield()),
		CurrentNkd:               ConvMoneyValueFromTkf(portfolioPosition.GetCurrentNkd()),
		AveragePositionPricePt:   ConvQuotationFromTkf(portfolioPosition.GetAveragePositionPricePt()),
		CurrentPrice:             ConvMoneyValueFromTkf(portfolioPosition.GetCurrentPrice()),
		AveragePositionPriceFifo: ConvMoneyValueFromTkf(portfolioPosition.GetAveragePositionPriceFifo()),
		QuantityLots:             ConvQuotationFromTkf(portfolioPosition.GetQuantityLots()),
	}
}

// Конвертация tkf.Operation в domain.Operation
func ConvOperation(operation *tkf.Operation) *domain.Operation {
	if operation == nil {
		return nil
	}

	date := operation.GetDate().AsTime()

	tkfTrades := operation.GetTrades()
	trades := make([]*domain.OperationTrade, 0, len(tkfTrades))
	for _, tkfTrade := range tkfTrades {
		dt := tkfTrade.GetDateTime().AsTime()
		trades = append(trades, &domain.OperationTrade{
			TradeId:  tkfTrade.GetTradeId(),
			DateTime: &dt,
			Quantity: tkfTrade.GetQuantity(),
			Price:    ConvMoneyValueFromTkf(tkfTrade.GetPrice()),
		})
	}

	return &domain.Operation{
		Id:                operation.GetId(),
		ParentOperationId: operation.GetParentOperationId(),
		Currency:          operation.GetCurrency(),
		Payment:           ConvMoneyValueFromTkf(operation.GetPayment()),
		Price:             ConvMoneyValueFromTkf(operation.GetPrice()),
		State:             domain.OperationState(operation.GetState()),
		Quantity:          operation.GetQuantity(),
		QuantityRest:      operation.GetQuantityRest(),
		Figi:              operation.GetFigi(),
		InstrumentType:    operation.GetInstrumentType(),
		Date:              &date,
		Type:              operation.GetType(),
		OperationType:     domain.OperationType(operation.GetOperationType()),
		Trades:            trades,
	}

}

// Конвертация tkf.OrderState в domain.OrderState
func ConvOrderState(orderState *tkf.OrderState) *domain.OrderState {
	if orderState == nil {
		return nil
	}

	tkfStages := orderState.GetStages()
	stages := make([]*domain.OrderStage, 0, len(tkfStages))
	for _, tkfStage := range tkfStages {
		stages = append(stages, &domain.OrderStage{
			Price:    ConvMoneyValueFromTkf(tkfStage.GetPrice()),
			Quantity: tkfStage.GetQuantity(),
			TradeId:  tkfStage.GetTradeId(),
		})
	}

	date := orderState.GetOrderDate().AsTime()

	return &domain.OrderState{
		OrderId:               orderState.GetOrderId(),
		ExecutionReportStatus: domain.OrderExecutionReportStatus(orderState.GetExecutionReportStatus()),
		LotsRequested:         orderState.GetLotsRequested(),
		LotsExecuted:          orderState.GetLotsExecuted(),
		InitialOrderPrice:     ConvMoneyValueFromTkf(orderState.GetInitialOrderPrice()),
		ExecutedOrderPrice:    ConvMoneyValueFromTkf(orderState.GetExecutedOrderPrice()),
		TotalOrderAmount:      ConvMoneyValueFromTkf(orderState.GetTotalOrderAmount()),
		AveragePositionPrice:  ConvMoneyValueFromTkf(orderState.GetAveragePositionPrice()),
		InitialCommission:     ConvMoneyValueFromTkf(orderState.GetInitialCommission()),
		ExecutedCommission:    ConvMoneyValueFromTkf(orderState.GetExecutedCommission()),
		Figi:                  orderState.GetFigi(),
		Direction:             domain.OrderDirection(orderState.GetDirection()),
		InitialSecurityPrice:  ConvMoneyValueFromTkf(orderState.GetInitialSecurityPrice()),
		Stages:                stages,
		ServiceCommission:     ConvMoneyValueFromTkf(orderState.GetServiceCommission()),
		Currency:              orderState.GetCurrency(),
		OrderType:             domain.OrderType(orderState.GetOrderType()),
		OrderDate:             &date,
	}
}

// Конвертация tkf.PostOrderResponse в domain.PostOrderResponse
func ConvPostOrderResponse(postOrderResponse *tkf.PostOrderResponse) *domain.PostOrderResponse { // TODO Избавиться от response
	if postOrderResponse == nil {
		return nil
	}

	return &domain.PostOrderResponse{
		OrderId:               postOrderResponse.GetOrderId(),
		ExecutionReportStatus: domain.OrderExecutionReportStatus(postOrderResponse.GetExecutionReportStatus()),
		LotsRequested:         postOrderResponse.GetLotsRequested(),
		LotsExecuted:          postOrderResponse.GetLotsExecuted(),
		InitialOrderPrice:     ConvMoneyValueFromTkf(postOrderResponse.GetInitialOrderPrice()),
		ExecutedOrderPrice:    ConvMoneyValueFromTkf(postOrderResponse.GetExecutedOrderPrice()),
		TotalOrderAmount:      ConvMoneyValueFromTkf(postOrderResponse.GetTotalOrderAmount()),
		InitialCommission:     ConvMoneyValueFromTkf(postOrderResponse.GetInitialCommission()),
		ExecutedCommission:    ConvMoneyValueFromTkf(postOrderResponse.GetExecutedCommission()),
		AciValue:              ConvMoneyValueFromTkf(postOrderResponse.GetAciValue()),
		Figi:                  postOrderResponse.GetFigi(),
		Direction:             domain.OrderDirection(postOrderResponse.GetDirection()),
		InitialSecurityPrice:  ConvMoneyValueFromTkf(postOrderResponse.GetInitialSecurityPrice()),
		OrderType:             domain.OrderType(postOrderResponse.GetOrderType()),
		Message:               postOrderResponse.GetMessage(),
		InitialOrderPricePt:   ConvQuotationFromTkf(postOrderResponse.GetInitialOrderPricePt()),
	}
}

// Конвертация tkf.Share в domain.Share
func ConvShare(share *tkf.Share) *domain.Share {
	if share == nil {
		return nil
	}

	ipoDate := share.GetIpoDate().AsTime()

	return &domain.Share{
		Figi:                  share.GetFigi(),
		Ticker:                share.GetTicker(),
		ClassCode:             share.GetClassCode(),
		Isin:                  share.GetIsin(),
		Lot:                   share.GetLot(),
		Currency:              share.GetCurrency(),
		Klong:                 ConvQuotationFromTkf(share.GetKlong()),
		Kshort:                ConvQuotationFromTkf(share.GetKshort()),
		Dlong:                 ConvQuotationFromTkf(share.GetDlong()),
		Dshort:                ConvQuotationFromTkf(share.GetDshort()),
		DlongMin:              ConvQuotationFromTkf(share.GetDlongMin()),
		DshortMin:             ConvQuotationFromTkf(share.GetDshortMin()),
		ShortEnabledFlag:      share.GetShortEnabledFlag(),
		Name:                  share.GetName(),
		Exchange:              share.GetExchange(),
		IpoDate:               &ipoDate,
		IssueSize:             share.GetIssueSize(),
		CountryOfRisk:         share.GetCountryOfRisk(),
		CountryOfRiskName:     share.GetCountryOfRiskName(),
		Sector:                share.GetSector(),
		IssueSizePlan:         share.GetIssueSizePlan(),
		Nominal:               ConvMoneyValueFromTkf(share.GetNominal()),
		TradingStatus:         domain.SecurityTradingStatus(share.GetTradingStatus()),
		OtcFlag:               share.GetOtcFlag(),
		BuyAvailableFlag:      share.GetBuyAvailableFlag(),
		SellAvailableFlag:     share.GetSellAvailableFlag(),
		DivYieldFlag:          share.GetDivYieldFlag(),
		ShareType:             domain.ShareType(share.GetShareType()),
		MinPriceIncrement:     ConvQuotationFromTkf(share.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: share.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Bond в domain.Bond
func ConvBond(bond *tkf.Bond) *domain.Bond {
	if bond == nil {
		return nil
	}

	maturityDate := bond.GetMaturityDate().AsTime()
	stateRegDate := bond.GetStateRegDate().AsTime()
	placementDate := bond.GetPlacementDate().AsTime()

	return &domain.Bond{
		Figi:                  bond.GetFigi(),
		Ticker:                bond.GetTicker(),
		ClassCode:             bond.GetClassCode(),
		Isin:                  bond.GetIsin(),
		Lot:                   bond.GetLot(),
		Currency:              bond.GetCurrency(),
		Klong:                 ConvQuotationFromTkf(bond.GetKlong()),
		Kshort:                ConvQuotationFromTkf(bond.GetKshort()),
		Dlong:                 ConvQuotationFromTkf(bond.GetDlong()),
		Dshort:                ConvQuotationFromTkf(bond.GetDshort()),
		DlongMin:              ConvQuotationFromTkf(bond.GetDlongMin()),
		DshortMin:             ConvQuotationFromTkf(bond.GetDshortMin()),
		ShortEnabled:          bond.GetShortEnabledFlag(),
		Name:                  bond.GetName(),
		Exchange:              bond.GetExchange(),
		CouponQuantityPerYear: bond.GetCouponQuantityPerYear(),
		MaturityDate:          &maturityDate,
		Nominal:               ConvMoneyValueFromTkf(bond.GetNominal()),
		StateRegDate:          &stateRegDate,
		PlacementDate:         &placementDate,
		PlacementPrice:        ConvMoneyValueFromTkf(bond.GetPlacementPrice()),
		AciValue:              ConvMoneyValueFromTkf(bond.GetAciValue()),
		CountryOfRisk:         bond.GetCountryOfRisk(),
		CountryOfRiskName:     bond.GetCountryOfRiskName(),
		Sector:                bond.GetSector(),
		IssueKind:             bond.GetIssueKind(),
		IssueSize:             bond.GetIssueSize(),
		TradingStatus:         domain.SecurityTradingStatus(bond.GetTradingStatus()),
		Otc:                   bond.GetOtcFlag(),
		BuyAvailable:          bond.GetBuyAvailableFlag(),
		SellAvailable:         bond.GetSellAvailableFlag(),
		FloatingCoupon:        bond.GetFloatingCouponFlag(),
		Perpetual:             bond.GetPerpetualFlag(),
		Amortization:          bond.GetAmortizationFlag(),
		MinPriceIncrement:     ConvQuotationFromTkf(bond.GetMinPriceIncrement()),
		ApiTradeAvailable:     bond.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Currency в domain.Currency
func ConvCurrency(currency *tkf.Currency) *domain.Currency {
	if currency == nil {
		return nil
	}

	return &domain.Currency{
		Figi:              currency.GetFigi(),
		Ticker:            currency.GetTicker(),
		ClassCode:         currency.GetClassCode(),
		Isin:              currency.GetIsin(),
		Lot:               currency.GetLot(),
		Currency:          currency.GetCurrency(),
		Klong:             ConvQuotationFromTkf(currency.GetKlong()),
		Kshort:            ConvQuotationFromTkf(currency.GetKshort()),
		Dlong:             ConvQuotationFromTkf(currency.GetDlong()),
		Dshort:            ConvQuotationFromTkf(currency.GetDshort()),
		DlongMin:          ConvQuotationFromTkf(currency.GetDlongMin()),
		DshortMin:         ConvQuotationFromTkf(currency.GetDshortMin()),
		ShortEnabled:      currency.GetShortEnabledFlag(),
		Name:              currency.GetName(),
		Exchange:          currency.GetExchange(),
		Nominal:           ConvMoneyValueFromTkf(currency.GetNominal()),
		CountryOfRisk:     currency.GetCountryOfRisk(),
		CountryOfRiskName: currency.GetCountryOfRiskName(),
		TradingStatus:     domain.SecurityTradingStatus(currency.GetTradingStatus()),
		Otc:               currency.GetOtcFlag(),
		SellAvailable:     currency.GetSellAvailableFlag(),
		IsoCurrencyName:   currency.GetIsoCurrencyName(),
		MinPriceIncrement: ConvQuotationFromTkf(currency.GetMinPriceIncrement()),
		ApiTradeAvailable: currency.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Etf в Etf
func ConvEtf(etf *tkf.Etf) *domain.Etf {
	if etf == nil {
		return nil
	}

	releasedDate := etf.ReleasedDate.AsTime()

	return &domain.Etf{
		Figi:              etf.GetFigi(),
		Ticker:            etf.GetTicker(),
		ClassCode:         etf.GetClassCode(),
		Isin:              etf.GetIsin(),
		Lot:               etf.GetLot(),
		Currency:          etf.GetCurrency(),
		Klong:             ConvQuotationFromTkf(etf.GetKlong()),
		Kshort:            ConvQuotationFromTkf(etf.GetKshort()),
		Dlong:             ConvQuotationFromTkf(etf.GetDlong()),
		Dshort:            ConvQuotationFromTkf(etf.GetDshort()),
		DlongMin:          ConvQuotationFromTkf(etf.GetDlongMin()),
		DshortMin:         ConvQuotationFromTkf(etf.GetDshortMin()),
		ShortEnabled:      etf.GetShortEnabledFlag(),
		Name:              etf.GetName(),
		Exchange:          etf.GetExchange(),
		FixedCommission:   ConvQuotationFromTkf(etf.GetFixedCommission()),
		FocusType:         etf.GetFocusType(),
		ReleasedDate:      &releasedDate,
		NumShares:         ConvQuotationFromTkf(etf.GetNumShares()),
		CountryOfRisk:     etf.GetCountryOfRisk(),
		CountryOfRiskName: etf.GetCountryOfRiskName(),
		Sector:            etf.GetSector(),
		RebalancingFreq:   etf.GetRebalancingFreq(),
		TradingStatus:     domain.SecurityTradingStatus(etf.GetTradingStatus()),
		Otc:               etf.GetOtcFlag(),
		BuyAvailable:      etf.GetBuyAvailableFlag(),
		SellAvailable:     etf.GetSellAvailableFlag(),
		MinPriceIncrement: ConvQuotationFromTkf(etf.GetMinPriceIncrement()),
		ApiTradeAvailable: etf.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Future в domain.Future
func ConvFuture(future *tkf.Future) *domain.Future {
	if future == nil {
		return nil
	}

	firstTradeDate := future.GetFirstTradeDate().AsTime()
	lastTradeDate := future.GetLastTradeDate().AsTime()
	expirationDate := future.GetExpirationDate().AsTime()

	return &domain.Future{
		Figi:              future.GetFigi(),
		Ticker:            future.GetTicker(),
		ClassCode:         future.GetClassCode(),
		Lot:               future.GetLot(),
		Currency:          future.GetCurrency(),
		Klong:             ConvQuotationFromTkf(future.GetKlong()),
		Kshort:            ConvQuotationFromTkf(future.GetKshort()),
		Dlong:             ConvQuotationFromTkf(future.GetDlong()),
		Dshort:            ConvQuotationFromTkf(future.GetDshort()),
		DlongMin:          ConvQuotationFromTkf(future.GetDlongMin()),
		DshortMin:         ConvQuotationFromTkf(future.GetDshortMin()),
		ShortEnabled:      future.GetShortEnabledFlag(),
		Name:              future.GetName(),
		Exchange:          future.GetExchange(),
		FirstTradeDate:    &firstTradeDate,
		LastTradeDate:     &lastTradeDate,
		FuturesType:       future.GetFuturesType(),
		AssetType:         future.GetAssetType(),
		BasicAsset:        future.GetBasicAsset(),
		BasicAssetSize:    ConvQuotationFromTkf(future.GetBasicAssetSize()),
		CountryOfRisk:     future.GetCountryOfRisk(),
		CountryOfRiskName: future.GetCountryOfRiskName(),
		Sector:            future.GetSector(),
		ExpirationDate:    &expirationDate,
		TradingStatus:     domain.SecurityTradingStatus(future.GetTradingStatus()),
		Otc:               future.GetOtcFlag(),
		BuyAvailable:      future.GetBuyAvailableFlag(),
		SellAvailable:     future.GetSellAvailableFlag(),
		MinPriceIncrement: ConvQuotationFromTkf(future.GetMinPriceIncrement()),
		ApiTradeAvailable: future.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.AssetSecurity в domain.AssetSecurity
func ConvAssetSecurity(assetSecurity *tkf.AssetSecurity) *domain.AssetSecurity {
	if assetSecurity == nil {
		return nil
	}

	// Конвертация tkf.AssetShare в tkf.AssetShare
	tkfShare := assetSecurity.GetShare()
	var share *domain.AssetShare
	if assetSecurity.GetType() == "share" && tkfShare != nil {
		share = ConvAssetShare(tkfShare)
	}

	// Конвертация tkf.AssetBond в AssetBond
	tkfBond := assetSecurity.GetBond()
	var bond *domain.AssetBond
	if assetSecurity.GetType() == "bond" && tkfBond != nil {
		bond = ConvAssetBond(tkfBond)
	}

	// Конвертация tkf.AssetStructuredProduct в AssetStructuredProduct
	tkfSP := assetSecurity.GetSp()
	var structuredProduct *domain.AssetStructuredProduct
	if assetSecurity.GetType() == "sp" && tkfSP != nil {
		structuredProduct = ConvAssetStructuredProduct(tkfSP)
	}

	// Конвертация tkf.AssetEtf в AssetEtf
	tkfEtf := assetSecurity.GetEtf()
	var etf *domain.AssetEtf
	if assetSecurity.GetType() == "etf" && tkfEtf != nil {
		etf = ConvAssetEtf(tkfEtf)
	}

	// Конвертация tkf.AssetClearingCertificate в AssetClearingCertificate
	tkfCC := assetSecurity.GetClearingCertificate()
	var clearingCertificate *domain.AssetClearingCertificate
	if assetSecurity.GetType() == "clearing_certificate" && tkfCC != nil {
		clearingCertificate = &domain.AssetClearingCertificate{
			Nominal:         *ConvQuotationFromTkf(tkfCC.GetNominal()),
			NominalCurrency: tkfCC.GetNominalCurrency(),
		}
	}

	return &domain.AssetSecurity{
		Isin:                assetSecurity.GetIsin(),
		Type:                assetSecurity.GetType(),
		Share:               share,
		Bond:                bond,
		Sp:                  structuredProduct,
		Etf:                 etf,
		ClearingCertificate: clearingCertificate,
	}
}

// Конвертация tkf.AssetShare в tkf.AssetShare
func ConvAssetShare(assetShare *tkf.AssetShare) *domain.AssetShare {
	if assetShare == nil {
		return nil
	}

	ipoDate := assetShare.GetIpoDate().AsTime()
	regestryDate := assetShare.GetRegistryDate().AsTime()
	placementDate := assetShare.GetPlacementDate().AsTime()

	return &domain.AssetShare{
		Type:               domain.ShareType(assetShare.GetType()),
		IssueSize:          *ConvQuotationFromTkf(assetShare.GetIssueSize()),
		Nominal:            *ConvQuotationFromTkf(assetShare.GetNominal()),
		NominalCurrency:    assetShare.GetNominalCurrency(),
		PrimaryIndex:       assetShare.GetPrimaryIndex(),
		DividendRate:       *ConvQuotationFromTkf(assetShare.GetDividendRate()),
		PreferredShareType: assetShare.GetPreferredShareType(),
		IpoDate:            &ipoDate,
		RegistryDate:       &regestryDate,
		DivYield:           assetShare.GetDivYieldFlag(),
		IssueKind:          assetShare.GetIssueKind(),
		PlacementDate:      &placementDate,
		RepresIsin:         assetShare.GetRepresIsin(),
		IssueSizePlan:      *ConvQuotationFromTkf(assetShare.GetIssueSizePlan()),
		TotalFloat:         *ConvQuotationFromTkf(assetShare.GetTotalFloat()),
	}
}

// Конвертация tkf.AssetBond в AssetBond
func ConvAssetBond(assetBond *tkf.AssetBond) *domain.AssetBond {
	if assetBond == nil {
		return nil
	}

	maturityDate := assetBond.GetMaturityDate().AsTime()
	stateRegDate := assetBond.GetStateRegDate().AsTime()
	placementDate := assetBond.GetPlacementDate().AsTime()

	return &domain.AssetBond{
		CurrentNominal:        *ConvQuotationFromTkf(assetBond.GetCurrentNominal()),
		BorrowName:            assetBond.GetBorrowName(),
		IssueSize:             *ConvQuotationFromTkf(assetBond.GetIssueSize()),
		Nominal:               *ConvQuotationFromTkf(assetBond.GetNominal()),
		NominalCurrency:       assetBond.GetNominalCurrency(),
		IssueKind:             assetBond.GetIssueKind(),
		InterestKind:          assetBond.GetInterestKind(),
		CouponQuantityPerYear: assetBond.GetCouponQuantityPerYear(),
		IndexedNominal:        assetBond.GetIndexedNominalFlag(),
		Subordinated:          assetBond.GetSubordinatedFlag(),
		Collateral:            assetBond.GetCollateralFlag(),
		TaxFree:               assetBond.GetTaxFreeFlag(),
		Amortization:          assetBond.GetAmortizationFlag(),
		FloatingCoupon:        assetBond.GetFloatingCouponFlag(),
		Perpetual:             assetBond.GetPerpetualFlag(),
		MaturityDate:          &maturityDate,
		ReturnCondition:       assetBond.GetReturnCondition(),
		StateRegDate:          &stateRegDate,
		PlacementDate:         &placementDate,
		PlacementPrice:        *ConvQuotationFromTkf(assetBond.GetPlacementPrice()),
		IssueSizePlan:         *ConvQuotationFromTkf(assetBond.GetIssueSizePlan()),
	}
}

// Конвертация tkf.AssetStructuredProduct в AssetStructuredProduct
func ConvAssetStructuredProduct(assetSP *tkf.AssetStructuredProduct) *domain.AssetStructuredProduct {
	if assetSP == nil {
		return nil
	}

	maturityDate := assetSP.GetMaturityDate().AsTime()
	placementDate := assetSP.GetPlacementDate().AsTime()

	return &domain.AssetStructuredProduct{
		BorrowName:      assetSP.GetBorrowName(),
		Nominal:         *ConvQuotationFromTkf(assetSP.GetNominal()),
		NominalCurrency: assetSP.GetNominalCurrency(),
		Type:            domain.StructuredProductType(assetSP.GetType()),
		LogicPortfolio:  assetSP.GetLogicPortfolio(),
		AssetType:       domain.AssetType(assetSP.GetAssetType()),
		BasicAsset:      assetSP.GetBasicAsset(),
		SafetyBarrier:   *ConvQuotationFromTkf(assetSP.GetSafetyBarrier()),
		MaturityDate:    &maturityDate,
		IssueSizePlan:   *ConvQuotationFromTkf(assetSP.GetIssueSizePlan()),
		IssueSize:       *ConvQuotationFromTkf(assetSP.GetIssueSize()),
		PlacementDate:   &placementDate,
		IssueKind:       assetSP.GetIssueKind(),
	}
}

// Конвертация tkf.AssetEtf в AssetEtf
func ConvAssetEtf(assetEtf *tkf.AssetEtf) *domain.AssetEtf {
	if assetEtf == nil {
		return nil
	}

	releasedDate := assetEtf.GetReleasedDate().AsTime()

	tkfRebalancingDates := assetEtf.GetRebalancingDates()
	rebalancingDates := make([]*time.Time, 0, len(tkfRebalancingDates))
	for _, v := range tkfRebalancingDates {
		tkfRT := v.AsTime()
		rebalancingDates = append(rebalancingDates, &tkfRT)
	}

	return &domain.AssetEtf{
		TotalExpense:              *ConvQuotationFromTkf(assetEtf.GetTotalExpense()),
		HurdleRate:                *ConvQuotationFromTkf(assetEtf.GetHurdleRate()),
		PerformanceFee:            *ConvQuotationFromTkf(assetEtf.GetPerformanceFee()),
		FixedCommission:           *ConvQuotationFromTkf(assetEtf.GetFixedCommission()),
		PaymentType:               assetEtf.GetPaymentType(),
		Watermark:                 assetEtf.GetWatermarkFlag(),
		BuyPremium:                *ConvQuotationFromTkf(assetEtf.GetBuyPremium()),
		SellDiscount:              *ConvQuotationFromTkf(assetEtf.GetSellDiscount()),
		Rebalancing:               assetEtf.GetRebalancingFlag(),
		RebalancingFreq:           assetEtf.GetRebalancingFreq(),
		ManagementType:            assetEtf.GetManagementType(),
		PrimaryIndex:              assetEtf.GetPrimaryIndex(),
		FocusType:                 assetEtf.GetFocusType(),
		Leveraged:                 assetEtf.GetLeveragedFlag(),
		NumShare:                  *ConvQuotationFromTkf(assetEtf.GetNumShare()),
		Ucits:                     assetEtf.GetUcitsFlag(),
		ReleasedDate:              &releasedDate,
		Description:               assetEtf.GetDescription(),
		PrimaryIndexDescription:   assetEtf.GetPrimaryIndexDescription(),
		PrimaryIndexCompany:       assetEtf.GetPrimaryIndexCompany(),
		IndexRecoveryPeriod:       *ConvQuotationFromTkf(assetEtf.GetIndexRecoveryPeriod()),
		InavCode:                  assetEtf.GetInavCode(),
		DivYield:                  assetEtf.GetDivYieldFlag(),
		ExpenseCommission:         *ConvQuotationFromTkf(assetEtf.GetExpenseCommission()),
		PrimaryIndexTrackingError: *ConvQuotationFromTkf(assetEtf.GetPrimaryIndexTrackingError()),
		RebalancingPlan:           assetEtf.GetRebalancingPlan(),
		TaxRate:                   assetEtf.GetTaxRate(),
		RebalancingDates:          rebalancingDates,
		IssueKind:                 assetEtf.GetIssueKind(),
		Nominal:                   *ConvQuotationFromTkf(assetEtf.GetNominal()),
		NominalCurrency:           assetEtf.GetNominalCurrency(),
	}
}

// Конвертация tkf.Brand в domain.Brand
func ConvBrand(brand *tkf.Brand) *domain.Brand {
	if brand == nil {
		return nil
	}

	return &domain.Brand{
		Uid:               brand.GetUid(),
		Name:              brand.GetName(),
		Description:       brand.GetDescription(),
		Info:              brand.GetInfo(),
		Company:           brand.GetCompany(),
		Sector:            brand.GetSector(),
		CountryOfRisk:     brand.GetCountryOfRisk(),
		CountryOfRiskName: brand.GetCountryOfRiskName(),
	}
}

// Конвертация []*tkf.AssetInstrument в []*domain.AssetInstrument
func ConvAssetInstrument(instrument *tkf.AssetInstrument) *domain.AssetInstrument {
	if instrument == nil {
		return nil
	}

	tkfLinks := instrument.GetLinks()
	links := make([]*domain.InstrumentLink, 0, len(tkfLinks))
	for _, tkfLink := range tkfLinks {
		link := &domain.InstrumentLink{
			Type:          tkfLink.GetType(),
			InstrumentUid: tkfLink.GetInstrumentUid(),
		}
		links = append(links, link)
	}

	return &domain.AssetInstrument{
		Uid:            instrument.GetUid(),
		Figi:           instrument.GetFigi(),
		InstrumentType: instrument.GetInstrumentType(),
		Ticker:         instrument.GetTicker(),
		ClassCode:      instrument.GetClassCode(),
		Links:          links,
	}
}

// Конвертация tkf.FavoriteInstrument в domain.FavoriteInstrument
func ConvFavoriteInstrument(instrument *tkf.FavoriteInstrument) *domain.FavoriteInstrument {
	if instrument == nil {
		return nil
	}

	return &domain.FavoriteInstrument{
		Figi:              instrument.GetFigi(),
		Ticker:            instrument.GetTicker(),
		ClassCode:         instrument.GetClassCode(),
		Isin:              instrument.GetIsin(),
		InstrumentType:    instrument.GetInstrumentType(),
		Otc:               instrument.GetOtcFlag(),
		ApiTradeAvailable: instrument.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.InstrumentShort domain. InstrumentShort
func ConvInstrumentShort(instrument *tkf.InstrumentShort) *domain.InstrumentShort {
	if instrument == nil {
		return nil
	}

	return &domain.InstrumentShort{
		Isin:              instrument.GetIsin(),
		Figi:              instrument.GetFigi(),
		Ticker:            instrument.GetTicker(),
		ClassCode:         instrument.GetClassCode(),
		InstrumentType:    instrument.GetInstrumentType(),
		Name:              instrument.GetName(),
		Uid:               instrument.GetUid(),
		PositionUid:       instrument.GetPositionUid(),
		ApiTradeAvailable: instrument.GetApiTradeAvailableFlag(),
	}
}

func ConvTimestamp(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp == nil {
		return nil
	}

	t := timestamp.AsTime()
	return &t
}
