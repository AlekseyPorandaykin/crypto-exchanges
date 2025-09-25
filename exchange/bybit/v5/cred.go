package v5

import "github.com/AlekseyPorandaykin/crypto-exchanges/exchange/bybit/v5/request"

const Name = "Dev"
const ApiKeyDev = "hLaVTvLRvMJ3WR87eX"
const ApiSecretDev = "Won7ATCagt5wNNMu2zTrfLlPeCNuYl1vij6R"

const NameIp = ""
const ApiKeyIp = "yoUdN3F0L4TH7YV9gz"
const ApiSecretIp = "U8HcR8csR3UBYELh9f3EjaU5v46YjNxe87wU"

const ApiKeyTrade = "3y5RGxTO2OvPKYtC3x"
const ApiSecretTrade = "u7upwMKAHO8YbNydkYkti6yilZ5e4hChd6cc"

var credParamDev = request.CredentialParam{ApiKey: ApiKeyDev, ApiSecret: ApiSecretDev}
var credParamIp = request.CredentialParam{ApiKey: ApiKeyIp, ApiSecret: ApiSecretIp}
var credParamTrade = request.CredentialParam{ApiKey: ApiKeyTrade, ApiSecret: ApiSecretTrade}
