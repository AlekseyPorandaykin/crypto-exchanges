package okx

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/okx/request"
	"testing"
)

var (
	phrase    = "DevCyprus@1"
	apikey    = "16158a0f-d52a-4ebc-aabe-885bff238ddb"
	secretkey = "33FB3E11525FB7DF3E9E86300799816B"
	IP        = ""
	Name      = "Dev"
	Access    = "Чтение"
)

var creadential = request.Credential{
	ApiKey:     apikey,
	Secret:     secretkey,
	PassPhrase: phrase,
}

func TestClient_Tickers(t *testing.T) {
	c, err := NewClient("https://www.okx.com/")
	if err != nil {
		return
	}
	c.Tickers(context.TODO())
}

func TestClient_AccountBalance(t *testing.T) {
	c, err := NewClient("https://www.okx.com/")
	if err != nil {
		return
	}
	data, err := c.FundingBalance(context.TODO(), creadential)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println((data))
}

func TestClient_TradingAccountBalance(t *testing.T) {
	c, err := NewClient("https://www.okx.com/")
	if err != nil {
		return
	}
	data, err := c.TradingAccountBalance(context.TODO(), creadential)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println((data.IsOk()))
}
