package store

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"

	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/services/solver/solver/store"

	"github.com/CoopHive/hive/pkg/jsonl"
)

type SolverStoreMemory struct {
	jobOfferMap      map[string]*dto.JobOfferContainer
	resourceOfferMap map[string]*dto.ResourceOfferContainer
	dealMap          map[string]*dto.DealContainer
	resultMap        map[string]*dto.Result
	matchDecisionMap map[string]*dto.MatchDecision
	mutex            sync.RWMutex
	logWriters       map[string]jsonl.Writer
}

func getMatchID(resourceOffer string, jobOffer string) string {
	return fmt.Sprintf("%s-%s", resourceOffer, jobOffer)
}

func NewSolverStoreMemory(conf *viper.Viper) (*SolverStoreMemory, error) {
	logWriters := make(map[string]jsonl.Writer)

	kinds := []string{"job_offers", "resource_offers", "deals", "decisions", "results"}
	for k := range kinds {
		logfile, err := os.OpenFile(fmt.Sprintf(conf.GetString(enums.APP_LOG_FILE_FORMAT), kinds[k]), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		logWriters[kinds[k]] = jsonl.NewWriter(logfile)
	}

	return &SolverStoreMemory{
		jobOfferMap:      map[string]*dto.JobOfferContainer{},
		resourceOfferMap: map[string]*dto.ResourceOfferContainer{},
		dealMap:          map[string]*dto.DealContainer{},
		resultMap:        map[string]*dto.Result{},
		matchDecisionMap: map[string]*dto.MatchDecision{},
		logWriters:       logWriters,
	}, nil
}

func (s *SolverStoreMemory) AddJobOffer(jobOffer dto.JobOfferContainer) (*dto.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.jobOfferMap[jobOffer.ID] = &jobOffer

	s.logWriters["job_offers"].Write(jobOffer)
	return &jobOffer, nil
}

func (s *SolverStoreMemory) AddResourceOffer(resourceOffer dto.ResourceOfferContainer) (*dto.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resourceOfferMap[resourceOffer.ID] = &resourceOffer

	s.logWriters["resource_offers"].Write(resourceOffer)
	return &resourceOffer, nil
}

func (s *SolverStoreMemory) AddDeal(deal dto.DealContainer) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.dealMap[deal.ID] = &deal
	s.logWriters["deals"].Write(deal)
	return &deal, nil
}

func (s *SolverStoreMemory) AddResult(result dto.Result) (*dto.Result, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[result.DealID] = &result
	s.logWriters["results"].Write(result)
	return &result, nil
}

func (s *SolverStoreMemory) AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*dto.MatchDecision, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := getMatchID(resourceOffer, jobOffer)
	_, ok := s.matchDecisionMap[id]
	if ok {
		return nil, fmt.Errorf("that match already exists")
	}
	decision := &dto.MatchDecision{
		ResourceOffer: resourceOffer,
		JobOffer:      jobOffer,
		Deal:          deal,
		Result:        result,
	}
	s.matchDecisionMap[id] = decision
	s.logWriters["decisions"].Write(decision)
	return decision, nil
}

func (s *SolverStoreMemory) GetJobOffers(query store.GetJobOffersQuery) ([]dto.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	jobOffers := []dto.JobOfferContainer{}
	for _, jobOffer := range s.jobOfferMap {
		matching := true
		if query.JobCreator != "" && jobOffer.JobCreator != query.JobCreator {
			matching = false
		}
		if query.NotMatched {
			if jobOffer.DealID != "" {
				matching = false
			}
		}
		if matching {
			jobOffers = append(jobOffers, *jobOffer)
		}
	}
	return jobOffers, nil
}

func (s *SolverStoreMemory) GetResourceOffers(query store.GetResourceOffersQuery) ([]dto.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	resourceOffers := []dto.ResourceOfferContainer{}
	for _, resourceOffer := range s.resourceOfferMap {
		matching := true
		if query.ResourceProvider != "" && resourceOffer.ResourceProvider != query.ResourceProvider {
			matching = false
		}
		if query.Active && !dto.IsActiveAgreementState(resourceOffer.State) {
			matching = false
		}
		if query.NotMatched {
			if resourceOffer.DealID != "" {
				matching = false
			}
		}
		if matching {
			resourceOffers = append(resourceOffers, *resourceOffer)
		}
	}
	return resourceOffers, nil
}

