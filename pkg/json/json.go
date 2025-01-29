package json

import (
    "encoding/json"
    "fmt"
    "net/http"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
)

// ParseJSON remains the same since it handles incoming JSON
func ParseJSON(r *http.Request, model any) error {
    if r.Body == nil {
        return fmt.Errorf("Missing request body")
    }
    return json.NewDecoder(r.Body).Decode(model)
}

// WriteJSON now handles both proto messages and regular structs
func WriteJSON(w http.ResponseWriter, status int, v any) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)

    var data []byte
    var err error

    // Check if the value is a proto message
    if protoMsg, ok := v.(proto.Message); ok {
        // Use protojson for proto messages
        marshaler := protojson.MarshalOptions{
            EmitUnpopulated: true,    // Include null values
            UseProtoNames:   true,    // Use original field names
        }
        data, err = marshaler.Marshal(protoMsg)
    } else {
        // Use standard json for non-proto types
        data, err = json.Marshal(v)
    }

    if err != nil {
        return err
    }

    _, err = w.Write(data)
    return err
}

// WriteError remains the same
func WriteError(w http.ResponseWriter, status int, err error) {
    WriteJSON(w, status, map[string]string{"error": err.Error()})
}
