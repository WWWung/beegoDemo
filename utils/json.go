package utils

import (
	"encoding/json"
	"test/throw"
)

//FromJSON ..
func FromJSON(jsonStr string, obj interface{}) {
	err := json.Unmarshal([]byte(jsonStr), obj)
	throw.CheckErr(err)
}
