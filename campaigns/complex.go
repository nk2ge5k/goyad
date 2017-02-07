package campaigns

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type StrategyWeeklyBudgetBase struct {
	support.ApiObject

	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyMaximumClicks struct {
	support.ApiObject
	StrategyWeeklyBudgetBase
}

type StrategyMaximumConversionRate struct {
	support.ApiObject
	StrategyWeeklyBudgetBase

	GoalId *int64 `json:"GoalId,omitempty"`
}

type StrategyMaximumAppInstalls struct {
	support.ApiObject
	StrategyWeeklyBudgetBase
}

type StrategyAverageCpc struct {
	support.ApiObject

	AverageCpc       *int64 `json:"AverageCpc,omitempty"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
}

type StrategyAverageCpa struct {
	support.ApiObject

	AverageCpa       *int64 `json:"AverageCpa,omitempty"`
	GoalId           *int64 `json:"GoalId,omitempty"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyAverageCpi struct {
	support.ApiObject

	AverageCpi       *int64 `json:"AverageCpi,omitempty"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyAverageRoi struct {
	support.ApiObject

	ReserveReturn    *int   `json:"ReserveReturn,omitempty"`
	RoiCoef          *int64 `json:"RoiCoef,omitempty"`
	GoalId           *int64 `json:"GoalId,omitempty"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
	Profitability    *int64 `json:"Profitability,omitempty"`
}

type StrategyWeeklyClickPackage struct {
	support.ApiObject

	ClicksPerWeek *int64 `json:"ClicksPerWeek,omitempty"`
	AverageCpc    *int64 `json:"AverageCpc,omitempty"`
	BidCeiling    *int64 `json:"BidCeiling,omitempty"`
}

type StrategyNetworkDefault struct {
	support.ApiObject

	LimitPercent *int `json:"LimitPercent,omitempty"`
	BidPercent   *int `json:"BidPercent,omitempty"`
}

type TextCampaignStrategyBase struct {
	support.ApiObject

	WbMaximumClicks         *StrategyMaximumClicks         `json:"WbMaximumClicks,omitempty"`
	WbMaximumConversionRate *StrategyMaximumConversionRate `json:"WbMaximumConversionRate,omitempty"`
	AverageCpc              *StrategyAverageCpc            `json:"AverageCpc,omitempty"`
	AverageCpa              *StrategyAverageCpa            `json:"AverageCpa,omitempty"`
	WeeklyClickPackage      *StrategyWeeklyClickPackage    `json:"WeeklyClickPackage,omitempty"`
	AverageRoi              *StrategyAverageRoi            `json:"AverageRoi,omitempty"`
}

type TextCampaignNetworkStrategy struct {
	support.ApiObject
	TextCampaignStrategyBase

	BiddingStrategyType TextCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
	NetworkDefault      *StrategyNetworkDefault             `json:"NetworkDefault,omitempty"`
}

type TextCampaignSearchStrategy struct {
	support.ApiObject
	TextCampaignStrategyBase

	BiddingStrategyType TextCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type MobileAppCampaignStrategyBase struct {
	support.ApiObject

	WbMaximumClicks      *StrategyMaximumClicks      `json:"WbMaximumClicks,omitempty"`
	WbMaximumAppInstalls *StrategyMaximumAppInstalls `json:"WbMaximumAppInstalls,omitempty"`
	AverageCpc           *StrategyAverageCpc         `json:"AverageCpc,omitempty"`
	AverageCpi           *StrategyAverageCpi         `json:"AverageCpi,omitempty"`
	WeeklyClickPackage   *StrategyWeeklyClickPackage `json:"WeeklyClickPackage,omitempty"`
}

type MobileAppCampaignSearchStrategy struct {
	support.ApiObject
	MobileAppCampaignStrategyBase

	BiddingStrategyType MobileAppCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type MobileAppCampaignNetworkStrategy struct {
	support.ApiObject
	MobileAppCampaignStrategyBase

	BiddingStrategyType MobileAppCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
	NetworkDefault      *StrategyNetworkDefault                  `json:"NetworkDefault,omitempty"`
}

type DynamicTextCampaignStrategyBase struct {
	support.ApiObject

	WbMaximumClicks         *StrategyMaximumClicks         `json:"WbMaximumClicks,omitempty"`
	WbMaximumConversionRate *StrategyMaximumConversionRate `json:"WbMaximumConversionRate,omitempty"`
	AverageCpc              *StrategyAverageCpc            `json:"AverageCpc,omitempty"`
	AverageCpa              *StrategyAverageCpa            `json:"AverageCpa,omitempty"`
	WeeklyClickPackage      *StrategyWeeklyClickPackage    `json:"WeeklyClickPackage,omitempty"`
	AverageRoi              *StrategyAverageRoi            `json:"AverageRoi,omitempty"`
}

type DynamicTextCampaignSearchStrategy struct {
	support.ApiObject
	DynamicTextCampaignStrategyBase

	BiddingStrategyType DynamicTextCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type DynamicTextCampaignNetworkStrategy struct {
	support.ApiObject

	BiddingStrategyType DynamicTextCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
}

type TextCampaignStrategy struct {
	support.ApiObject

	Search  *TextCampaignSearchStrategy  `json:"Search,omitempty"`
	Network *TextCampaignNetworkStrategy `json:"Network,omitempty"`
}

type MobileAppCampaignStrategy struct {
	support.ApiObject

	Search  *MobileAppCampaignSearchStrategy  `json:"Search,omitempty"`
	Network *MobileAppCampaignNetworkStrategy `json:"Network,omitempty"`
}

type DynamicTextCampaignStrategy struct {
	support.ApiObject

	Search  *DynamicTextCampaignSearchStrategy  `json:"Search,omitempty"`
	Network *DynamicTextCampaignNetworkStrategy `json:"Network,omitempty"`
}

type StrategyWeeklyBudgetAddBase struct {
	support.ApiObject

	WeeklySpendLimit int64  `json:"WeeklySpendLimit"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyMaximumClicksAdd struct {
	support.ApiObject
	StrategyWeeklyBudgetAddBase
}

type StrategyMaximumConversionRateAdd struct {
	support.ApiObject
	StrategyWeeklyBudgetAddBase

	GoalId int64 `json:"GoalId"`
}

type StrategyMaximumAppInstallsAdd struct {
	support.ApiObject
	StrategyWeeklyBudgetAddBase
}

type StrategyAverageCpcAdd struct {
	support.ApiObject

	AverageCpc       int64  `json:"AverageCpc"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
}

type StrategyAverageCpaAdd struct {
	support.ApiObject

	AverageCpa       int64  `json:"AverageCpa"`
	GoalId           int64  `json:"GoalId"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyAverageCpiAdd struct {
	support.ApiObject

	AverageCpi       int64  `json:"AverageCpi"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
}

type StrategyAverageRoiAdd struct {
	support.ApiObject

	ReserveReturn    int    `json:"ReserveReturn"`
	RoiCoef          int64  `json:"RoiCoef"`
	GoalId           int64  `json:"GoalId"`
	WeeklySpendLimit *int64 `json:"WeeklySpendLimit,omitempty"`
	BidCeiling       *int64 `json:"BidCeiling,omitempty"`
	Profitability    *int64 `json:"Profitability,omitempty"`
}

type StrategyWeeklyClickPackageAdd struct {
	support.ApiObject

	ClicksPerWeek int64  `json:"ClicksPerWeek"`
	AverageCpc    *int64 `json:"AverageCpc,omitempty"`
	BidCeiling    *int64 `json:"BidCeiling,omitempty"`
}

type StrategyNetworkDefaultAdd struct {
	support.ApiObject

	LimitPercent *int `json:"LimitPercent,omitempty"`
	BidPercent   *int `json:"BidPercent,omitempty"`
}

type TextCampaignStrategyAddBase struct {
	support.ApiObject

	WbMaximumClicks         *StrategyMaximumClicksAdd         `json:"WbMaximumClicks,omitempty"`
	WbMaximumConversionRate *StrategyMaximumConversionRateAdd `json:"WbMaximumConversionRate,omitempty"`
	AverageCpc              *StrategyAverageCpcAdd            `json:"AverageCpc,omitempty"`
	AverageCpa              *StrategyAverageCpaAdd            `json:"AverageCpa,omitempty"`
	WeeklyClickPackage      *StrategyWeeklyClickPackageAdd    `json:"WeeklyClickPackage,omitempty"`
	AverageRoi              *StrategyAverageRoiAdd            `json:"AverageRoi,omitempty"`
}

type TextCampaignNetworkStrategyAdd struct {
	support.ApiObject
	TextCampaignStrategyAddBase

	BiddingStrategyType TextCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
	NetworkDefault      *StrategyNetworkDefaultAdd          `json:"NetworkDefault,omitempty"`
}

type TextCampaignSearchStrategyAdd struct {
	support.ApiObject
	TextCampaignStrategyAddBase

	BiddingStrategyType TextCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type MobileAppCampaignStrategyAddBase struct {
	support.ApiObject

	WbMaximumClicks      *StrategyMaximumClicksAdd      `json:"WbMaximumClicks,omitempty"`
	WbMaximumAppInstalls *StrategyMaximumAppInstallsAdd `json:"WbMaximumAppInstalls,omitempty"`
	AverageCpc           *StrategyAverageCpcAdd         `json:"AverageCpc,omitempty"`
	AverageCpi           *StrategyAverageCpiAdd         `json:"AverageCpi,omitempty"`
	WeeklyClickPackage   *StrategyWeeklyClickPackageAdd `json:"WeeklyClickPackage,omitempty"`
}

type MobileAppCampaignSearchStrategyAdd struct {
	support.ApiObject
	MobileAppCampaignStrategyAddBase

	BiddingStrategyType MobileAppCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type MobileAppCampaignNetworkStrategyAdd struct {
	support.ApiObject
	MobileAppCampaignStrategyAddBase

	BiddingStrategyType MobileAppCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
	NetworkDefault      *StrategyNetworkDefaultAdd               `json:"NetworkDefault,omitempty"`
}

type DynamicTextCampaignStrategyAddBase struct {
	support.ApiObject

	WbMaximumClicks         *StrategyMaximumClicksAdd         `json:"WbMaximumClicks,omitempty"`
	WbMaximumConversionRate *StrategyMaximumConversionRateAdd `json:"WbMaximumConversionRate,omitempty"`
	AverageCpc              *StrategyAverageCpcAdd            `json:"AverageCpc,omitempty"`
	AverageCpa              *StrategyAverageCpaAdd            `json:"AverageCpa,omitempty"`
	WeeklyClickPackage      *StrategyWeeklyClickPackageAdd    `json:"WeeklyClickPackage,omitempty"`
	AverageRoi              *StrategyAverageRoiAdd            `json:"AverageRoi,omitempty"`
}

type DynamicTextCampaignSearchStrategyAdd struct {
	support.ApiObject
	DynamicTextCampaignStrategyAddBase

	BiddingStrategyType DynamicTextCampaignSearchStrategyTypeEnum `json:"BiddingStrategyType"`
}

type DynamicTextCampaignNetworkStrategyAdd struct {
	support.ApiObject

	BiddingStrategyType DynamicTextCampaignNetworkStrategyTypeEnum `json:"BiddingStrategyType"`
}

type TextCampaignStrategyAdd struct {
	support.ApiObject

	Search  TextCampaignSearchStrategyAdd  `json:"Search"`
	Network TextCampaignNetworkStrategyAdd `json:"Network"`
}

type MobileAppCampaignStrategyAdd struct {
	support.ApiObject

	Search  MobileAppCampaignSearchStrategyAdd  `json:"Search"`
	Network MobileAppCampaignNetworkStrategyAdd `json:"Network"`
}

type DynamicTextCampaignStrategyAdd struct {
	support.ApiObject

	Search  DynamicTextCampaignSearchStrategyAdd  `json:"Search"`
	Network DynamicTextCampaignNetworkStrategyAdd `json:"Network"`
}

type Notification struct {
	support.ApiObject

	SmsSettings   *SmsSettings   `json:"SmsSettings,omitempty"`
	EmailSettings *EmailSettings `json:"EmailSettings,omitempty"`
}

type SmsSettings struct {
	support.ApiObject

	Events   *[]SmsEventsEnum `json:"Events,omitempty"`
	TimeFrom *string          `json:"TimeFrom,omitempty"`
	TimeTo   *string          `json:"TimeTo,omitempty"`
}

type EmailSettings struct {
	support.ApiObject

	Email                 *string            `json:"Email,omitempty"`
	CheckPositionInterval *int               `json:"CheckPositionInterval,omitempty"`
	WarningBalance        *int               `json:"WarningBalance,omitempty"`
	SendAccountNews       *general.YesNoEnum `json:"SendAccountNews,omitempty"`
	SendWarnings          *general.YesNoEnum `json:"SendWarnings,omitempty"`
}

type TimeTargetingOnPublicHolidays struct {
	support.ApiObject

	SuspendOnHolidays general.YesNoEnum `json:"SuspendOnHolidays"`
	BidPercent        *int              `json:"BidPercent,omitempty"`
	StartHour         *int              `json:"StartHour,omitempty"`
	EndHour           *int              `json:"EndHour,omitempty"`
}

type RelevantKeywordsSetting struct {
	support.ApiObject

	BudgetPercent  *int                      `json:"BudgetPercent,omitempty"`
	Mode           *RelevantKeywordsModeEnum `json:"Mode,omitempty"`
	OptimizeGoalId *int64                    `json:"OptimizeGoalId,omitempty"`
}

type DailyBudget struct {
	support.ApiObject

	Amount int64               `json:"Amount"`
	Mode   DailyBudgetModeEnum `json:"Mode"`
}

type TextCampaignSetting struct {
	support.ApiObject

	Option TextCampaignSettingsEnum `json:"Option"`
	Value  general.YesNoEnum        `json:"Value"`
}

type TextCampaignSettingGet struct {
	support.ApiObject

	Option TextCampaignSettingsGetEnum `json:"Option"`
	Value  general.YesNoEnum           `json:"Value"`
}

type MobileAppCampaignSetting struct {
	support.ApiObject

	Option MobileAppCampaignSettingsEnum `json:"Option"`
	Value  general.YesNoEnum             `json:"Value"`
}

type MobileAppCampaignSettingGet struct {
	support.ApiObject

	Option MobileAppCampaignSettingsGetEnum `json:"Option"`
	Value  general.YesNoEnum                `json:"Value"`
}

type DynamicTextCampaignSetting struct {
	support.ApiObject

	Option DynamicTextCampaignSettingsEnum `json:"Option"`
	Value  general.YesNoEnum               `json:"Value"`
}

type DynamicTextCampaignSettingGet struct {
	support.ApiObject

	Option DynamicTextCampaignSettingsGetEnum `json:"Option"`
	Value  general.YesNoEnum                  `json:"Value"`
}

type CampaignFundsParam struct {
	support.ApiObject

	Sum                     int64  `json:"Sum"`
	Balance                 int64  `json:"Balance"`
	BalanceBonus            int64  `json:"BalanceBonus"`
	SumAvailableForTransfer *int64 `json:"SumAvailableForTransfer,omitempty"`
}

type SharedAccountFundsParam struct {
	support.ApiObject

	Refund *int64 `json:"Refund,omitempty"`
	Spend  *int64 `json:"Spend,omitempty"`
}

type FundsParam struct {
	support.ApiObject

	Mode               CampaignFundsEnum        `json:"Mode"`
	CampaignFunds      *CampaignFundsParam      `json:"CampaignFunds,omitempty"`
	SharedAccountFunds *SharedAccountFundsParam `json:"SharedAccountFunds,omitempty"`
}

type CampaignAssistant struct {
	support.ApiObject

	Manager *string `json:"Manager,omitempty"`
	Agency  *string `json:"Agency,omitempty"`
}

type TimeTargetingBase struct {
	support.ApiObject

	Schedule                *general.ArrayOfString `json:"Schedule,omitempty"`
	ConsiderWorkingWeekends general.YesNoEnum      `json:"ConsiderWorkingWeekends"`
}

type TimeTargeting struct {
	support.ApiObject
	TimeTargetingBase

	HolidaysSchedule *TimeTargetingOnPublicHolidays `json:"HolidaysSchedule,omitempty"`
}

type TimeTargetingAdd struct {
	support.ApiObject
	TimeTargetingBase

	HolidaysSchedule *TimeTargetingOnPublicHolidays `json:"HolidaysSchedule,omitempty"`
}

type CampaignBase struct {
	support.ApiObject

	ClientInfo   *string       `json:"ClientInfo,omitempty"`
	Notification *Notification `json:"Notification,omitempty"`
	TimeZone     *string       `json:"TimeZone,omitempty"`
}

type TextCampaignBase struct {
	support.ApiObject

	CounterIds       *general.ArrayOfInteger  `json:"CounterIds,omitempty"`
	RelevantKeywords *RelevantKeywordsSetting `json:"RelevantKeywords,omitempty"`
}

type DynamicTextCampaignBase struct {
	support.ApiObject

	CounterIds *general.ArrayOfInteger `json:"CounterIds,omitempty"`
}

type TextCampaignAddItem struct {
	support.ApiObject

	BiddingStrategy  TextCampaignStrategyAdd  `json:"BiddingStrategy"`
	Settings         *[]TextCampaignSetting   `json:"Settings,omitempty"`
	CounterIds       *general.ArrayOfInteger  `json:"CounterIds,omitempty"`
	RelevantKeywords *RelevantKeywordsSetting `json:"RelevantKeywords,omitempty"`
}

type MobileAppCampaignAddItem struct {
	support.ApiObject

	BiddingStrategy MobileAppCampaignStrategyAdd `json:"BiddingStrategy"`
	Settings        *[]MobileAppCampaignSetting  `json:"Settings,omitempty"`
}

type DynamicTextCampaignAddItem struct {
	support.ApiObject

	BiddingStrategy DynamicTextCampaignStrategyAdd `json:"BiddingStrategy"`
	Settings        *[]DynamicTextCampaignSetting  `json:"Settings,omitempty"`
	CounterIds      *general.ArrayOfInteger        `json:"CounterIds,omitempty"`
}

type CampaignAddItem struct {
	support.ApiObject
	CampaignBase

	Name                string                      `json:"Name"`
	StartDate           string                      `json:"StartDate"`
	DailyBudget         *DailyBudget                `json:"DailyBudget,omitempty"`
	EndDate             *string                     `json:"EndDate,omitempty"`
	NegativeKeywords    *general.ArrayOfString      `json:"NegativeKeywords,omitempty"`
	BlockedIps          *general.ArrayOfString      `json:"BlockedIps,omitempty"`
	ExcludedSites       *general.ArrayOfString      `json:"ExcludedSites,omitempty"`
	TextCampaign        *TextCampaignAddItem        `json:"TextCampaign,omitempty"`
	MobileAppCampaign   *MobileAppCampaignAddItem   `json:"MobileAppCampaign,omitempty"`
	DynamicTextCampaign *DynamicTextCampaignAddItem `json:"DynamicTextCampaign,omitempty"`
	TimeTargeting       *TimeTargetingAdd           `json:"TimeTargeting,omitempty"`
}

type TextCampaignUpdateItem struct {
	support.ApiObject
	TextCampaignBase

	BiddingStrategy *TextCampaignStrategy  `json:"BiddingStrategy,omitempty"`
	Settings        *[]TextCampaignSetting `json:"Settings,omitempty"`
}

type MobileAppCampaignUpdateItem struct {
	support.ApiObject

	BiddingStrategy *MobileAppCampaignStrategy  `json:"BiddingStrategy,omitempty"`
	Settings        *[]MobileAppCampaignSetting `json:"Settings,omitempty"`
}

type DynamicTextCampaignUpdateItem struct {
	support.ApiObject
	DynamicTextCampaignBase

	BiddingStrategy *DynamicTextCampaignStrategy  `json:"BiddingStrategy,omitempty"`
	Settings        *[]DynamicTextCampaignSetting `json:"Settings,omitempty"`
}

type CampaignUpdateItem struct {
	support.ApiObject
	CampaignBase

	Id                  int64                          `json:"Id"`
	Name                *string                        `json:"Name,omitempty"`
	StartDate           *string                        `json:"StartDate,omitempty"`
	DailyBudget         *DailyBudget                   `json:"DailyBudget,omitempty"`
	EndDate             *string                        `json:"EndDate,omitempty"`
	NegativeKeywords    *general.ArrayOfString         `json:"NegativeKeywords,omitempty"`
	BlockedIps          *general.ArrayOfString         `json:"BlockedIps,omitempty"`
	ExcludedSites       *general.ArrayOfString         `json:"ExcludedSites,omitempty"`
	TextCampaign        *TextCampaignUpdateItem        `json:"TextCampaign,omitempty"`
	MobileAppCampaign   *MobileAppCampaignUpdateItem   `json:"MobileAppCampaign,omitempty"`
	DynamicTextCampaign *DynamicTextCampaignUpdateItem `json:"DynamicTextCampaign,omitempty"`
	TimeTargeting       *TimeTargeting                 `json:"TimeTargeting,omitempty"`
}

type TextCampaignGetItem struct {
	support.ApiObject
	TextCampaignBase

	BiddingStrategy *TextCampaignStrategy     `json:"BiddingStrategy,omitempty"`
	Settings        *[]TextCampaignSettingGet `json:"Settings,omitempty"`
}

type MobileAppCampaignGetItem struct {
	support.ApiObject

	BiddingStrategy *MobileAppCampaignStrategy     `json:"BiddingStrategy,omitempty"`
	Settings        *[]MobileAppCampaignSettingGet `json:"Settings,omitempty"`
}

type DynamicTextCampaignGetItem struct {
	support.ApiObject
	DynamicTextCampaignBase

	BiddingStrategy *DynamicTextCampaignStrategy     `json:"BiddingStrategy,omitempty"`
	Settings        *[]DynamicTextCampaignSettingGet `json:"Settings,omitempty"`
}

type CampaignGetItem struct {
	support.ApiObject
	CampaignBase

	Id                  *int64                      `json:"Id,omitempty"`
	Name                *string                     `json:"Name,omitempty"`
	StartDate           *string                     `json:"StartDate,omitempty"`
	Type                *CampaignTypeGetEnum        `json:"Type,omitempty"`
	Status              *general.StatusEnum         `json:"Status,omitempty"`
	State               *CampaignStateGetEnum       `json:"State,omitempty"`
	StatusPayment       *CampaignStatusPaymentEnum  `json:"StatusPayment,omitempty"`
	StatusClarification *string                     `json:"StatusClarification,omitempty"`
	SourceId            *int64                      `json:"SourceId,omitempty"`
	Statistics          *general.Statistics         `json:"Statistics,omitempty"`
	Currency            *general.CurrencyEnum       `json:"Currency,omitempty"`
	Funds               *FundsParam                 `json:"Funds,omitempty"`
	RepresentedBy       *CampaignAssistant          `json:"RepresentedBy,omitempty"`
	DailyBudget         *DailyBudget                `json:"DailyBudget,omitempty"`
	EndDate             *string                     `json:"EndDate,omitempty"`
	NegativeKeywords    *general.ArrayOfString      `json:"NegativeKeywords,omitempty"`
	BlockedIps          *general.ArrayOfString      `json:"BlockedIps,omitempty"`
	ExcludedSites       *general.ArrayOfString      `json:"ExcludedSites,omitempty"`
	TextCampaign        *TextCampaignGetItem        `json:"TextCampaign,omitempty"`
	MobileAppCampaign   *MobileAppCampaignGetItem   `json:"MobileAppCampaign,omitempty"`
	DynamicTextCampaign *DynamicTextCampaignGetItem `json:"DynamicTextCampaign,omitempty"`
	TimeTargeting       *TimeTargeting              `json:"TimeTargeting,omitempty"`
}

type CampaignsSelectionCriteria struct {
	support.ApiObject

	Ids             *[]int64                       `json:"Ids,omitempty"`
	Types           *[]CampaignTypeEnum            `json:"Types,omitempty"`
	States          *[]CampaignStateEnum           `json:"States,omitempty"`
	Statuses        *[]CampaignStatusSelectionEnum `json:"Statuses,omitempty"`
	StatusesPayment *[]CampaignStatusPaymentEnum   `json:"StatusesPayment,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria             CampaignsSelectionCriteria      `json:"SelectionCriteria"`
	FieldNames                    []CampaignFieldEnum             `json:"FieldNames"`
	TextCampaignFieldNames        *[]TextCampaignFieldEnum        `json:"TextCampaignFieldNames,omitempty"`
	MobileAppCampaignFieldNames   *[]MobileAppCampaignFieldEnum   `json:"MobileAppCampaignFieldNames,omitempty"`
	DynamicTextCampaignFieldNames *[]DynamicTextCampaignFieldEnum `json:"DynamicTextCampaignFieldNames,omitempty"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Campaigns *[]CampaignGetItem `json:"Campaigns,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	Campaigns []CampaignAddItem `json:"Campaigns"`
}

type AddResponse struct {
	support.ApiObject

	AddResults *[]general.ActionResult `json:"AddResults,omitempty"`
}

type UpdateRequest struct {
	support.ApiObject

	Campaigns []CampaignUpdateItem `json:"Campaigns"`
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
