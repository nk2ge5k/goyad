package changes

import "encoding/json"
import "errors"
import "github.com/nk2ge5k/goyad"

type Service struct {
	Client *goyad.Client
}

func New(c *goyad.Client) Service {
	return Service{c}
}

func (s Service) CheckDictionaries(p CheckDictionariesRequest) (CheckDictionariesResponse, error) {
	if r, err := s.Client.Do("changes", "checkDictionaries", p); err != nil {
		return CheckDictionariesResponse{}, err
	} else {
		o := map[string]CheckDictionariesResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return CheckDictionariesResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return CheckDictionariesResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) CheckCampaigns(p CheckCampaignsRequest) (CheckCampaignsResponse, error) {
	if r, err := s.Client.Do("changes", "checkCampaigns", p); err != nil {
		return CheckCampaignsResponse{}, err
	} else {
		o := map[string]CheckCampaignsResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return CheckCampaignsResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return CheckCampaignsResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Check(p CheckRequest) (CheckResponse, error) {
	if r, err := s.Client.Do("changes", "check", p); err != nil {
		return CheckResponse{}, err
	} else {
		o := map[string]CheckResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return CheckResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return CheckResponse{}, errors.New("Unable to find result")
		}
	}

}
