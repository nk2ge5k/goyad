package keywords

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type KeywordsSelectionCriteria struct {
	support.ApiObject

	Ids             *[]int64                      `json:"Ids,omitempty"`
	AdGroupIds      *[]int64                      `json:"AdGroupIds,omitempty"`
	CampaignIds     *[]int64                      `json:"CampaignIds,omitempty"`
	States          *[]KeywordStateSelectionEnum  `json:"States,omitempty"`
	Statuses        *[]KeywordStatusSelectionEnum `json:"Statuses,omitempty"`
	ModifiedSince   *string                       `json:"ModifiedSince,omitempty"`
	ServingStatuses *[]general.ServingStatusEnum  `json:"ServingStatuses,omitempty"`
}

type KeywordAddItem struct {
	support.ApiObject

	Keyword          string                `json:"Keyword"`
	AdGroupId        int64                 `json:"AdGroupId"`
	Bid              *int64                `json:"Bid,omitempty"`
	ContextBid       *int64                `json:"ContextBid,omitempty"`
	StrategyPriority *general.PriorityEnum `json:"StrategyPriority,omitempty"`
	UserParam1       *string               `json:"UserParam1,omitempty"`
	UserParam2       *string               `json:"UserParam2,omitempty"`
}

type KeywordProductivity struct {
	support.ApiObject

	Value      *float32 `json:"Value,omitempty"`
	References *[]int   `json:"References,omitempty"`
}

type KeywordGetItem struct {
	support.ApiObject

	Id                *int64                     `json:"Id,omitempty"`
	Keyword           *string                    `json:"Keyword,omitempty"`
	AdGroupId         *int64                     `json:"AdGroupId,omitempty"`
	CampaignId        *int64                     `json:"CampaignId,omitempty"`
	Bid               *int64                     `json:"Bid,omitempty"`
	ContextBid        *int64                     `json:"ContextBid,omitempty"`
	StrategyPriority  *general.PriorityEnum      `json:"StrategyPriority,omitempty"`
	State             *general.StateEnum         `json:"State,omitempty"`
	Status            *general.StatusEnum        `json:"Status,omitempty"`
	UserParam1        *string                    `json:"UserParam1,omitempty"`
	UserParam2        *string                    `json:"UserParam2,omitempty"`
	Productivity      *KeywordProductivity       `json:"Productivity,omitempty"`
	StatisticsSearch  *general.Statistics        `json:"StatisticsSearch,omitempty"`
	StatisticsNetwork *general.Statistics        `json:"StatisticsNetwork,omitempty"`
	ServingStatus     *general.ServingStatusEnum `json:"ServingStatus,omitempty"`
}

type KeywordUpdateItem struct {
	support.ApiObject

	Id         int64   `json:"Id"`
	Keyword    *string `json:"Keyword,omitempty"`
	UserParam1 *string `json:"UserParam1,omitempty"`
	UserParam2 *string `json:"UserParam2,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	Keywords []KeywordAddItem `json:"Keywords"`
}

type AddResponse struct {
	support.ApiObject

	AddResults []general.ActionResult `json:"AddResults"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria KeywordsSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []KeywordFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Keywords *[]KeywordGetItem `json:"Keywords,omitempty"`
}

type UpdateRequest struct {
	support.ApiObject

	Keywords []KeywordUpdateItem `json:"Keywords"`
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
