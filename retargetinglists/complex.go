package retargetinglists

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type RetargetingListSelectionCriteria struct {
	support.ApiObject

	Ids *[]int64 `json:"Ids,omitempty"`
}

type RetargetingListRuleArgumentItem struct {
	support.ApiObject

	MembershipLifeSpan *int  `json:"MembershipLifeSpan,omitempty"`
	ExternalId         int64 `json:"ExternalId"`
}

type RetargetingListRuleItem struct {
	support.ApiObject

	Arguments []RetargetingListRuleArgumentItem `json:"Arguments"`
	Operator  RetargetingListRuleOperatorEnum   `json:"Operator"`
}

type RetargetingListBase struct {
	support.ApiObject

	Name        *string                    `json:"Name,omitempty"`
	Description *string                    `json:"Description,omitempty"`
	Rules       *[]RetargetingListRuleItem `json:"Rules,omitempty"`
}

type RetargetingListGetItem struct {
	support.ApiObject
	RetargetingListBase

	Id          *int64                    `json:"Id,omitempty"`
	IsAvailable *general.YesNoEnum        `json:"IsAvailable,omitempty"`
	Scope       *RetargetingListScopeEnum `json:"Scope,omitempty"`
}

type RetargetingListUpdateItem struct {
	support.ApiObject
	RetargetingListBase

	Id int64 `json:"Id"`
}

type RetargetingListAddItem struct {
	support.ApiObject

	Name        string                    `json:"Name"`
	Description *string                   `json:"Description,omitempty"`
	Rules       []RetargetingListRuleItem `json:"Rules"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria *RetargetingListSelectionCriteria `json:"SelectionCriteria,omitempty"`
	FieldNames        []RetargetingListFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	RetargetingLists *[]RetargetingListGetItem `json:"RetargetingLists,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	RetargetingLists []RetargetingListAddItem `json:"RetargetingLists"`
}

type AddResponse struct {
	support.ApiObject

	AddResults []general.ActionResult `json:"AddResults"`
}

type UpdateRequest struct {
	support.ApiObject

	RetargetingLists []RetargetingListUpdateItem `json:"RetargetingLists"`
}

type UpdateResponse struct {
	support.ApiObject

	UpdateResults []general.ActionResult `json:"UpdateResults"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults []general.ActionResult `json:"DeleteResults"`
}
