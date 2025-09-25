package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"testing"
)

func TestClient_TradeAmendOrderExternal(t *testing.T) {

	//client := bybit.NewClient().WithAuth(ApiKeyTrade, ApiSecretTrade)
	orderID := "1a24eeaf-3680-46a7-8d50-1ec0f63a8ff0"
	triggerPrice := "67000"
	//res, err := client.V5().Order().AmendOrder(bybit.V5AmendOrderParam{
	//	Category:     bybit.CategoryV5Linear,
	//	Symbol:       "BTCUSDT",
	//	OrderID:      &orderID,
	//	TriggerPrice: &triggerPrice,
	//})
	//fmt.Println(res, err)
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	param := request.AmendOrderParam{
		Category:     domain.LinearOrderCategory,
		Symbol:       "BTCUSDT",
		OrderID:      orderID,
		TriggerPrice: triggerPrice,
	}
	c.TradeAmendOrder(ctx, credParamTrade, param)
}
