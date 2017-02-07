package audiencettargets

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type AudienceTargetSelectionCriteria struct {
	support.ApiObject

	Ids                *[]int64                   `json:"Ids,omitempty"`
	AdGroupIds         *[]int64                   `json:"AdGroupIds,omitempty"`
	CampaignIds        *[]int64                   `json:"CampaignIds,omitempty"`
	RetargetingListIds *[]int64                   `json:"RetargetingListIds,omitempty"`
	States             *[]AudienceTargetStateEnum `json:"States,omitempty"`
}

type AudienceTargetBase struct {
	support.ApiObject

	ContextBid       *int64                `json:"ContextBid,omitempty"`
	StrategyPriority *general.PriorityEnum `json:"StrategyPriority,omitempty"`
}

type AudienceTargetAddItem struct {
	support.ApiObject
	AudienceTargetBase

	AdGroupId         int64  `json:"AdGroupId"`
	RetargetingListId *int64 `json:"RetargetingListId,omitempty"`
}

type AudienceTargetGetItem struct {
	support.ApiObject
	AudienceTargetBase

	Id                *int64             `json:"Id,omitempty"`
	AdGroupId         *int64             `json:"AdGroupId,omitempty"`
	CampaignId        *int64             `json:"CampaignId,omitempty"`
	RetargetingListId *int64             `json:"RetargetingListId,omitempty"`
	State             *general.StateEnum `json:"State,omitempty"`
}

type AudienceTargetSetBidsItem struct {
	support.ApiObject
	AudienceTargetBase

	Id         *int64 `json:"Id,omitempty"`
	AdGroupId  *int64 `json:"AdGroupId,omitempty"`
	CampaignId *int64 `json:"CampaignId,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria AudienceTargetSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []AudienceTargetFieldEnum       `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	AudienceTargets *[]AudienceTargetGetItem `json:"AudienceTargets,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	AudienceTargets []AudienceTargetAddItem `json:"AudienceTargets"`
}

type AddResponse struct {
	support.ApiObject

	AddResults []general.ActionResult `json:"AddResults"`
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

	Bids []AudienceTargetSetBidsItem `json:"Bids"`
}

type SetBidsResponse struct {
	support.ApiObject

	SetBidsResults []general.SetBidsActionResult `json:"SetBidsResults"`
}
