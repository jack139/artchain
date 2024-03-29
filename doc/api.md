##  应用层API



| 修改日期   | 修改内容                                                     |
| ---------- | ------------------------------------------------------------ |
| 2021-04-28 | 初版                                 |



###  一、 说明

​		应用层API与区块链节点一起部署，提供给客户端调用，进行基础的区块链功能操作。



### 二、 概念和定义

#### 1. 节点

​		节点是区块链上的一个业务处理和存储的单元，是一个具有独立处理区块链业务的服务程序。节点可以是一台物理服务器，也可以是多个节点共用一个物理服务器，通过不同端口提供各自节点的功能。

#### 2. 链用户

​		链用户是具有提交区块链交易权限的用户，线下可定义为交易所。每个链用户通过一对密钥识别，同时使用此密钥进行数据的加密解密操作，因此链用户的密钥需要妥善保管。

### 三、 API提供的功能

#### 业务处理接口


| 序号               | 接口功能                 | URI                        |
| :------------------: | -------------------------- | -------------------------- |
| 1            | 注册用户             | **/biz/user/register**         |
| 2         | 修改用户信息          | **/biz/user/modify**           |
| 3       | 恢复已注册用户        | /biz/user/restore              |
| 4            | 新建物品             | **/biz/item/new**         |
| 5         | 修改物品信息          | **/biz/item/modify**           |
| 6 | 上传图片到IPFS       | **/ipfs/upload/image**        |
| 7 | 删除图片             | **/ipfs/remove/image**        |
| 8 | 修改物品所有人 | **/biz/item/change_owner** |
| 9         | 生成物品NFT          | /biz/nft/new      |
| 10        | 添加物品评价          | **/biz/review/new**            |
| 11        | 修改物品评价          | **/biz/review/modify**         |
| 12 | 修改物品评价反馈信息   | /biz/review/feedback           |
| 13           | 发起拍卖             | **/biz/auction/new**           |
| 14       | 修改拍卖信息       | **/biz/auction/modify** |
| 15               | 出价                 | **/biz/auction/bid**           |
| 16       | 建立成交交易         | **/biz/trans/new**       |
| 17 | 审核用户 | **/biz/audit/user** |
| 18 | 审核物品 | **/biz/audit/item** |
| 19 | ~~审核照片~~ | ~~/biz/audit/image~~ |
| 20 | 审核评价 | **/biz/audit/review** |
| 21 | 审核拍卖请求 | **/biz/audit/auction** |
| 22 | 审核成交交易 | **/biz/audit/trans** |
| 23 | 撤销出价 | **/biz/auction/bid/withdraw** |



#### 查询接口

| 序号               | 接口功能                 | URI                        |
| :------------------: | -------------------------- | -------------------------- |
| 1 | 查询用户清单 | **/query/user/list** |
| 2 | 查询用户信息         | **/query/user/info**       |
| 3 | 验证用户身份          | **/query/user/verify**           |
| 4 | 查询物品清单         | **/query/item/list**       |
| 5 | 查询物品信息         | **/query/item/info**       |
| 6 | 验证物品NFT          | /query/nft/verify                |
| 7 | 查询物品评价清单     | **/query/review/list**     |
| 8 | 查询物品评价信息     | **/query/review/info**     |
| 9 | 查询拍卖行清单       | **/query/auction_house/list** |
| 10 | ~~查询拍卖行信息~~   | ~~/query/auction_house/info~~ |
| 11 | 查询拍卖清单         | **/query/auction/list**    |
| 12 | 查询拍卖信息         | **/query/auction/info**    |
| 13 | 查询出价清单      | **/query/bid/list**    |
| 14 | 查询最高出价         | **/query/bid/highest**     |
| 15 | 查询出价信息       | **/query/bid/info**    |
| 16 | 查询成交交易清单       | **/query/trans/list** |
| 17 | 查询成交交易信息     | **/query/trans/info** |
| 18 | 查询指定区块原始数据 | **/query/block/rawdata** |
| 19 | 查询用户通证 | **/query/user/credit_balance** |
| 20 | 指定状态的用户清单 | **/query/user/list_by_status** |
| 21 | 指定状态的物品清单 | **/query/item/list_by_status** |
| 22 | 指定状态的评价清单 | **/query/review/list_by_status** |
| 23 | 指定状态的拍卖申请清单 | **/query/auction/list_by_status** |
| 24 | 指定条件查询成交交易 | **/query/trans/list_by_condition** |
| 25 | 从IPFS下载数据 | **/ipfs/download** |



### 四、接口定义

#### 1. 全局接口定义

输入参数

| 参数      | 类型   | 说明                          | 示例        |
| --------- | ------ | ----------------------------- | ----------- |
| appid | string | 应用渠道编号                  |             |
| version   | string | 版本号                        | 1 |
| sign_type | string | 签名算法，目前使用SHA256算法 | SHA256 |
| sign_data | string | 签名数据，具体算法见下文      |             |
| timestamp | int    | unix时间戳（秒）              |             |
| data      | json   | 接口数据，详见各接口定义      |             |

> 签名/验签算法：
>
> 1. appid和app_secret均从线下获得。
> 2. 筛选，获取参数键值对，剔除sign_data参数。data参数按key升序排列进行json序列化。
> 3. 排序，按key升序排序；data中json也按key升序排序。
> 4. 拼接，按排序好的顺序拼接请求参数。
>
> ```key1=value1&key2=value2&...&key=appSecret```，key=app_secret固定拼接在参数串末尾。
>
> 4. 签名，使用制定的算法进行加签获取二进制字节，使用 16进制进行编码Hex.encode得到签名串，然后base64编码。
> 5. 验签，对收到的参数按1-4步骤签名，比对得到的签名串与提交的签名串是否一致。

签名示例：

```json
请求参数：
{
    "appid": "66A095861BAE55F8735199DBC45D3E8E", 
    "version": "1", 
    "data": {
        "test1": "test1", 
        "atest2": "test2", 
        "Atest2": "test2"
    }, 
    "timestamp": 1608904438, 
    "sign_type": "SHA256",  
    "sign_data": "..."
}

密钥：
app_secret="43E554621FF7BF4756F8C1ADF17F209C"

待加签串：
appid=66A095861BAE55F8735199DBC45D3E8E&data={"Atest2":"test2","atest2":"test2","test1":"test1"}&sign_type=SHA256&timestamp=1608948188&version=1&key=43E554621FF7BF4756F8C1ADF17F209C

SHA256加签结果：
"fa72d34eafea3639b0a207bdd7ceb49586f4be92e58ee97b6453b696b0edb781"

base64后结果：
"ZmE3MmQzNGVhZmVhMzYzOWIwYTIwN2JkZDdjZWI0OTU4NmY0YmU5MmU1OGVlOTdiNjQ1M2I2OTZiMGVkYjc4MQ=="
```

