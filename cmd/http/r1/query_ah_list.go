package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询拍卖行清单 */
func QueryAHList(ctx *fasthttp.RequestCtx) {
	log.Println("query_user_list")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	_, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	/*
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need page")
		return
	}
	limit, ok := (*reqData)["limit"].(float64)
	if !ok {
		helper.RespError(ctx, 9002, "need limit")
		return
	}
	*/

	// 查询链上数据
	respData2, err := queryAHListPage(ctx, uint64(1), uint64(1000))
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		newItem := map[string]interface{} {
			"id"         : item["id"],
			"chain_addr" : item["chainAddr"],
			"login_name" : item["name"],
			"user_type"  : item["userType"],
			"reg_date"   : item["regDate"],
			//"last_date"  : item["lastDate"],
			"status"     : item["status"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"ah_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryAHListPage(ctx *fasthttp.RequestCtx, page uint64, limit uint64) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	pageReq := query.PageRequest{
		Key:        []byte(""),
		Offset:     (page - 1) * limit,
		Limit:      limit,
		CountTotal: true,
	}

	params := &persontypes.QueryGetUserByUserTypeRequest{
		Pagination: &pageReq,
		UserType: "AH",
	}

	res, err := queryClient.UserByUserType(context.Background(), params)
	if err != nil {
		return nil, err
	}

	//log.Printf("%T\n", res)

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	// 转换输出
	clientCtx.PrintProto(res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	userMapList := respData["User"].([]interface{})

	return &(userMapList), nil
}