func (s *SolverStoreMemory) GetDeals(query store.GetDealsQuery) ([]dto.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deals := []dto.DealContainer{}
	queryState := uint8(0)
	if query.State != "" {
		parsedState, err := dto.GetAgreementState(query.State)
		if err != nil {
			return nil, err
		}
		queryState = parsedState
	}
	for _, deal := range s.dealMap {
		matching := true
		if query.JobCreator != "" && deal.JobCreator != query.JobCreator {
			matching = false
		}
		if query.ResourceProvider != "" && deal.ResourceProvider != query.ResourceProvider {
			matching = false
		}
		if query.Mediator != "" && deal.Mediator != query.Mediator {
			matching = false
		}
		if query.State != "" && deal.State != queryState {
			matching = false
		}
		if matching {
			deals = append(deals, *deal)
		}
	}
	return deals, nil
}

func (s *SolverStoreMemory) GetJobOffer(id string) (*dto.JobOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, nil
	}
	return jobOffer, nil
}

func (s *SolverStoreMemory) GetResourceOffer(id string) (*dto.ResourceOfferContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, nil
	}
	return resourceOffer, nil
}

func (s *SolverStoreMemory) GetDeal(id string) (*dto.DealContainer, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, nil
	}
	return deal, nil
}

func (s *SolverStoreMemory) GetResult(id string) (*dto.Result, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result, ok := s.resultMap[id]
	if !ok {
		return nil, nil
	}
	return result, nil
}

func (s *SolverStoreMemory) GetMatchDecision(resourceOffer string, jobOffer string) (*dto.MatchDecision, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	id := getMatchID(resourceOffer, jobOffer)
	decision, ok := s.matchDecisionMap[id]
	if !ok {
		return nil, nil
	}
	return decision, nil
}

func (s *SolverStoreMemory) UpdateJobOfferState(id string, dealID string, state uint8) (*dto.JobOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	jobOffer, ok := s.jobOfferMap[id]
	if !ok {
		return nil, fmt.Errorf("job offer not found: %s", id)
	}
	jobOffer.DealID = dealID
	jobOffer.State = state
	s.jobOfferMap[id] = jobOffer
	return jobOffer, nil
}

func (s *SolverStoreMemory) UpdateResourceOfferState(id string, dealID string, state uint8) (*dto.ResourceOfferContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	resourceOffer, ok := s.resourceOfferMap[id]
	if !ok {
		return nil, fmt.Errorf("resource offer not found: %s", id)
	}
	resourceOffer.DealID = dealID
	resourceOffer.State = state
	s.resourceOfferMap[id] = resourceOffer
	return resourceOffer, nil
}

func (s *SolverStoreMemory) UpdateDealState(id string, state uint8) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.State = state
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealMediator(id string, mediator string) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	deal.Mediator = mediator
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsResourceProvider(id string, data dto.DealTransactionsResourceProvider) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.ResourceProvider
	if data.Agree != "" {
		txs.Agree = data.Agree
	}
	if data.AddResult != "" {
		txs.AddResult = data.AddResult
	}
	if data.TimeoutAgree != "" {
		txs.TimeoutAgree = data.TimeoutAgree
	}
	if data.TimeoutJudgeResult != "" {
		txs.TimeoutJudgeResult = data.TimeoutJudgeResult
	}
	if data.TimeoutMediateResult != "" {
		txs.TimeoutMediateResult = data.TimeoutMediateResult
	}
	return deal, nil
}
func (s *SolverStoreMemory) UpdateDealTransactionsJobCreator(id string, data dto.DealTransactionsJobCreator) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.JobCreator
	if data.Agree != "" {
		txs.Agree = data.Agree
	}
	if data.AcceptResult != "" {
		txs.AcceptResult = data.AcceptResult
	}
	if data.CheckResult != "" {
		txs.CheckResult = data.CheckResult
	}
	if data.TimeoutAgree != "" {
		txs.TimeoutAgree = data.TimeoutAgree
	}
	if data.TimeoutSubmitResult != "" {
		txs.TimeoutSubmitResult = data.TimeoutSubmitResult
	}
	if data.TimeoutMediateResult != "" {
		txs.TimeoutMediateResult = data.TimeoutMediateResult
	}
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) UpdateDealTransactionsMediator(id string, data dto.DealTransactionsMediator) (*dto.DealContainer, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	deal, ok := s.dealMap[id]
	if !ok {
		return nil, fmt.Errorf("deal not found: %s", id)
	}
	txs := &deal.Transactions.Mediator
	if data.MediationAcceptResult != "" {
		txs.MediationAcceptResult = data.MediationAcceptResult
	}
	if data.MediationRejectResult != "" {
		txs.MediationRejectResult = data.MediationRejectResult
	}
	s.dealMap[id] = deal
	return deal, nil
}

func (s *SolverStoreMemory) RemoveJobOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.jobOfferMap, id)
	return nil
}

func (s *SolverStoreMemory) RemoveResourceOffer(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.resourceOfferMap, id)
	return nil
}

// Compile-time interface check:
var _ store.SolverStore = (*SolverStoreMemory)(nil)
