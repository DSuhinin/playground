package app

import (
	"github.com/KWRI/demo-service/src/dao"
)

//
// ControllerProvider provides an interface to work with Deal controller.
//
type ControllerProvider interface {

	//
	// Get returns the deal by deal ID.
	//
	Get(dealID string) (*DealResponse, error)

	//
	// GetList returns the list of deals.
	//
	GetList() ([]*DealResponse, error)
}

//
// DealController represents the Deals entity controller.
//
type DealController struct {
	validator  ValidatorProvider
	repository dao.DealRepositoryProvider
}

//
// NewDeal creates new instance of Deals entity controller.
//
func NewDeal(validator ValidatorProvider, repository dao.DealRepositoryProvider) *DealController {
	return &DealController{
		validator:  validator,
		repository: repository,
	}
}

//
// Get returns the deal by deal ID.
//
func (c DealController) Get(dealID string) (*DealResponse, error) {

	if err := c.validator.ValidateGetDealRequest(dealID); err != nil {
		return nil, err
	}

	deal, err := c.repository.Get(dealID)
	if err != nil {
		return nil, ErrInternalError.WithMessage("error while getting deal: %+v", err)
	}

	return NewDealResponse(deal), nil
}

//
// GetList returns the list of deals.
//
func (c DealController) GetList() ([]*DealResponse, error) {

	dealList, err := c.repository.GetList()
	if err != nil {
		return nil, ErrInternalError.WithMessage("error while getting deals: %+v", err)
	}

	return NewDealListResponse(dealList), nil
}
