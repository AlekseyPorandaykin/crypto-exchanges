package v5

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/sender"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strconv"
	"sync"
	"testing"
	"time"
)

var testLogger = zap.L()

func init() {
	l, err := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		Encoding:          "json",
		DisableStacktrace: false,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "ts",
			LevelKey:    "level",
			MessageKey:  "message",
			LineEnding:  zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime:  zapcore.RFC3339TimeEncoder,
		},
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stdout",
			"stderr",
		},
	}.Build()
	if err != nil {
		panic(err)
	}
	testLogger = l
}

func TestClient_SpotTicker(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	resL, _ := c.MarketLinearTicker(context.TODO())
	_ = resL
	res, _ := c.MarketSpotTicker(context.TODO())
	for _, price := range res.Result.List {
		if price.Symbol == "STRKUSDT" {
			fmt.Println(price)
		}
	}
	//c.MarketLinearTicker(context.TODO())
	//c.MarketInverseTicker(context.TODO())
	//c.MarketOptionTicker(context.TODO())
}
func TestClient_MarketInstrumentsInfo(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	spotInfo, _ := c.MarketInstrumentsSpotInfo(context.TODO())
	linearInfo, _ := c.MarketInstrumentsLinearInfo(context.TODO())
	inverseInfo, _ := c.MarketInstrumentsInverseInfo(context.TODO())
	optionInfo, _ := c.MarketInstrumentsOptionInfo(context.TODO())
	fmt.Println(spotInfo, linearInfo, inverseInfo, optionInfo)
}

func TestClient_SpotCoinsBalance(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	c.WithLogger(testLogger)
	ctx := context.TODO()
	for i := 0; i < 5; i++ {
		go func() {
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
			c.FundAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
			c.OptionAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
			c.ContractAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
			c.UnifiedAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
			//fmt.Println(spotAssets, fundAssets, optionAssets, contractAssets, unifiedAssets)
		}()
	}
	spotAssets, _ := c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
	fundAssets, _ := c.FundAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
	optionAssets, _ := c.OptionAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
	contractAssets, _ := c.ContractAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
	unifiedAssets, _ := c.UnifiedAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
	fmt.Println(spotAssets, fundAssets, optionAssets, contractAssets, unifiedAssets)
}
func TestClient_GetUIDWalletType(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	data, err := c.GetUIDWalletType(context.TODO(), ApiKeyDev, ApiSecretDev)
	fmt.Println(data.Result.Accounts, err)
}
func TestClient_GetApiKey(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	c.GetApiKey(context.TODO(), "ApiKeyDev", ApiSecretDev)
}
func TestClient_GetCoinsBalance(t *testing.T) {
	ctx := context.TODO()
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	contract, err := c.AssetContractCoinsBalance(ctx, ApiKeyDev, ApiSecretDev)
	unified, err := c.AssetUnifiedCoinsBalance(ctx, ApiKeyDev, ApiSecretDev, nil)
	fund, err := c.AssetFundCoinsBalance(ctx, ApiKeyDev, ApiSecretDev)

	fmt.Println(contract, unified, fund)
}

func TestClient_PositionInfo(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.PositionInfo(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}, request.PositionInfoParam{
		Category: domain.LinearOrderCategory,
		Symbol:   "BTCUSDT",
	})
}
func TestClient_PositionMoveHistory(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	req := request.MovePositionHistoryParam{}
	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-7 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.PositionMoveHistory(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}

}