返回结果

| 参数      | 类型    | 说明                                                         |
| --------- | ------- | ------------------------------------------------------------ |
| code      | int   | 状态代码，0 表示成功，非0 表示出错                                 |
| msg   | string | 成功时返回success；出错时，返回出错信息                                                     |
| data      | json    | 成功时返回结果数据，详见具体接口                |

返回示例

```json
{
    "code": 0, 
    "msg": "success", 
    "data": {
    }
}
```

全局出错代码

| 编码 | 说明                               |
| ---- | ---------------------------------- |
| 9000 | 签名错误                           |



#### 2. 业务处理接口

##### 2.1 注册用户

请求URL

> http://<host>:<port>/api/<version>/biz/user/register

请求方式

> POST

输入参数（data字段下）

| 参数          | 类型   | 必填 | 说明               |
| ------------- | ------ | ---- | ------------------ |
| caller_addr   | string | Y    | 调用者的链地址     |
| login_name    | string | Y    | 登录名             |
| user_type     | string | Y    | 注册用户类型       |
| bank_acc_name | string |      | 银行开户姓名       |
| bank_name     | string |      | 银行名称           |
| bank_acc_no   | string |      | 银行账号           |
| address       | string |      | 联系地址           |
| phone         | string |      | 联系电话           |
| email         | string |      | 电子邮件           |
| referrer      | string |      | 推荐人的chain_addr |

> user_type 取值：
>
> TRD 交易人
>
> AH 拍卖行（需审核）
>
> DEL 经销商（需审核）
>
> ART 艺术家（需审核）
>
> REV 评论家（需审核）



返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户链地址、密码字符串                  |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr":"bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s",
        "login_name": "test1", 
        "user_type": "TRD", 
        "email": "111111@qq.com", 
        "referrer": "bid1art111111111"
    }, 
    "timestamp": 1619679437, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MmRkNDBjYmJiN2E0ZDI0OTRjNzE2NzMxYWUxODE3ZTlhMDM1YWMxZDhiZTgzM2E4NGU4YzJlMWI1YWJmMzA1MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "chain_addr":"bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s",
        "height":"1259",
        "mystery":"quantum erupt grit prefer decline rather horror patch cricket muffin method wet will mistake glance scene yard damage tag bundle owner method case paddle"
    },
    "msg":"success"
}
```



##### 2.2 修改用户信息

请求URL

> http://<host>:<port>/api/<version>/biz/user/modify

请求方式

> POST

输入参数（data字段下）

| 参数          | 类型   | 必填 | 说明           |
| ------------- | ------ | ---- | -------------- |
| caller_addr   | string | Y    | 调用者的链地址 |
| chain_addr    | string | Y    | 用户链地址     |
| bank_acc_name | string |      | 银行开户姓名   |
| bank_name     | string |      | 银行名称       |
| bank_acc_no   | string |      | 银行账号       |
| address       | string |      | 联系地址       |
| phone         | string |      | 联系电话       |
| email         | string |      | 电子邮件       |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块的高度                          |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "chain_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "login_name": "test3", 
        "bank_acc_name": "1testbank", 
    }, 
    "timestamp": 1619767850, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "M2MyMDU1ZmIzZTE5YjY1ZDc3YjBhYWVkZjczNTVjYTIxYWRiZDdmN2VlNjQxODUyYmJlNTQxNzIyMWY4NzlmNA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"1010"
    },
    "msg":"success"
}
```



##### 2.4 新建物品

请求URL

> http://<host>:<port>/api/<version>/biz/item/new

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                                   |
| ----------- | ------ | ---- | -------------------------------------- |
| caller_addr | string | Y    | 调用者的链地址                         |
| owner_addr  | string | Y    | 所有者的链地址                         |
| desc        | string | Y    | 物品名称                               |
| detail      | string |      | 物品描述                               |
| date        | string | Y    | 出现年代                               |
| type        | string |      | 类型：原作、复制品                     |
| subject     | string |      | 主题：古代、现代、风景、雕塑、人像，等 |
| media       | string |      | 材质：石头、金属、瓷器、油画、素描、等 |
| size        | string |      | 尺寸描述                               |
| base_price  | string | Y    | 最近交易价格                           |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "desc": "测试物品", 
        "date": "1911s", 
        "base_price": 
        "$2001", 
        "owner_addr": "bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7"
    }, 
    "timestamp": 1619791906, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZDliZTI3YmZiYjhhZjI4MTFiNjhlNDFlODRlODIzNDdlM2Q3MGZhZDYwNmM1NmYyYjM4NTczNjZhMmFhZWJiYQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"2692"
    },
    "msg":"success"
}
```



##### 2.5 修改物品信息

请求URL

> http://<host>:<port>/api/<version>/biz/item/modify

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明             |
| ----------- | ------ | ---- | ---------------- |
| caller_addr | string | Y    | 调用者的链地址   |
| id          | string | Y    | 物品id           |
| desc        | string |      | 新的物品名称     |
| detail      | string |      | 新的物品描述     |
| date        | string |      | 新的出现年代     |
| type        | string |      | 新的类型         |
| subject     | string |      | 新的主题         |
| media       | string |      | 新的材质         |
| size        | string |      | 新的尺寸描述     |
| base_price  | string |      | 新的最近交易价格 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "0", 
        "base_price": "$1001", 
        "caller_addr": "bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7"
    }, 
    "timestamp": 1619791789, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NTA0NzQ3NGQzODkxYjFmMWUyMGIzZjVkMDM5MGM4YzhkYTJkNDAwNmYwMjQ2YTk1ZjI0Y2IwOTUwMmYwNzQyZg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"2576"
    },"msg":"success"
}
```



##### 2.6 上传图片到IPFS

请求URL

