package app

import (
	"github.com/KWRI/demo-service/src/dao"
)

//
// ValidatorProvider provides an interface to work with validator.
//
type ValidatorProvider interface {
	//
	// ValidateGetDealRequest makes validation of GET /deal/{deal_id} request.
	//
	ValidateGetDealRequest(dealId string) error
}

var (
	//
	// validateID validates Deal ID.
	//
	validateID = func(ID string, repository dao.DealRepositoryProvider) error {

		if _, err := repository.Get(ID); err != nil {
			if err == dao.ErrorEntityNotFound {
				return ErrDealNotFound
			}

			return ErrInternalError.WithMessage("database error:%+v", err)
		}

		return nil
	}
)

//
// Validator represents the request validator.
//
type Validator struct {
	repository dao.DealRepositoryProvider
}

//
// NewValidator creates new validator instance.
//
func NewValidator(repository dao.DealRepositoryProvider) *Validator {
	return &Validator{
		repository: repository,
	}
}

//
// ValidateGetDealRequest makes validation of GET /deals/{deal_id} request.
//
func (v Validator) ValidateGetDealRequest(dealID string) error {

	return validateID(dealID, v.repository)
}
