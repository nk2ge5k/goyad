package adgroups

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type AdGroupsSelectionCriteria struct {
	support.ApiObject

	CampaignIds     *[]int64                             `json:"CampaignIds,omitempty"`
	Ids             *[]int64                             `json:"Ids,omitempty"`
	Types           *[]AdGroupTypesEnum                  `json:"Types,omitempty"`
	Statuses        *[]AdGroupStatusSelectionEnum        `json:"Statuses,omitempty"`
	TagIds          *[]int64                             `json:"TagIds,omitempty"`
	Tags            *[]string                            `json:"Tags,omitempty"`
	AppIconStatuses *[]AdGroupAppIconStatusSelectionEnum `json:"AppIconStatuses,omitempty"`
	ServingStatuses *[]general.ServingStatusEnum         `json:"ServingStatuses,omitempty"`
}

type AdGroupBase struct {
	support.ApiObject

	RegionIds        *[]int64               `json:"RegionIds,omitempty"`
	NegativeKeywords *general.ArrayOfString `json:"NegativeKeywords,omitempty"`
	TrackingParams   *string                `json:"TrackingParams,omitempty"`
}

type AdGroupAddItem struct {
	support.ApiObject

	Name               string                 `json:"Name"`
	CampaignId         int64                  `json:"CampaignId"`
	RegionIds          []int64                `json:"RegionIds"`
	NegativeKeywords   *general.ArrayOfString `json:"NegativeKeywords,omitempty"`
	TrackingParams     *string                `json:"TrackingParams,omitempty"`
	MobileAppAdGroup   *MobileAppAdGroupAdd   `json:"MobileAppAdGroup,omitempty"`
	DynamicTextAdGroup *DynamicTextAdGroup    `json:"DynamicTextAdGroup,omitempty"`
}

type AdGroupGetItem struct {
	support.ApiObject
	AdGroupBase

	Id                     *int64                     `json:"Id,omitempty"`
	Name                   *string                    `json:"Name,omitempty"`
	CampaignId             *int64                     `json:"CampaignId,omitempty"`
	Status                 *general.StatusEnum        `json:"Status,omitempty"`
	Type                   *AdGroupTypesEnum          `json:"Type,omitempty"`
	Subtype                *AdGroupSubtypeEnum        `json:"Subtype,omitempty"`
	MobileAppAdGroup       *MobileAppAdGroupGet       `json:"MobileAppAdGroup,omitempty"`
	DynamicTextAdGroup     *DynamicTextAdGroupGet     `json:"DynamicTextAdGroup,omitempty"`
	DynamicTextFeedAdGroup *DynamicTextFeedAdGroupGet `json:"DynamicTextFeedAdGroup,omitempty"`
	ServingStatus          *general.ServingStatusEnum `json:"ServingStatus,omitempty"`
}

type MobileAppAdGroupGet struct {
	support.ApiObject

	StoreUrl                     *string                                `json:"StoreUrl,omitempty"`
	TargetDeviceType             *[]TargetDeviceTypeEnum                `json:"TargetDeviceType,omitempty"`
	TargetCarrier                *TargetCarrierEnum                     `json:"TargetCarrier,omitempty"`
	TargetOperatingSystemVersion *string                                `json:"TargetOperatingSystemVersion,omitempty"`
	AppIconModeration            *general.ExtensionModeration           `json:"AppIconModeration,omitempty"`
	AppOperatingSystemType       *general.MobileOperatingSystemTypeEnum `json:"AppOperatingSystemType,omitempty"`
	AppAvailabilityStatus        *AppAvailabilityStatusEnum             `json:"AppAvailabilityStatus,omitempty"`
}

type MobileAppAdGroupUpdate struct {
	support.ApiObject

	TargetDeviceType             *[]TargetDeviceTypeEnum `json:"TargetDeviceType,omitempty"`
	TargetCarrier                *TargetCarrierEnum      `json:"TargetCarrier,omitempty"`
	TargetOperatingSystemVersion *string                 `json:"TargetOperatingSystemVersion,omitempty"`
}

type MobileAppAdGroupAdd struct {
	support.ApiObject

	StoreUrl                     string                 `json:"StoreUrl"`
	TargetDeviceType             []TargetDeviceTypeEnum `json:"TargetDeviceType"`
	TargetCarrier                TargetCarrierEnum      `json:"TargetCarrier"`
	TargetOperatingSystemVersion string                 `json:"TargetOperatingSystemVersion"`
}

type DynamicTextAdGroup struct {
	support.ApiObject

	DomainUrl string `json:"DomainUrl"`
}

type DynamicTextAdGroupGet struct {
	support.ApiObject

	DomainUrl                 *string                     `json:"DomainUrl,omitempty"`
	DomainUrlProcessingStatus *SourceProcessingStatusEnum `json:"DomainUrlProcessingStatus,omitempty"`
}

type DynamicSourceGet struct {
	support.ApiObject

	Source                 *string                     `json:"Source,omitempty"`
	SourceType             *SourceTypeGetEnum          `json:"SourceType,omitempty"`
	SourceProcessingStatus *SourceProcessingStatusEnum `json:"SourceProcessingStatus,omitempty"`
}

type DynamicTextFeedAdGroupGet struct {
	support.ApiObject
	DynamicSourceGet
}

type AdGroupUpdateItem struct {
	support.ApiObject
	AdGroupBase

	Id                 int64                   `json:"Id"`
	Name               *string                 `json:"Name,omitempty"`
	MobileAppAdGroup   *MobileAppAdGroupUpdate `json:"MobileAppAdGroup,omitempty"`
	DynamicTextAdGroup *DynamicTextAdGroup     `json:"DynamicTextAdGroup,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria                AdGroupsSelectionCriteria          `json:"SelectionCriteria"`
	FieldNames                       []AdGroupFieldEnum                 `json:"FieldNames"`
	MobileAppAdGroupFieldNames       *[]MobileAppAdGroupFieldEnum       `json:"MobileAppAdGroupFieldNames,omitempty"`
	DynamicTextAdGroupFieldNames     *[]DynamicTextAdGroupFieldEnum     `json:"DynamicTextAdGroupFieldNames,omitempty"`
	DynamicTextFeedAdGroupFieldNames *[]DynamicTextFeedAdGroupFieldEnum `json:"DynamicTextFeedAdGroupFieldNames,omitempty"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	AdGroups *[]AdGroupGetItem `json:"AdGroups,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	AdGroups []AdGroupAddItem `json:"AdGroups"`
}

type AddResponse struct {
	support.ApiObject

	AddResults []general.ActionResult `json:"AddResults"`
}

type UpdateRequest struct {
	support.ApiObject

	AdGroups []AdGroupUpdateItem `json:"AdGroups"`
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