> http://<host>:<port>/api/ipfs/upload/image

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                 |
| ----------- | ------ | ---- | -------------------- |
| caller_addr | string | Y    | 调用者的链地址       |
| item_id     | string | Y    | 物品id               |
| image       | string | Y    | base64编码的图片数据 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | IPFS文件hash值                          |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art17qppfv5k29r9txqu8sj3l6vfwtt90rr82r9gt7", 
        "item_id": "1", 
        "image": "iVBSUhEUgAABKIAAAXqCAIAAABMelVoAA...AAAAABJRU5ErkJggg=="
    }, 
    "timestamp": 1622624546, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZmZjY2U5Mjg3OTE3NDc1ODZkYWM2ZjM3OTU1ZmFjODdhNTc2NjExYjNiMTQyNmVjOTI0OTQ0MDcyZDAzZmQzYQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "hash":"QmbkAewP7KN9VMKwYbx6xj45Yc82dPraRnARFjww4MLBt3",
        "height":"49975"
    },
    "msg":"success"
}
```



##### 2.7 从IPFS删除图片

请求URL

> http://<host>:<port>/api/ipfs/remove/image

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| item_id     | string | Y    | 物品id         |
| hash        | string | Y    | 图片hash       |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | IPFS文件hash值                          |

请求示例

```json

```

返回示例

```json

```



##### 2.8 修改物品所有人

请求URL

> http://<host>:<port>/api/<version>/biz/item/change_owner

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 物品id         |
| owner_addr  | string | Y    | 新的物品所有人 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "0", 
        "owner_addr": "bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7", 
        "caller_addr": "bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7"
    }, 
    "timestamp": 1619791789, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NTA0NzQ3NGQzODkxYjFmMWUyMGIzZjVkMDM5MGM4YzhkYTJkNDAwNmYwMjQ2YTk1ZjI0Y2IwOTUwMmYwNzQyZg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"2576"
    },"msg":"success"
}
```





##### 2.10 添加物品评价

请求URL

> http://<host>:<port>/api/<version>/biz/review/new

请求方式

> POST

输入参数（data字段下）

| 参数          | 类型   | 必填 | 说明           |
| ------------- | ------ | ---- | -------------- |
| caller_addr   | string | Y    | 调用者的链地址 |
| item_id       | string | Y    | 被评论的物品id |
| reviewer_addr | string | Y    | 评论者的链地址 |
| detail        | string | Y    | 评论内容       |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "item_id": "3", 
        "reviewer_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "detail": "aaaaabbbbbb \u54c8\u54c8"
    }, 
    "timestamp": 1619913711, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NzRlZWJhNmU3NDc0YzlmMmQwYmU0ZTEzODc0YThhNTc2ODVkZjcyMDBhYWQyZDVhYzVjNjg0OTliYzNjNjhlYQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"17186"
    },
    "msg":"success"
}
```



##### 2.11 修改物品评价

请求URL

> http://<host>:<port>/api/<version>/biz/review/modify

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 评论id         |
| item_id     | string | Y    | 被评论的物品id |
| detail      | string | Y    | 修改的评论内容 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "2", 
        "item_id": "3", 
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "detail": "aaaaabbbbbb \u54c8\u54c81111111"
    }, 
    "timestamp": 1619914366, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "Njc4MzgwMTBkMTBiMjFlZmJmYWE5YzViNmI0YmMzMTc3ODJlYzllNzIxOWY1Mjk4YTNjYmIxNjc2MTkzOTZmYw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"17835"
    },
    "msg":"success"
}
```



##### 2.12 修改物品评价反馈信息

请求URL

> http://<host>:<port>/api/<version>/biz/review/feedback

请求方式

> POST

输入参数（data字段下）

| 参数          | 类型   | 必填 | 说明                    |
| ------------- | ------ | ---- | ----------------------- |
| caller_addr   | string | Y    | 调用者的链地址          |
| review_id     | string | Y    | 评论id                  |
| reviewer_addr | string | Y    | 反馈者的链地址          |
| detail        | string | Y    | 反馈信息：1 赞同 0 反对 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json

```

返回示例

```json

```



##### 2.13 发起拍卖

请求URL

> http://<host>:<port>/api/<version>/biz/auction/new

请求方式

> POST

输入参数（data字段下）

| 参数             | 类型   | 必填 | 说明                         |
| ---------------- | ------ | ---- | ---------------------------- |
| caller_addr      | string | Y    | 调用者的链地址               |
| seller_addr      | string | Y    | 卖家的链地址，需与所有者一致 |
| auction_house_id | string | Y    | 拍卖行ID                     |
| item_id          | string | Y    | 物品id                       |
| reserved_price   | string | Y    | 底价                         |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "seller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "auction_house_id": "1", 
        "item_id": "2", 
        "reserved_price": "1000"
    }, 
    "timestamp": 1619946244, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NjZmNzcxMGQwZDFhMDlhNjdhNDQ4NDQ5ZWI3Yzg0MjUyZjIyNTg2OWUwMDMwZjdmY2I4YTRhMmEzODczY2I2MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"41812"
    },
    "msg":"success"
}
```



##### 2.14 修改拍卖信息

只有当拍卖还在WAIT状态时才可修改

请求URL

> http://<host>:<port>/api/<version>/biz/auction/modify

请求方式

> POST

输入参数（data字段下）

| 参数             | 类型   | 必填 | 说明           |
| ---------------- | ------ | ---- | -------------- |
| caller_addr      | string | Y    | 调用者的链地址 |
| id               | string | Y    | 拍卖id         |
| auction_house_id | string |      | 修改的拍卖行ID |
| reserved_price   | string |      | 修改的底价     |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "id": "0", 
        "auction_house_id": "", 
        "reserved_price": "2000"
    }, 
    "timestamp": 1619946339, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MTBhMDNjNzgzOTNjNjdiNWNlNWI5Mzg1OTc3MTA3ODFlN2FjZDBhN2NkOWZiZGFjZmY2YzVhOGQ5NmQyMDQ4Mg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"41907"
    },
    "msg":"success"
}
```



##### 2.15 出价

请求URL

> http://<host>:<port>/api/<version>/biz/auction/bid

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明                             |
| ----------- | ------ | ---- | -------------------------------- |
| caller_addr | string | Y    | 调用者的链地址                   |
| buyer_addr  | string | Y    | 出价者的链地址，不能是物品所有者 |
| auction_id  | string | Y    | 拍卖ID                           |
| bid_price   | string | Y    | 出价                             |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json

```

返回示例

```json

```





##### 2.16 建立成交交易

请求URL

> http://<host>:<port>/api/<version>/biz/trans/new

请求方式

> POST

输入参数（data字段下）

