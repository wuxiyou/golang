package packsCom

import (
	"encoding/json"
	"../tools"
)
// todo
func login(data string) string {
	var params map[string]interface{}
	json.Unmarshal([]byte(data), &params)

	mustParams := []string{"usrName","passWord","time", "sign"}
	for _, val := range mustParams {
		res := tools.Is_exits_key(params, val)
		if res != true {
			break
		}
	}


}
