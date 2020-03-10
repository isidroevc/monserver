package services

import (
	"database/sql"
	"encoding/json"

	"github.com/isidroevc/monserver/models"
)

type StoreService struct {
	database *sql.DB
}

func NewStoreService(database *sql.DB) *StoreService {
	storeService := new(StoreService)
	storeService.database = database
	return storeService
}

func (self *StoreService) UpdateStats(statsMessage *models.StatMessage) error {
	const sql = `
	REPLACE INTO
		stats
	SET
		node_id = ?,
		community_chain = ?,
		total_memory = ?,
		used_memory = ?,
		used_memory_percentage = ?,
		processors = ?,
		most_used_processor_percentage = ?,
		net_interfaces = ?,
		most_used_interface_income_bytes = ?,
		most_used_interface_outcome_bytes = ?
	`
	usedMemoryPercentaje := float64(statsMessage.UsedMemory) / float64(statsMessage.TotalMemory)
	jsonProcessorsInfo, _ := json.Marshal(statsMessage.Processors)
	mostUsedProcessorPercentage := float64(-1)
	netInterfaces, _ := json.Marshal(statsMessage.Processors)
	mostUsedInterfaceIncomeBytes := int64(0)
	for _, netInterface := range statsMessage.NetInterfaces {
		if netInterface.IncomeBytes > mostUsedInterfaceIncomeBytes {
			mostUsedInterfaceIncomeBytes = netInterface.IncomeBytes
		}
	}

	mostUsedInterfaceOutcomeBytes := int64(0)
	for _, netInterface := range statsMessage.NetInterfaces {
		if netInterface.IncomeBytes > mostUsedInterfaceOutcomeBytes {
			mostUsedInterfaceOutcomeBytes = netInterface.IncomeBytes
		}
	}
	for _, processor := range statsMessage.Processors {
		if processor > mostUsedProcessorPercentage {
			mostUsedProcessorPercentage = processor
		}
	}
	_, sqlError := self.database.Exec(sql,
		statsMessage.NodeId,
		statsMessage.CommunityChain,
		statsMessage.TotalMemory,
		statsMessage.UsedMemory,
		usedMemoryPercentaje,
		jsonProcessorsInfo,
		mostUsedProcessorPercentage,
		netInterfaces,
		mostUsedInterfaceIncomeBytes,
		mostUsedInterfaceOutcomeBytes,
	)
	if sqlError != nil {
		return sqlError
	}
	return nil
}