| 参数         | 类型   | 必填 | 说明           |
| ------------ | ------ | ---- | -------------- |
| caller_addr  | string | Y    | 调用者的链地址 |
| buyer_addr   | string | Y    | 买家的链地址   |
| auction_id   | string | Y    | 拍卖ID         |
| item_id      | string | Y    | 物品ID         |
| trans_type   | string | Y    | 交易类型       |
| hammer_time  | string | Y    | 成交时间       |
| hammer_price | string | Y    | 成交价格       |
| details      | string |      | 交易细节描述   |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "buyer_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "auction_id": "1", 
        "item_id": "2", 
        "trans_type": "BID", 
        "hammer_time": "2021-01-01", 
        "hammer_price": "1000", 
        "details": "\u6d4b\u8bd5\u6d4b\u8bd5\u6d4b\u8bd5"
    }, 
    "timestamp": 1619955365, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "Y2M2MTc4ZTk5MTk2NzFkZmMzZWY2M2U1NmQ2ODMzNzQyODI5MGEwMjE3ZjVkYjg1NTBkOTU5ZDllNmZhZDZhZA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"46674"
    },
    "msg":"success"
}
```



##### 2.17 审核用户

请求URL

> http://<host>:<port>/api/<version>/biz/audit/user

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| chain_addr  | string | Y    | 用户链地址     |
| status      | string | Y    | 用户状态       |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块的高度                          |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "chain_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "status": "ACTIVE", 
    }, 
    "timestamp": 1619767850, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "M2MyMDU1ZmIzZTE5YjY1ZDc3YjBhYWVkZjczNTVjYTIxYWRiZDdmN2VlNjQxODUyYmJlNTQxNzIyMWY4NzlmNA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"1010"
    },
    "msg":"success"
}
```





##### 2.18 审核物品

请求URL

> http://<host>:<port>/api/<version>/biz/audit/item

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 物品id         |
| status      | string | Y    | 状态           |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "0", 
        "status": "ACTIVE", 
        "caller_addr": "bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7"
    }, 
    "timestamp": 1619791789, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NTA0NzQ3NGQzODkxYjFmMWUyMGIzZjVkMDM5MGM4YzhkYTJkNDAwNmYwMjQ2YTk1ZjI0Y2IwOTUwMmYwNzQyZg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"2576"
    },"msg":"success"
}
```





##### 2.20 审核物品评价

请求URL

> http://<host>:<port>/api/<version>/biz/audit/review

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 评论id         |
| item_id     | string | Y    | 被评论的物品id |
| status      | string | Y    | 状态           |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "2", 
        "item_id": "3", 
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "status": "ACTIVE"
    }, 
    "timestamp": 1619914366, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "Njc4MzgwMTBkMTBiMjFlZmJmYWE5YzViNmI0YmMzMTc3ODJlYzllNzIxOWY1Mjk4YTNjYmIxNjc2MTkzOTZmYw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"17835"
    },
    "msg":"success"
}
```



##### 2.21 审核拍卖请求

请求URL

> http://<host>:<port>/api/<version>/biz/audit/auction

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 拍卖id         |
| status      | string | Y    | 状态           |
| open_date   | string |      | 拍卖开始时间   |
| close_date  | string |      | 拍卖截止时间   |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "id": "0", 
        "status": "ACTIVE", 
    }, 
    "timestamp": 1619946339, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MTBhMDNjNzgzOTNjNjdiNWNlNWI5Mzg1OTc3MTA3ODFlN2FjZDBhN2NkOWZiZGFjZmY2YzVhOGQ5NmQyMDQ4Mg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"41907"
    },
    "msg":"success"
}
```



##### 2.22 审核成交交易

请求URL

> http://<host>:<port>/api/<version>/biz/audit/trans

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| id          | string | Y    | 交易id         |
| status      | string | Y    | 状态           |
| action      | string |      | 用于操作日志   |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "caller_addr": "bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004", 
        "id": "0", 
        "status": "ACTIVE", 
    }, 
    "timestamp": 1619946339, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MTBhMDNjNzgzOTNjNjdiNWNlNWI5Mzg1OTc3MTA3ODFlN2FjZDBhN2NkOWZiZGFjZmY2YzVhOGQ5NmQyMDQ4Mg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "height":"41907"
    },
    "msg":"success"
}
```





##### 2.23 撤销出价

请求URL

> http://<host>:<port>/api/<version>/biz/auction/bid/withdraw

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明           |
| ----------- | ------ | ---- | -------------- |
| caller_addr | string | Y    | 调用者的链地址 |
| auction_id  | string | Y    | 拍卖ID         |
| id          | string | Y    | 出价ID         |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易区块高度                            |

请求示例

```json

```

返回示例

```json

```



#### 3. 查询接口



##### 3.1 查询用户清单

请求URL

> http://<host>:<port>/api/<version>/query/user/list

请求方式

> POST

输入参数（data字段下）

| 参数  | 类型 | 必填 | 说明              |
| ----- | ---- | ---- | ----------------- |
| page  | uint | Y    | 第几页，最小为1   |
| limit | uint | Y    | 每页数量，最小为1 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10
    }, 
    "timestamp": 1620440241, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NDkyOGI0NzBkYWMwMTJjMWVmOWM0NWVmMWY2YTk2ZWYyMmRiNWZkYTVmNWQwOTdhNzQzZDhhZTQzODM1Njc3ZA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "user_list":[
            {
                "chain_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "id":"0",
                "last_date":"2021-04-30 15:57:03",
                "login_name":"test2",
                "reg_date":"2021-04-30 15:14:38",
                "status":"ACTIVE",
                "user_type":"TRD"
            },{
                "chain_addr":"bid1art16zs5zpmsw5wezyrpnls76ytdy7ws2zpqan9ey9",
                "id":"1",
                "last_date":"",
                "login_name":"test3",
                "reg_date":"2021-05-07 13:50:25",
                "status":"ACTIVE",
                "user_type":"TRD"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.2 查询用户信息

请求URL

> http://<host>:<port>/api/<version>/query/user/info

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "chain_addr": "bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s"
    }, 
    "timestamp": 1618295472, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MzI1YzE5ZWFkM2NmNTMzNjFiMWVmYTMwM2ZhZmU2MDQwMWU0NzJkM2QzMDA1OWM1YWI0ZjY5NjUwODQwMzg0ZA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "user":{
            "address":"",
            "bank_acc_name":"",
            "bank_acc_no":"",
            "bank_name":"",
            "chain_addr":"bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s",
            "email":"111111@qq.com",
            "login_name":"test1",
            "user_type":"TRD",
            "phone":"",
            "status":"ACTIVE",
            "reg_date":"2021-04-30 10:42:16",
            "last_date":"2021-04-30 15:57:03"
        }
    },
    "msg":"success"
}
```





##### 3.3 验证用户身份

请求URL

> http://<host>:<port>/api/<version>/query/user/verify

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |
| mystery    | string | Y    | 密码字符串 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 是否验证通过                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "chain_addr": "bid1art16zs5zpmsw5wezyrpnls76ytdy7ws2zpqan9ey9", 
        "mystery": "aware race riot apart mesh release cloud asset obey noble poet throw pigeon unveil demise expose nature flat badge gentle emotion bulb claim away"
    }, 
    "timestamp": 1620368483, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NGU0MjY2Yjk5ZWY2ZDM5ZWRhZmE1ZTViODAxZWY1ZGExNjU0NjEzMzdkZjE2ZGRiYjg2NGYyOTZhOTBmNWVjMA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "verified":false, /* true 验证通过； false 验证失败 */
    },
    "msg":"success"
}
```



