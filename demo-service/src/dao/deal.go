package dao

import (
	"github.com/jmoiron/sqlx"

	"github.com/KWRI/demo-service/core/log"
	"github.com/KWRI/demo-service/src/dao/adapter"
	"github.com/KWRI/demo-service/src/model"
)

//
// DealRepositoryProvider provides an interface to work with Deal repository.
//
type DealRepositoryProvider interface {
	//
	// Get retrieves an existing Deal entity by Deal ID.
	//
	Get(dealID string) (*model.Deal, error)

	//
	// GetList retrieves deals list.
	//
	GetList() ([]model.Deal, error)
}

//
// DealRepository represents the repository to work with Deal entity.
//
type DealRepository struct {
	logger   log.Logger
	database *sqlx.DB
}

//
// NewDealRepository returns a new repository instance.
//
func NewDealRepository(logger log.Logger, database adapter.Provider) *DealRepository {

	return &DealRepository{
		logger:   logger,
		database: database.GetDB(),
	}
}

//
// Get retrieves an existing Deal entity by Deal ID.
//
func (r DealRepository) Get(dealID string) (*model.Deal, error) {

	deal := model.Deal{}
	if err := r.database.Get(&deal, `
		SELECT
			deal_id,
			opportunity_id,
			contract_id,
			kw_uid,
			kw_uid_name,
			mc_id,
			mc_key,
			checklist_ids,
			deal_owner,
			deal_owner_name,
			deal_name
		FROM OM_Deals
		WHERE deal_id = ?`, dealID,
	); err != nil {
		if err.Error() == ErrorEntityNotFound.Error() {
			return nil, ErrorEntityNotFound
		}

		return nil, ErrorInternalError.WithMessage("impossible to fetch deal:%+v", err)
	}

	return &deal, nil
}

//
// GetList retrieves deals list.
//
func (r DealRepository) GetList() ([]model.Deal, error) {

	dealList := []model.Deal{}
	if err := r.database.Select(&dealList, `
		SELECT
			deal_id,
			opportunity_id,
			contract_id,
			kw_uid,
			kw_uid_name,
			mc_id,
			mc_key,
			checklist_ids,
			deal_owner,
			deal_owner_name,
			deal_name 
		FROM OM_Deals
		LIMIT 3`,
	); err != nil {
		return nil, ErrorInternalError.WithMessage("impossible to fetch deals:%+v", err)
	}

	return dealList, nil
}
