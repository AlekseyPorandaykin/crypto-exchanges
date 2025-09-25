package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"testing"
)

func TestClient_MarketGetKline(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.MarketGetLinearKlineMinute(ctx, "BTCUSDT", 60)
	c.MarketGetLinearKlineDay(ctx, "SOLUSDT")

}
func TestClient_MarketGetOrderBook(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	res, _ := c.MarketGetOrderBook(ctx, request.MarketGetOrderBookParam{Category: domain.LinearOrderCategory, Symbol: "BTCUSDT", Limit: 500})
	res.Result.Asks()
	res.Result.Bids()

}
