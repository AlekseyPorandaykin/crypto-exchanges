package v3

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v3/request"
	"testing"
	"time"
)

func TestClient_AssetWithdrawRecords(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	//res, err := c.AssetWithdrawRecords(context.TODO(), credParamDev, request.AssetWithdrawParam{WithdrawType: "2"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	req := request.AssetWithdrawParam{WithdrawType: "2"}
	end := time.Now()
	for end.After(time.Now().Add(-360 * 24 * time.Hour)) {
		start := end.Add(-30 * 24 * time.Hour)
		req.StartTime = start
		req.EndTime = end
		records, err := c.AssetWithdrawRecords(context.TODO(), credParamDev, req)
		if err != nil {
			fmt.Println(err)
			return
		}
		_ = records
		end = start
	}
}
