package ads

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
	if r, err := s.Client.Do("ads", "add", p); err != nil {
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

func (s Service) Update(p UpdateRequest) (UpdateResponse, error) {
	if r, err := s.Client.Do("ads", "update", p); err != nil {
		return UpdateResponse{}, err
	} else {
		o := map[string]UpdateResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return UpdateResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return UpdateResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Get(p GetRequest) (GetResponse, error) {
	if r, err := s.Client.Do("ads", "get", p); err != nil {
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
	if r, err := s.Client.Do("ads", "delete", p); err != nil {
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

func (s Service) Archive(p ArchiveRequest) (ArchiveResponse, error) {
	if r, err := s.Client.Do("ads", "archive", p); err != nil {
		return ArchiveResponse{}, err
	} else {
		o := map[string]ArchiveResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return ArchiveResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return ArchiveResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Unarchive(p UnarchiveRequest) (UnarchiveResponse, error) {
	if r, err := s.Client.Do("ads", "unarchive", p); err != nil {
		return UnarchiveResponse{}, err
	} else {
		o := map[string]UnarchiveResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return UnarchiveResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return UnarchiveResponse{}, errors.New("Unable to find result")
		}
	}

}

func (s Service) Suspend(p SuspendRequest) (SuspendResponse, error) {
	if r, err := s.Client.Do("ads", "suspend", p); err != nil {
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
	if r, err := s.Client.Do("ads", "resume", p); err != nil {
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

func (s Service) Moderate(p ModerateRequest) (ModerateResponse, error) {
	if r, err := s.Client.Do("ads", "moderate", p); err != nil {
		return ModerateResponse{}, err
	} else {
		o := map[string]ModerateResponse{}
		if err := json.Unmarshal(r, &o); err != nil {
			return ModerateResponse{}, err
		}
		if res, ok := o["result"]; ok {
			return res, nil
		} else {
			return ModerateResponse{}, errors.New("Unable to find result")
		}
	}

}
