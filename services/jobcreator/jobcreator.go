package jobcreator

import (
	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/services/dealmaker"
)

type service struct {
	dealMakerService *dealmaker.Service
	*genesis.Service
}
