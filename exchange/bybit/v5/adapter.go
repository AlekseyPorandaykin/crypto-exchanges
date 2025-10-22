package v5

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/AlekseyPorandaykin/crypto-exchanges/client"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"github.com/pkg/errors"
)

const exchangeName = "bybit"

type Adapter struct {
	client *Client
}

func NewAdapter(c *Client) client.ExchangeClient {
	return &Adapter{c}
}

func (a *Adapter) ExchangeName() string {
	return exchangeName
}

func (a *Adapter) LastMinuteCandlesticks(ctx context.Context, symbol string, minutes int) ([]client.Candlestick, error) {
	return a.lastCandlesticks(ctx, symbol, strconv.Itoa(minutes))
}

func (a *Adapter) LastHourCandlesticks(ctx context.Context, symbol string, hours int) ([]client.Candlestick, error) {
	minutes := hours * 60
	return a.lastCandlesticks(ctx, symbol, strconv.Itoa(minutes))
}

func (a *Adapter) LastDayCandlesticks(ctx context.Context, symbol string) ([]client.Candlestick, error) {
	return a.lastCandlesticks(ctx, symbol, "D")
}

func (a *Adapter) LastWeekCandlesticks(ctx context.Context, symbol string) ([]client.Candlestick, error) {
	return a.lastCandlesticks(ctx, symbol, "W")
}

func (a *Adapter) LastMonthCandlesticks(ctx context.Context, symbol string) ([]client.Candlestick, error) {
	return a.lastCandlesticks(ctx, symbol, "M")
}

func (a *Adapter) lastCandlesticks(ctx context.Context, symbol, interval string) ([]client.Candlestick, error) {
	resp, err := a.client.MarketGetKline(ctx, request.MarketGetKlineParam{
		Symbol:   symbol,
		Interval: interval,
		Limit:    100,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get MarketGetKline")
	}
	if !resp.IsOk() {
		return nil, errors.New("incorrect response lastCandlesticks")
	}
	data := resp.Result.Candlesticks()
	result := make([]client.Candlestick, 0, len(data))
	for _, item := range data {
		startTimeMs, errST := strconv.Atoi(item.StartTime)
		if errST != nil {
			continue
		}
		result = append(result, client.Candlestick{
			StartTime:  time.UnixMilli(int64(startTimeMs)).In(time.UTC),
			OpenPrice:  item.OpenPrice,
			HighPrice:  item.HighPrice,
			LowPrice:   item.LowPrice,
			ClosePrice: item.ClosePrice,
			Volume:     item.Volume,
		})
	}

	return result, nil
}

func (a *Adapter) Prices(ctx context.Context) ([]client.Price, error) {
	tickerResp, err := a.client.MarketSpotTicker(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}
	if !tickerResp.IsOk() {
		return nil, fmt.Errorf("incorrect response: %s", tickerResp.Message)
	}
	currentTime := time.UnixMilli(tickerResp.Time)
	if currentTime.Year() != time.Now().Year() {
		currentTime = time.Now()
	}
	result := make([]client.Price, 0, len(tickerResp.Result.List))
	for _, data := range tickerResp.Result.List {
		result = append(result, client.Price{
			Symbol:    data.Symbol,
			Exchange:  exchangeName,
			Value:     data.LastPrice,
			CreatedAt: currentTime,
		})
	}
	return result, nil
}

func (a *Adapter) SymbolInfo(ctx context.Context) ([]client.SymbolInfo, error) {
	var result []client.SymbolInfo
	spotInfo, err := a.spotSymbolInfo(ctx)
	if err != nil {
		return nil, err
	}
	linearInfo, err := a.linearSymbolInfo(ctx)
	if err != nil {
		return nil, err
	}
	inverseInfo, err := a.inverseSymbolInfo(ctx)
	if err != nil {
		return nil, err
	}
	optionInfo, err := a.optionSymbolInfo(ctx)
	if err != nil {
		return nil, err
	}

	result = make([]client.SymbolInfo, 0, len(spotInfo)+len(linearInfo)+len(inverseInfo)+len(optionInfo))
	result = append(result, spotInfo...)
	result = append(result, linearInfo...)
	result = append(result, inverseInfo...)
	result = append(result, optionInfo...)
	return result, nil
}

func (a *Adapter) spotSymbolInfo(ctx context.Context) ([]client.SymbolInfo, error) {
	instrumentInfo, err := a.client.MarketInstrumentsSpotInfo(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]client.SymbolInfo, 0, len(instrumentInfo.Result.List))
	for _, item := range instrumentInfo.Result.List {
		result = append(result, client.SymbolInfo{
			Symbol:     item.Symbol,
			Exchange:   exchangeName,
			BaseAsset:  item.BaseCoin,
			QuoteAsset: item.QuoteCoin,
			Category:   client.SymbolCategorySpot,
		})
	}
	tickerData, err := a.client.MarketSpotTicker(ctx)
	if err != nil {
		return nil, err
	}

	return a.enrichFundingRate(result, tickerData.Result.List), nil
}

