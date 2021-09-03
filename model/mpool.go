package model

type MPoolInfo struct {
	Total *MPoolStat            `json:"total"`
	Stats map[string]*MPoolStat `json:"stats"`

	WalletAddrs []string `json:"wallet_addrs"`
}

type MPoolStat struct {
	Past   uint `json:"past"`
	Cur    uint `json:"cur"`
	Future uint `json:"future"`

	GasLimit    int64 `json:"gas_limit"`
	BelowCurrBF uint  `json:"below_curr_bf"`
}
