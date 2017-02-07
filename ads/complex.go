package ads

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/ext"
import "github.com/nk2ge5k/goyad/general"

type ArrayOfAdCategoryEnum struct {
	support.ApiObject

	Items []AdCategoryEnum `json:"Items"`
}

type AdExtensionAdGetItem struct {
	support.ApiObject

	AdExtensionId int64                   `json:"AdExtensionId"`
	Type          ext.AdExtensionTypeEnum `json:"Type"`
}

type MobileAppAdFeatureItem struct {
	support.ApiObject

	Feature MobileAppFeatureEnum `json:"Feature"`
	Enabled general.YesNoEnum    `json:"Enabled"`
}

type MobileAppAdFeatureGetItem struct {
	support.ApiObject
	MobileAppAdFeatureItem

	IsAvailable general.YesNoUnknownEnum `json:"IsAvailable"`
}

type MobileAppAdBase struct {
	support.ApiObject

	Title       *string                        `json:"Title,omitempty"`
	Text        *string                        `json:"Text,omitempty"`
	TrackingUrl *string                        `json:"TrackingUrl,omitempty"`
	Action      *general.MobileAppAdActionEnum `json:"Action,omitempty"`
	AdImageHash *string                        `json:"AdImageHash,omitempty"`
}

type AdAddItemBase struct {
	support.ApiObject

	AdGroupId int64 `json:"AdGroupId"`
}

type TextAdAddBase struct {
	support.ApiObject

	VCardId        *int64   `json:"VCardId,omitempty"`
	AdImageHash    *string  `json:"AdImageHash,omitempty"`
	SitelinkSetId  *int64   `json:"SitelinkSetId,omitempty"`
	AdExtensionIds *[]int64 `json:"AdExtensionIds,omitempty"`
}

type AdAddItem struct {
	support.ApiObject
	AdAddItemBase

	TextAd           *TextAdAdd           `json:"TextAd,omitempty"`
	DynamicTextAd    *DynamicTextAdAdd    `json:"DynamicTextAd,omitempty"`
	MobileAppAd      *MobileAppAdAdd      `json:"MobileAppAd,omitempty"`
	TextImageAd      *TextImageAdAdd      `json:"TextImageAd,omitempty"`
	MobileAppImageAd *MobileAppImageAdAdd `json:"MobileAppImageAd,omitempty"`
}

type TextAdAdd struct {
	support.ApiObject
	TextAdAddBase

	Text           string            `json:"Text"`
	Title          string            `json:"Title"`
	Href           *string           `json:"Href,omitempty"`
	Mobile         general.YesNoEnum `json:"Mobile"`
	DisplayUrlPath *string           `json:"DisplayUrlPath,omitempty"`
}

type DynamicTextAdAdd struct {
	support.ApiObject
	TextAdAddBase

	Text string `json:"Text"`
}

type MobileAppAdAdd struct {
	support.ApiObject

	AdImageHash *string                       `json:"AdImageHash,omitempty"`
	Text        string                        `json:"Text"`
	Title       string                        `json:"Title"`
	TrackingUrl *string                       `json:"TrackingUrl,omitempty"`
	Action      general.MobileAppAdActionEnum `json:"Action"`
	Features    *[]MobileAppAdFeatureItem     `json:"Features,omitempty"`
	AgeLabel    *MobAppAgeLabelEnum           `json:"AgeLabel,omitempty"`
}

type ImageAdAddBase struct {
	support.ApiObject

	AdImageHash string `json:"AdImageHash"`
}

type TextImageAdAdd struct {
	support.ApiObject
	ImageAdAddBase

	Href string `json:"Href"`
}

type MobileAppImageAdAdd struct {
	support.ApiObject
	ImageAdAddBase

	TrackingUrl *string `json:"TrackingUrl,omitempty"`
}

type AdUpdateItem struct {
	support.ApiObject

	Id               int64                   `json:"Id"`
	TextAd           *TextAdUpdate           `json:"TextAd,omitempty"`
	DynamicTextAd    *DynamicTextAdUpdate    `json:"DynamicTextAd,omitempty"`
	MobileAppAd      *MobileAppAdUpdate      `json:"MobileAppAd,omitempty"`
	TextImageAd      *TextImageAdUpdate      `json:"TextImageAd,omitempty"`
	MobileAppImageAd *MobileAppImageAdUpdate `json:"MobileAppImageAd,omitempty"`
}

