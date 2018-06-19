package serialization

import "encoding/json"

func SerializeJSON(in interface{}) error {
	_, err := json.Marshal(in)
	return err
}

func DeserializeJSON(in []byte) error {
	var out interface{}
	return json.Unmarshal(in, out)
}