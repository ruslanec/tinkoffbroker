package service

import (
	"time"

	domain "github.com/ruslanec/tinkoffbroker"
	tkf "github.com/ruslanec/tinkoffbroker/service/proto"
)

// Конвертация tkf.MoneyValue в domain.MoneyValue
func convMoneyValue(moneyValue *tkf.MoneyValue) *domain.MoneyValue {
	return &domain.MoneyValue{
		Currency: moneyValue.GetCurrency(),
		Units:    moneyValue.GetUnits(),
		Nano:     moneyValue.GetNano(),
	}
}

// Конвертация tkf.Quotation в domain.Quotation
func convQuotation(quotation *tkf.Quotation) *domain.Quotation {
	return &domain.Quotation{
		Units: quotation.GetUnits(),
		Nano:  quotation.GetNano(),
	}
}

// Конвертация tkf.PortfolioPosition в domain.PortfolioPosition
func convPortfolioPosition(portfolioPosition *tkf.PortfolioPosition) *domain.PortfolioPosition {
	return &domain.PortfolioPosition{
		Figi:                     portfolioPosition.GetFigi(),
		InstrumentType:           portfolioPosition.GetInstrumentType(),
		Quantity:                 convQuotation(portfolioPosition.GetQuantity()),
		AveragePositionPrice:     convMoneyValue(portfolioPosition.GetAveragePositionPrice()),
		ExpectedYield:            convQuotation(portfolioPosition.GetExpectedYield()),
		CurrentNkd:               convMoneyValue(portfolioPosition.GetCurrentNkd()),
		AveragePositionPricePt:   convQuotation(portfolioPosition.GetAveragePositionPricePt()),
		CurrentPrice:             convMoneyValue(portfolioPosition.GetCurrentPrice()),
		AveragePositionPriceFifo: convMoneyValue(portfolioPosition.GetAveragePositionPriceFifo()),
		QuantityLots:             convQuotation(portfolioPosition.GetQuantityLots()),
	}
}

