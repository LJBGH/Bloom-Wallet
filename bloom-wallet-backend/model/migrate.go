package model

func MigrateModels() []interface{} {
	return []interface{}{
		&ProjectContract{},
		&ProjectContractEvent{},
	}
}
