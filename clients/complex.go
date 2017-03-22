package clients

import "github.com/nk2ge5k/goyad/support"
import "github.com/nk2ge5k/goyad/gc"

type GetRequest struct {
	support.ApiObject

	FieldNames []ClientFieldEnum `json:"FieldNames"`
}

type GetResponse struct {
	support.ApiObject

	Clients *[]gc.ClientGetItem `json:"Clients,omitempty"`
}
