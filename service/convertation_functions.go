package service

import (
	"time"

	"github.com/ruslanec/tinkoffbroker/domain"
	tkf "github.com/ruslanec/tinkoffbroker/proto"
)

// Конвертация tkf.MoneyValue в MoneyValue
func ConvMoneyValue(moneyValue *tkf.MoneyValue) *domain.MoneyValue {
	return &domain.MoneyValue{
		Currency: moneyValue.GetCurrency(),
		Units:    moneyValue.GetUnits(),
		Nano:     moneyValue.GetNano(),
	}
}

// Конвертация tkf.Quotation в Quotation
func ConvQuotation(quotation *tkf.Quotation) *domain.Quotation {
	return &domain.Quotation{
		Units: quotation.GetUnits(),
		Nano:  quotation.GetNano(),
	}
}

// Конвертация tkf.PortfolioPosition в PortfolioPosition
func ConvPortfolioPosition(portfolioPosition *tkf.PortfolioPosition) *domain.PortfolioPosition {
	return &domain.PortfolioPosition{
		Figi:                     portfolioPosition.GetFigi(),
		InstrumentType:           portfolioPosition.GetInstrumentType(),
		Quantity:                 ConvQuotation(portfolioPosition.GetQuantity()),
		AveragePositionPrice:     ConvMoneyValue(portfolioPosition.GetAveragePositionPrice()),
		ExpectedYield:            ConvQuotation(portfolioPosition.GetExpectedYield()),
		CurrentNkd:               ConvMoneyValue(portfolioPosition.GetCurrentNkd()),
		AveragePositionPricePt:   ConvQuotation(portfolioPosition.GetAveragePositionPricePt()),
		CurrentPrice:             ConvMoneyValue(portfolioPosition.GetCurrentPrice()),
		AveragePositionPriceFifo: ConvMoneyValue(portfolioPosition.GetAveragePositionPriceFifo()),
		QuantityLots:             ConvQuotation(portfolioPosition.GetQuantityLots()),
	}
}