##### 3.4 查询物品清单

请求URL

> http://<host>:<port>/api/<version>/query/item/list

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明              |
| ---------- | ------ | ---- | ----------------- |
| page       | uint   | Y    | 第几页，最小为1   |
| limit      | uint   | Y    | 每页数量，最小为1 |
| owner_addr | string |      | 物品所有人链地址  |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 物品清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10
    },
    "timestamp": 1619834379, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ODA2MjNiY2M3MzM2NjhkNjUwMzZiZDNkMmY1NmM3ZmE5YzFlMGE5YWU4NjUzNzg5MGUyODRhY2JjNTA2YjcxYg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "item_list":[
            {
                "id":"1",
                "base_price":"$1001",
                "date":"1900s",
                "desc":"测试物品",
                "detail":"",
                "last_date":"2021-04-30 22:09:49",
                "media":"",
                "owner_addr":"bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7",
                "size":"",
                "subject":"",
                "status":"WAIT",
                "type":""
            },
            {
                "id":"2",
                "base_price":"$2001",
                "date":"1911s",
                "desc":"测试物品2",
                "detail":"",
                "last_date":"2021-04-30 22:11:46",
                "media":"",
                "owner_addr":"bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7",
                "size":"",
                "subject":"",
                "status":"WAIT",
                "type":""
            }
        ]
    },
    "msg":"success"
}
```



##### 3.5 查询物品信息

请求URL

> http://<host>:<port>/api/<version>/query/item/info

请求方式

> POST

输入参数（data字段下）

| 参数 | 类型   | 必填 | 说明   |
| ---- | ------ | ---- | ------ |
| id   | string | Y    | 物品id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 物品信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256",
    "data": {
        "id": "0"
    }, 
    "timestamp": 1619791976, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "YTgyZTBmNmIyZTNjYTVlZmU2ZWFiMTI3ZmMzOGVhZjYyZTIzNjkzZjlmYzQ3OTYxNTU5Nzg3MGVkYWI3MDQ1OQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "item":{
            "id":"1",
            "base_price":"$1001",
            "date":"1900s",
            "desc":"测试物品",
            "detail":"",
            "last_date":"2021-04-30 22:09:49",
            "media":"",
            "owner_addr":"bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7",
            "size":"",
            "subject":"",
            "status":"WAIT",
            "type":""
        }
    },
    "msg":"success"
}
```



##### 3.7 查询评价清单

请求URL

> http://<host>:<port>/api/<version>/query/review/list

请求方式

> POST

输入参数（data字段下）

| 参数    | 类型   | 必填 | 说明              |
| ------- | ------ | ---- | ----------------- |
| item_id | string | Y    | 物品id            |
| status  | string | Y    | 状态              |
| page    | uint   | Y    | 第几页，最小为1   |
| limit   | uint   | Y    | 每页数量，最小为1 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 物品评价清单数据                        |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10, 
        "item_id": "3",
        "status": "ACTIVE",
    }, 
    "timestamp": 1619914700, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NDU5OWU4MmFlMzM0MjNkY2Q5ZWJkYWU4OWM4YzY2ZjFmZTViZDRjOTY0MjM2OTBmMWJjNjY4ODE5MWYyZjE4MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "review_list":[
            {
                "id":"1",
                "detail":"aaaaa 哈哈",
                "item_id":"3",
                "last_date":"",
                "review_date":"2021-05-02 08:01:41",
                "reviewer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":"WAIT"
            },{
                "id":"2",
                "detail":"aaaaabbbbbb 哈哈1111111",
                "item_id":"3",
                "last_date":"2021-05-02 08:12:47",
                "review_date":"2021-05-02 08:01:51",
                "reviewer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":""
            }
        ]
    },
    "msg":"success"
}
```



##### 3.8 查询评价信息

请求URL

> http://<host>:<port>/api/<version>/query/review/info

请求方式

> POST

输入参数（data字段下）

| 参数    | 类型   | 必填 | 说明   |
| ------- | ------ | ---- | ------ |
| id      | string | Y    | 评价id |
| item_id | string | Y    | 物品id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 评价信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "3", 
        "item_id": "0", 
    }, 
    "timestamp": 1621934310, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZTJjM2RlYjRhOGVlOTY4NTFiN2U1OTczNjU1MzViYjU3Mzc1OTYyNWUwMGZjZDExZWZkODdhMTA0NjI3Yzg2ZQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "review":{
            "detail":"xxxxxxxxxxxxxx123",
            "id":"3",
            "item_id":"0",
            "last_date":[
                {
                    "act":"new",
                    "caller":"bid1art1p25qdr4y7fclldmjdre4yss0fkltfteyz78p3h",
                    "date":"2021-05-25 16:30:53"
                },{
                    "act":"edit",
                    "caller":"bid1art1p25qdr4y7fclldmjdre4yss0fkltfteyz78p3h",
                    "date":"2021-05-25 17:08:35"
                }
            ],
            "review_date":"2021-05-25 16:30:53",
            "reviewer_addr":"bid1art1p25qdr4y7fclldmjdre4yss0fkltfteyz78p3h",
            "status":"WAIT"
        }
    },
    "msg":"success"
}
```



##### 3.9 查询拍卖行清单

请求URL

> http://<host>:<port>/api/<version>/query/auction_house/list

请求方式

> POST

输入参数（data字段下）

