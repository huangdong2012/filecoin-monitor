package model

type PackageKind string

const (
	PackageKind_Lotus   PackageKind = "lotus"
	PackageKind_Miner   PackageKind = "miner"
	PackageKind_Worker  PackageKind = "worker"
	PackageKind_Storage PackageKind = "storage"
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
