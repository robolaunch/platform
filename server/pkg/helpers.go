package swagger

import (
	"encoding/json"
)

func structToJSONByteArray(obj interface{}) []byte {
	b, err := json.Marshal(obj)
	if err != nil {
		return []byte{}
	}
	return b
}
