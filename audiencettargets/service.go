package audiencettargets

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
	if r, err := s.Client.Do("audiencettargets", "get", p); err != nil {
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

func (s Service) Add(p AddRequest) (AddResponse, error) {
	if r, err := s.Client.Do("audiencettargets", "add", p); err != nil {
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

func (s Service) Delete(p DeleteRequest) (DeleteResponse, error) {
	if r, err := s.Client.Do("audiencettargets", "delete", p); err != nil {
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

func (s Service) Suspend(p SuspendRequest) (SuspendResponse, error) {
	if r, err := s.Client.Do("audiencettargets", "suspend", p); err != nil {
		return SuspendResponse{}, err
	} else {
		o := map[string]SuspendResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return SuspendResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return SuspendResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Resume(p ResumeRequest) (ResumeResponse, error) {
	if r, err := s.Client.Do("audiencettargets", "resume", p); err != nil {
		return ResumeResponse{}, err
	} else {
		o := map[string]ResumeResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return ResumeResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return ResumeResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) SetBids(p SetBidsRequest) (SetBidsResponse, error) {
	if r, err := s.Client.Do("audiencettargets", "setBids", p); err != nil {
		return SetBidsResponse{}, err
	} else {
		o := map[string]SetBidsResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return SetBidsResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return SetBidsResponse{}, errors.New("Unable to find result")
		}
	}

}
