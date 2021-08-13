package model

type PackageKind string

const (
	PackageKind_Lotus        PackageKind = "lotus"
	PackageKind_Miner        PackageKind = "miner"
	PackageKind_Worker       PackageKind = "worker"
	PackageKind_WorkerWdPost PackageKind = "worker_wdpost"
	PackageKind_WorkerWnPost PackageKind = "worker_wnpost"
	PackageKind_Storage      PackageKind = "storage"
	PackageKind_Agent        PackageKind = "agent"
)

type NodeStatus int32

const (
	NodeStatus_Online NodeStatus = iota
	NodeStatus_Offline
)

const (
	TaskStatus_IDLE = iota
	TaskStatus_Running
	TaskStatus_Finish
	TaskStatus_Error
)
