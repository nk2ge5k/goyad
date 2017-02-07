package dictionaries

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type CurrenciesItem struct {
	support.ApiObject

	Currency   string          `json:"Currency"`
	Properties []ConstantsItem `json:"Properties"`
}

type MetroStationsItem struct {
	support.ApiObject

	GeoRegionId      int64  `json:"GeoRegionId"`
	MetroStationId   int64  `json:"MetroStationId"`
	MetroStationName string `json:"MetroStationName"`
}

type GeoRegionsItem struct {
	support.ApiObject

	GeoRegionId   int64  `json:"GeoRegionId"`
	GeoRegionName string `json:"GeoRegionName"`
	GeoRegionType string `json:"GeoRegionType"`
	ParentId      *int64 `json:"ParentId"`
}

type TimeZonesItem struct {
	support.ApiObject

	TimeZone     string `json:"TimeZone"`
	TimeZoneName string `json:"TimeZoneName"`
	UtcOffset    int    `json:"UtcOffset"`
}

type ConstantsItem struct {
	support.ApiObject

	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type AdCategoriesItem struct {
	support.ApiObject

	AdCategory  string `json:"AdCategory"`
	Description string `json:"Description"`
	Message     string `json:"Message"`
}

type OperationSystemVersionsItem struct {
	support.ApiObject

	OsName    string `json:"OsName"`
	OsVersion string `json:"OsVersion"`
}

type ProductivityAssertionsItem struct {
	support.ApiObject

	Reference      int    `json:"Reference"`
	Title          string `json:"Title"`
	Recommendation string `json:"Recommendation"`
}

type SupplySidePlatformsItem struct {
	support.ApiObject

	Title string `json:"Title"`
}

type InterestsItem struct {
	support.ApiObject

	InterestId   int64             `json:"InterestId"`
	ParentId     *int64            `json:"ParentId"`
	Name         string            `json:"Name"`
	IsTargetable general.YesNoEnum `json:"IsTargetable"`
}

type GetRequest struct {
	support.ApiObject

	DictionaryNames []DictionaryNameEnum `json:"DictionaryNames"`
}

type GetResponse struct {
	support.ApiObject

	Currencies              *[]CurrenciesItem              `json:"Currencies,omitempty"`
	MetroStations           *[]MetroStationsItem           `json:"MetroStations,omitempty"`
	GeoRegions              *[]GeoRegionsItem              `json:"GeoRegions,omitempty"`
	TimeZones               *[]TimeZonesItem               `json:"TimeZones,omitempty"`
	Constants               *[]ConstantsItem               `json:"Constants,omitempty"`
	AdCategories            *[]AdCategoriesItem            `json:"AdCategories,omitempty"`
	OperationSystemVersions *[]OperationSystemVersionsItem `json:"OperationSystemVersions,omitempty"`
	ProductivityAssertions  *[]ProductivityAssertionsItem  `json:"ProductivityAssertions,omitempty"`
	SupplySidePlatforms     *[]SupplySidePlatformsItem     `json:"SupplySidePlatforms,omitempty"`
	Interests               *[]InterestsItem               `json:"Interests,omitempty"`
}
