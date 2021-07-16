package model

import "time"

type NodeInfo struct {
	Version   string       `json:"version"`    //版本号
	Status    WorkerStatus `json:"status"`     //运行状态
	Desc      string       `json:"desc"`       //描述
	StartTime time.Time    `json:"start_time"` //启动时间
}

type LotusInfo struct {
	*NodeInfo
	PeerID  string       `json:"peer_id"` //libp2p地址
	Epoch   int64        `json:"epoch"`   //当前高度
	Wallets []WalletInfo `json:"wallets"` //导入的钱包列表
}

type MinerInfo struct {
	*NodeInfo
	Miner       string       `json:"miner"`        //矿工账号 如:t01000
	PeerID      string       `json:"peer_id"`      //libp2p地址
	LotusPeerID string       `json:"lotus"`        //所属的链节点的libp2p地址
	SectorInfo  *SectorInfo  `json:"sector_info"`  //扇区信息
	PowerInfo   *PowerInfo   `json:"power_info"`   //算力信息
	WalletInfo  *WalletInfo  `json:"wallet_info"`  //钱包信息
	BalanceInfo *BalanceInfo `json:"balance_info"` //余额信息
}

type WorkerInfo struct {
	*NodeInfo
	LotusPeerID string           `json:"lotus"`        //所属的链节点的libp2p地址
	TotalCounts map[string]int64 `json:"task_counts"`  //总任务数 key:p1/p2/c1/c2等 value:数量
	ErrorCounts map[string]int64 `json:"error_counts"` //失败的任务数 key:p1/p2/c1/c2等 value:数量
}

type StorageInfo struct {
	*NodeInfo
}

type WalletInfo struct {
	Address string  `json:"address"` //地址
	Balance float64 `json:"balance"` //余额
}

type SectorInfo struct {
	Size   string           `json:"size"`   //扇区尺寸
	Total  int64            `json:"total"`  //已密封扇区总数
	Counts map[string]int32 `json:"counts"` //key: 1h/24h/error等 value: 密封扇区数量
}

type PowerInfo struct {
	Raw     string `json:"raw"`     //原始算力
	Quality string `json:"quality"` //有效算力
}

type BalanceInfo struct {
	MinerBalance   float64 `json:"miner_balance"`
	MinerVesting   float64 `json:"miner_vesting"`
	MinerPerCommit float64 `json:"miner_per_commit"`
	MinerAvailable float64 `json:"miner_available"`
	MinerPledge    float64 `json:"miner_pledge"`

	MarketBalance   float64 `json:"market_balance"`
	MarketLocked    float64 `json:"market_locked"`
	MarketAvailable float64 `json:"market_available"`

	WorkerBalance float64 `json:"worker_balance"`
	WorkerControl float64 `json:"worker_control"`

	TotalSpendable float64 `json:"total_spendable"`
}
