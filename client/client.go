package client

import "context"

//TODO: Описываем интерфейс , который будет реализовывать клиенты через адаптеры

type ExchangeClient interface {
	LastMinuteCandlesticks(ctx context.Context, symbol string, minutes int) ([]Candlestick, error)
	LastHourCandlesticks(ctx context.Context, symbol string, hours int) ([]Candlestick, error)
	LastDayCandlesticks(ctx context.Context, symbol string) ([]Candlestick, error)
	LastWeekCandlesticks(ctx context.Context, symbol string) ([]Candlestick, error)
	LastMonthCandlesticks(ctx context.Context, symbol string) ([]Candlestick, error)
	Prices(ctx context.Context) ([]Price, error)
	SymbolInfo(ctx context.Context) ([]SymbolInfo, error)
}
