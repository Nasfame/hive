package solver

import (
	"archive/tar"
	"encoding/json"

	"fmt"
	"io"
	corehttp "net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/services/solver/solver/store"
)

/*
*
*
*

# Lists

*
*
*
*/
func (solverServer *solverServer) getJobOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]dto.JobOfferContainer, error) {
	query := store.GetJobOffersQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if notMatched := req.URL.Query().Get("not_matched"); notMatched == "true" {
		query.NotMatched = true
	}
	return solverServer.store.GetJobOffers(query)
}

func (solverServer *solverServer) getResourceOffers(res corehttp.ResponseWriter, req *corehttp.Request) ([]dto.ResourceOfferContainer, error) {
	query := store.GetResourceOffersQuery{}
	// if there is a job_creator query param then assign it
	if resourceProvider := req.URL.Query().Get("resource_provider"); resourceProvider != "" {
		query.ResourceProvider = resourceProvider
	}
	if active := req.URL.Query().Get("active"); active == "true" {
		query.Active = true
	}
	if notMatched := req.URL.Query().Get("not_matched"); notMatched == "true" {
		query.NotMatched = true
	}
	return solverServer.store.GetResourceOffers(query)
}

func (solverServer *solverServer) getDeals(res corehttp.ResponseWriter, req *corehttp.Request) ([]dto.DealContainer, error) {
	query := store.GetDealsQuery{}
	// if there is a job_creator query param then assign it
	if jobCreator := req.URL.Query().Get("job_creator"); jobCreator != "" {
		query.JobCreator = jobCreator
	}
	if resourceProvider := req.URL.Query().Get("resource_provider"); resourceProvider != "" {
		query.ResourceProvider = resourceProvider
	}
	if state := req.URL.Query().Get("state"); state != "" {
		query.State = state
	}
	return solverServer.store.GetDeals(query)
}

/*
*
*
*

	Getters

*
*
*
*/
func (solverServer *solverServer) getDeal(res corehttp.ResponseWriter, req *corehttp.Request) (dto.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		return dto.DealContainer{}, err
	}
	if deal == nil {
		return dto.DealContainer{}, fmt.Errorf("deal not found")
	}
	return *deal, nil
}

func (solverServer *solverServer) getResult(res corehttp.ResponseWriter, req *corehttp.Request) (dto.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	result, err := solverServer.store.GetResult(id)
	if err != nil {
		return dto.Result{}, err
	}
	if result == nil {
		return dto.Result{}, fmt.Errorf("result not found")
	}
	return *result, nil
}

/*
*
*
*

	Adders

*
*
*
*/
func (solverServer *solverServer) addJobOffer(jobOffer dto.JobOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.JobOfferContainer, error) {
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != jobOffer.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	err = dto.CheckJobOffer(jobOffer)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking job offer")
		return nil, err
	}
	return solverServer.controller.addJobOffer(jobOffer)
}

func (solverServer *solverServer) addResourceOffer(resourceOffer dto.ResourceOffer, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.ResourceOfferContainer, error) {
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != resourceOffer.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	err = dto.CheckResourceOffer(resourceOffer)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking resource offer")
		return nil, err
	}
	return solverServer.controller.addResourceOffer(resourceOffer)
}

func (solverServer *solverServer) addResult(results dto.Result, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.Result, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the resource provider can add a result
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	err = dto.CheckResult(results)
	if err != nil {
		log.Error().Err(err).Msgf("Error checking resource offer")
		return nil, err
	}
	results.DealID = id
	return solverServer.store.AddResult(results)
}

/*
*
*
*

	Updaters

*
*
*
*/
func (solverServer *solverServer) updateTransactionsResourceProvider(payload dto.DealTransactionsResourceProvider, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != deal.ResourceProvider {
		return nil, fmt.Errorf("resource provider address does not match signer address")
	}
	return solverServer.controller.updateDealTransactionsResourceProvider(id, payload)
}

func (solverServer *solverServer) updateTransactionsJobCreator(payload dto.DealTransactionsJobCreator, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != deal.JobCreator {
		return nil, fmt.Errorf("job creator address does not match signer address")
	}
	return solverServer.controller.updateDealTransactionsJobCreator(id, payload)
}

func (solverServer *solverServer) updateTransactionsMediator(payload dto.DealTransactionsMediator, res corehttp.ResponseWriter, req *corehttp.Request) (*dto.DealContainer, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	deal, err := solverServer.store.GetDeal(id)
	if err != nil {
		log.Error().Err(err).Msgf("error loading deal")
		return nil, err
	}
	if deal == nil {
		log.Error().Err(err).Msgf("deal not found")
		return nil, fmt.Errorf("deal not found")
	}
	signerAddress, err := http.GetAddressFromHeaders(req)
	if err != nil {
		log.Error().Err(err).Msgf("have error parsing user address")
		return nil, err
	}
	// only the job creator can post a job offer
	if signerAddress != deal.Mediator {
		return nil, fmt.Errorf("job creator address does not match mediator address")
	}
	return solverServer.controller.updateDealTransactionsMediator(id, payload)
}

/*
*
*
*

	Files

*
*
*
*/

func (solverServer *solverServer) downloadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() *http.HTTPError {
		deal, err := solverServer.store.GetDeal(id)
		if err != nil {
			log.Error().Err(err).Msgf("error loading deal")
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusInternalServerError,
			}
		}
		if deal == nil {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusNotFound,
			}
		}
		filesPath := GetDealsFilePath(id)
		// check if the filesPath directory exists
		if _, err := os.Stat(filesPath); os.IsNotExist(err) {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusNotFound,
			}
		}
		buf, err := system.GetTarBuffer(filesPath)
		if err != nil {
			return &http.HTTPError{
				Message:    err.Error(),
				StatusCode: corehttp.StatusInternalServerError,
			}
		}
		res.Header().Set("Content-Disposition", "attachment; filename=archive.tar")
		res.Header().Set("Content-Type", "application/x-tar")
		io.Copy(res, buf)
		return nil
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), err.StatusCode)
		return
	}
}

func (solverServer *solverServer) uploadFiles(res corehttp.ResponseWriter, req *corehttp.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := func() error {
		deal, err := solverServer.store.GetDeal(id)
		if err != nil {
			log.Error().Err(err).Msgf("error loading deal")
			return err
		}
		if deal == nil {
			log.Error().Msgf("deal not found")
			return err
		}
		signerAddress, err := http.GetAddressFromHeaders(req)
		if err != nil {
			log.Error().Err(err).Msgf("have error parsing user address")
			return err
		}
		// only the resource provider can add a result
		if signerAddress != deal.ResourceProvider {
			return fmt.Errorf("resource provider address does not match signer address")
		}
		tr := tar.NewReader(req.Body)
		uploadPath, err := EnsureDealsFilePath(id)
		if err != nil {
			return err
		}
		for {
			header, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			target := filepath.Join(uploadPath, header.Name)
			switch header.Typeflag {
			case tar.TypeDir:
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			case tar.TypeReg:
				f, err := os.Create(target)
				if err != nil {
					return err
				}
				defer f.Close()
				if _, err := io.Copy(f, tr); err != nil {
					return err
				}
			}
		}
		return nil
	}()

	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(dto.Result{
		// TODO: we need to be putting this in IPFS and calculating the CID
		DataID: id,
	})
	if err != nil {
		log.Ctx(req.Context()).Error().Msgf("error for json encoding: %s", err.Error())
		corehttp.Error(res, err.Error(), corehttp.StatusInternalServerError)
		return
	}
}
