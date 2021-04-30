package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"strings"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询用户信息 */
func QueryUserInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_user_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	chainAddr, ok := (*reqData)["chain_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need chain_addr")
		return
	}

	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 检查 用户地址 是否存在
	_, err = helper.FetchKey(clientCtx.Keyring, chainAddr)
	if err != nil {
		helper.RespError(ctx, 9001, "invalid chain_addr")
		return
	}

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	params := &persontypes.QueryGetUserByChainAddrRequest{
		ChainAddr: chainAddr,
	}

	res, err := queryClient.UserByChainAddr(context.Background(), params)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	log.Printf("%t\n", res)

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
		helper.RespError(ctx, 9004, err.Error())
		return
	}

	// 处理data字段
	respData2, err := unmarshalUser(&respData, chainAddr)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}

	resp := map[string] interface{} {
		"user" : *respData2,
	}

	helper.RespJson(ctx, &resp)
}


/* userInfo字段是已序列化的json串，反序列化一下，针对数据列表 */
func unmarshalUser(reqData *map[string]interface{}, user string) (*map[string]interface{}, error) {
	var respData map[string]interface{}

	item := (*reqData)["User"].(map[string]interface{})

	// 检查 userInfo 字段是否正常
	_, ok := item["userInfo"]
	if !ok {
		return &respData, nil
	}
	if !strings.HasPrefix(item["userInfo"].(string), "{") {
		return &respData, nil
	}

	// 反序列化
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(item["userInfo"].(string)), &data); err != nil {
		return nil, err
	}
	item["userInfo"] = data

	// 返回的数据结构
	respData = map[string] interface{} {
		"chain_addr": item["chainAddr"],
		"login_name": item["name"],
		"bank_acc_name": (item["userInfo"].(map[string]interface{}))["bank_acc_name"],
		"bank_name": (item["userInfo"].(map[string]interface{}))["bank_name"],
		"bank_acc_no": (item["userInfo"].(map[string]interface{}))["bank_acc_no"],
		"address": (item["userInfo"].(map[string]interface{}))["contact_address"],
		"phone": (item["userInfo"].(map[string]interface{}))["phone"],
		"email": (item["userInfo"].(map[string]interface{}))["email"],
		"referrer": (item["userInfo"].(map[string]interface{}))["referrer"],
		"reg_date": item["regDate"],
	}

	return &respData, nil
}

