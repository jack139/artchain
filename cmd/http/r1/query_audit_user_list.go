package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询待审核用户清单 */
func QueryAuditUserList(ctx *fasthttp.RequestCtx) {
	log.Println("query_audit_user_list")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need status")
		return
	}

	// 查询链上数据
	respData2, err := queryAuditUserListPage(ctx, status)
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
			"status"     : item["status"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"user_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryAuditUserListPage(ctx *fasthttp.RequestCtx, status string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	params := &persontypes.QueryGetUserByStatusRequest{
		Status: status,
	}

	res, err := queryClient.UserByStatus(context.Background(), params)
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

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	userMapList := respData["User"].([]interface{})

	return &(userMapList), nil
}
