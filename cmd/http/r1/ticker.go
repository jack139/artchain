package r1

import (
	"log"
	"time"
	//"bytes"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
)


// 检查拍卖时间进行状态转换
func CheckAuction() error {

	limit := uint64(1000)

	// 查询拍卖信息, 一次处理1000个
	respData2, err := queryAuctionListByStatusPage(1, uint64(limit), "INIT|OPEN")
	if err!=nil{
		return err
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")

	dataList := *respData2

	for _, item0 := range dataList {
		newStatus := ""
		item := item0.(map[string]interface{})

		//log.Printf("%v", item)

		if item["status"]=="INIT" { 
			if item["openDate"].(string) < nowTime { // 开始拍卖
				newStatus = "OPEN"
				log.Printf("auction --> OPEN: %v", item["id"])
			}
		} else if item["status"]=="OPEN" {
			if item["closeDate"].(string) < nowTime { // 停止拍卖
				newStatus = "CLOSE"
				log.Printf("auction --> CLOSE: %v", item["id"])
			}			
		} else {
			continue
		}

		// 修改 拍卖状态
		if newStatus!="" {

			// 检查 lastDate 字段是否正常
			if _, ok := item["lastDate"]; !ok {
				return fmt.Errorf("lastDate empty") // 不应该发生
			}
			if !strings.HasPrefix(item["lastDate"].(string), "[") {
				return fmt.Errorf("lastDate broken") // 不应该发生
			}

			// 反序列化
			var data2 []map[string]interface{}
			if err := json.Unmarshal([]byte(item["lastDate"].(string)), &data2); err != nil {
				return err
			}
			item["lastDate"] = data2

			auctionId, err := strconv.ParseUint(item["id"].(string), 10, 64)
			if err != nil {
				return err
			}

			// 修改链上数据
			_, err = auctionModify(&item, 
				item["creator"].(string), auctionId, 
				item["auctionHouseId"].(string), 
				item["reservePrice"].(string), 
				newStatus,
			)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