| 参数 | 类型 | 必填 | 说明 |
| ---- | ---- | ---- | ---- |
|      |      |      |      |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 拍卖行清单数据                          |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {}, 
    "timestamp": 1621231367, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NTQyNzEyYjNkMDY2MzczODE5NDhjN2Y2NWI1ZDY3Mzg4M2UwOWRlYjM4MjdhYTRmODJmZjU4NDc4NDhkZDRlNw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "ah_list":[
            {
                "chain_addr":"bid1art18xx25laenhg07hjqcrys05v0svpjzumqpzshag",
                "id":"6",
                "last_date":"",
                "login_name":"ah",
                "reg_date":"2021-05-17 13:58:41",
                "status":"WAIT",
                "user_type":"AH|AUC_OP|TRANS_OP"
            }
        ]
    },
    "msg":"success"
}
```





##### 3.11 查询拍卖清单

请求URL

> http://<host>:<port>/api/<version>/query/auction/list

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明              |
| ----------- | ------ | ---- | ----------------- |
| page        | uint   | Y    | 第几页，最小为1   |
| limit       | uint   | Y    | 每页数量，最小为1 |
| seller_addr | string |      | 卖家的链地址      |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 拍卖清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10
    }, 
    "timestamp": 1619947049, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NzJhOTliNGZmMWYyMmVjYjA3NTIyNjM3ZTJkOWM0N2QyZTQxOTEwYzZmMDM1Y2QwNWFjZDQxN2E4YzlhYTM0MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "auction_list":[
            {
                "auction_house_id":"1",
                "close_date":"",
                "id":"0",
                "item_id":"2",
                "last_date":"2021-05-02 17:05:40",
                "open_date":"",
                "req_date":"2021-05-02 17:04:04",
                "reserved_price":"2000",
                "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":"WAIT"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.12 查询拍卖信息

请求URL

> http://<host>:<port>/api/<version>/query/auction/info

请求方式

> POST

输入参数（data字段下）

| 参数 | 类型   | 必填 | 说明   |
| ---- | ------ | ---- | ------ |
| id   | string | Y    | 拍卖id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 拍卖信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "0", 
    }, 
    "timestamp": 1619946364, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MzUxZmJiMzNkYjRhZjk0ZGQ1ZDZkYmE2MTkyYjRlMDhjMTlhZWMzZTYyYzU0ZGQ5NWZhYTQ2NmQxNDA4ZGI2MA=="
}

```

返回示例

```json
{
    "code":0,
    "data":{
        "auction":{
            "auction_house_id":"1",
            "close_date":"",
            "id":"0",
            "item_id":"2",
            "last_date":"2021-05-02 17:05:40",
            "open_date":"",
            "req_date":"2021-05-02 17:04:04",
            "reserved_price":"2000",
            "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            "status":"WAIT"
        }
    },
    "msg":"success"
}
```



##### 3.13 查询出价清单

请求URL

> http://<host>:<port>/api/<version>/query/bid/list

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明              |
| ---------- | ------ | ---- | ----------------- |
| auction_id | string | Y    | 拍卖id            |
| status     | string | Y    | 状态              |
| page       | uint   | Y    | 第几页，最小为1   |
| limit      | uint   | Y    | 每页数量，最小为1 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 出价评价清单数据                        |

请求示例

```json

```

返回示例

```json

```



##### 3.14 查询最高出价

请求URL

> http://<host>:<port>/api/<version>/query/bid/highest

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明   |
| ---------- | ------ | ---- | ------ |
| auction_id | string | Y    | 拍卖id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 出价信息数据                            |

请求示例

```json

```

返回示例

```json

```





##### 3.15 查询出价信息

请求URL

> http://<host>:<port>/api/<version>/query/bid/info

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明   |
| ---------- | ------ | ---- | ------ |
| id         | string | Y    | 出价id |
| auction_id | string | Y    | 拍卖id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 出价信息数据                            |

请求示例

```json

```

返回示例

```json

```





##### 3.16 查询交易清单

请求URL

> http://<host>:<port>/api/<version>/query/trans/list

请求方式

> POST

输入参数（data字段下）

