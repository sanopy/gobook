// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func Pack(ptr interface{}) string {
	isFirst := true
	var buf bytes.Buffer
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		switch v.Field(i).Kind() {
		case reflect.String, reflect.Array, reflect.Slice:
			if v.Field(i).Len() == 0 {
				continue
			}
		}
		if isFirst {
			buf.WriteString("?")
		} else {
			buf.WriteString("&")
		}
		isFirst = false
		buf.WriteString(toString(name, v.Field(i)))
	}

	return buf.String()
}

func toString(name string, v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return fmt.Sprintf("%s=true", name)
		} else {
			return fmt.Sprintf("%s=false", name)
		}

	case reflect.String:
		return fmt.Sprintf("%s=%s", name, v.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%s=%d", name, v.Int())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%s=%g", name, v.Float())

	case reflect.Array, reflect.Slice:
		var buf bytes.Buffer
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteString("&")
			}
			buf.WriteString(toString(name, v.Index(i)))
		}
		return buf.String()

	default:
		panic(fmt.Sprintf("unsupported type: %s", v.Type()))
	}
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
