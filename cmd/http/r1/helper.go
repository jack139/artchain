package r1

import (
	"github.com/jack139/artchain/cmd/ipfs"

	"encoding/json"
	"strings"
)

/*
	{
 		"Contract":[
 			{
 				"creator":"artchain1ycd5htx4emg686w05rhpfpqj5hv35g7sy2qmae",
 				"id":"0",
 				"artchainNo":"123",
 				"partyA":"artchain1ghfcl0hm5pxu0q0jgnl2nw3hhmrkklgyh3lgvx",
 				"partyB":"artchain1ghfcl0hm5pxu0q0jgnl2nw3hhmrkklgyh3lgvx",
 				"action":"\u000b",
 				"data":"{\"image\":\"QmUaLpY8SX68Cop5UDZbYGpaDu3JCK7euFsnDzw5xpN1sz\"}"
 			}
 		]
 	}

 							#
 							#
 						    V

 	[
		{
			"id":"0",
			"assets_id":"123",
			"exchange_id":"artchain1ghfcl0hm5pxu0q0jgnl2nw3hhmrkklgyh3lgvx",
			"action":11,
			"image":"zzzzzzz",
			"type":"DEAL",
			"refer":"",
		}
 	]
*/

/* 处理一个交易的数据，为了兼容旧的字段名 */
func processData(item0 *map[string]interface{}, user string) (*map[string]interface{}, error) {
	item := *item0

	var data map[string]interface{}
	// 检查query用户是否是相关者
	if (item["partyA"] != user) && (item["partyB"] != user) {
		// 不相关，不返回data数据
		data = make(map[string]interface{})
		data["image"] = ""
	} else {
		// 相关，解析 data内容
		if err := json.Unmarshal([]byte(item["data"].(string)), &data); err != nil {
			return nil, err
		}

		// 处理image 字段，从ipfs读取
		_, ok := data["image"]
		if ok && len(data["image"].(string)) > 0 {
			image_data, err := ipfs.Get(data["image"].(string))
			if err != nil {
				return nil, err
			}
			data["image"] = string(image_data)
		}
	}

	// 建立返回的数据
	new_item := map[string]interface{}{
		"id":          item["id"],
		"exchange_id": user,
		"userkey_a":   item["partyA"],
		"userkey_b":   item["partyB"],
		"assets_id":   item["artchainNo"],
		"action":      item["action"],
		"type":        "DEAL",
		"refer":       "",
		"data":        data,
	}

	return &new_item, nil
}

/* data字段是已序列化的json串，反序列化一下， 只对一个数据 */
func unmarshalData(reqData *map[string]interface{}, user string) (*map[string]interface{}, error) {
	var respData map[string]interface{}

	data := (*reqData)["User"].(map[string]interface{})

	// 检查 data 字段是否正常
	_, ok := data["data"]
	if !ok {
		return &respData, nil
	}
	if !strings.HasPrefix(data["data"].(string), "{") {
		return &respData, nil
	}

	return processData(&data, user)
}

