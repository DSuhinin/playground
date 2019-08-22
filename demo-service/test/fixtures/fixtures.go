package fixtures

//
// Fixtures for Functional, Integration tests.
//
const (
	DealID = "bafbd6a1506c41eeac30659fe190ba3b"
)

//
// Predefined responses for different Functional, Integration tests.
//
const (
	DealGetResponse = `{
	  "deal_id": 888,
	  "opportunity_id": "",
	  "contract_id": 0,
	  "kw_uid": "556396",
	  "kw_uid_name": "Nicole Burton",
	  "mc_id": 2,
	  "mc_key": 0,
	  "checklist_ids": "",
	  "deal_owner": "556396",
	  "deal_owner_name": "Nicole Burton",
	  "deal_name": "Buyer Deal"
	}`

	DealsListResponse = `[
	  {
		"deal_id": 761,
		"opportunity_id": "O0-761",
		"contract_id": 0,
		"kw_uid": "556396",
		"kw_uid_name": "Nicole Burton",
		"mc_id": 2,
		"mc_key": 0,
		"checklist_ids": "",
		"deal_owner": "615826",
		"deal_owner_name": "",
		"deal_name": "Buyer Deal"
	  },
	  {
		"deal_id": 762,
		"opportunity_id": "O0-762",
		"contract_id": 0,
		"kw_uid": "556396",
		"kw_uid_name": "Nicole Burton",
		"mc_id": 2,
		"mc_key": 0,
		"checklist_ids": "",
		"deal_owner": "556396",
		"deal_owner_name": "Nicole Burton",
		"deal_name": "Buyer Deal"
	  },
	  {
		"deal_id": 763,
		"opportunity_id": "O0-763",
		"contract_id": 0,
		"kw_uid": "556396",
		"kw_uid_name": "Nicole Burton",
		"mc_id": 2,
		"mc_key": 0,
		"checklist_ids": "",
		"deal_owner": "556396",
		"deal_owner_name": "Nicole Burton",
		"deal_name": "Buyer Deal"
	  }
	]`
)
