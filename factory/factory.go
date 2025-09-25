package factory

import (
	"fmt"
	"net/http"

	"github.com/AlekseyPorandaykin/crypto-exchanges/client"
	"github.com/AlekseyPorandaykin/crypto-exchanges/config"
	v5 "github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5"
	"github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/sender"
	"go.uber.org/zap"
)

type DefaultParams struct {
	Logger           *zap.Logger
	HttpRoundTripper http.RoundTripper
}

type Option func(*DefaultParams)

func WithLogger(l *zap.Logger) Option {
	return func(p *DefaultParams) {
		p.Logger = l
	}
}

func WithHttpRoundTripper(rt http.RoundTripper) Option {
	return func(p *DefaultParams) {
		if rt == nil {
			return
		}
		p.HttpRoundTripper = rt
	}
}

func NewExchangeClient(cfg client.ExchangeConfig, options ...Option) (client.ExchangeClient, error) {
	defaultParams := &DefaultParams{
		Logger: zap.NewNop(),
	}
	for _, option := range options {
		option(defaultParams)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("exhchange=%s, invalid config: %w", cfg.Name(), err)
	}

	switch cfg.(type) {
	case config.BybitV5Config:
		bybitCfg := cfg.(config.BybitV5Config)
		basic := sender.NewBasic()
		if defaultParams.HttpRoundTripper != nil {
			basic.WithHttpTransport(defaultParams.HttpRoundTripper)
		}
		var senderImpl sender.Sender = basic
		if bybitCfg.AllowRequestLogger {
			senderImpl = sender.NewRequestLogger(senderImpl, defaultParams.Logger.Named("request"))
		}
		if bybitCfg.AllowWaitAdder {
			senderImpl = sender.NewWaitAdder(senderImpl)
		}
		c, err := v5.NewClient(bybitCfg.BaseUrl, senderImpl)
		if err != nil {
			return nil, err
		}
		if bybitCfg.AllowLogger {
			c.WithLogger(defaultParams.Logger.Named("bybit_v5"))
		}
		return v5.NewAdapter(c), nil
	}
	return nil, nil
}
