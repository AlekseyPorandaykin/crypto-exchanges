package main

import (
	"context"
	"fmt"

	"github.com/AlekseyPorandaykin/crypto-exchanges/config"
	"github.com/AlekseyPorandaykin/crypto-exchanges/factory"
	"go.uber.org/zap"
)

func main() {
	cfg := config.BybitV5Config{
		BaseUrl:            "https://api.bybit.com",
		AllowLogger:        true,
		AllowRequestLogger: true,
		AllowWaitAdder:     true,
	}

	c, err := factory.NewExchangeClient(cfg, factory.WithLogger(zap.Must(zap.NewDevelopment())))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.SymbolInfo(context.TODO()))
	fmt.Println(c.Prices(context.TODO()))
	fmt.Println(c.LastMonthCandlesticks(context.TODO(), "BTCUSDT"))
}
