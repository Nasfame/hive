package solver

import (
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/utils"
)

func LogSolverEvent(badge string, ev SolverEvent) {
	switch ev.EventType {
	case JobOfferAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobOfferAdded", badge), fmt.Sprintf("%+v", *ev.JobOffer)).
			Msgf("")
	case ResourceOfferAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceOfferAdded", badge), fmt.Sprintf("%+v", *ev.ResourceOffer)).
			Msgf("")
	case DealAdded:
		log.Debug().
			Str(fmt.Sprintf("%s -> DealAdded", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobOfferStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceOfferStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceOfferStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case DealStateUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> DealStateUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case ResourceProviderTransactionsUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> ResourceProviderTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	case JobCreatorTransactionsUpdated:
		log.Debug().
			Str(fmt.Sprintf("%s -> JobCreatorTransactionsUpdated", badge), fmt.Sprintf("%+v", ev)).
			Msgf("")
	}
}

func ServiceLogSolverEvent(service system.Service, ev SolverEvent) {
	LogSolverEvent(system.GetServiceBadge(service), ev)
}

func GetDealsFilePath(id string) string {
	p := filepath.Join(config.Conf.GetString(enums.BACALHAU_JOBS_DIR), id)
	utils.EnsureDir(p)
	return p
}

func EnsureDealsFilePath(id string) (string, error) {
	return utils.EnsureDir(filepath.Join(config.Conf.GetString(enums.BACALHAU_JOBS_DIR), id))
}

func GetDownloadsFilePath(id string) string {
	downloadDir := config.Conf.GetString(enums.DOWNlOADS_DIR)
	return filepath.Join(downloadDir, id)
}

func EnsureDownloadsFilePath(id string) (string, error) {
	downloadDir := config.Conf.GetString(enums.DOWNlOADS_DIR)
	return utils.EnsureDir(filepath.Join(downloadDir, id))
}
