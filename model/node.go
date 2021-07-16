package model

import "time"

type NodeInfo struct {
	HostIP    string     `json:"host_ip"`    //主机IP
	HostNo    string     `json:"host_no"`    //主机编号
	Version   string     `json:"version"`    //版本号
	Status    NodeStatus `json:"status"`     //运行状态
	Desc      string     `json:"desc"`       //描述
	StartTime time.Time  `json:"start_time"` //启动时间
}

type LotusInfo struct {
	*NodeInfo
	PeerID       string       `json:"peer_id"`       //libp2p地址
	MainEpoch    int64        `json:"main_epoch"`    //主网高度
	CurrentEpoch int64        `json:"current_epoch"` //当前链节点高度
	Wallets      []WalletInfo `json:"wallets"`       //导入的钱包列表
}

type MinerInfo struct {
	*NodeInfo
	MinerID     string       `json:"miner_id"`     //矿工账号 如:t01000
	PeerID      string       `json:"peer_id"`      //libp2p地址
	LotusPeerID string       `json:"lotus"`        //所属的链节点的libp2p地址
	PledgeLoop  bool         `json:"pledge_loop"`  //是否开启刷单
	PowerInfo   *PowerInfo   `json:"power_info"`   //算力信息
	BalanceInfo *BalanceInfo `json:"balance_info"` //余额信息
}

type WorkerInfo struct {
	*NodeInfo
	MinerID string `json:"lotus"` //所属的矿工节点ID 如t01000
}

type StorageInfo struct {
	*NodeInfo
	//todo 挂载action信息
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

	ExpectBlockRate   float64 `json:"expect_block_rate"`    //预期出块率
	ExpectBlocksByDay float32 `json:"expect_blocks_by_day"` //预期每天出多少块
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
