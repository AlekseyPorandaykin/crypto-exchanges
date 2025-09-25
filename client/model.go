package client

import "time"

// TODO: Описываем модели, которые будут использоваться в интерфейсе
type Price struct {
	Symbol    string
	Exchange  string
	Value     string
	CreatedAt time.Time
}

type Candlestick struct {
	StartTime  time.Time
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string
	Volume     string
}

type SymbolCategory string

const (
	SymbolCategorySpot   SymbolCategory = "spot"
	SymbolCategoryFuture SymbolCategory = "future"
	SymbolCategoryOther  SymbolCategory = "other"
)

type SymbolInfo struct {
	Symbol          string
	Exchange        string
	BaseAsset       string
	QuoteAsset      string
	Category        SymbolCategory
	FundingRate     float32
	NextFundingTime time.Time
}
