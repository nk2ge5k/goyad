package support

type JsonFields []string

func (j JsonFields) Has(v string) bool {
	for _, e := range j {
		if e == v {
			return true
		}
	}
	return false
}

type MappedObjectInterface interface {
	Get() (JsonFields, JsonFields)
}

type ApiObject struct {
	ForceFields JsonFields `json:"-"`
	NullFields  JsonFields `json:"-"`
}

func (o ApiObject) Get() (JsonFields, JsonFields) {
	return o.NullFields, o.ForceFields
}