type TextAdUpdateBase struct {
	support.ApiObject

	VCardId        *int64                  `json:"VCardId,omitempty"`
	AdImageHash    *string                 `json:"AdImageHash,omitempty"`
	SitelinkSetId  *int64                  `json:"SitelinkSetId,omitempty"`
	CalloutSetting *ext.AdExtensionSetting `json:"CalloutSetting,omitempty"`
}

type TextAdUpdate struct {
	support.ApiObject
	TextAdUpdateBase

	Text           *string       `json:"Text,omitempty"`
	Title          *string       `json:"Title,omitempty"`
	Href           *string       `json:"Href,omitempty"`
	AgeLabel       *AgeLabelEnum `json:"AgeLabel,omitempty"`
	DisplayUrlPath *string       `json:"DisplayUrlPath,omitempty"`
}

type DynamicTextAdUpdate struct {
	support.ApiObject
	TextAdUpdateBase

	Text *string `json:"Text,omitempty"`
}

type MobileAppAdUpdate struct {
	support.ApiObject
	MobileAppAdBase

	Features *[]MobileAppAdFeatureItem `json:"Features,omitempty"`
	AgeLabel *MobAppAgeLabelEnum       `json:"AgeLabel,omitempty"`
}

type ImageAdUpdateBase struct {
	support.ApiObject

	AdImageHash *string `json:"AdImageHash,omitempty"`
}

type TextImageAdUpdate struct {
	support.ApiObject
	ImageAdUpdateBase

	Href *string `json:"Href,omitempty"`
}

type MobileAppImageAdUpdate struct {
	support.ApiObject
	ImageAdUpdateBase

	TrackingUrl *string `json:"TrackingUrl,omitempty"`
}

type AdsSelectionCriteria struct {
	support.ApiObject

	Ids                         *[]int64                                `json:"Ids,omitempty"`
	States                      *[]AdStateSelectionEnum                 `json:"States,omitempty"`
	Statuses                    *[]AdStatusSelectionEnum                `json:"Statuses,omitempty"`
	CampaignIds                 *[]int64                                `json:"CampaignIds,omitempty"`
	AdGroupIds                  *[]int64                                `json:"AdGroupIds,omitempty"`
	Types                       *[]AdTypeEnum                           `json:"Types,omitempty"`
	Mobile                      *general.YesNoEnum                      `json:"Mobile,omitempty"`
	VCardIds                    *[]int64                                `json:"VCardIds,omitempty"`
	SitelinkSetIds              *[]int64                                `json:"SitelinkSetIds,omitempty"`
	AdImageHashes               *[]string                               `json:"AdImageHashes,omitempty"`
	VCardModerationStatuses     *[]general.ExtensionStatusSelectionEnum `json:"VCardModerationStatuses,omitempty"`
	SitelinksModerationStatuses *[]general.ExtensionStatusSelectionEnum `json:"SitelinksModerationStatuses,omitempty"`
	AdImageModerationStatuses   *[]general.ExtensionStatusSelectionEnum `json:"AdImageModerationStatuses,omitempty"`
	AdExtensionIds              *[]int64                                `json:"AdExtensionIds,omitempty"`
}

type TextAdGetBase struct {
	support.ApiObject

	VCardId             *int64                       `json:"VCardId,omitempty"`
	AdImageHash         *string                      `json:"AdImageHash,omitempty"`
	SitelinkSetId       *int64                       `json:"SitelinkSetId,omitempty"`
	VCardModeration     *general.ExtensionModeration `json:"VCardModeration,omitempty"`
	SitelinksModeration *general.ExtensionModeration `json:"SitelinksModeration,omitempty"`
	AdImageModeration   *general.ExtensionModeration `json:"AdImageModeration,omitempty"`
	AdExtensions        *[]AdExtensionAdGetItem      `json:"AdExtensions,omitempty"`
}

type TextAdGet struct {
	support.ApiObject
	TextAdGetBase

	Text                     *string                      `json:"Text,omitempty"`
	Title                    *string                      `json:"Title,omitempty"`
	Href                     *string                      `json:"Href,omitempty"`
	Mobile                   *general.YesNoEnum           `json:"Mobile,omitempty"`
	DisplayDomain            *string                      `json:"DisplayDomain,omitempty"`
	DisplayUrlPath           *string                      `json:"DisplayUrlPath,omitempty"`
	DisplayUrlPathModeration *general.ExtensionModeration `json:"DisplayUrlPathModeration,omitempty"`
}

