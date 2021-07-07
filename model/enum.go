package model

type Role string

const (
	RoleLotus   Role = "lotus"
	RoleMiner   Role = "miner"
	RoleWorker  Role = "worker"
	RoleStorage Role = "storage"
)
