# goyad
Yande Direct API Go library (v5)

# Basic usage 

```go
import (
	"github.com/nk2ge5k/goyad"
	"github.com/nk2ge5k/goyad/campaigns"
)

func main() {

	tf := &[]campaigns.TextCampaignFieldEnum{
		"BiddingStrategy",
		"CounterIds",
		"RelevantKeywords",
		"Settings",
	}

	cs := campaigns.GetRequest{
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
		TextCampaignFieldNames: tf,
	}
  
	c := goyad.NewClient(CLIENT_LOGIN, API_TOKEN)
	s := campaigns.New(&c)

	if resul, err := s.Get(cs);
}

```