type DynamicTextAdGet struct {
	support.ApiObject
	TextAdGetBase

	Text *string `json:"Text,omitempty"`
}

type MobileAppAdGet struct {
	support.ApiObject
	MobileAppAdBase

	Features          *[]MobileAppAdFeatureGetItem `json:"Features,omitempty"`
	AdImageModeration *general.ExtensionModeration `json:"AdImageModeration,omitempty"`
}

type ImageAdGetBase struct {
	support.ApiObject

	AdImageHash *string `json:"AdImageHash,omitempty"`
}

type TextImageAdGet struct {
	support.ApiObject
	ImageAdGetBase

	Href *string `json:"Href,omitempty"`
}

type MobileAppImageAdGet struct {
	support.ApiObject
	ImageAdGetBase

	TrackingUrl *string `json:"TrackingUrl,omitempty"`
}

type AdGetItem struct {
	support.ApiObject

	Id                  *int64                 `json:"Id,omitempty"`
	CampaignId          *int64                 `json:"CampaignId,omitempty"`
	AdGroupId           *int64                 `json:"AdGroupId,omitempty"`
	Status              *general.StatusEnum    `json:"Status,omitempty"`
	State               *general.StateEnum     `json:"State,omitempty"`
	StatusClarification *string                `json:"StatusClarification,omitempty"`
	AdCategories        *ArrayOfAdCategoryEnum `json:"AdCategories,omitempty"`
	AgeLabel            *AgeLabelEnum          `json:"AgeLabel,omitempty"`
	Type                *AdTypeEnum            `json:"Type,omitempty"`
	Subtype             *AdSubtypeEnum         `json:"Subtype,omitempty"`
	TextAd              *TextAdGet             `json:"TextAd,omitempty"`
	DynamicTextAd       *DynamicTextAdGet      `json:"DynamicTextAd,omitempty"`
	MobileAppAd         *MobileAppAdGet        `json:"MobileAppAd,omitempty"`
	TextImageAd         *TextImageAdGet        `json:"TextImageAd,omitempty"`
	MobileAppImageAd    *MobileAppImageAdGet   `json:"MobileAppImageAd,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria          AdsSelectionCriteria         `json:"SelectionCriteria"`
	FieldNames                 []AdFieldEnum                `json:"FieldNames"`
	TextAdFieldNames           *[]TextAdFieldEnum           `json:"TextAdFieldNames,omitempty"`
	MobileAppAdFieldNames      *[]MobileAppAdFieldEnum      `json:"MobileAppAdFieldNames,omitempty"`
	DynamicTextAdFieldNames    *[]DynamicTextAdFieldEnum    `json:"DynamicTextAdFieldNames,omitempty"`
	TextImageAdFieldNames      *[]TextImageAdFieldEnum      `json:"TextImageAdFieldNames,omitempty"`
	MobileAppImageAdFieldNames *[]MobileAppImageAdFieldEnum `json:"MobileAppImageAdFieldNames,omitempty"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Ads *[]AdGetItem `json:"Ads,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	Ads []AdAddItem `json:"Ads"`
}

type AddResponse struct {
	support.ApiObject

	AddResults *[]general.ActionResult `json:"AddResults,omitempty"`
}

type UpdateRequest struct {
	support.ApiObject

	Ads []AdUpdateItem `json:"Ads"`
}

type UpdateResponse struct {
	support.ApiObject

	UpdateResults *[]general.ActionResult `json:"UpdateResults,omitempty"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults *[]general.ActionResult `json:"DeleteResults,omitempty"`
}

type ArchiveRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type ArchiveResponse struct {
	support.ApiObject

	ArchiveResults *[]general.ActionResult `json:"ArchiveResults,omitempty"`
}

type UnarchiveRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type UnarchiveResponse struct {
	support.ApiObject

	UnarchiveResults *[]general.ActionResult `json:"UnarchiveResults,omitempty"`
}

type SuspendRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type SuspendResponse struct {
	support.ApiObject

	SuspendResults *[]general.ActionResult `json:"SuspendResults,omitempty"`
}

type ResumeRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type ResumeResponse struct {
	support.ApiObject

	ResumeResults *[]general.ActionResult `json:"ResumeResults,omitempty"`
}

type ModerateRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type ModerateResponse struct {
	support.ApiObject

	ModerateResults *[]general.ActionResult `json:"ModerateResults,omitempty"`
}
