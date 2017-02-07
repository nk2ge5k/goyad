package ext

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type Callout struct {
	support.ApiObject

	CalloutText string `json:"CalloutText"`
}

type AdExtensionBase struct {
	support.ApiObject

	Type                *AdExtensionTypeEnum `json:"Type,omitempty"`
	Callout             *Callout             `json:"Callout,omitempty"`
	State               *general.StateEnum   `json:"State,omitempty"`
	Status              *general.StatusEnum  `json:"Status,omitempty"`
	StatusClarification *string              `json:"StatusClarification,omitempty"`
}

type AdExtension struct {
	support.ApiObject
	AdExtensionBase

	AdExtensionId *int64 `json:"AdExtensionId,omitempty"`
}

type AdExtensionSettingItem struct {
	support.ApiObject

	AdExtensionId int64                 `json:"AdExtensionId"`
	Operation     general.OperationEnum `json:"Operation"`
}

type AdExtensionSetting struct {
	support.ApiObject

	AdExtensions []AdExtensionSettingItem `json:"AdExtensions"`
}
