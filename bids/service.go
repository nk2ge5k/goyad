package bids

import "encoding/json"
import "errors"
import "github.com/nk2ge5k/goyad"

type Service struct {
	Client *goyad.Client
}

func New(c *goyad.Client) Service {
	return Service{c}
}

func (s Service) Get(p GetRequest) (GetResponse, error) {
	if r, err := s.Client.Do("bids", "get", p); err != nil {
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

func (s Service) Set(p SetRequest) (SetResponse, error) {
	if r, err := s.Client.Do("bids", "set", p); err != nil {
		return SetResponse{}, err
	} else {
		o := map[string]SetResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return SetResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return SetResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) SetAuto(p SetAutoRequest) (SetAutoResponse, error) {
	if r, err := s.Client.Do("bids", "setAuto", p); err != nil {
		return SetAutoResponse{}, err
	} else {
		o := map[string]SetAutoResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return SetAutoResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return SetAutoResponse{}, errors.New("Unable to find result")
		}
	}

}
