package sitelinks

import "encoding/json"
import "errors"
import "github.com/nk2ge5k/goyad"

type Service struct {
	Client *goyad.Client
}

func New(c *goyad.Client) Service {
	return Service{c}
}

func (s Service) Add(p AddRequest) (AddResponse, error) {
	if r, err := s.Client.Do("sitelinks", "add", p); err != nil {
		return AddResponse{}, err
	} else {
		o := map[string]AddResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return AddResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return AddResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Get(p GetRequest) (GetResponse, error) {
	if r, err := s.Client.Do("sitelinks", "get", p); err != nil {
		return GetResponse{}, err
	} else {
		o := map[string]GetResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return GetResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return GetResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Delete(p DeleteRequest) (DeleteResponse, error) {
	if r, err := s.Client.Do("sitelinks", "delete", p); err != nil {
		return DeleteResponse{}, err
	} else {
		o := map[string]DeleteResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return DeleteResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return DeleteResponse{}, errors.New("Unable to find result")
		}
	}

}
