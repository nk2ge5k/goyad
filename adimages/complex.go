package adimages

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type AdImageAddItem struct {
	support.ApiObject

	ImageData base64Binary `json:"ImageData"`
	Name      string       `json:"Name"`
}

type AdImageGetItem struct {
	support.ApiObject

	AdImageHash *string             `json:"AdImageHash,omitempty"`
	Name        *string             `json:"Name,omitempty"`
	Associated  *general.YesNoEnum  `json:"Associated,omitempty"`
	Type        *AdImageTypeEnum    `json:"Type,omitempty"`
	Subtype     *AdImageSubtypeEnum `json:"Subtype,omitempty"`
	OriginalUrl *string             `json:"OriginalUrl,omitempty"`
	PreviewUrl  *string             `json:"PreviewUrl,omitempty"`
}

type AdImageSelectionCriteria struct {
	support.ApiObject

	AdImageHashes *[]string          `json:"AdImageHashes,omitempty"`
	Associated    *general.YesNoEnum `json:"Associated,omitempty"`
}

type AdImageHashesCriteria struct {
	support.ApiObject

	AdImageHashes []string `json:"AdImageHashes"`
}

type AdImageActionResult struct {
	support.ApiObject
	general.ActionResultBase

	AdImageHash *string `json:"AdImageHash,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria *AdImageSelectionCriteria `json:"SelectionCriteria,omitempty"`
	FieldNames        []AdImageFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	AdImages *[]AdImageGetItem `json:"AdImages,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	AdImages []AdImageAddItem `json:"AdImages"`
}

type AddResponse struct {
	support.ApiObject

	AddResults *[]AdImageActionResult `json:"AddResults,omitempty"`
}

type DeleteRequest struct {
	support.ApiObject

	SelectionCriteria AdImageHashesCriteria `json:"SelectionCriteria"`
}

type DeleteResponse struct {
	support.ApiObject

	DeleteResults *[]AdImageActionResult `json:"DeleteResults,omitempty"`
}
