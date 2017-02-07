package changes

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type CampaignChangesItem struct {
	support.ApiObject

	CampaignId int64                   `json:"CampaignId"`
	ChangesIn  []CampaignChangesInEnum `json:"ChangesIn"`
}

type CheckResponseModified struct {
	support.ApiObject

	CampaignIds   *[]int64            `json:"CampaignIds,omitempty"`
	AdGroupIds    *[]int64            `json:"AdGroupIds,omitempty"`
	AdIds         *[]int64            `json:"AdIds,omitempty"`
	CampaignsStat *[]CampaignStatItem `json:"CampaignsStat,omitempty"`
}

type CampaignStatItem struct {
	support.ApiObject

	CampaignId int64  `json:"CampaignId"`
	BorderDate string `json:"BorderDate"`
}

type CheckResponseIds struct {
	support.ApiObject

	CampaignIds *[]int64 `json:"CampaignIds,omitempty"`
	AdGroupIds  *[]int64 `json:"AdGroupIds,omitempty"`
	AdIds       *[]int64 `json:"AdIds,omitempty"`
}

type CheckDictionariesRequest struct {
	support.ApiObject

	Timestamp *string `json:"Timestamp,omitempty"`
}

type CheckDictionariesResponse struct {
	support.ApiObject

	TimeZonesChanged *general.YesNoEnum `json:"TimeZonesChanged,omitempty"`
	RegionsChanged   *general.YesNoEnum `json:"RegionsChanged,omitempty"`
	InterestsChanged *general.YesNoEnum `json:"InterestsChanged,omitempty"`
	Timestamp        string             `json:"Timestamp"`
}

type CheckCampaignsRequest struct {
	support.ApiObject

	Timestamp string `json:"Timestamp"`
}

type CheckCampaignsResponse struct {
	support.ApiObject

	Campaigns *[]CampaignChangesItem `json:"Campaigns,omitempty"`
	Timestamp string                 `json:"Timestamp"`
}

type CheckRequest struct {
	support.ApiObject

	CampaignIds *[]int64         `json:"CampaignIds,omitempty"`
	AdGroupIds  *[]int64         `json:"AdGroupIds,omitempty"`
	AdIds       *[]int64         `json:"AdIds,omitempty"`
	Timestamp   string           `json:"Timestamp"`
	FieldNames  []CheckFieldEnum `json:"FieldNames"`
}

type CheckResponse struct {
	support.ApiObject

	Modified    *CheckResponseModified `json:"Modified,omitempty"`
	NotFound    *CheckResponseIds      `json:"NotFound,omitempty"`
	Unprocessed *CheckResponseIds      `json:"Unprocessed,omitempty"`
	Timestamp   *string                `json:"Timestamp,omitempty"`
}
