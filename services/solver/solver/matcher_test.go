package solver

import (
	"testing"

	"github.com/CoopHive/hive/pkg/dto"
)

func TestDoOffersMatch(t *testing.T) {
	services := dto.ServiceConfig{
		Solver:   "oranges",
		Mediator: []string{"apples"},
	}

	basicResourceOffer := dto.ResourceOffer{
		Spec: dto.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
		},
		DefaultPricing: dto.DealPricing{
			InstructionPrice: 10,
		},
		Mode:     dto.FixedPrice,
		Services: services,
	}

	basicJobOffer := dto.JobOffer{
		Spec: dto.MachineSpec{
			CPU: 1000,
			GPU: 1000,
			RAM: 1024,
		},
		Mode:     dto.MarketPrice,
		Services: services,
	}

	testCases := []struct {
		name          string
		resourceOffer func(offer dto.ResourceOffer) dto.ResourceOffer
		jobOffer      func(offer dto.JobOffer) dto.JobOffer
		shouldMatch   bool
	}{
		{
			name: "Basic match",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "CPU mis-match",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				offer.Spec.CPU = 2000
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Empty mediators",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				offer.Services.Mediator = []string{}
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Mis-matched mediators",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				offer.Services.Mediator = []string{"apples2"}
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Different but matching mediators",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				offer.Services.Mediator = []string{"apples2", "apples"}
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				return offer
			},
			shouldMatch: true,
		},
		{
			name: "Different solver",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				offer.Services.Solver = "pears"
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - too expensive",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				offer.Mode = dto.FixedPrice
				offer.Pricing.InstructionPrice = 9
				return offer
			},
			shouldMatch: false,
		},
		{
			name: "Fixed price - can afford",
			resourceOffer: func(offer dto.ResourceOffer) dto.ResourceOffer {
				return offer
			},
			jobOffer: func(offer dto.JobOffer) dto.JobOffer {
				offer.Mode = dto.FixedPrice
				offer.Pricing.InstructionPrice = 11
				return offer
			},
			shouldMatch: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := doOffersMatch(tc.resourceOffer(basicResourceOffer), tc.jobOffer(basicJobOffer))
			if result != tc.shouldMatch {
				t.Errorf("Expected match to be %v, but got %v", tc.shouldMatch, result)
			}
		})
	}
}