func (a *Adapter) enrichFundingRate(data []client.SymbolInfo, tickerData []SymbolTick) []client.SymbolInfo {
	tickerBySymbol := make(map[string]SymbolTick, len(tickerData))
	for _, tickerItem := range tickerData {
		tickerBySymbol[tickerItem.Symbol] = tickerItem
	}
	for i := range data {
		tickerItem, has := tickerBySymbol[data[i].Symbol]
		if !has {
			continue
		}
		if tickerItem.FundingRate != "" {
			fundingRate, _ := strconv.ParseFloat(tickerItem.FundingRate, 32)
			data[i].FundingRate = float32(fundingRate)
		}
		nextFundingTime, errNFT := strconv.ParseInt(tickerItem.NextFundingTime, 10, 64)
		if errNFT == nil {
			data[i].NextFundingTime = time.UnixMilli(nextFundingTime).In(time.UTC)
		}
	}

	return data
}
func (a *Adapter) linearSymbolInfo(ctx context.Context) ([]client.SymbolInfo, error) {
	instrumentInfo, err := a.client.MarketInstrumentsLinearInfo(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]client.SymbolInfo, 0, len(instrumentInfo.Result.List))
	for _, item := range instrumentInfo.Result.List {
		result = append(result, client.SymbolInfo{
			Symbol:     item.Symbol,
			Exchange:   exchangeName,
			BaseAsset:  item.BaseCoin,
			QuoteAsset: item.QuoteCoin,
			Category:   client.SymbolCategoryFuture,
		})
	}
	tickerData, err := a.client.MarketLinearTicker(ctx)
	if err != nil {
		return nil, err
	}

	return a.enrichFundingRate(result, tickerData.Result.List), nil
}

func (a *Adapter) inverseSymbolInfo(ctx context.Context) ([]client.SymbolInfo, error) {
	instrumentInfo, err := a.client.MarketInstrumentsInverseInfo(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]client.SymbolInfo, 0, len(instrumentInfo.Result.List))
	for _, item := range instrumentInfo.Result.List {
		result = append(result, client.SymbolInfo{
			Symbol:     item.Symbol,
			Exchange:   exchangeName,
			BaseAsset:  item.BaseCoin,
			QuoteAsset: item.QuoteCoin,
			Category:   client.SymbolCategoryOther,
		})
	}
	tickerData, err := a.client.MarketInverseTicker(ctx)
	if err != nil {
		return nil, err
	}

	return a.enrichFundingRate(result, tickerData.Result.List), nil
}

func (a *Adapter) optionSymbolInfo(ctx context.Context) ([]client.SymbolInfo, error) {
	instrumentInfo, err := a.client.MarketInstrumentsOptionInfo(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]client.SymbolInfo, 0, len(instrumentInfo.Result.List))
	for _, item := range instrumentInfo.Result.List {
		result = append(result, client.SymbolInfo{
			Symbol:     item.Symbol,
			Exchange:   exchangeName,
			BaseAsset:  item.BaseCoin,
			QuoteAsset: item.QuoteCoin,
			Category:   client.SymbolCategoryOther,
		})
	}
	tickerData, err := a.client.MarketOptionTicker(ctx)
	if err != nil {
		return nil, err
	}

	return a.enrichFundingRate(result, tickerData.Result.List), nil
}
