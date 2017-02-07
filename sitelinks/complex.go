package sitelinks

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type SitelinksSetAddItem struct {
	support.ApiObject

	Sitelinks []Sitelink `json:"Sitelinks"`
}

type SitelinksSetGetItem struct {
	support.ApiObject

	Id        *int64      `json:"Id,omitempty"`
	Sitelinks *[]Sitelink `json:"Sitelinks,omitempty"`
}

type Sitelink struct {
	support.ApiObject

	Title       string  `json:"Title"`
	Href        string  `json:"Href"`
	Description *string `json:"Description,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria *general.IdsCriteria    `json:"SelectionCriteria,omitempty"`
	FieldNames        []SitelinksSetFieldEnum `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	SitelinksSets *[]SitelinksSetGetItem `json:"SitelinksSets,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	SitelinksSets []SitelinksSetAddItem `json:"SitelinksSets"`
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
