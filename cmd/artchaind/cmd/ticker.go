package cmd

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctiontypes "github.com/jack139/artchain/x/auction/types"
	release1 "github.com/jack139/artchain/cmd/http/r1"

	"log"
	"time"
	"bytes"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)


// 检查拍卖时间进行状态转换
func checkAuction() error {

	limit := uint64(1000)

	// 查询拍卖信息, 一次处理1000个
	respData2, err := release1.QueryAuctionListByStatusPage(1, uint64(limit), "INIT|OPEN")
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

			// 构建lastDate
			lastDateMap := item["lastDate"].([]map[string]interface{})
			lastDateMap = append(lastDateMap, map[string]interface{}{
				"caller": item["creator"].(string),
				"act":  "auto",
				"date": time.Now().Format("2006-01-02 15:04:05"),
			})
			lastDate, err := json.Marshal(lastDateMap)
			if err != nil {
				return err
			}

			// 设置 caller_addr
			originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
			if err != nil {
				return err
			}
			helper.HttpCmd.Flags().Set(flags.FlagFrom, item["creator"].(string))  // 设置 --from 地址
			defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

			// 获取 ctx 上下文
			clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(item["id"].(string), 10, 64)
			if err != nil {
				return err
			}

			// 数据上链
			msg := auctiontypes.NewMsgUpdateRequest(
				item["creator"].(string), //creator string, 
				auctionId, //id uint64, 
				item["recType"].(string), //recType string, 
				item["itemId"].(string), //itemId string, 
				item["auctionHouseId"].(string), //auctionHouseId string, 
				item["SellerId"].(string), //SellerId string, 
				item["requestDate"].(string), //requestDate string, 
				item["reservePrice"].(string), //reservePrice string, 
				newStatus, //status string, 
				item["openDate"].(string), //openDate string, 
				item["closeDate"].(string), //closeDate string,
				string(lastDate), // lastDate
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// 设置 接收输出
			buf := new(bytes.Buffer)
			clientCtx.Output = buf

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, helper.HttpCmd.Flags(), msg)
			if err != nil {
				return err
			}

			// 结果输出
			respBytes := []byte(buf.String())

			log.Println("output: ", buf.String())

			// 转换成map, 生成返回数据
			var respData map[string]interface{}

			if err := json.Unmarshal(respBytes, &respData); err != nil {
				return err
			}

			// code==0 提交成功
			if respData["code"].(float64)!=0 { 
				return fmt.Errorf("fail: %s", buf.String())
			}

		}

	}

	return nil
}
