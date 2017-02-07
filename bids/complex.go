package bids

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type BidActionResult struct {
	support.ApiObject
	BidBase

	Warnings *[]general.ExceptionNotification `json:"Warnings,omitempty"`
	Errors   *[]general.ExceptionNotification `json:"Errors,omitempty"`
}

type BidsSelectionCriteria struct {
	support.ApiObject

	CampaignIds     *[]int64                     `json:"CampaignIds,omitempty"`
	AdGroupIds      *[]int64                     `json:"AdGroupIds,omitempty"`
	KeywordIds      *[]int64                     `json:"KeywordIds,omitempty"`
	ServingStatuses *[]general.ServingStatusEnum `json:"ServingStatuses,omitempty"`
}

type SearchPrices struct {
	support.ApiObject

	Position *general.PositionEnum `json:"Position,omitempty"`
	Price    *int64                `json:"Price,omitempty"`
}

type ContextCoverage struct {
	support.ApiObject

	Items *[]ContextCoverageItem `json:"Items,omitempty"`
}

type ContextCoverageItem struct {
	support.ApiObject

	Probability float32 `json:"Probability"`
	Price       int64   `json:"Price"`
}

type BidBase struct {
	support.ApiObject

	CampaignId *int64 `json:"CampaignId,omitempty"`
	AdGroupId  *int64 `json:"AdGroupId,omitempty"`
	KeywordId  *int64 `json:"KeywordId,omitempty"`
}

type BidGetItem struct {
	support.ApiObject
	BidBase

	Bid                *int64                     `json:"Bid,omitempty"`
	ContextBid         *int64                     `json:"ContextBid,omitempty"`
	StrategyPriority   *general.PriorityEnum      `json:"StrategyPriority,omitempty"`
	CompetitorsBids    *[]int64                   `json:"CompetitorsBids,omitempty"`
	SearchPrices       *[]SearchPrices            `json:"SearchPrices,omitempty"`
	ContextCoverage    *ContextCoverage           `json:"ContextCoverage,omitempty"`
	MinSearchPrice     *int64                     `json:"MinSearchPrice,omitempty"`
	CurrentSearchPrice *int64                     `json:"CurrentSearchPrice,omitempty"`
	AuctionBids        *[]AuctionBidItem          `json:"AuctionBids,omitempty"`
	ServingStatus      *general.ServingStatusEnum `json:"ServingStatus,omitempty"`
}

type AuctionBidItem struct {
	support.ApiObject

	Position string `json:"Position"`
	Bid      int64  `json:"Bid"`
	Price    int64  `json:"Price"`
}

type BidSetItem struct {
	support.ApiObject

	CampaignId       *int64                `json:"CampaignId,omitempty"`
	AdGroupId        *int64                `json:"AdGroupId,omitempty"`
	KeywordId        *int64                `json:"KeywordId,omitempty"`
	Bid              *int64                `json:"Bid,omitempty"`
	ContextBid       *int64                `json:"ContextBid,omitempty"`
	StrategyPriority *general.PriorityEnum `json:"StrategyPriority,omitempty"`
}

type BidSetAutoItem struct {
	support.ApiObject

	CampaignId      *int64                `json:"CampaignId,omitempty"`
	AdGroupId       *int64                `json:"AdGroupId,omitempty"`
	KeywordId       *int64                `json:"KeywordId,omitempty"`
	MaxBid          *int64                `json:"MaxBid,omitempty"`
	Position        *general.PositionEnum `json:"Position,omitempty"`
	IncreasePercent *int                  `json:"IncreasePercent,omitempty"`
	CalculateBy     *CalculateByEnum      `json:"CalculateBy,omitempty"`
	ContextCoverage *int                  `json:"ContextCoverage,omitempty"`
	Scope           []general.ScopeEnum   `json:"Scope"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria BidsSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []BidFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Bids *[]BidGetItem `json:"Bids,omitempty"`
}

type SetRequest struct {
	support.ApiObject

	Bids []BidSetItem `json:"Bids"`
}

type SetResponse struct {
	support.ApiObject

	SetResults *[]BidActionResult `json:"SetResults,omitempty"`
}

type SetAutoRequest struct {
	support.ApiObject

	Bids []BidSetAutoItem `json:"Bids"`
}

type SetAutoResponse struct {
	support.ApiObject

	SetAutoResults *[]BidActionResult `json:"SetAutoResults,omitempty"`
}
