package config

import (
	"github.com/AlekseyPorandaykin/crypto-exchanges/client"
	v5 "github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5"
	"github.com/pkg/errors"
)

var _ client.ExchangeConfig = (*BybitV5Config)(nil)

type BybitV5Config struct {
	BaseUrl            string `env:"BYBIT_BASE_URL" envDefault:"https://api.bybit.com"`
	AllowLogger        bool   `env:"BYBIT_ALLOW_LOGGER" envDefault:"false"`
	AllowRequestLogger bool   `env:"BYBIT_ALLOW_REQUEST_LOGGER" envDefault:"false"`
	AllowWaitAdder     bool   `env:"BYBIT_ALLOW_WAIT_ADDER" envDefault:"true"`
}

func (b BybitV5Config) Validate() error {
	if b.BaseUrl == "" {
		return errors.New("base url is required")
	}
	return nil
}

func (b BybitV5Config) Name() string {
	return v5.ExchangeName
}
