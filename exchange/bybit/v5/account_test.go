package v5

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"
	"testing"
	"time"
)

func TestClient_AccountFreeRate(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AccountFreeRate(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}, request.AccountFeeRateParam{
		Category: domain.LinearOrderCategory,
	})
}

func TestClient_AccountGetAccountInfo(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AccountGetAccountInfo(ctx, request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev})
}
func TestClient_AccountWalletBalance(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AccountWalletBalance(ctx, ApiKeyDev, ApiSecretDev, domain.UnifiedAccountType)
	c.AccountWalletBalance(ctx, ApiKeyDev, ApiSecretDev, domain.FundAccountType)
	c.AccountWalletBalance(ctx, ApiKeyDev, ApiSecretDev, domain.SpotAccountType)
}

func TestClient_AccountTransactionLog(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	req := request.AccountTransactionLogParam{}
	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-7 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.AccountTransactionLog(ctx, credParamIp, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}
}
