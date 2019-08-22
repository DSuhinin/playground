package app

import (
	"github.com/KWRI/demo-service/src/model"
)

// DealResponse wrapper for GET /deals/{deal_id} response.
//
type DealResponse struct {
	DealId        int32  `json:"deal_id"`
	OpportunityId string `json:"opportunity_id"`
	ContractId    int32  `json:"contract_id"`
	KwUid         string `json:"kw_uid"`
	KwUidName     string `json:"kw_uid_name"`
	McId          int32  `json:"mc_id"`
	McKey         int32  `json:"mc_key"`
	ChecklistIds  string `json:"checklist_ids"`
	DealOwner     string `json:"deal_owner"`
	DealOwnerName string `json:"deal_owner_name"`
	DealName      string `json:"deal_name"`
}

//
// NewDealResponse creates new Deal response.
//
func NewDealResponse(deal *model.Deal) *DealResponse {
	return &DealResponse{
		DealId:        deal.DealId,
		OpportunityId: deal.OpportunityId.String,
		ContractId:    deal.ContractId,
		KwUid:         deal.KwUid,
		KwUidName:     deal.KwUidName,
		McId:          deal.McId,
		McKey:         deal.McKey,
		ChecklistIds:  deal.ChecklistIds.String,
		DealOwner:     deal.DealOwner.String,
		DealOwnerName: deal.DealOwnerName.String,
		DealName:      deal.DealName,
	}
}

func NewDealListResponse(dealList []model.Deal) []*DealResponse {
	dealListResponse := []*DealResponse{}
	for _, deal := range dealList {
		dealListResponse = append(dealListResponse, NewDealResponse(&deal))
	}

	return dealListResponse
}