func TestClient_PositionClosedPnL(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	req := request.ClosedPnlParam{Category: "linear"}
	end := time.Now().Add(-14 * 24 * time.Hour)

	startPeriod := time.Date(2024, 6, 20, 0, 0, 0, 0, time.UTC)
	c.PositionClosedPnL(ctx, credParamDev, request.ClosedPnlParam{
		Category:  domain.LinearOrderCategory,
		StartTime: startPeriod,
		EndTime:   startPeriod.Add(7 * 24 * time.Hour),
	})
	var sumPnl float64
	for end.After(time.Now().Add(-21 * 24 * time.Hour)) {
		start := end.Add(-7 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.PositionClosedPnL(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, record := range records.Result.List {
			pnl, err := strconv.ParseFloat(record.ClosedPnl, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			sumPnl += pnl
		}
		_ = records
		end = start
	}
	fmt.Println(sumPnl)
}

func TestClient_AssetCoinExchangeRecords(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AssetCoinExchangeRecords(ctx, ApiKeyIp, ApiSecretIp)
}
func TestClient_AssetInternalTransferRecords(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AssetInternalTransferRecords(ctx, ApiKeyIp, ApiSecretIp)
}
func TestClient_GetWithdrawalRecords(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	req := request.AssetWithdrawalRecordsParam{WithdrawType: 2}

	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-30 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.AssetWithdrawalRecords(ctx, credParamDev, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}
}

func TestClient_GetDepositRecords(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	req := request.GetDepositRecordParam{}
	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-30 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.AssetDepositRecords(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}
}

func TestClient_Case(t *testing.T) {
	var Qty float64 = 0.000000234000000000
	res := strconv.FormatFloat(Qty, 'f', -1, 64)
	fmt.Println(res)
}

func TestClient_LoadCheck(t *testing.T) {
	//s := sender.NewWaitAdder(sender.NewRequestLogger(sender.NewBasic(), testLogger))
	s := sender.NewWaitAdder(sender.NewBasic())
	c, err := NewClient("https://api.bybit.com/", s)
	if err != nil {
		return
	}
	c.WithLogger(testLogger)
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	fmt.Println("start")
	//Market запросы
	for i := 0; i < 100; i++ {
		wg.Add(14)
		go func() {
			defer wg.Done()
			c.MarketSpotTicker(ctx)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.MarketLinearTicker(ctx)
		}()
		go func() {
			defer wg.Done()
			c.MarketInverseTicker(ctx)
		}()
		go func() {
			defer wg.Done()
			c.MarketOptionTicker(ctx)
		}()
		go func() {
			defer wg.Done()
			c.MarketGetLinearKlineMinute(ctx, "BTCUSDT", 60)
		}()
		go func() {
			defer wg.Done()
			c.MarketGetLinearKlineDay(ctx, "BTCUSDT")
		}()
		go func() {
			defer wg.Done()
			c.MarketGetLinearKlineWeek(ctx, "BTCUSDT")
		}()
		go func() {
			defer wg.Done()
			c.MarketGetLinearKlineMonth(ctx, "BTCUSDT")
		}()
		go func() {
			defer wg.Done()
			c.MarketInstrumentsSpotInfo(ctx)
		}()
		go func() {
			defer wg.Done()
			c.MarketInstrumentsLinearInfo(context.TODO())
		}()
		go func() {
			defer wg.Done()
			c.MarketInstrumentsInverseInfo(context.TODO())
		}()
		go func() {
			defer wg.Done()
			c.MarketInstrumentsOptionInfo(context.TODO())
		}()
		go func() {
			defer wg.Done()
			c.MarketGetOrderBook(ctx, request.MarketGetOrderBookParam{Category: domain.LinearOrderCategory, Symbol: "BTCUSDT", Limit: 500})
		}()
	}
	wg.Wait()
	fmt.Println("done")
}

func TestClient_LoadWithLimitCheck(t *testing.T) {
	//s := sender.NewWaitAdder(sender.NewRequestLogger(sender.NewBasic(), testLogger))
	s := sender.NewWaitAdder(sender.NewBasic())
	c, err := NewClient("https://api.bybit.com/", s)
	if err != nil {
		return
	}
	c.WithLogger(testLogger)
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	fmt.Println("start")
	//Market запросы
	for i := 0; i < 100; i++ {
		wg.Add(14)
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
		go func() {
			defer wg.Done()
			c.SpotAssetInfo(ctx, ApiKeyDev, ApiSecretDev)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
