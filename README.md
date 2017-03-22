# goyad
Yandex Direct API Go library (v5)

# Basic usage 

```go
import (
	"github.com/nk2ge5k/goyad"
	"github.com/nk2ge5k/goyad/campaigns"
)

func main() {
	client := goyad.NewClient(CLIENT_LOGIN, API_TOKEN)

	request := campaigns.GetRequest{
		FieldNames: []campaigns.CampaignFieldEnum{
			"BlockedIps",
			"ClientInfo",
			"Currency",
			"DailyBudget",
			"EndDate",
			"ExcludedSites",
			"Funds",
			"Id",
			"Name",
			"NegativeKeywords",
			"Notification",
			"RepresentedBy",
			"SourceId",
			"StartDate",
			"State",
			"Statistics",
			"Status",
			"StatusClarification",
			"StatusPayment",
			"TimeTargeting",
			"TimeZone",
			"Type",
		},
		TextCampaignFieldNames: &[]campaigns.TextCampaignFieldEnum{
			"BiddingStrategy",
			"CounterIds",
			"RelevantKeywords",
			"Settings",
		}
	}
 
	service := campaigns.New(&client)

	result, err := service.Get(request);
	// Do something
}

```
