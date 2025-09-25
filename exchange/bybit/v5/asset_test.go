package v5

import (
	"context"
	"testing"
)

func TestClient_AssetUniversalTransferRecords(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AssetUniversalTransferRecords(ctx, ApiKeyIp, ApiSecretIp)
}

func TestClient_AssetCoinInfo(t *testing.T) {
	ctx := context.TODO()
	c, err := DefaultClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.AssetCoinsInfo(ctx, ApiKeyDev, ApiSecretDev)
}
