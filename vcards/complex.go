package vcards

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type Phone struct {
	support.ApiObject

	CountryCode string  `json:"CountryCode"`
	CityCode    string  `json:"CityCode"`
	PhoneNumber string  `json:"PhoneNumber"`
	Extension   *string `json:"Extension,omitempty"`
}

type InstantMessenger struct {
	support.ApiObject

	MessengerClient string `json:"MessengerClient"`
	MessengerLogin  string `json:"MessengerLogin"`
}

type MapPoint struct {
	support.ApiObject

	X  float32 `json:"X"`
	Y  float32 `json:"Y"`
	X1 float32 `json:"X1"`
	Y1 float32 `json:"Y1"`
	X2 float32 `json:"X2"`
	Y2 float32 `json:"Y2"`
}

type VCardAddItem struct {
	support.ApiObject

	CampaignId       int64             `json:"CampaignId"`
	Country          string            `json:"Country"`
	City             string            `json:"City"`
	CompanyName      string            `json:"CompanyName"`
	WorkTime         string            `json:"WorkTime"`
	Phone            Phone             `json:"Phone"`
	Street           *string           `json:"Street,omitempty"`
	House            *string           `json:"House,omitempty"`
	Building         *string           `json:"Building,omitempty"`
	Apartment        *string           `json:"Apartment,omitempty"`
	InstantMessenger *InstantMessenger `json:"InstantMessenger,omitempty"`
	ExtraMessage     *string           `json:"ExtraMessage,omitempty"`
	ContactEmail     *string           `json:"ContactEmail,omitempty"`
	Ogrn             *string           `json:"Ogrn,omitempty"`
	MetroStationId   *int64            `json:"MetroStationId,omitempty"`
	PointOnMap       *MapPoint         `json:"PointOnMap,omitempty"`
	ContactPerson    *string           `json:"ContactPerson,omitempty"`
}

type VCardGetItem struct {
	support.ApiObject

	Id               *int64            `json:"Id,omitempty"`
	CampaignId       *int64            `json:"CampaignId,omitempty"`
	Country          *string           `json:"Country,omitempty"`
	City             *string           `json:"City,omitempty"`
	WorkTime         *string           `json:"WorkTime,omitempty"`
	Phone            *Phone            `json:"Phone,omitempty"`
	Street           *string           `json:"Street,omitempty"`
	House            *string           `json:"House,omitempty"`
	Building         *string           `json:"Building,omitempty"`
	Apartment        *string           `json:"Apartment,omitempty"`
	InstantMessenger *InstantMessenger `json:"InstantMessenger,omitempty"`
	CompanyName      *string           `json:"CompanyName,omitempty"`
	ExtraMessage     *string           `json:"ExtraMessage,omitempty"`
	ContactEmail     *string           `json:"ContactEmail,omitempty"`
	Ogrn             *string           `json:"Ogrn,omitempty"`
	MetroStationId   *int64            `json:"MetroStationId,omitempty"`
	PointOnMap       *MapPoint         `json:"PointOnMap,omitempty"`
	ContactPerson    *string           `json:"ContactPerson,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria *general.IdsCriteria `json:"SelectionCriteria,omitempty"`
	FieldNames        []VCardFieldEnum     `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	VCards *[]VCardGetItem `json:"VCards,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	VCards []VCardAddItem `json:"VCards"`
}

type AddResponse struct {
	support.ApiObject

	AddResults *[]general.ActionResult `json:"AddResults,omitempty"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria general.IdsCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults *[]general.ActionResult `json:"DeleteResults,omitempty"`
}
