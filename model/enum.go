package model

type Role string

const (
	Role_Lotus   Role = "lotus"
	Role_Miner   Role = "miner"
	Role_Worker  Role = "worker"
	Role_Storage Role = "storage"
)

type WorkerStatus int32

const (
	WorkerStatus_IDLE WorkerStatus = iota
	WorkerStatus_Running
	WorkerStatus_Finish
	WorkerStatus_Error
)
