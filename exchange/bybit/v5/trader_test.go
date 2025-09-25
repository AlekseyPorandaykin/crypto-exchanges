package v5

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/sender"
	"testing"
	"time"
)

func TestClient_TradeHistory(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}

	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-7 * 24 * time.Hour)
		//c.TradeHistory(ctx, credParamDev, request.TradeHistoryParam{Category: domain.SpotOrderCategory, StartTime: start, EndTime: end})
		c.TradeHistory(ctx, credParamDev, request.TradeHistoryParam{Category: domain.LinearOrderCategory, StartTime: start, EndTime: end})
		//c.TradeHistory(ctx, credParamDev, request.TradeHistoryParam{Category: domain.InverseOrderCategory, StartTime: start, EndTime: end})
		end = start
	}

}
func TestClient_TradePlaceOrder(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	param := request.PlaceOrderParam{
		Category:    domain.SpotOrderCategory,
		Symbol:      "ETHUSDT",
		PositionIdx: "0",
		Side:        domain.BuySide,
		OrderType:   domain.MarketOrderType,
		Qty:         0.004,
	}
	c.TradePlaceOrder(ctx, credParamTrade, param)
}
func TestClient_TradeAmendOrder(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	param := request.AmendOrderParam{
		Category:     domain.LinearOrderCategory,
		Symbol:       "BTCUSDT",
		OrderID:      "1a24eeaf-3680-46a7-8d50-1ec0f63a8ff0",
		TriggerPrice: "67100",
		//TpslMode:     "Full",
	}
	c.TradeAmendOrder(ctx, credParamTrade, param)
}
func TestClient_TradeSpotHistory(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.TradeSpotHistory(ctx, credParamTrade)
	c.TradeLinearHistory(ctx, credParamTrade)
	c.TradeInverseHistory(ctx, credParamTrade)
	c.TradeOptionHistory(ctx, credParamTrade)
}

func TestClient_TradeOpenOrders(t *testing.T) {
	ctx := context.TODO()
	c, err := NewClient("https://api.bybit.com/", sender.NewBasic())
	if err != nil {
		return
	}
	c.TradeOpenOrders(ctx, ApiKeyDev, ApiSecretDev, request.TradeOpenOrdersParam{
		Category: domain.LinearOrderCategory,
		Symbol:   "BTCUSDT",
	})
	c.TradeLinearOpenOrders(ctx, ApiKeyDev, ApiSecretDev)
	c.TradeSpotOpenOrders(ctx, ApiKeyDev, ApiSecretDev)
	c.TradeInverseOpenOrders(ctx, ApiKeyDev, ApiSecretDev)
	c.TradeOptionOpenOrders(ctx, ApiKeyDev, ApiSecretDev)
}

func TestClient_TradeOrderHistory(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	spotOrders, _ := c.TradeOrderHistory(ctx, credParamDev, request.TradeOrderHistoryParam{Category: domain.SpotOrderCategory})
	linearOrders, _ := c.TradeLinearOrderHistory(ctx, credParamDev)
	inverseOrders, _ := c.TradeInverseOrderHistory(ctx, credParamDev)
	optionOrder, _ := c.TradeOptionOrderHistory(ctx, credParamDev)
	fmt.Println(spotOrders, linearOrders, inverseOrders, optionOrder)

	req := request.TradeOrderHistoryParam{Category: domain.LinearOrderCategory}
	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-7 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.TradeOrderHistory(ctx, credParamDev, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}

}
