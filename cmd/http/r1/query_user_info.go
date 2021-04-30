package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"github.com/cosmos/cosmos-sdk/client"

	"fmt"
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


	// 查询链上数据
	respData2, err := queryUserInfoByChainAddr(ctx, chainAddr)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	userMap := *respData2
	userInfo := userMap["userInfo"].(map[string]interface{})

	// 构建返回结构
	respData := map[string] interface{} {
		"chain_addr": userMap["chainAddr"],
		"login_name": userMap["name"],
		"bank_acc_name": userInfo["bank_acc_name"],
		"bank_name": userInfo["bank_name"],
		"bank_acc_no": userInfo["bank_acc_no"],
		"address": userInfo["contact_address"],
		"phone": userInfo["phone"],
		"email": userInfo["email"],
		"referrer": userInfo["referrer"],
		"reg_date": userMap["regDate"],
		"last_date": userMap["lastDate"],
	}

	resp := map[string] interface{} {
		"user" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 User map
func queryUserInfoByChainAddr(ctx *fasthttp.RequestCtx, chainAddr string) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 检查 用户地址 是否存在
	_, err := helper.FetchKey(clientCtx.Keyring, chainAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid chain_addr")
	}

	// 准备查询
	queryClient := persontypes.NewQueryClient(clientCtx)

	params := &persontypes.QueryGetUserByChainAddrRequest{
		ChainAddr: chainAddr,
	}

	res, err := queryClient.UserByChainAddr(context.Background(), params)
	if err != nil {
		return nil, err
	}

	//log.Printf("%t\n", res)

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

	// 处理data字段
	respData2, err := unmarshalUser(&respData, chainAddr)
	if err!=nil{
		return nil, err
	}

	return respData2, nil
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

	return &item, nil
}