// Конвертация tkf.Operation в domain.Operation
func convOperation(operation *tkf.Operation) *domain.Operation {
	date := operation.GetDate().AsTime()
	var trades []*domain.OperationTrade
	for _, v := range operation.GetTrades() {
		dt := v.GetDateTime().AsTime()
		trades = append(trades, &domain.OperationTrade{
			TradeId:  v.GetTradeId(),
			DateTime: &dt,
			Quantity: operation.GetQuantity(),
			Price:    convMoneyValue(v.GetPrice()),
		})
	}

	return &domain.Operation{
		Id:                operation.GetId(),
		ParentOperationId: operation.GetParentOperationId(),
		Currency:          operation.GetCurrency(),
		Payment:           convMoneyValue(operation.GetPayment()),
		Price:             convMoneyValue(operation.GetPrice()),
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
func convOrderState(orderState *tkf.OrderState) *domain.OrderState {
	var stages []*domain.OrderStage
	for _, v := range orderState.GetStages() {
		stages = append(stages, &domain.OrderStage{
			Price:    convMoneyValue(v.GetPrice()),
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
		InitialOrderPrice:     convMoneyValue(orderState.GetInitialOrderPrice()),
		ExecutedOrderPrice:    convMoneyValue(orderState.GetExecutedOrderPrice()),
		TotalOrderAmount:      convMoneyValue(orderState.GetTotalOrderAmount()),
		AveragePositionPrice:  convMoneyValue(orderState.GetAveragePositionPrice()),
		InitialCommission:     convMoneyValue(orderState.GetInitialCommission()),
		ExecutedCommission:    convMoneyValue(orderState.GetExecutedCommission()),
		Figi:                  orderState.GetFigi(),
		Direction:             domain.OrderDirection(orderState.GetDirection()),
		InitialSecurityPrice:  convMoneyValue(orderState.GetInitialSecurityPrice()),
		Stages:                stages,
		ServiceCommission:     convMoneyValue(orderState.GetServiceCommission()),
		Currency:              orderState.GetCurrency(),
		OrderType:             domain.OrderType(orderState.GetOrderType()),
		OrderDate:             &date,
	}
}

// Конвертация tkf.PostOrderResponse в domain.PostOrderResponse
func convPostOrderResponse(postOrderResponse *tkf.PostOrderResponse) *domain.PostOrderResponse {
	return &domain.PostOrderResponse{
		OrderId:               postOrderResponse.GetOrderId(),
		ExecutionReportStatus: domain.OrderExecutionReportStatus(postOrderResponse.GetExecutionReportStatus()),
		LotsRequested:         postOrderResponse.GetLotsRequested(),
		LotsExecuted:          postOrderResponse.GetLotsExecuted(),
		InitialOrderPrice:     convMoneyValue(postOrderResponse.GetInitialOrderPrice()),
		ExecutedOrderPrice:    convMoneyValue(postOrderResponse.GetExecutedOrderPrice()),
		TotalOrderAmount:      convMoneyValue(postOrderResponse.GetTotalOrderAmount()),
		InitialCommission:     convMoneyValue(postOrderResponse.GetInitialCommission()),
		ExecutedCommission:    convMoneyValue(postOrderResponse.GetExecutedCommission()),
		AciValue:              convMoneyValue(postOrderResponse.GetAciValue()),
		Figi:                  postOrderResponse.GetFigi(),
		Direction:             domain.OrderDirection(postOrderResponse.GetDirection()),
		InitialSecurityPrice:  convMoneyValue(postOrderResponse.GetInitialSecurityPrice()),
		OrderType:             domain.OrderType(postOrderResponse.GetOrderType()),
		Message:               postOrderResponse.GetMessage(),
		InitialOrderPricePt:   convQuotation(postOrderResponse.GetInitialOrderPricePt()),
	}
}

// Конвертация tkf.Share в domain.Share
func convShare(share *tkf.Share) *domain.Share {
	ipoDate := share.GetIpoDate().AsTime()
	return &domain.Share{
		Figi:                  share.GetFigi(),
		Ticker:                share.GetTicker(),
		ClassCode:             share.GetClassCode(),
		Isin:                  share.GetIsin(),
		Lot:                   share.GetLot(),
		Currency:              share.GetCurrency(),
		Klong:                 convQuotation(share.GetKlong()),
		Kshort:                convQuotation(share.GetKshort()),
		Dlong:                 convQuotation(share.GetDlong()),
		Dshort:                convQuotation(share.GetDshort()),
		DlongMin:              convQuotation(share.GetDlongMin()),
		DshortMin:             convQuotation(share.GetDshortMin()),
		ShortEnabledFlag:      share.GetShortEnabledFlag(),
		Name:                  share.GetName(),
		Exchange:              share.GetExchange(),
		IpoDate:               &ipoDate,
		IssueSize:             share.GetIssueSize(),
		CountryOfRisk:         share.GetCountryOfRisk(),
		CountryOfRiskName:     share.GetCountryOfRiskName(),
		Sector:                share.GetSector(),
		IssueSizePlan:         share.GetIssueSizePlan(),
		Nominal:               convMoneyValue(share.GetNominal()),
		TradingStatus:         domain.SecurityTradingStatus(share.GetTradingStatus()),
		OtcFlag:               share.GetOtcFlag(),
		BuyAvailableFlag:      share.GetBuyAvailableFlag(),
		SellAvailableFlag:     share.GetSellAvailableFlag(),
		DivYieldFlag:          share.GetDivYieldFlag(),
		ShareType:             domain.ShareType(share.GetShareType()),
		MinPriceIncrement:     convQuotation(share.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: share.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Bond в domain.Bond
func convBond(bond *tkf.Bond) *domain.Bond {
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
		Klong:                 convQuotation(bond.GetKlong()),
		Kshort:                convQuotation(bond.GetKshort()),
		Dlong:                 convQuotation(bond.GetDlong()),
		Dshort:                convQuotation(bond.GetDshort()),
		DlongMin:              convQuotation(bond.DlongMin),
		DshortMin:             convQuotation(bond.GetDshortMin()),
		ShortEnabledFlag:      bond.GetShortEnabledFlag(),
		Name:                  bond.GetName(),
		Exchange:              bond.GetExchange(),
		CouponQuantityPerYear: bond.GetCouponQuantityPerYear(),
		MaturityDate:          &maturityDate,
		Nominal:               convMoneyValue(bond.GetNominal()),
		StateRegDate:          &stateRegDate,
		PlacementDate:         &placementDate,
		PlacementPrice:        convMoneyValue(bond.GetPlacementPrice()),
		AciValue:              convMoneyValue(bond.GetAciValue()),
		CountryOfRisk:         bond.GetCountryOfRisk(),
		CountryOfRiskName:     bond.GetCountryOfRiskName(),
		Sector:                bond.GetSector(),
		IssueKind:             bond.GetIssueKind(),
		IssueSize:             bond.GetIssueSize(),
		TradingStatus:         domain.SecurityTradingStatus(bond.GetTradingStatus()),
		OtcFlag:               bond.GetOtcFlag(),
		BuyAvailableFlag:      bond.GetBuyAvailableFlag(),
		SellAvailableFlag:     bond.GetSellAvailableFlag(),
		FloatingCouponFlag:    bond.GetFloatingCouponFlag(),
		PerpetualFlag:         bond.GetPerpetualFlag(),
		AmortizationFlag:      bond.GetAmortizationFlag(),
		MinPriceIncrement:     convQuotation(bond.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: bond.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Currency в domain.Currency
func convCurrency(currency *tkf.Currency) *domain.Currency {
	return &domain.Currency{
		Figi:                  currency.GetFigi(),
		Ticker:                currency.GetTicker(),
		ClassCode:             currency.GetClassCode(),
		Isin:                  currency.GetIsin(),
		Lot:                   currency.GetLot(),
		Currency:              currency.GetCurrency(),
		Klong:                 convQuotation(currency.GetKlong()),
		Kshort:                convQuotation(currency.GetKshort()),
		Dlong:                 convQuotation(currency.GetDlong()),
		Dshort:                convQuotation(currency.GetDshort()),
		DlongMin:              convQuotation(currency.GetDlongMin()),
		DshortMin:             convQuotation(currency.GetDshortMin()),
		ShortEnabledFlag:      currency.GetShortEnabledFlag(),
		Name:                  currency.GetName(),
		Exchange:              currency.GetExchange(),
		Nominal:               convMoneyValue(currency.GetNominal()),
		CountryOfRisk:         currency.GetCountryOfRisk(),
		CountryOfRiskName:     currency.GetCountryOfRiskName(),
		TradingStatus:         domain.SecurityTradingStatus(currency.GetTradingStatus()),
		OtcFlag:               currency.GetOtcFlag(),
		SellAvailableFlag:     currency.GetSellAvailableFlag(),
		IsoCurrencyName:       currency.GetIsoCurrencyName(),
		MinPriceIncrement:     convQuotation(currency.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: currency.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Etf в domain.Etf
func convEtf(etf *tkf.Etf) *domain.Etf {
	releasedDate := etf.ReleasedDate.AsTime()
	return &domain.Etf{
		Figi:                  etf.GetFigi(),
		Ticker:                etf.GetTicker(),
		ClassCode:             etf.GetClassCode(),
		Isin:                  etf.GetIsin(),
		Lot:                   etf.GetLot(),
		Currency:              etf.GetCurrency(),
		Klong:                 convQuotation(etf.GetKlong()),
		Kshort:                convQuotation(etf.GetKshort()),
		Dlong:                 convQuotation(etf.GetDlong()),
		Dshort:                convQuotation(etf.GetDshort()),
		DlongMin:              convQuotation(etf.GetDlongMin()),
		DshortMin:             convQuotation(etf.GetDshortMin()),
		ShortEnabledFlag:      etf.GetShortEnabledFlag(),
		Name:                  etf.GetName(),
		Exchange:              etf.GetExchange(),
		FixedCommission:       convQuotation(etf.GetFixedCommission()),
		FocusType:             etf.GetFocusType(),
		ReleasedDate:          &releasedDate,
		NumShares:             convQuotation(etf.GetNumShares()),
		CountryOfRisk:         etf.GetCountryOfRisk(),
		CountryOfRiskName:     etf.GetCountryOfRiskName(),
		Sector:                etf.GetSector(),
		RebalancingFreq:       etf.GetRebalancingFreq(),
		TradingStatus:         domain.SecurityTradingStatus(etf.GetTradingStatus()),
		OtcFlag:               etf.GetOtcFlag(),
		BuyAvailableFlag:      etf.GetBuyAvailableFlag(),
		SellAvailableFlag:     etf.GetSellAvailableFlag(),
		MinPriceIncrement:     convQuotation(etf.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: etf.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Future в domain.Future
func convFuture(future *tkf.Future) *domain.Future {
	firstTradeDate := future.GetFirstTradeDate().AsTime()
	lastTradeDate := future.GetLastTradeDate().AsTime()
	expirationDate := future.GetExpirationDate().AsTime()
	return &domain.Future{
		Figi:                  future.GetFigi(),
		Ticker:                future.GetTicker(),
		ClassCode:             future.GetClassCode(),
		Lot:                   future.GetLot(),
		Currency:              future.GetCurrency(),
		Klong:                 convQuotation(future.GetKlong()),
		Kshort:                convQuotation(future.GetKshort()),
		Dlong:                 convQuotation(future.GetDlong()),
		Dshort:                convQuotation(future.GetDshort()),
		DlongMin:              convQuotation(future.GetDlongMin()),
		DshortMin:             convQuotation(future.GetDshortMin()),
		ShortEnabledFlag:      future.GetShortEnabledFlag(),
		Name:                  future.GetName(),
		Exchange:              future.GetExchange(),
		FirstTradeDate:        &firstTradeDate,
		LastTradeDate:         &lastTradeDate,
		FuturesType:           future.GetFuturesType(),
		AssetType:             future.GetAssetType(),
		BasicAsset:            future.GetBasicAsset(),
		BasicAssetSize:        convQuotation(future.GetBasicAssetSize()),
		CountryOfRisk:         future.GetCountryOfRisk(),
		CountryOfRiskName:     future.GetCountryOfRiskName(),
		Sector:                future.GetSector(),
		ExpirationDate:        &expirationDate,
		TradingStatus:         domain.SecurityTradingStatus(future.GetTradingStatus()),
		OtcFlag:               future.GetOtcFlag(),
		BuyAvailableFlag:      future.GetBuyAvailableFlag(),
		SellAvailableFlag:     future.GetSellAvailableFlag(),
		MinPriceIncrement:     convQuotation(future.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: future.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.AssetSecurity в domain.AssetSecurity
func convAssetSecurity(tkfAS *tkf.AssetSecurity) *domain.AssetSecurity {
	// Конвертация tkf.AssetShare в domain.tkf.AssetShare
	tkfShare := tkfAS.GetShare()
	var share *domain.AssetShare
	if tkfAS.GetType() == "share" && tkfShare != nil {
		share = convAssetShare(tkfShare)
	}

	// Конвертация tkf.AssetBond в domain.AssetBond
	tkfBond := tkfAS.GetBond()
	var bond *domain.AssetBond
	if tkfAS.GetType() == "bond" && tkfBond != nil {
		bond = convAssetBond(tkfBond)
	}

	// Конвертация tkf.AssetStructuredProduct в domain.AssetStructuredProduct
	tkfSP := tkfAS.GetSp()
	var structuredProduct *domain.AssetStructuredProduct
	if tkfAS.GetType() == "sp" && tkfSP != nil {
		structuredProduct = convAssetStructuredProduct(tkfSP)
	}

	// Конвертация tkf.AssetEtf в domain.AssetEtf
	tkfEtf := tkfAS.GetEtf()
	var etf *domain.AssetEtf
	if tkfAS.GetType() == "etf" && tkfEtf != nil {
		etf = convAssetEtf(tkfEtf)
	}

	// Конвертация tkf.AssetClearingCertificate в domain.AssetClearingCertificate
	tkfCC := tkfAS.GetClearingCertificate()
	var clearingCertificate *domain.AssetClearingCertificate
	if tkfAS.GetType() == "clearing_certificate" && tkfCC != nil {
		clearingCertificate = &domain.AssetClearingCertificate{
			Nominal:         *convQuotation(tkfCC.GetNominal()),
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

// Конвертация tkf.AssetShare в domain.tkf.AssetShare
func convAssetShare(assetShare *tkf.AssetShare) *domain.AssetShare {
	if assetShare == nil {
		return nil
	}

	ipoDate := assetShare.GetIpoDate().AsTime()
	regestryDate := assetShare.GetRegistryDate().AsTime()
	placementDate := assetShare.GetPlacementDate().AsTime()
	return &domain.AssetShare{
		Type:               domain.ShareType(assetShare.GetType()),
		IssueSize:          *convQuotation(assetShare.GetIssueSize()),
		Nominal:            *convQuotation(assetShare.GetNominal()),
		NominalCurrency:    assetShare.GetNominalCurrency(),
		PrimaryIndex:       assetShare.GetPrimaryIndex(),
		DividendRate:       *convQuotation(assetShare.GetDividendRate()),
		PreferredShareType: assetShare.GetPreferredShareType(),
		IpoDate:            &ipoDate,
		RegistryDate:       &regestryDate,
		DivYield:           assetShare.GetDivYieldFlag(),
		IssueKind:          assetShare.GetIssueKind(),
		PlacementDate:      &placementDate,
		RepresIsin:         assetShare.GetRepresIsin(),
		IssueSizePlan:      *convQuotation(assetShare.GetIssueSizePlan()),
		TotalFloat:         *convQuotation(assetShare.GetTotalFloat()),
	}
}

// Конвертация tkf.AssetBond в domain.AssetBond
func convAssetBond(assetBond *tkf.AssetBond) *domain.AssetBond {
	if assetBond == nil {
		return nil
	}

	maturityDate := assetBond.GetMaturityDate().AsTime()
	stateRegDate := assetBond.GetStateRegDate().AsTime()
	placementDate := assetBond.GetPlacementDate().AsTime()
	return &domain.AssetBond{
		CurrentNominal:        *convQuotation(assetBond.GetCurrentNominal()),
		BorrowName:            assetBond.GetBorrowName(),
		IssueSize:             *convQuotation(assetBond.GetIssueSize()),
		Nominal:               *convQuotation(assetBond.GetNominal()),
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
		PlacementPrice:        *convQuotation(assetBond.GetPlacementPrice()),
		IssueSizePlan:         *convQuotation(assetBond.GetIssueSizePlan()),
	}
}

// Конвертация tkf.AssetStructuredProduct в domain.AssetStructuredProduct
func convAssetStructuredProduct(assetSP *tkf.AssetStructuredProduct) *domain.AssetStructuredProduct {
	if assetSP == nil {
		return nil
	}

	maturityDate := assetSP.GetMaturityDate().AsTime()
	placementDate := assetSP.GetPlacementDate().AsTime()
	return &domain.AssetStructuredProduct{
		BorrowName:      assetSP.GetBorrowName(),
		Nominal:         *convQuotation(assetSP.GetNominal()),
		NominalCurrency: assetSP.GetNominalCurrency(),
		Type:            domain.StructuredProductType(assetSP.GetType()),
		LogicPortfolio:  assetSP.GetLogicPortfolio(),
		AssetType:       domain.AssetType(assetSP.GetAssetType()),
		BasicAsset:      assetSP.GetBasicAsset(),
		SafetyBarrier:   *convQuotation(assetSP.GetSafetyBarrier()),
		MaturityDate:    &maturityDate,
		IssueSizePlan:   *convQuotation(assetSP.GetIssueSizePlan()),
		IssueSize:       *convQuotation(assetSP.GetIssueSize()),
		PlacementDate:   &placementDate,
		IssueKind:       assetSP.GetIssueKind(),
	}
}

// Конвертация tkf.AssetEtf в domain.AssetEtf
func convAssetEtf(assetEtf *tkf.AssetEtf) *domain.AssetEtf {
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
		TotalExpense:              *convQuotation(assetEtf.GetTotalExpense()),
		HurdleRate:                *convQuotation(assetEtf.GetHurdleRate()),
		PerformanceFee:            *convQuotation(assetEtf.GetPerformanceFee()),
		FixedCommission:           *convQuotation(assetEtf.GetFixedCommission()),
		PaymentType:               assetEtf.GetPaymentType(),
		Watermark:                 assetEtf.GetWatermarkFlag(),
		BuyPremium:                *convQuotation(assetEtf.GetBuyPremium()),
		SellDiscount:              *convQuotation(assetEtf.GetSellDiscount()),
		Rebalancing:               assetEtf.GetRebalancingFlag(),
		RebalancingFreq:           assetEtf.GetRebalancingFreq(),
		ManagementType:            assetEtf.GetManagementType(),
		PrimaryIndex:              assetEtf.GetPrimaryIndex(),
		FocusType:                 assetEtf.GetFocusType(),
		Leveraged:                 assetEtf.GetLeveragedFlag(),
		NumShare:                  *convQuotation(assetEtf.GetNumShare()),
		Ucits:                     assetEtf.GetUcitsFlag(),
		ReleasedDate:              &releasedDate,
		Description:               assetEtf.GetDescription(),
		PrimaryIndexDescription:   assetEtf.GetPrimaryIndexDescription(),
		PrimaryIndexCompany:       assetEtf.GetPrimaryIndexCompany(),
		IndexRecoveryPeriod:       *convQuotation(assetEtf.GetIndexRecoveryPeriod()),
		InavCode:                  assetEtf.GetInavCode(),
		DivYield:                  assetEtf.GetDivYieldFlag(),
		ExpenseCommission:         *convQuotation(assetEtf.GetExpenseCommission()),
		PrimaryIndexTrackingError: *convQuotation(assetEtf.GetPrimaryIndexTrackingError()),
		RebalancingPlan:           assetEtf.GetRebalancingPlan(),
		TaxRate:                   assetEtf.GetTaxRate(),
		RebalancingDates:          rebalancingDates,
		IssueKind:                 assetEtf.GetIssueKind(),
		Nominal:                   *convQuotation(assetEtf.GetNominal()),
		NominalCurrency:           assetEtf.GetNominalCurrency(),
	}
}

// Конвертация tkf.Brand в domain.Brand
func convBrand(brand *tkf.Brand) *domain.Brand {
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
func convAssetInstrument(instrument *tkf.AssetInstrument) *domain.AssetInstrument {
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

// Конвертация tkf.FavoriteInstrument в domain.FavoriteInstrument
func convFavoriteInstrument(instrument *tkf.FavoriteInstrument) *domain.FavoriteInstrument {
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

// Конвертация tkf.InstrumentShort в domain.InstrumentShort
func convInstrumentShort(instrument *tkf.InstrumentShort) *domain.InstrumentShort {
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
