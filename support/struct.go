package support

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	TagName string = "json"
)

type Struct struct {
	// Struct source
	source interface{}
	// Struct reflection
	value reflect.Value
}

type MapConfig struct {
	Name       string
	isString   bool
	isNull     bool
	isRequired bool
	omitEmpty  bool
}

type TagOptions []string

func (o TagOptions) Has(k string) bool {
	for _, v := range o {
		if v == k {
			return true
		}
	}
	return false
}

// Constructor for Struct
func New(s interface{}) *Struct {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic(
			fmt.Errorf(
				"Argument must be kind a struct, %s givven",
				v.Kind().String(),
			),
		)
	}

	return &Struct{
		source: s,
		value:  v,
	}
}

func (s *Struct) Map(nullFields, forceFields JsonFields) map[string]interface{} {

	m := make(map[string]interface{})
	s.Fill(m, nullFields, forceFields)

	return m
}

func (s *Struct) Fill(out map[string]interface{}, nullFields, forceFields JsonFields) {
	fields := s.Fields()

	for _, f := range fields {
		fn := f.Name
		sv := s.value.FieldByName(fn)

		if sv.Kind() == reflect.Ptr {
			sv = sv.Elem()
		}

		if f.Anonymous {
			sub := New(sv.Interface())
			sub.Fill(out, nullFields, forceFields)
		} else {
			var isEmpty bool

			conf := MapConfig{
				Name:       fn,
				isString:   false,
				isNull:     false,
				isRequired: false,
				omitEmpty:  false,
			}

			isEmpty = IsEmpty(sv) || reflect.DeepEqual(
				sv.Interface(),
				reflect.Zero(sv.Type()).Interface(),
			)

			tag := f.Tag.Get(TagName)

			if tag != "" {
				tfName, tOpts := parseTag(tag)

				if tfName != "" {
					conf.Name = tfName
				}

				conf.isNull = nullFields.Has(conf.Name) || nullFields.Has(fn)
				conf.isRequired = forceFields.Has(conf.Name) || forceFields.Has(fn)

				if !isEmpty && conf.isNull {
					panic(
						fmt.Errorf(
							"Field %s not empty and null",
							fn,
						),
					)
				}

				if tOpts.Has("omitempty") {
					if isEmpty && !conf.isNull && !conf.isRequired {
						continue
					}
				}

				if tOpts.Has("string") {
					if conf.isNull {
						out[conf.Name] = nil
					} else {
						out[conf.Name] = fmt.Sprintf("%v", sv.Interface())
					}
					continue
				}
			}

			if isEmpty && conf.isNull {
				out[conf.Name] = nil
				continue
			}

			out[conf.Name] = mapValue(sv)
		}
	}

}

func (s *Struct) Fields() []reflect.StructField {
	var f []reflect.StructField

	t := s.value.Type()
	l := t.NumField()

	for i := 0; i < l; i++ {
		c := t.Field(i)

		if c.PkgPath != "" {
			continue
		}

		if c.Tag.Get(TagName) == "-" {
			continue
		}

		f = append(f, c)
	}

	return f
}

func mapValue(v reflect.Value) interface{} {

	switch v.Kind() {
	case reflect.Struct:
		m := make(map[string]interface{})
		sub := New(v.Interface())
		sub.Fill(m, JsonFields{}, JsonFields{})

		return m
	case reflect.Array, reflect.Slice:
		a := make([]interface{}, v.Len())

		for u := 0; u < v.Len(); u++ {
			a[u] = mapValue(v.Index(u))
		}

		return a
	default:
		return v.Interface()
	}

}

func parseTag(t string) (string, TagOptions) {
	r := strings.Split(t, ",")
	return r[0], r[1:]
}

func IsEmpty(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false

}
