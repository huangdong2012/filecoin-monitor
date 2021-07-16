package model

type Role string

const (
	Role_Lotus   Role = "lotus"
	Role_Miner   Role = "miner"
	Role_Worker  Role = "worker"
	Role_Storage Role = "storage"
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
