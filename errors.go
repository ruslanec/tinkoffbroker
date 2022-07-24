package tinkoffbroker

import "errors"

var (
	ErrArgCandleUnspecified = errors.New("error candle interval") // Неправильный интервал торговой свечи
)
var (
	ErrArgTimeInterval = errors.New("error period bounds")
)

var (
	ErrSvcNotImplemented = errors.New("service not implemented")
)

var (
	ErrArgEmptyID       = errors.New("empty ID")
	ErrArgEmptyAccounID = errors.New("empty account ID")
	ErrArgEmptyFigi     = errors.New("empty figi")
	ErrRetEmptyField    = errors.New("returned empty field")
	ErrArgEmptyQuery    = errors.New("empty query argument")
)
