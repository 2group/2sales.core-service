package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"
	"unicode"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ParseJSON(r *http.Request, model any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}
	return json.NewDecoder(r.Body).Decode(model)
}

type jsonNumberMarshaler struct {
	json.Number
}

func (j jsonNumberMarshaler) MarshalJSON() ([]byte, error) {
	if j.Number == "" {
		return []byte("0"), nil
	}
	i, err := j.Int64()
	if err != nil {
		return []byte(fmt.Sprintf("%v", j.Number)), nil
	}
	return []byte(fmt.Sprintf("%d", i)), nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var data []byte
	var err error

	// Special handling for protobuf messages.
	if protoMsg, ok := v.(proto.Message); ok {
		marshaler := protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true, // forces snake_case keys (as defined in the proto)
			UseEnumNumbers:  true,
		}

		// Marshal the protobuf message to JSON.
		data, err = marshaler.Marshal(protoMsg)
		if err != nil {
			return err
		}

		// Decode the JSON into a map.
		var objMap map[string]interface{}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.UseNumber()
		if err := decoder.Decode(&objMap); err != nil {
			return err
		}

		// Convert any number representations.
		objMap = convertNumbers(objMap).(map[string]interface{})

		// Recursively fill missing fields with nil using proto reflection.
		fillMissingFields(protoMsg.ProtoReflect(), objMap)

		// Normalize all keys to snake_case.
		objMap = normalizeKeysMap(objMap)

		// Marshal the modified map back to JSON.
		data, err = json.Marshal(objMap)
		if err != nil {
			return err
		}
	} else {
		data, err = json.Marshal(v)
		if err != nil {
			return err
		}
	}

	_, err = w.Write(data)
	return err
}

// fillMissingFields recursively walks through the proto message (m)
// and ensures every field defined in the descriptor exists in objMap.
// If a field is missing, it adds it with a nil value.
// For message fields, it recurses into the nested map.
func fillMissingFields(m protoreflect.Message, objMap map[string]interface{}) {
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fieldDesc := fields.Get(i)
		jsonName := fieldDesc.JSONName()

		// Debug: show field name and whether it is present.

		if _, exists := objMap[jsonName]; !exists {
			objMap[jsonName] = nil
		} else {
			// If the field is a message, process recursively.
			if fieldDesc.Kind() == protoreflect.MessageKind {
				if m.Has(fieldDesc) {
					subMsg := m.Get(fieldDesc).Message()
					if subMap, ok := objMap[jsonName].(map[string]interface{}); ok {
						fillMissingFields(subMsg, subMap)
					}
				}
			}
		}
	}
}

// normalizeKeysMap creates a new map where every key is converted to snake_case.
// It recurses into nested maps and slices.
func normalizeKeysMap(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range m {
		newKey := toSnakeCase(k)
		switch val := v.(type) {
		case map[string]interface{}:
			newMap[newKey] = normalizeKeysMap(val)
		case []interface{}:
			newSlice := make([]interface{}, len(val))
			for i, item := range val {
				if subMap, ok := item.(map[string]interface{}); ok {
					newSlice[i] = normalizeKeysMap(subMap)
				} else {
					newSlice[i] = item
				}
			}
			newMap[newKey] = newSlice
		default:
			newMap[newKey] = v
		}
	}
	return newMap
}

// toSnakeCase converts a given string from camelCase (or mixed) to snake_case.
func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func convertNumbers(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		for k, v := range x {
			x[k] = convertNumbers(v)
		}
		return x
	case []interface{}:
		for i, v := range x {
			x[i] = convertNumbers(v)
		}
		return x
	case string:
		if num, err := strconv.ParseInt(x, 10, 64); err == nil {
			return json.Number(strconv.FormatInt(num, 10))
		}
		return x
	default:
		return v
	}
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
