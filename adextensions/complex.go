package adextensions

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/ext"
import "github.com/nk2ge5k/goyad/general"

type AdExtensionGetItem struct {
	support.ApiObject
	ext.AdExtensionBase

	Id         *int64             `json:"Id,omitempty"`
	Associated *general.YesNoEnum `json:"Associated,omitempty"`
}

type AdExtensionAddItem struct {
	support.ApiObject

	Callout *ext.Callout `json:"Callout,omitempty"`
}

type AdExtensionsSelectionCriteria struct {
	support.ApiObject

	Ids           *[]int64                                `json:"Ids,omitempty"`
	Types         *[]ext.AdExtensionTypeEnum              `json:"Types,omitempty"`
	States        *[]ext.AdExtensionStateSelectionEnum    `json:"States,omitempty"`
	Statuses      *[]general.ExtensionStatusSelectionEnum `json:"Statuses,omitempty"`
	ModifiedSince *string                                 `json:"ModifiedSince,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria AdExtensionsSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []AdExtensionFieldEnum        `json:"FieldNames"`
	CalloutFieldNames *[]CalloutFieldEnum           `json:"CalloutFieldNames,omitempty"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	AdExtensions *[]AdExtensionGetItem `json:"AdExtensions,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	AdExtensions []AdExtensionAddItem `json:"AdExtensions"`
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