| 参数  | 类型 | 必填 | 说明              |
| ----- | ---- | ---- | ----------------- |
| page  | uint | Y    | 第几页，最小为1   |
| limit | uint | Y    | 每页数量，最小为1 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10
    }, 
    "timestamp": 1619955520, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MjQ5N2YxYjQzNzM1Y2ZmNGU1YjQ3N2U0Njg2NWM1OTEyYmY2ODc3YjAyMjAzZWYwODE1ZjljYjViZGUzMzlkZg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "trans_list":[
            {
                "auction_id":"1",
                "buyer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "details":"测试测试测试",
                "hammer_price":"1000",
                "hammer_time":"2021-01-01",
                "id":"0",
                "item_id":"2",
                "last_date":"",
                "status":"WAIT",
                "trans_date":"2021-05-02 19:36:05",
                "trans_type":"BID"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.17 查询交易信息

请求URL

> http://<host>:<port>/api/<version>/query/trans/info

请求方式

> POST

输入参数（data字段下）

| 参数 | 类型   | 必填 | 说明   |
| ---- | ------ | ---- | ------ |
| id   | string | Y    | 交易id |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易信息数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "id": "0"
    }, 
    "timestamp": 1619955444, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "M2M3NmI4MTUzMTRmYzQxYmZkZmQzMWY4ODk4OWZhZGUzYWYxZTJmYzkzYjU5YTdhODk0MGYwN2FmMmQ5YzBjNw=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "trans":{
            "auction_id":"1",
            "buyer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
            "details":"测试测试测试",
            "hammer_price":"1000",
            "hammer_time":"2021-01-01",
            "id":"0",
            "item_id":"2",
            "last_date":"",
            "status":"WAIT",
            "trans_date":"2021-05-02 19:36:05",
            "trans_type":"BID"
        }
    },
    "msg":"success"
}
```



##### 3.18 查询指定区块原始数据

请求URL

> http://<host>:<port>/api/<version>/query/block/rawdata

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明     |
| ---------- | ------ | ---- | -------- |
| chain_addr | string | Y    | 用户公钥 |
| height     | string | Y    | 区块高度 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 指定区块的原始区块数据                  |

> 说明：
>
> 按区块height查询时没有限制链用户范围。

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "chain_addr": "bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s", 
        "height": '210274'
    }, 
    "timestamp": 1618284344, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "OWMxZDZlMGYxNDY2Y2Q1YWQyN2JlZGQzYzcxY2Y0ZGNlYmNmOTBmODRjNjM5MzA4ZmYyZDg0MWY2Y2FlZTFjYQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "blcok":{
            "block":{
                "data":{
                    "txs":[
                        "CqwCCqkCCiYvamFjazEzOS5hcnRjaGFpbi5wZXJzb24uTXNnQ3JlYXRlVXNlchL+AQouYmlkMWFydDEzZTlqdjU4enNtZnh3dWRobjdtczR3amwycWhmY2pzODZmcHIzaxIEVVNFUhoFdGVzdDIiA1RSRCpteyJiYW5rX2FjY19uYW1lIjoiIiwiYmFua19hY2Nfbm8iOiIiLCJiYW5rX25hbWUiOiIiLCJjb250YWN0X2FkZHJlc3MiOiIiLCJlbWFpbCI6IiIsInBob25lIjoiIiwicmVmZXJyZXIiOiIifTIGQUNUSVZFOhMyMDIxLTA0LTMwIDExOjAzOjQ4Qi5iaWQxYXJ0MWdhcHFnMzdjZDd5ZWNrcWhkc2o3Z2pkc21wYWZkajQwNGhqcGUwElgKUApGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQLdwByqNVe+df4QZR/Yq8t2UoXYFaWxQ0munbSo1mkdoBIECgIIARgCEgQQwJoMGkCnfSDF4ZLtXXGP1F+tMGThtkboS3iCIhc0oaOTYjg77RL76b4WurfAObKDZbR+QsUqFGoVt7OyO0s3tLQbK9di"
                    ]
                },
                "evidence":{
                    "evidence":null
                },
                "header":{
                    "app_hash":"AF4548DF024EE021AC7A6E35648976E6552488509F701E06265CBD3E4081FEDA",
                    "chain_id":"artchain",
                    "consensus_hash":"048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
                    "data_hash":"07978E0B458CFB2C5F062FA919C2CD0F92C098AB54AF521B47EBE549F372DA07",
                    "evidence_hash":"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
                    "height":"1259",
                    "last_block_id":{
                        "hash":"FC84079FA9CC64F429CF06FBABF415F110970FC91DCCCE86B4F2ABCAFF253AAE",
                        "parts":{
                            "hash":"7009EA8E692ECAB0B348B895E04E8D210C7304B5E76BE5A8A1D276C355652894",
                            "total":1
                        }
                    },
                    "last_commit_hash":"4AF3B0FDAEFED4A1EAD4356495696A4EE9A055452642CAD7E1CAE0BDD14BFEFA",
                    "last_results_hash":"14D9AFD75C7C2F0341E2EBB0E77B75F5A6B5919EB008DB544DB3D38AAC067A89",
                    "next_validators_hash":"B908A599F87E13F27706AFD4E3FC1121EF6DF4DC4DBFD9C49A579091413A0A7D",
                    "proposer_address":"698C00E52A944F1C7A198ECAA9AA3FD9326CADAB",
                    "time":"2021-04-30T03:03:48.446392688Z",
                    "validators_hash":"B908A599F87E13F27706AFD4E3FC1121EF6DF4DC4DBFD9C49A579091413A0A7D",
                    "version":{"block":"11"}
                },
                "last_commit":{
                    "block_id":{
                        "hash":"FC84079FA9CC64F429CF06FBABF415F110970FC91DCCCE86B4F2ABCAFF253AAE",
                        "parts":{
                            "hash":"7009EA8E692ECAB0B348B895E04E8D210C7304B5E76BE5A8A1D276C355652894",
                            "total":1
                        }
                    },
                    "height":"1258",
                    "round":0,
                    "signatures":[
                        {
                            "block_id_flag":2,
                            "signature":"c+7wxW2QLxzwKdWshsZBQ3bko5pl2hqLRy9/gpXwKZDwrp5lVjJx+ltfTiwpUuvmw3K2P+QMgQlG9tAOSb7HDQ==",
                            "timestamp":"2021-04-30T03:03:48.446392688Z",
                            "validator_address":"698C00E52A944F1C7A198ECAA9AA3FD9326CADAB"
                        }
                    ]
                }
            },
            "block_id":{
                "hash":"E7835037816A9FF2835C0DAFFF794216B397478254D067FC7B4C04EB3AE9E583",
                "parts":{
                    "hash":"7E139A4B6A98100E0938C99CCB1889A9E5FE93DB6429EB6ACFD39BD4C14F9C06",
                    "total":1
                }
            }
        }
    },
    "msg":"success"
}
```



##### 3.19 查询用户通证

请求URL

> http://<host>:<port>/api/<version>/query/user/credit_balance

请求方式

> POST

输入参数（data字段下）

| 参数       | 类型   | 必填 | 说明       |
| ---------- | ------ | ---- | ---------- |
| chain_addr | string | Y    | 用户链地址 |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 通证余额信息                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "chain_addr": "bid1art1rv0rvemwkw9m7tpcng53ez3ngdzttxgmtrxx3s"
    }, 
    "timestamp": 1618295472, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MzI1YzE5ZWFkM2NmNTMzNjFiMWVmYTMwM2ZhZmU2MDQwMWU0NzJkM2QzMDA1OWM1YWI0ZjY5NjUwODQwMzg0ZA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "balance":{
            "amount":"1",
            "denom":"credit"
        }
    },
    "msg":"success"
}
```



##### 3.20 指定状态的用户清单

请求URL

> http://<host>:<port>/api/<version>/query/user/list_by_status

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明              |
| ------ | ------ | ---- | ----------------- |
| page   | uint   | Y    | 第几页，最小为1   |
| limit  | uint   | Y    | 每页数量，最小为1 |
| status | string | Y    | 用户状态          |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 用户清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10,
        "status": "WAIT",
    }, 
    "timestamp": 1620440241, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NDkyOGI0NzBkYWMwMTJjMWVmOWM0NWVmMWY2YTk2ZWYyMmRiNWZkYTVmNWQwOTdhNzQzZDhhZTQzODM1Njc3ZA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "user_list":[
            {
                "chain_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "id":"0",
                "last_date":"2021-04-30 15:57:03",
                "login_name":"test2",
                "reg_date":"2021-04-30 15:14:38",
                "status":"ACTIVE",
                "user_type":"TRD"
            },{
                "chain_addr":"bid1art16zs5zpmsw5wezyrpnls76ytdy7ws2zpqan9ey9",
                "id":"1",
                "last_date":"",
                "login_name":"test3",
                "reg_date":"2021-05-07 13:50:25",
                "status":"ACTIVE",
                "user_type":"TRD"
            }
        ]
    },
    "msg":"success"
}
```





##### 3.21 指定状态的物品清单

请求URL

