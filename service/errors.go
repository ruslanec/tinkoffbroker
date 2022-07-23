package service

import "errors"

var (
	ErrCandleInterval = errors.New("error candle interval") // Неправильный интервал торговой свечи
	ErrInputArgument  = errors.New("error input argument")  // Неправильный входной аргумент
)
