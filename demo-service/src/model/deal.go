package model

import "database/sql"

//
// Deal represents Deal entity.
//
type Deal struct {
	DealId        int32          `db:"deal_id"`
	OpportunityId sql.NullString `db:"opportunity_id"`
	ContractId    int32          `db:"contract_id"`
	KwUid         string         `db:"kw_uid"`
	KwUidName     string         `db:"kw_uid_name"`
	McId          int32          `db:"mc_id"`
	McKey         int32          `db:"mc_key"`
	ChecklistIds  sql.NullString `db:"checklist_ids"`
	DealOwner     sql.NullString `db:"deal_owner"`
	DealOwnerName sql.NullString `db:"deal_owner_name"`
	DealName      string         `db:"deal_name"`
}
