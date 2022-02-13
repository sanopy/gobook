package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// JsonMarshal encodes a Go value in Json form.
func JsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := jsonencode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// jsonencode writes to buf an Json representation of v.
func jsonencode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprint(buf, "true")
		} else {
			fmt.Fprint(buf, "false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return jsonencode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // [value, ...]
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := jsonencode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // {name: value, ...}
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "\"%s\": ", v.Type().Field(i).Name)
			if err := jsonencode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case reflect.Map: // {key: value, ...}
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := jsonencode(buf, key); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := jsonencode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	default: // chan, func, interface, complex
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
