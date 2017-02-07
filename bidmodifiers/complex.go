package bidmodifiers

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type BidModifiersSelectionCriteria struct {
	support.ApiObject

	CampaignIds *[]int64               `json:"CampaignIds,omitempty"`
	AdGroupIds  *[]int64               `json:"AdGroupIds,omitempty"`
	Ids         *[]int64               `json:"Ids,omitempty"`
	Types       *[]BidModifierTypeEnum `json:"Types,omitempty"`
	Levels      []BidModifierLevelEnum `json:"Levels"`
}

type BidModifierGetItem struct {
	support.ApiObject

	CampaignId             *int64                     `json:"CampaignId,omitempty"`
	AdGroupId              *int64                     `json:"AdGroupId,omitempty"`
	Id                     *int64                     `json:"Id,omitempty"`
	Level                  *BidModifierLevelEnum      `json:"Level,omitempty"`
	Type                   *BidModifierTypeEnum       `json:"Type,omitempty"`
	MobileAdjustment       *MobileAdjustmentGet       `json:"MobileAdjustment,omitempty"`
	DemographicsAdjustment *DemographicsAdjustmentGet `json:"DemographicsAdjustment,omitempty"`
	RetargetingAdjustment  *RetargetingAdjustmentGet  `json:"RetargetingAdjustment,omitempty"`
}

type MobileAdjustmentGet struct {
	support.ApiObject

	BidModifier *int `json:"BidModifier,omitempty"`
}

type DemographicsAdjustmentGet struct {
	support.ApiObject

	Gender      *general.GenderEnum   `json:"Gender,omitempty"`
	Age         *general.AgeRangeEnum `json:"Age,omitempty"`
	BidModifier *int                  `json:"BidModifier,omitempty"`
	Enabled     *general.YesNoEnum    `json:"Enabled,omitempty"`
}

type RetargetingAdjustmentGet struct {
	support.ApiObject

	RetargetingConditionId *int64             `json:"RetargetingConditionId,omitempty"`
	BidModifier            *int               `json:"BidModifier,omitempty"`
	Accessible             *general.YesNoEnum `json:"Accessible,omitempty"`
	Enabled                *general.YesNoEnum `json:"Enabled,omitempty"`
}

type BidModifierAddBase struct {
	support.ApiObject

	CampaignId *int64 `json:"CampaignId,omitempty"`
	AdGroupId  *int64 `json:"AdGroupId,omitempty"`
}

type MobileAdjustmentAdd struct {
	support.ApiObject

	BidModifier int `json:"BidModifier"`
}

type DemographicsAdjustmentAdd struct {
	support.ApiObject

	Gender      *general.GenderEnum   `json:"Gender,omitempty"`
	Age         *general.AgeRangeEnum `json:"Age,omitempty"`
	BidModifier int                   `json:"BidModifier"`
}

type RetargetingAdjustmentAdd struct {
	support.ApiObject

	RetargetingConditionId int64 `json:"RetargetingConditionId"`
	BidModifier            int   `json:"BidModifier"`
}

type BidModifierAddItem struct {
	support.ApiObject
	BidModifierAddBase

	MobileAdjustment        *MobileAdjustmentAdd         `json:"MobileAdjustment,omitempty"`
	DemographicsAdjustments *[]DemographicsAdjustmentAdd `json:"DemographicsAdjustments,omitempty"`
	RetargetingAdjustments  *[]RetargetingAdjustmentAdd  `json:"RetargetingAdjustments,omitempty"`
}

type ToggleResult struct {
	support.ApiObject
	general.ActionResultBase

	CampaignId *int64                     `json:"CampaignId,omitempty"`
	AdGroupId  *int64                     `json:"AdGroupId,omitempty"`
	Type       *BidModifierToggleTypeEnum `json:"Type,omitempty"`
}

type BidModifierSetItem struct {
	support.ApiObject

	Id          int64 `json:"Id"`
	BidModifier int   `json:"BidModifier"`
}

type BidModifierToggleItem struct {
	support.ApiObject

	CampaignId *int64                    `json:"CampaignId,omitempty"`
	AdGroupId  *int64                    `json:"AdGroupId,omitempty"`
	Type       BidModifierToggleTypeEnum `json:"Type"`
	Enabled    general.YesNoEnum         `json:"Enabled"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria                BidModifiersSelectionCriteria      `json:"SelectionCriteria"`
	FieldNames                       []BidModifierFieldEnum             `json:"FieldNames"`
	MobileAdjustmentFieldNames       *[]MobileAdjustmentFieldEnum       `json:"MobileAdjustmentFieldNames,omitempty"`
	DemographicsAdjustmentFieldNames *[]DemographicsAdjustmentFieldEnum `json:"DemographicsAdjustmentFieldNames,omitempty"`
	RetargetingAdjustmentFieldNames  *[]RetargetingAdjustmentFieldEnum  `json:"RetargetingAdjustmentFieldNames,omitempty"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	BidModifiers *[]BidModifierGetItem `json:"BidModifiers,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	BidModifiers []BidModifierAddItem `json:"BidModifiers"`
}

type AddResponse struct {
	support.ApiObject

	AddResults *[]general.MultiIdsActionResult `json:"AddResults,omitempty"`
}

type SetRequest struct {
	support.ApiObject

	BidModifiers []BidModifierSetItem `json:"BidModifiers"`
}

type SetResponse struct {
	support.ApiObject

	SetResults *[]general.ActionResult `json:"SetResults,omitempty"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults *[]general.ActionResult `json:"DeleteResults,omitempty"`
}

type ToggleRequest struct {
	support.ApiObject

	BidModifierToggleItems []BidModifierToggleItem `json:"BidModifierToggleItems"`
}

type ToggleResponse struct {
	support.ApiObject

	ToggleResults *[]ToggleResult `json:"ToggleResults,omitempty"`
}
