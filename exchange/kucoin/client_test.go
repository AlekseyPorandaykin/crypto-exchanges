package kucoin

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/kucoin/request"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/kucoin/sender"
	"testing"
)

func TestClient_GetAllTickers(t *testing.T) {
	c, err := NewClient("https://api.kucoin.com/", sender.New())
	if err != nil {
		return
	}
	c.GetAllTickers(context.TODO())
}
func TestClient_GetAccountSummaryInfo(t *testing.T) {
	c, err := NewClient("https://api.kucoin.com/", sender.New())
	if err != nil {
		return
	}
	data, err := c.GetAccountSummaryInfo(context.TODO(), request.Credential{
		ApiKey:     "658d014403d7180001d524eb",
		Secret:     "6ae6e3e6-1360-4dc0-bc4d-203a428327fb",
		PassPhrase: "DevCyprus",
	})
	_ = data
	fmt.Println(err)
}
func TestClient_GetAccountList(t *testing.T) {
	c, err := NewClient("https://api.kucoin.com/", sender.New())
	if err != nil {
		return
	}
	data, err := c.GetAccountList(context.TODO(), request.Credential{
		ApiKey:     "658d014403d7180001d524eb",
		Secret:     "6ae6e3e6-1360-4dc0-bc4d-203a428327fb",
		PassPhrase: "DevCyprus",
	})
	_ = data
	fmt.Println(err)
}
