package clients

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/general"

type EmailSubscriptionItem struct {
	support.ApiObject

	Option EmailSubscriptionEnum `json:"Option"`
	Value  general.YesNoEnum     `json:"Value"`
}

type Notification struct {
	support.ApiObject

	Email              string                  `json:"Email"`
	EmailSubscriptions []EmailSubscriptionItem `json:"EmailSubscriptions"`
	Lang               general.LangEnum        `json:"Lang"`
	SmsPhoneNumber     string                  `json:"SmsPhoneNumber"`
}

type ClientRestrictionItem struct {
	support.ApiObject

	Element ClientRestrictionEnum `json:"Element"`
	Value   int                   `json:"Value"`
}

type ClientSettingItemGet struct {
	support.ApiObject

	Option ClientSettingGetEnum `json:"Option"`
	Value  general.YesNoEnum    `json:"Value"`
}

type Grant struct {
	support.ApiObject

	Agency    *string           `json:"Agency,omitempty"`
	Privilege PrivilegeEnum     `json:"Privilege"`
	Value     general.YesNoEnum `json:"Value"`
}

type Representative struct {
	support.ApiObject

	Email string                         `json:"Email"`
	Login string                         `json:"Login"`
	Role  general.RepresentativeRoleEnum `json:"Role"`
}

type ClientBaseItem struct {
	support.ApiObject

	ClientInfo *string `json:"ClientInfo,omitempty"`
	Phone      *string `json:"Phone,omitempty"`
}

type ClientGetItem struct {
	support.ApiObject
	ClientBaseItem

	AccountQuality        *float32                 `json:"AccountQuality,omitempty"`
	Archived              *general.YesNoEnum       `json:"Archived,omitempty"`
	ClientId              *int64                   `json:"ClientId,omitempty"`
	CountryId             *int                     `json:"CountryId,omitempty"`
	CreatedAt             *string                  `json:"CreatedAt,omitempty"`
	Currency              *general.CurrencyEnum    `json:"Currency,omitempty"`
	Grants                *[]Grant                 `json:"Grants,omitempty"`
	Login                 *string                  `json:"Login,omitempty"`
	Notification          *Notification            `json:"Notification,omitempty"`
	OverdraftSumAvailable *int64                   `json:"OverdraftSumAvailable,omitempty"`
	Representatives       *[]Representative        `json:"Representatives,omitempty"`
	Restrictions          *[]ClientRestrictionItem `json:"Restrictions,omitempty"`
	Settings              *[]ClientSettingItemGet  `json:"Settings,omitempty"`
	Type                  *string                  `json:"Type,omitempty"`
	VatRate               *float32                 `json:"VatRate,omitempty"`
}

type GetRequest struct {
	support.ApiObject

	FieldNames []ClientFieldEnum `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject

	Clients *[]ClientGetItem `json:"Clients,omitempty"`
}
