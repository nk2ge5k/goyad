package dynamictextadtargets

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type WebpagesSelectionCriteria struct {
	support.ApiObject

	Ids         *[]int64                     `json:"Ids,omitempty"`
	AdGroupIds  *[]int64                     `json:"AdGroupIds,omitempty"`
	CampaignIds *[]int64                     `json:"CampaignIds,omitempty"`
	States      *[]WebpageStateSelectionEnum `json:"States,omitempty"`
}

type WebpageAddItem struct {
	support.ApiObject

	Name             string                `json:"Name"`
	AdGroupId        int64                 `json:"AdGroupId"`
	Conditions       *[]WebpageCondition   `json:"Conditions,omitempty"`
	Bid              *int64                `json:"Bid,omitempty"`
	ContextBid       *int64                `json:"ContextBid,omitempty"`
	StrategyPriority *general.PriorityEnum `json:"StrategyPriority,omitempty"`
}

type WebpageCondition struct {
	support.ApiObject

	Operand   WebpageConditionOperandEnum `json:"Operand"`
	Operator  StringConditionOperatorEnum `json:"Operator"`
	Arguments []string                    `json:"Arguments"`
}

type WebpageGetItem struct {
	support.ApiObject

	Id                  *int64                `json:"Id,omitempty"`
	AdGroupId           *int64                `json:"AdGroupId,omitempty"`
	CampaignId          *int64                `json:"CampaignId,omitempty"`
	Name                *string               `json:"Name,omitempty"`
	Bid                 *int64                `json:"Bid,omitempty"`
	ContextBid          *int64                `json:"ContextBid,omitempty"`
	StrategyPriority    *general.PriorityEnum `json:"StrategyPriority,omitempty"`
	State               *general.StateEnum    `json:"State,omitempty"`
	StatusClarification *string               `json:"StatusClarification,omitempty"`
	Conditions          *[]WebpageCondition   `json:"Conditions,omitempty"`
	ConditionType       *WebpageTypeEnum      `json:"ConditionType,omitempty"`
}

type SetBidsItem struct {
	support.ApiObject

	CampaignId       *int64                `json:"CampaignId,omitempty"`
	AdGroupId        *int64                `json:"AdGroupId,omitempty"`
	Id               *int64                `json:"Id,omitempty"`
	Bid              *int64                `json:"Bid,omitempty"`
	ContextBid       *int64                `json:"ContextBid,omitempty"`
	StrategyPriority *general.PriorityEnum `json:"StrategyPriority,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	Webpages []WebpageAddItem `json:"Webpages"`
}

type AddResponse struct {
	support.ApiObject

	AddResults []general.ActionResult `json:"AddResults"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria WebpagesSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []WebpageFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Webpages *[]WebpageGetItem `json:"Webpages,omitempty"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults []general.ActionResult `json:"DeleteResults"`
}

type SuspendRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type SuspendResponse struct {
	support.ApiObject

	SuspendResults []general.ActionResult `json:"SuspendResults"`
}

type ResumeRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type ResumeResponse struct {
	support.ApiObject

	ResumeResults []general.ActionResult `json:"ResumeResults"`
}

type SetBidsRequest struct {
	support.ApiObject

	Bids []SetBidsItem `json:"Bids"`
}

type SetBidsResponse struct {
	support.ApiObject

	SetBidsResults *[]general.SetBidsActionResult `json:"SetBidsResults,omitempty"`
}