// Конвертация tkf.Operation в Operation
func ConvOperation(operation *tkf.Operation) *domain.Operation {
	date := operation.GetDate().AsTime()
	var trades []*domain.OperationTrade
	for _, v := range operation.GetTrades() {
		dt := v.GetDateTime().AsTime()
		trades = append(trades, &domain.OperationTrade{
			TradeId:  v.GetTradeId(),
			DateTime: &dt,
			Quantity: operation.GetQuantity(),
			Price:    ConvMoneyValue(v.GetPrice()),
		})
	}

	return &domain.Operation{
		Id:                operation.GetId(),
		ParentOperationId: operation.GetParentOperationId(),
		Currency:          operation.GetCurrency(),
		Payment:           ConvMoneyValue(operation.GetPayment()),
		Price:             ConvMoneyValue(operation.GetPrice()),
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

// Конвертация tkf.OrderState в OrderState
func ConvOrderState(orderState *tkf.OrderState) *domain.OrderState {
	var stages []*domain.OrderStage
	for _, v := range orderState.GetStages() {
		stages = append(stages, &domain.OrderStage{
			Price:    ConvMoneyValue(v.GetPrice()),
			Quantity: v.GetQuantity(),
			TradeId:  v.GetTradeId(),
		})
	}
	date := orderState.GetOrderDate().AsTime()
	return &domain.OrderState{
		OrderId:               orderState.GetOrderId(),
		ExecutionReportStatus: domain.OrderExecutionReportStatus(orderState.GetExecutionReportStatus()),
		LotsRequested:         orderState.GetLotsRequested(),
		LotsExecuted:          orderState.GetLotsExecuted(),
		InitialOrderPrice:     ConvMoneyValue(orderState.GetInitialOrderPrice()),
		ExecutedOrderPrice:    ConvMoneyValue(orderState.GetExecutedOrderPrice()),
		TotalOrderAmount:      ConvMoneyValue(orderState.GetTotalOrderAmount()),
		AveragePositionPrice:  ConvMoneyValue(orderState.GetAveragePositionPrice()),
		InitialCommission:     ConvMoneyValue(orderState.GetInitialCommission()),
		ExecutedCommission:    ConvMoneyValue(orderState.GetExecutedCommission()),
		Figi:                  orderState.GetFigi(),
		Direction:             domain.OrderDirection(orderState.GetDirection()),
		InitialSecurityPrice:  ConvMoneyValue(orderState.GetInitialSecurityPrice()),
		Stages:                stages,
		ServiceCommission:     ConvMoneyValue(orderState.GetServiceCommission()),
		Currency:              orderState.GetCurrency(),
		OrderType:             domain.OrderType(orderState.GetOrderType()),
		OrderDate:             &date,
	}
}

// Конвертация tkf.PostOrderResponse в PostOrderResponse
func ConvPostOrderResponse(postOrderResponse *tkf.PostOrderResponse) *domain.PostOrderResponse {
	return &domain.PostOrderResponse{
		OrderId:               postOrderResponse.GetOrderId(),
		ExecutionReportStatus: domain.OrderExecutionReportStatus(postOrderResponse.GetExecutionReportStatus()),
		LotsRequested:         postOrderResponse.GetLotsRequested(),
		LotsExecuted:          postOrderResponse.GetLotsExecuted(),
		InitialOrderPrice:     ConvMoneyValue(postOrderResponse.GetInitialOrderPrice()),
		ExecutedOrderPrice:    ConvMoneyValue(postOrderResponse.GetExecutedOrderPrice()),
		TotalOrderAmount:      ConvMoneyValue(postOrderResponse.GetTotalOrderAmount()),
		InitialCommission:     ConvMoneyValue(postOrderResponse.GetInitialCommission()),
		ExecutedCommission:    ConvMoneyValue(postOrderResponse.GetExecutedCommission()),
		AciValue:              ConvMoneyValue(postOrderResponse.GetAciValue()),
		Figi:                  postOrderResponse.GetFigi(),
		Direction:             domain.OrderDirection(postOrderResponse.GetDirection()),
		InitialSecurityPrice:  ConvMoneyValue(postOrderResponse.GetInitialSecurityPrice()),
		OrderType:             domain.OrderType(postOrderResponse.GetOrderType()),
		Message:               postOrderResponse.GetMessage(),
		InitialOrderPricePt:   ConvQuotation(postOrderResponse.GetInitialOrderPricePt()),
	}
}

// Конвертация tkf.Share в Share
func ConvShare(share *tkf.Share) *domain.Share {
	ipoDate := share.GetIpoDate().AsTime()
	return &domain.Share{
		Figi:                  share.GetFigi(),
		Ticker:                share.GetTicker(),
		ClassCode:             share.GetClassCode(),
		Isin:                  share.GetIsin(),
		Lot:                   share.GetLot(),
		Currency:              share.GetCurrency(),
		Klong:                 ConvQuotation(share.GetKlong()),
		Kshort:                ConvQuotation(share.GetKshort()),
		Dlong:                 ConvQuotation(share.GetDlong()),
		Dshort:                ConvQuotation(share.GetDshort()),
		DlongMin:              ConvQuotation(share.GetDlongMin()),
		DshortMin:             ConvQuotation(share.GetDshortMin()),
		ShortEnabledFlag:      share.GetShortEnabledFlag(),
		Name:                  share.GetName(),
		Exchange:              share.GetExchange(),
		IpoDate:               &ipoDate,
		IssueSize:             share.GetIssueSize(),
		CountryOfRisk:         share.GetCountryOfRisk(),
		CountryOfRiskName:     share.GetCountryOfRiskName(),
		Sector:                share.GetSector(),
		IssueSizePlan:         share.GetIssueSizePlan(),
		Nominal:               ConvMoneyValue(share.GetNominal()),
		TradingStatus:         domain.SecurityTradingStatus(share.GetTradingStatus()),
		OtcFlag:               share.GetOtcFlag(),
		BuyAvailableFlag:      share.GetBuyAvailableFlag(),
		SellAvailableFlag:     share.GetSellAvailableFlag(),
		DivYieldFlag:          share.GetDivYieldFlag(),
		ShareType:             domain.ShareType(share.GetShareType()),
		MinPriceIncrement:     ConvQuotation(share.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: share.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Bond в Bond
func ConvBond(bond *tkf.Bond) *domain.Bond {
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
		Klong:                 ConvQuotation(bond.GetKlong()),
		Kshort:                ConvQuotation(bond.GetKshort()),
		Dlong:                 ConvQuotation(bond.GetDlong()),
		Dshort:                ConvQuotation(bond.GetDshort()),
		DlongMin:              ConvQuotation(bond.DlongMin),
		DshortMin:             ConvQuotation(bond.GetDshortMin()),
		ShortEnabled:          bond.GetShortEnabledFlag(),
		Name:                  bond.GetName(),
		Exchange:              bond.GetExchange(),
		CouponQuantityPerYear: bond.GetCouponQuantityPerYear(),
		MaturityDate:          &maturityDate,
		Nominal:               ConvMoneyValue(bond.GetNominal()),
		StateRegDate:          &stateRegDate,
		PlacementDate:         &placementDate,
		PlacementPrice:        ConvMoneyValue(bond.GetPlacementPrice()),
		AciValue:              ConvMoneyValue(bond.GetAciValue()),
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
		MinPriceIncrement:     ConvQuotation(bond.GetMinPriceIncrement()),
		ApiTradeAvailable:     bond.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Currency в Currency
func ConvCurrency(currency *tkf.Currency) *domain.Currency {
	return &domain.Currency{
		Figi:              currency.GetFigi(),
		Ticker:            currency.GetTicker(),
		ClassCode:         currency.GetClassCode(),
		Isin:              currency.GetIsin(),
		Lot:               currency.GetLot(),
		Currency:          currency.GetCurrency(),
		Klong:             ConvQuotation(currency.GetKlong()),
		Kshort:            ConvQuotation(currency.GetKshort()),
		Dlong:             ConvQuotation(currency.GetDlong()),
		Dshort:            ConvQuotation(currency.GetDshort()),
		DlongMin:          ConvQuotation(currency.GetDlongMin()),
		DshortMin:         ConvQuotation(currency.GetDshortMin()),
		ShortEnabled:      currency.GetShortEnabledFlag(),
		Name:              currency.GetName(),
		Exchange:          currency.GetExchange(),
		Nominal:           ConvMoneyValue(currency.GetNominal()),
		CountryOfRisk:     currency.GetCountryOfRisk(),
		CountryOfRiskName: currency.GetCountryOfRiskName(),
		TradingStatus:     domain.SecurityTradingStatus(currency.GetTradingStatus()),
		Otc:               currency.GetOtcFlag(),
		SellAvailable:     currency.GetSellAvailableFlag(),
		IsoCurrencyName:   currency.GetIsoCurrencyName(),
		MinPriceIncrement: ConvQuotation(currency.GetMinPriceIncrement()),
		ApiTradeAvailable: currency.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Etf в Etf
func ConvEtf(etf *tkf.Etf) *domain.Etf {
	releasedDate := etf.ReleasedDate.AsTime()
	return &domain.Etf{
		Figi:              etf.GetFigi(),
		Ticker:            etf.GetTicker(),
		ClassCode:         etf.GetClassCode(),
		Isin:              etf.GetIsin(),
		Lot:               etf.GetLot(),
		Currency:          etf.GetCurrency(),
		Klong:             ConvQuotation(etf.GetKlong()),
		Kshort:            ConvQuotation(etf.GetKshort()),
		Dlong:             ConvQuotation(etf.GetDlong()),
		Dshort:            ConvQuotation(etf.GetDshort()),
		DlongMin:          ConvQuotation(etf.GetDlongMin()),
		DshortMin:         ConvQuotation(etf.GetDshortMin()),
		ShortEnabled:      etf.GetShortEnabledFlag(),
		Name:              etf.GetName(),
		Exchange:          etf.GetExchange(),
		FixedCommission:   ConvQuotation(etf.GetFixedCommission()),
		FocusType:         etf.GetFocusType(),
		ReleasedDate:      &releasedDate,
		NumShares:         ConvQuotation(etf.GetNumShares()),
		CountryOfRisk:     etf.GetCountryOfRisk(),
		CountryOfRiskName: etf.GetCountryOfRiskName(),
		Sector:            etf.GetSector(),
		RebalancingFreq:   etf.GetRebalancingFreq(),
		TradingStatus:     domain.SecurityTradingStatus(etf.GetTradingStatus()),
		Otc:               etf.GetOtcFlag(),
		BuyAvailable:      etf.GetBuyAvailableFlag(),
		SellAvailable:     etf.GetSellAvailableFlag(),
		MinPriceIncrement: ConvQuotation(etf.GetMinPriceIncrement()),
		ApiTradeAvailable: etf.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Future в Future
func ConvFuture(future *tkf.Future) *domain.Future {
	firstTradeDate := future.GetFirstTradeDate().AsTime()
	lastTradeDate := future.GetLastTradeDate().AsTime()
	expirationDate := future.GetExpirationDate().AsTime()
	return &domain.Future{
		Figi:              future.GetFigi(),
		Ticker:            future.GetTicker(),
		ClassCode:         future.GetClassCode(),
		Lot:               future.GetLot(),
		Currency:          future.GetCurrency(),
		Klong:             ConvQuotation(future.GetKlong()),
		Kshort:            ConvQuotation(future.GetKshort()),
		Dlong:             ConvQuotation(future.GetDlong()),
		Dshort:            ConvQuotation(future.GetDshort()),
		DlongMin:          ConvQuotation(future.GetDlongMin()),
		DshortMin:         ConvQuotation(future.GetDshortMin()),
		ShortEnabled:      future.GetShortEnabledFlag(),
		Name:              future.GetName(),
		Exchange:          future.GetExchange(),
		FirstTradeDate:    &firstTradeDate,
		LastTradeDate:     &lastTradeDate,
		FuturesType:       future.GetFuturesType(),
		AssetType:         future.GetAssetType(),
		BasicAsset:        future.GetBasicAsset(),
		BasicAssetSize:    ConvQuotation(future.GetBasicAssetSize()),
		CountryOfRisk:     future.GetCountryOfRisk(),
		CountryOfRiskName: future.GetCountryOfRiskName(),
		Sector:            future.GetSector(),
		ExpirationDate:    &expirationDate,
		TradingStatus:     domain.SecurityTradingStatus(future.GetTradingStatus()),
		Otc:               future.GetOtcFlag(),
		BuyAvailable:      future.GetBuyAvailableFlag(),
		SellAvailable:     future.GetSellAvailableFlag(),
		MinPriceIncrement: ConvQuotation(future.GetMinPriceIncrement()),
		ApiTradeAvailable: future.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.AssetSecurity в AssetSecurity
func ConvAssetSecurity(tkfAS *tkf.AssetSecurity) *domain.AssetSecurity {
	// Конвертация tkf.AssetShare в tkf.AssetShare
	tkfShare := tkfAS.GetShare()
	var share *domain.AssetShare
	if tkfAS.GetType() == "share" && tkfShare != nil {
		share = ConvAssetShare(tkfShare)
	}

	// Конвертация tkf.AssetBond в AssetBond
	tkfBond := tkfAS.GetBond()
	var bond *domain.AssetBond
	if tkfAS.GetType() == "bond" && tkfBond != nil {
		bond = ConvAssetBond(tkfBond)
	}

	// Конвертация tkf.AssetStructuredProduct в AssetStructuredProduct
	tkfSP := tkfAS.GetSp()
	var structuredProduct *domain.AssetStructuredProduct
	if tkfAS.GetType() == "sp" && tkfSP != nil {
		structuredProduct = ConvAssetStructuredProduct(tkfSP)
	}

	// Конвертация tkf.AssetEtf в AssetEtf
	tkfEtf := tkfAS.GetEtf()
	var etf *domain.AssetEtf
	if tkfAS.GetType() == "etf" && tkfEtf != nil {
		etf = ConvAssetEtf(tkfEtf)
	}

	// Конвертация tkf.AssetClearingCertificate в AssetClearingCertificate
	tkfCC := tkfAS.GetClearingCertificate()
	var clearingCertificate *domain.AssetClearingCertificate
	if tkfAS.GetType() == "clearing_certificate" && tkfCC != nil {
		clearingCertificate = &domain.AssetClearingCertificate{
			Nominal:         *ConvQuotation(tkfCC.GetNominal()),
			NominalCurrency: tkfCC.GetNominalCurrency(),
		}
	}

	return &domain.AssetSecurity{
		Isin:                tkfAS.GetIsin(),
		Type:                tkfAS.GetType(),
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
		IssueSize:          *ConvQuotation(assetShare.GetIssueSize()),
		Nominal:            *ConvQuotation(assetShare.GetNominal()),
		NominalCurrency:    assetShare.GetNominalCurrency(),
		PrimaryIndex:       assetShare.GetPrimaryIndex(),
		DividendRate:       *ConvQuotation(assetShare.GetDividendRate()),
		PreferredShareType: assetShare.GetPreferredShareType(),
		IpoDate:            &ipoDate,
		RegistryDate:       &regestryDate,
		DivYield:           assetShare.GetDivYieldFlag(),
		IssueKind:          assetShare.GetIssueKind(),
		PlacementDate:      &placementDate,
		RepresIsin:         assetShare.GetRepresIsin(),
		IssueSizePlan:      *ConvQuotation(assetShare.GetIssueSizePlan()),
		TotalFloat:         *ConvQuotation(assetShare.GetTotalFloat()),
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
		CurrentNominal:        *ConvQuotation(assetBond.GetCurrentNominal()),
		BorrowName:            assetBond.GetBorrowName(),
		IssueSize:             *ConvQuotation(assetBond.GetIssueSize()),
		Nominal:               *ConvQuotation(assetBond.GetNominal()),
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
		PlacementPrice:        *ConvQuotation(assetBond.GetPlacementPrice()),
		IssueSizePlan:         *ConvQuotation(assetBond.GetIssueSizePlan()),
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
		Nominal:         *ConvQuotation(assetSP.GetNominal()),
		NominalCurrency: assetSP.GetNominalCurrency(),
		Type:            domain.StructuredProductType(assetSP.GetType()),
		LogicPortfolio:  assetSP.GetLogicPortfolio(),
		AssetType:       domain.AssetType(assetSP.GetAssetType()),
		BasicAsset:      assetSP.GetBasicAsset(),
		SafetyBarrier:   *ConvQuotation(assetSP.GetSafetyBarrier()),
		MaturityDate:    &maturityDate,
		IssueSizePlan:   *ConvQuotation(assetSP.GetIssueSizePlan()),
		IssueSize:       *ConvQuotation(assetSP.GetIssueSize()),
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
	var rebalancingDates []*time.Time
	for _, v := range assetEtf.GetRebalancingDates() {
		tkfRT := v.AsTime()
		rebalancingDates = append(rebalancingDates, &tkfRT)
	}
	return &domain.AssetEtf{
		TotalExpense:              *ConvQuotation(assetEtf.GetTotalExpense()),
		HurdleRate:                *ConvQuotation(assetEtf.GetHurdleRate()),
		PerformanceFee:            *ConvQuotation(assetEtf.GetPerformanceFee()),
		FixedCommission:           *ConvQuotation(assetEtf.GetFixedCommission()),
		PaymentType:               assetEtf.GetPaymentType(),
		Watermark:                 assetEtf.GetWatermarkFlag(),
		BuyPremium:                *ConvQuotation(assetEtf.GetBuyPremium()),
		SellDiscount:              *ConvQuotation(assetEtf.GetSellDiscount()),
		Rebalancing:               assetEtf.GetRebalancingFlag(),
		RebalancingFreq:           assetEtf.GetRebalancingFreq(),
		ManagementType:            assetEtf.GetManagementType(),
		PrimaryIndex:              assetEtf.GetPrimaryIndex(),
		FocusType:                 assetEtf.GetFocusType(),
		Leveraged:                 assetEtf.GetLeveragedFlag(),
		NumShare:                  *ConvQuotation(assetEtf.GetNumShare()),
		Ucits:                     assetEtf.GetUcitsFlag(),
		ReleasedDate:              &releasedDate,
		Description:               assetEtf.GetDescription(),
		PrimaryIndexDescription:   assetEtf.GetPrimaryIndexDescription(),
		PrimaryIndexCompany:       assetEtf.GetPrimaryIndexCompany(),
		IndexRecoveryPeriod:       *ConvQuotation(assetEtf.GetIndexRecoveryPeriod()),
		InavCode:                  assetEtf.GetInavCode(),
		DivYield:                  assetEtf.GetDivYieldFlag(),
		ExpenseCommission:         *ConvQuotation(assetEtf.GetExpenseCommission()),
		PrimaryIndexTrackingError: *ConvQuotation(assetEtf.GetPrimaryIndexTrackingError()),
		RebalancingPlan:           assetEtf.GetRebalancingPlan(),
		TaxRate:                   assetEtf.GetTaxRate(),
		RebalancingDates:          rebalancingDates,
		IssueKind:                 assetEtf.GetIssueKind(),
		Nominal:                   *ConvQuotation(assetEtf.GetNominal()),
		NominalCurrency:           assetEtf.GetNominalCurrency(),
	}
}

// Конвертация tkf.Brand в Brand
func ConvBrand(brand *tkf.Brand) *domain.Brand {
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

// Конвертация []*tkf.AssetInstrument в []*AssetInstrument
func ConvAssetInstrument(instrument *tkf.AssetInstrument) *domain.AssetInstrument {
	if instrument == nil {
		return nil
	}

	var links []*domain.InstrumentLink
	for _, tkfLink := range instrument.GetLinks() {
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

// Конвертация tkf.FavoriteInstrument в FavoriteInstrument
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

// Конвертация tkf.InstrumentShort в InstrumentShort
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