> http://<host>:<port>/api/<version>/query/item/list_by_status

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明              |
| ------ | ------ | ---- | ----------------- |
| page   | uint   | Y    | 第几页，最小为1   |
| limit  | uint   | Y    | 每页数量，最小为1 |
| status | string | Y    | 状态              |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 物品清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10,
        "status": "WAIT", 
    },
    "timestamp": 1619834379, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ODA2MjNiY2M3MzM2NjhkNjUwMzZiZDNkMmY1NmM3ZmE5YzFlMGE5YWU4NjUzNzg5MGUyODRhY2JjNTA2YjcxYg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "item_list":[
            {
                "id":"1",
                "base_price":"$1001",
                "date":"1900s",
                "desc":"测试物品",
                "detail":"",
                "last_date":"2021-04-30 22:09:49",
                "media":"",
                "owner_addr":"bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7",
                "size":"",
                "subject":"",
                "status":"WAIT",
                "type":""
            },
            {
                "id":"2",
                "base_price":"$2001",
                "date":"1911s",
                "desc":"测试物品2",
                "detail":"",
                "last_date":"2021-04-30 22:11:46",
                "media":"",
                "owner_addr":"bid1art1jv8z6e3507g2eeanep29dpx5m8qn83023gx3g7",
                "size":"",
                "subject":"",
                "status":"WAIT",
                "type":""
            }
        ]
    },
    "msg":"success"
}
```





##### 3.22 指定状态的评价清单

请求URL

> http://<host>:<port>/api/<version>/query/review/list_by_status

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明              |
| ------ | ------ | ---- | ----------------- |
| page   | uint   | Y    | 第几页，最小为1   |
| limit  | uint   | Y    | 每页数量，最小为1 |
| status | string | Y    | 状态              |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 物品评价清单数据                        |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10,
        "Status": "WAIT", 
    }, 
    "timestamp": 1619914700, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NDU5OWU4MmFlMzM0MjNkY2Q5ZWJkYWU4OWM4YzY2ZjFmZTViZDRjOTY0MjM2OTBmMWJjNjY4ODE5MWYyZjE4MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "review_list":[
            {
                "id":"1",
                "detail":"aaaaa 哈哈",
                "item_id":"3",
                "last_date":"",
                "review_date":"2021-05-02 08:01:41",
                "reviewer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":"WAIT"
            },{
                "id":"2",
                "detail":"aaaaabbbbbb 哈哈1111111",
                "item_id":"3",
                "last_date":"2021-05-02 08:12:47",
                "review_date":"2021-05-02 08:01:51",
                "reviewer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":""
            }
        ]
    },
    "msg":"success"
}
```





##### 3.23 指定状态的拍卖申请清单

请求URL

> http://<host>:<port>/api/<version>/query/auction/list_by_status

请求方式

> POST

输入参数（data字段下）

| 参数   | 类型   | 必填 | 说明              |
| ------ | ------ | ---- | ----------------- |
| page   | uint   | Y    | 第几页，最小为1   |
| limit  | uint   | Y    | 每页数量，最小为1 |
| status | string | Y    | 状态              |

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 拍卖清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "page": 1, 
        "limit": 10,
        "status": "WAIT", 
    }, 
    "timestamp": 1619947049, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "NzJhOTliNGZmMWYyMmVjYjA3NTIyNjM3ZTJkOWM0N2QyZTQxOTEwYzZmMDM1Y2QwNWFjZDQxN2E4YzlhYTM0MQ=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "auction_list":[
            {
                "auction_house_id":"1",
                "close_date":"",
                "id":"0",
                "item_id":"2",
                "last_date":"2021-05-02 17:05:40",
                "open_date":"",
                "req_date":"2021-05-02 17:04:04",
                "reserved_price":"2000",
                "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "status":"WAIT"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.24 指定条件查询交易清单

请求URL

> http://<host>:<port>/api/<version>/query/trans/list_by_condition

请求方式

> POST

输入参数（data字段下）

| 参数      | 类型   | 必填 | 说明              |
| --------- | ------ | ---- | ----------------- |
| cate      | string | Y    | 查询类别          |
| condition | string | Y    | 查询条件          |
| page      | uint   | Y    | 第几页，最小为1   |
| limit     | uint   | Y    | 每页数量，最小为1 |

> cate取值： seller, buyer, item, status

返回结果

| 参数 | 类型   | 说明                                    |
| ---- | ------ | --------------------------------------- |
| code | int    | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string | 成功时返回success；出错时，返回出错信息 |
| data | json   | 交易清单数据                            |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "cate": "status",
        "condition": "WAIT|PAID",
        "page": 1, 
        "limit": 10
    }, 
    "timestamp": 1619955520, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MjQ5N2YxYjQzNzM1Y2ZmNGU1YjQ3N2U0Njg2NWM1OTEyYmY2ODc3YjAyMjAzZWYwODE1ZjljYjViZGUzMzlkZg=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "trans_list":[
            {
                "auction_id":"1",
                "buyer_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "seller_addr":"bid1art18e3jj0yyzvu9vsg5d09fz6tz44kuc0r88uv004",
                "details":"测试测试测试",
                "hammer_price":"1000",
                "hammer_time":"2021-01-01",
                "id":"0",
                "item_id":"2",
                "last_date":"",
                "status":"WAIT",
                "trans_date":"2021-05-02 19:36:05",
                "trans_type":"BID"
            }
        ]
    },
    "msg":"success"
}
```



##### 3.25 从IPFS下载数据

请求URL

> http://<host>:<port>/api/ipfs/download

请求方式

> POST

输入参数（data字段下）

| 参数        | 类型   | 必填 | 说明             |
| ----------- | ------ | ---------------- | ---------------- |
| hash | string | Y | IPFS文件hash值 |

返回结果

| 参数 | 类型           | 说明                                    |
| ---- | -------------- | --------------------------------------- |
| code | int            | 状态代码，0 表示成功，非0 表示出错      |
| msg  | string         | 成功时返回success；出错时，返回出错信息 |
| data | base64编码的数据 |                                         |

请求示例

```json
{
    "version": "1", 
    "sign_type": "SHA256", 
    "data": {
        "hash": "QmbkAewP7KN9VMKwYbx6xj45Yc82dPraRnARFjww4MLBt3"
    }, 
    "timestamp": 1622626276, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "ZDUxMzhiNTQ2MjYyNTRiZWY1YWFhZjIwZDEzNDg4OTA1YTdiYjkzZDllOTIyYzYzNTUzZWFmMTI2N2NkZDVkNA=="
}
```

返回示例

```json
{
    "code":0,
    "data":{
        "data":"iVBORw0KGgoAAAANSUhEUgAA...AAAABJRU5ErkJggg=="},
    "msg":"success"
}
```


