package general

import "github.com/nk2ge5k/goyad/support"

type ArrayOfString struct {
	support.ApiObject

	Items []string `json:"Items"`
}

type ArrayOfInteger struct {
	support.ApiObject

	Items []int `json:"Items"`
}

type ExceptionNotification struct {
	support.ApiObject

	Code    int     `json:"Code"`
	Message string  `json:"Message"`
	Details *string `json:"Details,omitempty"`
}

type LimitOffset struct {
	support.ApiObject

	Limit  *int64 `json:"Limit,omitempty"`
	Offset *int64 `json:"Offset,omitempty"`
}

type Statistics struct {
	support.ApiObject

	Clicks      int64 `json:"Clicks"`
	Impressions int64 `json:"Impressions"`
}

type IdsCriteria struct {
	support.ApiObject

	Ids []int64 `json:"Ids"`
}

type GetRequestGeneral struct {
	support.ApiObject

	Page *LimitOffset `json:"Page,omitempty"`
}

type GetResponseGeneral struct {
	support.ApiObject

	LimitedBy *int64 `json:"LimitedBy,omitempty"`
}

type ApiExceptionMessage struct {
	support.ApiObject

	requestId   string  `json:"requestId"`
	errorCode   int     `json:"errorCode"`
	errorDetail *string `json:"errorDetail,omitempty"`
}

type ActionResultBase struct {
	support.ApiObject

	Warnings *[]ExceptionNotification `json:"Warnings,omitempty"`
	Errors   *[]ExceptionNotification `json:"Errors,omitempty"`
}

type ActionResult struct {
	support.ApiObject
	ActionResultBase

	Id *int64 `json:"Id,omitempty"`
}

type SetBidsActionResult struct {
	support.ApiObject
	ActionResult

	CampaignId *int64 `json:"CampaignId,omitempty"`
	AdGroupId  *int64 `json:"AdGroupId,omitempty"`
}

type MultiIdsActionResult struct {
	support.ApiObject
	ActionResultBase

	Ids *[]int64 `json:"Ids,omitempty"`
}

type ExtensionModeration struct {
	support.ApiObject

	Status              StatusEnum `json:"Status"`
	StatusClarification *string    `json:"StatusClarification,omitempty"`
}

type FaultResponse ApiExceptionMessage
