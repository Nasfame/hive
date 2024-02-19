package solver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/services/solver/solver/store"

	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/system"
)

type SolverClient struct {
	options         http.ClientOptions
	solverEventSubs []func(SolverEvent)
}

func NewSolverClient(
	options http.ClientOptions,
) (*SolverClient, error) {
	client := &SolverClient{
		options:         options,
		solverEventSubs: []func(SolverEvent){},
	}
	return client, nil
}

// connect the websocket to the solver server
func (client *SolverClient) Start(ctx context.Context, cm *system.CleanupManager) error {
	websocketEventChannel := make(chan []byte)
	go func() {
		for {
			select {
			case evBytes := <-websocketEventChannel:
				// parse the ev into a new SolverEvent
				var ev SolverEvent
				if err := json.Unmarshal(evBytes, &ev); err != nil {
					log.Error().Msgf("Error unmarshalling event: %s", err.Error())
					continue
				}
				// loop over each event channel and write the event to it
				for _, handler := range client.solverEventSubs {
					go handler(ev)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	http.ConnectWebSocket(
		http.WebsocketURL(client.options, config.WEBSOCKET_SUB_PATH),
		websocketEventChannel,
		ctx,
	)
	return nil
}

func (client *SolverClient) SubscribeEvents(handler func(SolverEvent)) {
	client.solverEventSubs = append(client.solverEventSubs, handler)
}

func (client *SolverClient) GetJobOffers(query store.GetJobOffersQuery) ([]dto.JobOfferContainer, error) {
	queryParams := map[string]string{}
	if query.JobCreator != "" {
		queryParams["job_creator"] = query.JobCreator
	}
	if query.NotMatched {
		queryParams["not_matched"] = "true"
	}
	return http.GetRequest[[]dto.JobOfferContainer](client.options, "/job_offers", queryParams)
}

func (client *SolverClient) GetResourceOffers(query store.GetResourceOffersQuery) ([]dto.ResourceOfferContainer, error) {
	queryParams := map[string]string{}
	if query.ResourceProvider != "" {
		queryParams["resource_provider"] = query.ResourceProvider
	}
	if query.Active {
		queryParams["active"] = "true"
	}
	if query.NotMatched {
		queryParams["not_matched"] = "true"
	}
	return http.GetRequest[[]dto.ResourceOfferContainer](client.options, "/resource_offers", queryParams)
}

func (client *SolverClient) GetDeals(query store.GetDealsQuery) ([]dto.DealContainer, error) {
	queryParams := map[string]string{}
	if query.JobCreator != "" {
		queryParams["job_creator"] = query.JobCreator
	}
	if query.ResourceProvider != "" {
		queryParams["resource_provider"] = query.ResourceProvider
	}
	if query.State != "" {
		queryParams["state"] = query.State
	}
	return http.GetRequest[[]dto.DealContainer](client.options, "/deals", queryParams)
}

func (client *SolverClient) GetDeal(id string) (dto.DealContainer, error) {
	return http.GetRequest[dto.DealContainer](client.options, fmt.Sprintf("/deals/%s", id), map[string]string{})
}

func (client *SolverClient) GetResult(id string) (dto.Result, error) {
	return http.GetRequest[dto.Result](client.options, fmt.Sprintf("/deals/%s/result", id), map[string]string{})
}

func (client *SolverClient) GetDealsWithFilter(query store.GetDealsQuery, filter func(dto.DealContainer) bool) ([]dto.DealContainer, error) {
	deals, err := client.GetDeals(query)
	if err != nil {
		return nil, err
	}
	ret := []dto.DealContainer{}
	for _, deal := range deals {
		if filter(deal) {
			ret = append(ret, deal)
		}
	}
	return ret, nil
}

func (client *SolverClient) AddJobOffer(jobOffer dto.JobOffer) (dto.JobOfferContainer, error) {
	return http.PostRequest[dto.JobOffer, dto.JobOfferContainer](client.options, "/job_offers", jobOffer)
}

func (client *SolverClient) AddResourceOffer(resourceOffer dto.ResourceOffer) (dto.ResourceOfferContainer, error) {
	return http.PostRequest[dto.ResourceOffer, dto.ResourceOfferContainer](client.options, "/resource_offers", resourceOffer)
}

func (client *SolverClient) AddResult(result dto.Result) (dto.Result, error) {
	return http.PostRequest[dto.Result, dto.Result](client.options, fmt.Sprintf("/deals/%s/result", result.DealID), result)
}

func (client *SolverClient) UpdateTransactionsResourceProvider(id string, payload dto.DealTransactionsResourceProvider) (dto.DealContainer, error) {
	return http.PostRequest[dto.DealTransactionsResourceProvider, dto.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/resource_provider", id), payload)
}

func (client *SolverClient) UpdateTransactionsJobCreator(id string, payload dto.DealTransactionsJobCreator) (dto.DealContainer, error) {
	return http.PostRequest[dto.DealTransactionsJobCreator, dto.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/job_creator", id), payload)
}

func (client *SolverClient) UpdateTransactionsMediator(id string, payload dto.DealTransactionsMediator) (dto.DealContainer, error) {
	return http.PostRequest[dto.DealTransactionsMediator, dto.DealContainer](client.options, fmt.Sprintf("/deals/%s/txs/mediator", id), payload)
}

func (client *SolverClient) UploadResultFiles(id string, localPath string) (dto.Result, error) {
	buf, err := system.GetTarBuffer(localPath)
	if err != nil {
		return dto.Result{}, err
	}
	return http.PostRequestBuffer[dto.Result](client.options, fmt.Sprintf("/deals/%s/files", id), buf)
}

func (client *SolverClient) DownloadResultFiles(id string, localPath string) error {
	buf, err := http.GetRequestBuffer(client.options, fmt.Sprintf("/deals/%s/files", id), map[string]string{})
	if err != nil {
		return err
	}
	return system.ExpandTarBuffer(buf, localPath)
}
