package conf

import (
	"encoding/json"
	"io/ioutil"
	"test/throw"
)

//Config ..
var Config AppConfig

//Init ..
func Init() {
	b, err := ioutil.ReadFile("conf/config.json")
	throw.CheckErr(err)
	jsonStr := string(b)
	var tmp AppConfig
	err = json.Unmarshal([]byte(jsonStr), &tmp)
	throw.CheckErr(err)
	Config = tmp
}
