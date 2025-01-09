package config

import (
	"math/big"
	"sc-service/smartcontract/ethbc"
)

var ETHClientInfo *ethbc.ETHClientInfo
var ETHClient ethbc.ETH_Client

func InitETHClient() {
	// URL dan private key cari di database

	ETHClientInfo = &ethbc.ETHClientInfo{
		NetURL:   "http://localhost:8545",
		PvKeyHex: "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		ChainID:  big.NewInt(1337),
	}
	ETHClient = ethbc.NewETHClient(*ETHClientInfo) // Memanggil fungsi InitClient di dalam pkg untuk membuat koneksi
}
func InitQuorumClient() {
	// URL dan private key cari di database

	ETHClientInfo = &ethbc.ETHClientInfo{
		NetURL:   "http://localhost:8545",
		PvKeyHex: "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		ChainID:  big.NewInt(1337),
	}
	ETHClient = ethbc.NewETHClient(*ETHClientInfo) // Memanggil fungsi InitClient di dalam pkg untuk membuat koneksi
}
func InitHyperLedgerClient() {
	// URL dan private key cari di database

	ETHClientInfo = &ethbc.ETHClientInfo{
		NetURL:   "http://localhost:8545",
		PvKeyHex: "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		ChainID:  big.NewInt(1337),
	}
	ETHClient = ethbc.NewETHClient(*ETHClientInfo) // Memanggil fungsi InitClient di dalam pkg untuk membuat koneksi
}
