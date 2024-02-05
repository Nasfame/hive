package store

import (
	"github.com/CoopHive/hive/pkg/dto"
)

type GetJobOffersQuery struct {
	JobCreator string `json:"job_creator"`
	// this means job offers that have not been matched at all yet
	// the solver will use this to load only non matched resource offers

	// we use the DealID property of the jobOfferContainer to tell if it's been matched
	NotMatched bool `json:"not_matched"`
}

type GetResourceOffersQuery struct {
	ResourceProvider string `json:"resource_provider"`

	// this means "currently occupied" - any free floating resource offers count
	// as active (because they could be matched any moment)
	// any resource offers of the following states are considered active:
	// - DealNegotiating
	// - DealAgreed
	// if we hit results submitted (or anything after that point)
	// then the resource offer is no longer considered active
	// (because the compute side is done and now we are onto payment & mediation)
	// this flag is used by the resource provider to ask "give me all my active resource offers"
	// so that it knows when to post more reosurce offers to the solver
	Active bool `json:"active"`

	// this means resource offers that have not been matched at all yet
	// the solver will use this to load only non matched resource offers

	// we use the DealID property of the resourceOfferContainer to tell if it's been matched
	NotMatched bool `json:"not_matched"`
}

type GetDealsQuery struct {
	JobCreator       string `json:"job_creator"`
	ResourceProvider string `json:"resource_provider"`
	Mediator         string `json:"mediator"`

	// only deals that are in this state will be returned
	State string `json:"state"`
}

type SolverStore interface {
	AddJobOffer(jobOffer dto.JobOfferContainer) (*dto.JobOfferContainer, error)
	AddResourceOffer(jobOffer dto.ResourceOfferContainer) (*dto.ResourceOfferContainer, error)
	AddDeal(deal dto.DealContainer) (*dto.DealContainer, error)
	AddResult(result dto.Result) (*dto.Result, error)
	AddMatchDecision(resourceOffer string, jobOffer string, deal string, result bool) (*dto.MatchDecision, error)
	GetJobOffers(query GetJobOffersQuery) ([]dto.JobOfferContainer, error)
	GetResourceOffers(query GetResourceOffersQuery) ([]dto.ResourceOfferContainer, error)
	GetDeals(query GetDealsQuery) ([]dto.DealContainer, error)
	GetJobOffer(id string) (*dto.JobOfferContainer, error)
	GetResourceOffer(id string) (*dto.ResourceOfferContainer, error)
	GetDeal(id string) (*dto.DealContainer, error)
	GetResult(id string) (*dto.Result, error)
	GetMatchDecision(resourceOffer string, jobOffer string) (*dto.MatchDecision, error)
	UpdateJobOfferState(id string, dealID string, state uint8) (*dto.JobOfferContainer, error)
	UpdateResourceOfferState(id string, dealID string, state uint8) (*dto.ResourceOfferContainer, error)
	UpdateDealState(id string, state uint8) (*dto.DealContainer, error)
	UpdateDealMediator(id string, mediator string) (*dto.DealContainer, error)
	UpdateDealTransactionsJobCreator(id string, data dto.DealTransactionsJobCreator) (*dto.DealContainer, error)
	UpdateDealTransactionsResourceProvider(id string, data dto.DealTransactionsResourceProvider) (*dto.DealContainer, error)
	UpdateDealTransactionsMediator(id string, data dto.DealTransactionsMediator) (*dto.DealContainer, error)
	RemoveJobOffer(id string) error
	RemoveResourceOffer(id string) error
}
