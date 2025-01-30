package json

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/encoding/protojson"
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

    if protoMsg, ok := v.(proto.Message); ok {
        marshaler := protojson.MarshalOptions{
            EmitUnpopulated: true,
            UseProtoNames: true,
            UseEnumNumbers: true,
        }

        // Marshal Protobuf to JSON
        data, err = marshaler.Marshal(protoMsg)
        if err != nil {
            return err
        }

        var objMap map[string]interface{}
        decoder := json.NewDecoder(bytes.NewReader(data))
        decoder.UseNumber()

        if err := decoder.Decode(&objMap); err != nil {
            return err
        }


        objMap = convertNumbers(objMap).(map[string]interface{})


        data, err = json.Marshal(objMap)
        if err != nil {
            return err
        }
    } else {
        data, err = json.Marshal(v)
    }

    if err != nil {
        return err
    }

    _, err = w.Write(data)
    return err
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
