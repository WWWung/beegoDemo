package utils

import "fmt"

// //FromJSON ..
// func FromJSON(jsonStr string, obj interface{}) {
// 	err := json.Unmarshal([]byte(jsonStr), obj)
// 	throw.CheckErr(err)
// }

//InterfaceToStr ..
func InterfaceToStr(obj interface{}) string {
	msg := fmt.Sprint(obj)
	return msg
}
