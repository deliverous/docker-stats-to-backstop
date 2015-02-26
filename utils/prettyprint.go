package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("%#v", v)
	}
	return string(bytes)
}
