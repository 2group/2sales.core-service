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

// ParseJSON decodes the request body into the provided model.
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

// WriteJSON writes the data v as JSON with the specified HTTP status.
// For protobuf messages, it marshals using protojson with EmitUnpopulated,
// then post-processes the JSON to (1) convert numbers (while skipping "bin"),
// (2) fill missing fields using proto reflection without overwriting valid values,
// and (3) normalize keys to snake_case.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var data []byte
	var err error

	if protoMsg, ok := v.(proto.Message); ok {
		marshaler := protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true, // forces keys defined in the proto (ideally snake_case)
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

		// Convert any number representations (skipping the key "bin").
		objMap = convertNumbers(objMap).(map[string]interface{})

		// Recursively fill missing fields using proto reflection,
		// but do not override valid values.
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

// fillMissingFields recursively processes the proto message m and ensures that
// for each field the normalized key (snake_case) exists in objMap. It will not
// overwrite a value that is already non-nil. If a duplicate camelCase key exists,
// its non-nil value is used if the normalized key is missing or nil.
func fillMissingFields(m protoreflect.Message, objMap map[string]interface{}) {
	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fieldDesc := fields.Get(i)
		normalizedKey := toSnakeCase(fieldDesc.JSONName())
		camelKey := fieldDesc.JSONName()

		// 1) Ensure the key exists at all, defaulting to nil
		if existing, exists := objMap[normalizedKey]; !exists || existing == nil {
			if val, existsCamel := objMap[camelKey]; existsCamel && val != nil {
				objMap[normalizedKey] = val
			} else if !exists {
				objMap[normalizedKey] = nil
			}
		}
		// drop the duplicate camelCase key
		if camelKey != normalizedKey {
			delete(objMap, camelKey)
		}

		// 2) Handle nested messages (singular and repeated)
		if fieldDesc.Kind() == protoreflect.MessageKind {
			if fieldDesc.IsList() {
				// For each element in the JSON array, recurse
				listIface, ok := objMap[normalizedKey].([]interface{})
				if !ok {
					continue
				}
				pfList := m.Get(fieldDesc).List()
				for idx := 0; idx < pfList.Len() && idx < len(listIface); idx++ {
					subMap, ok := listIface[idx].(map[string]interface{})
					if !ok {
						continue
					}
					// pfList.Get(idx).Message() is the proto sub-message
					fillMissingFields(pfList.Get(idx).Message(), subMap)
				}
			} else {
				// Singular nested message: recurse into its map
				subMap, ok := objMap[normalizedKey].(map[string]interface{})
				if ok {
					fillMissingFields(m.Get(fieldDesc).Message(), subMap)
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

// convertNumbers recursively converts string values to json.Number when appropriate,
// except for keys that should remain as strings (for example, "bin").
func convertNumbers(v interface{}) interface{} {
	switch x := v.(type) {
	case map[string]interface{}:
		for k, val := range x {
			// Skip conversion for key "bin".
			if k == "bin" {
				continue
			}
			x[k] = convertNumbers(val)
		}
		return x
	case []interface{}:
		for i, val := range x {
			x[i] = convertNumbers(val)
		}
		return x
	case string:
		// Attempt conversion only if the string represents an integer.
		if num, err := strconv.ParseInt(x, 10, 64); err == nil {
			return json.Number(strconv.FormatInt(num, 10))
		}
		return x
	default:
		return x
	}
}

// WriteError sends an error message as JSON.
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
