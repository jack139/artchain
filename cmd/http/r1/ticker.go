package r1

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
)

// 检查拍卖时间进行状态转换
func CheckAuction() error {

	limit := uint64(1000)

	// 查询拍卖信息, 一次处理1000个
	respData2, err := queryAuctionListByStatusPage(1, uint64(limit), "INIT|OPEN")
	if err != nil {
		return err
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")

	dataList := *respData2

	for _, item0 := range dataList {
		newStatus := ""
		item := item0.(map[string]interface{})

		//log.Printf("%v", item)

		if item["status"] == "INIT" {
			if item["openDate"].(string) < nowTime { // 开始拍卖
				newStatus = "OPEN"
				log.Printf("auction --> OPEN: %v", item["id"])
			}
		} else if item["status"] == "OPEN" {
			if item["closeDate"].(string) < nowTime { // 停止拍卖
				newStatus = "CLOSE"
				log.Printf("auction --> CLOSE: %v", item["id"])
			}
		} else {
			continue
		}

		// 修改 拍卖状态
		if newStatus != "" {

			// 检查 lastDate 字段是否正常
			if _, ok := item["lastDate"]; !ok {
				log.Print("ERROR: lastDate empty") // 不应该发生
				continue
			}
			if !strings.HasPrefix(item["lastDate"].(string), "[") {
				log.Print("ERROR: lastDate broken") // 不应该发生
				continue
			}

			// 反序列化
			var data2 []map[string]interface{}
			if err := json.Unmarshal([]byte(item["lastDate"].(string)), &data2); err != nil {
				log.Println("ERROR: ", err.Error())
				continue
			}
			item["lastDate"] = data2

			auctionId, err := strconv.ParseUint(item["id"].(string), 10, 64)
			if err != nil {
				log.Println("ERROR: ", err.Error())
				continue
			}

			// 修改链上数据
			_, err = auctionModify(&item, item["auctionHouseId"].(string), auctionId,
				"\x00", "\x00", "\x00", "\x00", newStatus, "robot")
			if err != nil {
				log.Println("ERROR: ", err.Error())
				continue
			}
		}

		if newStatus == "CLOSE" {
			log.Printf("New trans --> %v", item["id"])

			// 从拍卖叫价中获取最高价
			bidMap, err := queryBidHighest(item["id"].(string))
			if err != nil {
				log.Println("ERROR: ", err.Error())
				continue
			}

			if bidMap == nil { // 拍卖结束时没有最高价
				log.Println("ERROR: ", "NO hammer price!")
				continue
			}

			// 生成交易订单
			_, err = transNew(item["auctionHouseId"].(string),
				item["id"].(string),
				item["itemId"].(string),
				"BID",
				(*bidMap)["buyerId"].(string),
				item["SellerId"].(string),
				(*bidMap)["bidTime"].(string),
				(*bidMap)["bidPrice"].(string),
				"BID id = "+(*bidMap)["id"].(string), // 记录出价 id
				"robot")
			if err != nil {
				log.Println("ERROR: ", err.Error())
				continue
			}
		}
	}

	return nil
}
