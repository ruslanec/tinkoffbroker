package tinkoffbroker

import (
	"time"

	tkf "github.com/ruslanec/tinkoffbroker/proto"
)

// Конвертация tkf.MoneyValue в MoneyValue
func convMoneyValue(moneyValue *tkf.MoneyValue) *MoneyValue {
	return &MoneyValue{
		Currency: moneyValue.GetCurrency(),
		Units:    moneyValue.GetUnits(),
		Nano:     moneyValue.GetNano(),
	}
}

// Конвертация tkf.Quotation в Quotation
func convQuotation(quotation *tkf.Quotation) *Quotation {
	return &Quotation{
		Units: quotation.GetUnits(),
		Nano:  quotation.GetNano(),
	}
}

// Конвертация tkf.PortfolioPosition в PortfolioPosition
func convPortfolioPosition(portfolioPosition *tkf.PortfolioPosition) *PortfolioPosition {
	return &PortfolioPosition{
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

// Конвертация tkf.Operation в Operation
func convOperation(operation *tkf.Operation) *Operation {
	date := operation.GetDate().AsTime()
	var trades []*OperationTrade
	for _, v := range operation.GetTrades() {
		dt := v.GetDateTime().AsTime()
		trades = append(trades, &OperationTrade{
			TradeId:  v.GetTradeId(),
			DateTime: &dt,
			Quantity: operation.GetQuantity(),
			Price:    convMoneyValue(v.GetPrice()),
		})
	}

	return &Operation{
		Id:                operation.GetId(),
		ParentOperationId: operation.GetParentOperationId(),
		Currency:          operation.GetCurrency(),
		Payment:           convMoneyValue(operation.GetPayment()),
		Price:             convMoneyValue(operation.GetPrice()),
		State:             OperationState(operation.GetState()),
		Quantity:          operation.GetQuantity(),
		QuantityRest:      operation.GetQuantityRest(),
		Figi:              operation.GetFigi(),
		InstrumentType:    operation.GetInstrumentType(),
		Date:              &date,
		Type:              operation.GetType(),
		OperationType:     OperationType(operation.GetOperationType()),
		Trades:            trades,
	}

}

// Конвертация tkf.OrderState в OrderState
func convOrderState(orderState *tkf.OrderState) *OrderState {
	var stages []*OrderStage
	for _, v := range orderState.GetStages() {
		stages = append(stages, &OrderStage{
			Price:    convMoneyValue(v.GetPrice()),
			Quantity: v.GetQuantity(),
			TradeId:  v.GetTradeId(),
		})
	}
	date := orderState.GetOrderDate().AsTime()
	return &OrderState{
		OrderId:               orderState.GetOrderId(),
		ExecutionReportStatus: OrderExecutionReportStatus(orderState.GetExecutionReportStatus()),
		LotsRequested:         orderState.GetLotsRequested(),
		LotsExecuted:          orderState.GetLotsExecuted(),
		InitialOrderPrice:     convMoneyValue(orderState.GetInitialOrderPrice()),
		ExecutedOrderPrice:    convMoneyValue(orderState.GetExecutedOrderPrice()),
		TotalOrderAmount:      convMoneyValue(orderState.GetTotalOrderAmount()),
		AveragePositionPrice:  convMoneyValue(orderState.GetAveragePositionPrice()),
		InitialCommission:     convMoneyValue(orderState.GetInitialCommission()),
		ExecutedCommission:    convMoneyValue(orderState.GetExecutedCommission()),
		Figi:                  orderState.GetFigi(),
		Direction:             OrderDirection(orderState.GetDirection()),
		InitialSecurityPrice:  convMoneyValue(orderState.GetInitialSecurityPrice()),
		Stages:                stages,
		ServiceCommission:     convMoneyValue(orderState.GetServiceCommission()),
		Currency:              orderState.GetCurrency(),
		OrderType:             OrderType(orderState.GetOrderType()),
		OrderDate:             &date,
	}
}

// Конвертация tkf.PostOrderResponse в PostOrderResponse
func convPostOrderResponse(postOrderResponse *tkf.PostOrderResponse) *PostOrderResponse {
	return &PostOrderResponse{
		OrderId:               postOrderResponse.GetOrderId(),
		ExecutionReportStatus: OrderExecutionReportStatus(postOrderResponse.GetExecutionReportStatus()),
		LotsRequested:         postOrderResponse.GetLotsRequested(),
		LotsExecuted:          postOrderResponse.GetLotsExecuted(),
		InitialOrderPrice:     convMoneyValue(postOrderResponse.GetInitialOrderPrice()),
		ExecutedOrderPrice:    convMoneyValue(postOrderResponse.GetExecutedOrderPrice()),
		TotalOrderAmount:      convMoneyValue(postOrderResponse.GetTotalOrderAmount()),
		InitialCommission:     convMoneyValue(postOrderResponse.GetInitialCommission()),
		ExecutedCommission:    convMoneyValue(postOrderResponse.GetExecutedCommission()),
		AciValue:              convMoneyValue(postOrderResponse.GetAciValue()),
		Figi:                  postOrderResponse.GetFigi(),
		Direction:             OrderDirection(postOrderResponse.GetDirection()),
		InitialSecurityPrice:  convMoneyValue(postOrderResponse.GetInitialSecurityPrice()),
		OrderType:             OrderType(postOrderResponse.GetOrderType()),
		Message:               postOrderResponse.GetMessage(),
		InitialOrderPricePt:   convQuotation(postOrderResponse.GetInitialOrderPricePt()),
	}
}

// Конвертация tkf.Share в Share
func convShare(share *tkf.Share) *Share {
	ipoDate := share.GetIpoDate().AsTime()
	return &Share{
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
		TradingStatus:         SecurityTradingStatus(share.GetTradingStatus()),
		OtcFlag:               share.GetOtcFlag(),
		BuyAvailableFlag:      share.GetBuyAvailableFlag(),
		SellAvailableFlag:     share.GetSellAvailableFlag(),
		DivYieldFlag:          share.GetDivYieldFlag(),
		ShareType:             ShareType(share.GetShareType()),
		MinPriceIncrement:     convQuotation(share.GetMinPriceIncrement()),
		ApiTradeAvailableFlag: share.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Bond в Bond
func convBond(bond *tkf.Bond) *Bond {
	maturityDate := bond.GetMaturityDate().AsTime()
	stateRegDate := bond.GetStateRegDate().AsTime()
	placementDate := bond.GetPlacementDate().AsTime()
	return &Bond{
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
		ShortEnabled:          bond.GetShortEnabledFlag(),
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
		TradingStatus:         SecurityTradingStatus(bond.GetTradingStatus()),
		Otc:                   bond.GetOtcFlag(),
		BuyAvailable:          bond.GetBuyAvailableFlag(),
		SellAvailable:         bond.GetSellAvailableFlag(),
		FloatingCoupon:        bond.GetFloatingCouponFlag(),
		Perpetual:             bond.GetPerpetualFlag(),
		Amortization:          bond.GetAmortizationFlag(),
		MinPriceIncrement:     convQuotation(bond.GetMinPriceIncrement()),
		ApiTradeAvailable:     bond.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Currency в Currency
func convCurrency(currency *tkf.Currency) *Currency {
	return &Currency{
		Figi:              currency.GetFigi(),
		Ticker:            currency.GetTicker(),
		ClassCode:         currency.GetClassCode(),
		Isin:              currency.GetIsin(),
		Lot:               currency.GetLot(),
		Currency:          currency.GetCurrency(),
		Klong:             convQuotation(currency.GetKlong()),
		Kshort:            convQuotation(currency.GetKshort()),
		Dlong:             convQuotation(currency.GetDlong()),
		Dshort:            convQuotation(currency.GetDshort()),
		DlongMin:          convQuotation(currency.GetDlongMin()),
		DshortMin:         convQuotation(currency.GetDshortMin()),
		ShortEnabled:      currency.GetShortEnabledFlag(),
		Name:              currency.GetName(),
		Exchange:          currency.GetExchange(),
		Nominal:           convMoneyValue(currency.GetNominal()),
		CountryOfRisk:     currency.GetCountryOfRisk(),
		CountryOfRiskName: currency.GetCountryOfRiskName(),
		TradingStatus:     SecurityTradingStatus(currency.GetTradingStatus()),
		Otc:               currency.GetOtcFlag(),
		SellAvailable:     currency.GetSellAvailableFlag(),
		IsoCurrencyName:   currency.GetIsoCurrencyName(),
		MinPriceIncrement: convQuotation(currency.GetMinPriceIncrement()),
		ApiTradeAvailable: currency.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Etf в Etf
func convEtf(etf *tkf.Etf) *Etf {
	releasedDate := etf.ReleasedDate.AsTime()
	return &Etf{
		Figi:              etf.GetFigi(),
		Ticker:            etf.GetTicker(),
		ClassCode:         etf.GetClassCode(),
		Isin:              etf.GetIsin(),
		Lot:               etf.GetLot(),
		Currency:          etf.GetCurrency(),
		Klong:             convQuotation(etf.GetKlong()),
		Kshort:            convQuotation(etf.GetKshort()),
		Dlong:             convQuotation(etf.GetDlong()),
		Dshort:            convQuotation(etf.GetDshort()),
		DlongMin:          convQuotation(etf.GetDlongMin()),
		DshortMin:         convQuotation(etf.GetDshortMin()),
		ShortEnabled:      etf.GetShortEnabledFlag(),
		Name:              etf.GetName(),
		Exchange:          etf.GetExchange(),
		FixedCommission:   convQuotation(etf.GetFixedCommission()),
		FocusType:         etf.GetFocusType(),
		ReleasedDate:      &releasedDate,
		NumShares:         convQuotation(etf.GetNumShares()),
		CountryOfRisk:     etf.GetCountryOfRisk(),
		CountryOfRiskName: etf.GetCountryOfRiskName(),
		Sector:            etf.GetSector(),
		RebalancingFreq:   etf.GetRebalancingFreq(),
		TradingStatus:     SecurityTradingStatus(etf.GetTradingStatus()),
		Otc:               etf.GetOtcFlag(),
		BuyAvailable:      etf.GetBuyAvailableFlag(),
		SellAvailable:     etf.GetSellAvailableFlag(),
		MinPriceIncrement: convQuotation(etf.GetMinPriceIncrement()),
		ApiTradeAvailable: etf.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.Future в Future
func convFuture(future *tkf.Future) *Future {
	firstTradeDate := future.GetFirstTradeDate().AsTime()
	lastTradeDate := future.GetLastTradeDate().AsTime()
	expirationDate := future.GetExpirationDate().AsTime()
	return &Future{
		Figi:              future.GetFigi(),
		Ticker:            future.GetTicker(),
		ClassCode:         future.GetClassCode(),
		Lot:               future.GetLot(),
		Currency:          future.GetCurrency(),
		Klong:             convQuotation(future.GetKlong()),
		Kshort:            convQuotation(future.GetKshort()),
		Dlong:             convQuotation(future.GetDlong()),
		Dshort:            convQuotation(future.GetDshort()),
		DlongMin:          convQuotation(future.GetDlongMin()),
		DshortMin:         convQuotation(future.GetDshortMin()),
		ShortEnabled:      future.GetShortEnabledFlag(),
		Name:              future.GetName(),
		Exchange:          future.GetExchange(),
		FirstTradeDate:    &firstTradeDate,
		LastTradeDate:     &lastTradeDate,
		FuturesType:       future.GetFuturesType(),
		AssetType:         future.GetAssetType(),
		BasicAsset:        future.GetBasicAsset(),
		BasicAssetSize:    convQuotation(future.GetBasicAssetSize()),
		CountryOfRisk:     future.GetCountryOfRisk(),
		CountryOfRiskName: future.GetCountryOfRiskName(),
		Sector:            future.GetSector(),
		ExpirationDate:    &expirationDate,
		TradingStatus:     SecurityTradingStatus(future.GetTradingStatus()),
		Otc:               future.GetOtcFlag(),
		BuyAvailable:      future.GetBuyAvailableFlag(),
		SellAvailable:     future.GetSellAvailableFlag(),
		MinPriceIncrement: convQuotation(future.GetMinPriceIncrement()),
		ApiTradeAvailable: future.GetApiTradeAvailableFlag(),
	}
}

// Конвертация tkf.AssetSecurity в AssetSecurity
func convAssetSecurity(tkfAS *tkf.AssetSecurity) *AssetSecurity {
	// Конвертация tkf.AssetShare в tkf.AssetShare
	tkfShare := tkfAS.GetShare()
	var share *AssetShare
	if tkfAS.GetType() == "share" && tkfShare != nil {
		share = convAssetShare(tkfShare)
	}

	// Конвертация tkf.AssetBond в AssetBond
	tkfBond := tkfAS.GetBond()
	var bond *AssetBond
	if tkfAS.GetType() == "bond" && tkfBond != nil {
		bond = convAssetBond(tkfBond)
	}

	// Конвертация tkf.AssetStructuredProduct в AssetStructuredProduct
	tkfSP := tkfAS.GetSp()
	var structuredProduct *AssetStructuredProduct
	if tkfAS.GetType() == "sp" && tkfSP != nil {
		structuredProduct = convAssetStructuredProduct(tkfSP)
	}

	// Конвертация tkf.AssetEtf в AssetEtf
	tkfEtf := tkfAS.GetEtf()
	var etf *AssetEtf
	if tkfAS.GetType() == "etf" && tkfEtf != nil {
		etf = convAssetEtf(tkfEtf)
	}

	// Конвертация tkf.AssetClearingCertificate в AssetClearingCertificate
	tkfCC := tkfAS.GetClearingCertificate()
	var clearingCertificate *AssetClearingCertificate
	if tkfAS.GetType() == "clearing_certificate" && tkfCC != nil {
		clearingCertificate = &AssetClearingCertificate{
			Nominal:         *convQuotation(tkfCC.GetNominal()),
			NominalCurrency: tkfCC.GetNominalCurrency(),
		}
	}

	return &AssetSecurity{
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
func convAssetShare(assetShare *tkf.AssetShare) *AssetShare {
	if assetShare == nil {
		return nil
	}

	ipoDate := assetShare.GetIpoDate().AsTime()
	regestryDate := assetShare.GetRegistryDate().AsTime()
	placementDate := assetShare.GetPlacementDate().AsTime()
	return &AssetShare{
		Type:               ShareType(assetShare.GetType()),
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

// Конвертация tkf.AssetBond в AssetBond
func convAssetBond(assetBond *tkf.AssetBond) *AssetBond {
	if assetBond == nil {
		return nil
	}

	maturityDate := assetBond.GetMaturityDate().AsTime()
	stateRegDate := assetBond.GetStateRegDate().AsTime()
	placementDate := assetBond.GetPlacementDate().AsTime()
	return &AssetBond{
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

// Конвертация tkf.AssetStructuredProduct в AssetStructuredProduct
func convAssetStructuredProduct(assetSP *tkf.AssetStructuredProduct) *AssetStructuredProduct {
	if assetSP == nil {
		return nil
	}

	maturityDate := assetSP.GetMaturityDate().AsTime()
	placementDate := assetSP.GetPlacementDate().AsTime()
	return &AssetStructuredProduct{
		BorrowName:      assetSP.GetBorrowName(),
		Nominal:         *convQuotation(assetSP.GetNominal()),
		NominalCurrency: assetSP.GetNominalCurrency(),
		Type:            StructuredProductType(assetSP.GetType()),
		LogicPortfolio:  assetSP.GetLogicPortfolio(),
		AssetType:       AssetType(assetSP.GetAssetType()),
		BasicAsset:      assetSP.GetBasicAsset(),
		SafetyBarrier:   *convQuotation(assetSP.GetSafetyBarrier()),
		MaturityDate:    &maturityDate,
		IssueSizePlan:   *convQuotation(assetSP.GetIssueSizePlan()),
		IssueSize:       *convQuotation(assetSP.GetIssueSize()),
		PlacementDate:   &placementDate,
		IssueKind:       assetSP.GetIssueKind(),
	}
}

// Конвертация tkf.AssetEtf в AssetEtf
func convAssetEtf(assetEtf *tkf.AssetEtf) *AssetEtf {
	if assetEtf == nil {
		return nil
	}

	releasedDate := assetEtf.GetReleasedDate().AsTime()
	var rebalancingDates []*time.Time
	for _, v := range assetEtf.GetRebalancingDates() {
		tkfRT := v.AsTime()
		rebalancingDates = append(rebalancingDates, &tkfRT)
	}
	return &AssetEtf{
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

// Конвертация tkf.Brand в Brand
func convBrand(brand *tkf.Brand) *Brand {
	return &Brand{
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
func convAssetInstrument(instrument *tkf.AssetInstrument) *AssetInstrument {
	if instrument == nil {
		return nil
	}

	var links []*InstrumentLink
	for _, tkfLink := range instrument.GetLinks() {
		link := &InstrumentLink{
			Type:          tkfLink.GetType(),
			InstrumentUid: tkfLink.GetInstrumentUid(),
		}
		links = append(links, link)
	}

	return &AssetInstrument{
		Uid:            instrument.GetUid(),
		Figi:           instrument.GetFigi(),
		InstrumentType: instrument.GetInstrumentType(),
		Ticker:         instrument.GetTicker(),
		ClassCode:      instrument.GetClassCode(),
		Links:          links,
	}
}

// Конвертация tkf.FavoriteInstrument в FavoriteInstrument
func convFavoriteInstrument(instrument *tkf.FavoriteInstrument) *FavoriteInstrument {
	if instrument == nil {
		return nil
	}

	return &FavoriteInstrument{
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
func convInstrumentShort(instrument *tkf.InstrumentShort) *InstrumentShort {
	if instrument == nil {
		return nil
	}

	return &InstrumentShort{
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
