package agencyclients

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"
import "github.com/nk2ge5k/goyad/gc"

type AgencyClientsSelectionCriteria struct {
	support.ApiObject

	Logins   *[]string          `json:"Logins,omitempty"`
	Archived *general.YesNoEnum `json:"Archived,omitempty"`
}

type AgencyClientsAddItem struct {
	support.ApiObject

	Login        string                        `json:"Login"`
	FirstName    string                        `json:"FirstName"`
	LastName     string                        `json:"LastName"`
	CountryId    int                           `json:"CountryId"`
	Currency     general.CurrencyEnum          `json:"Currency"`
	Notification gc.NotificationGeneralClients `json:"Notification"`
	Settings     *[]gc.ClientSettingItemGet    `json:"Settings,omitempty"`
	Grants       *[]gc.Grant                   `json:"Grants,omitempty"`
}

type GetRequest struct {
	support.ApiObject
	general.GetRequestGeneral

	SelectionCriteria AgencyClientsSelectionCriteria `json:"SelectionCriteria"`
	FieldNames        []AgencyClientFieldEnum        `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject
	general.GetResponseGeneral

	Clients *[]gc.ClientGetItem `json:"Clients,omitempty"`
}

type AddRequest struct {
	support.ApiObject

	Clients []AgencyClientsAddItem `json:"Clients"`
}

type AddResponse struct {
	support.ApiObject
	general.ActionResultBase

	Login    *string `json:"Login,omitempty"`
	Password *string `json:"Password,omitempty"`
	Email    *string `json:"Email,omitempty"`
	ClientId *int64  `json:"ClientId,omitempty"`
}
