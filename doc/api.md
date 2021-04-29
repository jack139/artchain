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


| 接口功能                 | URI                        |
| -------------------- | -------------------------- |
| 注册用户             | /biz/user/register             |
| 修改用户信息          | /biz/user/modify               |
| 恢复已注册用户        | /biz/user/restore              |
| 新建物品             | /biz/item/newquery             |
| 修改物品信息          | /biz/item/modify               |
| 生成物品NFT          | /biz/nft/new                   |
| 添加物品评价          | /biz/review/new                |
| 修改物品评价          | /biz/review/modify             |
| 修改物品评价反馈信息   | /biz/review/feedback           |
| 发起拍卖             | /biz/auction/new               |
| 修改拍卖状态         | /biz/auction/change_status     |
| 出价                 | /biz/auction/bid               |
| 建立成交交易         | /biz/transaction/new           |
| 查询用户信息         | /query/user/info           |
| 验证用户身份          | /query/user/verify               |
| 查询物品清单         | /query/item/list           |
| 查询物品信息         | /query/item/info           |
| 验证物品NFT          | /query/nft/verify                |
| 查询物品评价清单     | /query/review/list         |
| 查询物品评价信息     | /query/review/info         |
| 查询拍卖行清单       | /query/auction_house/list  |
| 查询拍卖行信息       | /query/auction_house/info  |
| 查询拍卖清单         | /query/auction/list        |
| 查询拍卖信息         | /query/auction/info        |
| 查询出价信息         | /query/bid/info            |
| 查询最高出价         | /query/bid/highest         |
| 查询出价清单         | /query/bid/list            |
| 查询成交交易         | /query/transaction/list    |
| 查询成交交易信息     | /query/transaction/info    |
| 查询用户通证         | /query/user/credit_balance |
| 查询指定区块原始数据 | /query/block/rawdata |





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
    'code': 0, 
    'data': {
        'chain_addr': 'bid1art16207z4zw2ksmq6rjyxyvuugqlwcr8p8fw6lnv4', 
        'height': '406', 
        'mystery': 'mixture crew embark owner carpet topic dinosaur code rescue section gravity upset love ritual clog fat essence orient pride moral never illness access nasty'
    }, 
    'msg': 'success'
}
```






#### 3. 查询接口



##### 3.4 查询指定区块原始数据

请求URL

> http://<host>:<port>/api/<version>/query/block/rawdata

请求方式

> POST

输入参数（data字段下）

| 参数    | 类型   | 说明     |
| ------- | ------ | -------- |
| userkey | string | 用户公钥 |
| height  | string | 区块高度 |

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
        "userkey": "contract1ghfcl0hm5pxu0q0jgnl2nw3hhmrkklgyh3lgvx", 
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
    'code': 0, 
    'data': {
        'blcok': {
            'block': {
                'data': {
                    'txs': ['Co0CCooCCiwvamFjazEzOS5jb250cmFjdC5jb250cmFjdC5Nc2dDcmVhdGVDb250cmFjdBLZAQovY29udHJhY3QxOHpmZHNqem44dDUwMHF1NDNjcHhocjlzenMyMjZyODh1OGhseTkSBDEyMzQaL2NvbnRyYWN0MWxhbnJ2enhkOTl4eTAwempneGZqbTVwZHFoczVqdjZoNXo5bWV4Ii9jb250cmFjdDFnaGZjbDBobTVweHUwcTBqZ25sMm53M2hobXJra2xneWgzbGd2eCoCMTIyOnsiaW1hZ2UiOiJRbVFXdVg3bXdFNUxEdGpja2M3M3ZpV1o3TEJpVTRnRW5YcXRKdkMyN1JSUnlZIn0SWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA/6+qwsJW0bHd0OaCqa2Mfxr1lRQGE9NtS/+66lG9EO7EgQKAggBGAMSBBDAmgwaQAbZAl1RUh8EdAAWDWqx+MKKPrZ9JRW0PxgdTAdVWKT1OE8xw3Cq3wuUiuTCAajsEmCKjKeqcvz3UJC5eL1O93U=']
                }, 
                'evidence': {'evidence': None}, 
                'header': {
                    'app_hash': '767B1E1A1F42FB08187284E3831E4065AE42F2EA04C55B7EE84C8856E95BECD0', 
                    'chain_id': 'contract', 
                    'consensus_hash': '048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F', 
                    'data_hash': '71FC654A49D737F32092B64E32DE7569A072EAAEEB0DC0E4C2C16331D0672414', 
                    'evidence_hash': 'E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855', 
                    'height': '210274', 
                    'last_block_id': {
                        'hash': '4BB4962E7A1462DC665A888652652D652C21E24DF994E9260578093C8F7794A5', 
                        'parts': {
                            'hash': '35C561892C4B3F789D0AA714F8B62EBA1AE97B1A2106523C20F91111021CB3D0', 
                            'total': 1
                        }
                    }, 
                    'last_commit_hash': '3654C6C1A7BB84C44FFDCE14ACD2F3983FA4BE6678E13B99ED890845C1709E66', 
                    'last_results_hash': 'E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855', 
                    'next_validators_hash': '98B670043EDB03D07C4096C7DE5BC389EA11DB382A7E501F4635F7B73482C078', 
                    'proposer_address': '627C9E0096F61C1A40A980781B38B3CFC7B32E93', 
                    'time': '2021-04-09T09:03:37.86117745Z', 
                    'validators_hash': '98B670043EDB03D07C4096C7DE5BC389EA11DB382A7E501F4635F7B73482C078', 
                    'version': {'block': '11'}
                }, 
                'last_commit': {
                    'block_id': {
                        'hash': '4BB4962E7A1462DC665A888652652D652C21E24DF994E9260578093C8F7794A5', 
                        'parts': {
                            'hash': '35C561892C4B3F789D0AA714F8B62EBA1AE97B1A2106523C20F91111021CB3D0', 
                            'total': 1
                        }
                    }, 
                    'height': '210273', 
                    'round': 0, 
                    'signatures': [
                        {
                            'block_id_flag': 2, 
                            'signature': 'SwbBqVdczrfxTrvwCCUykvl2ZBM48MYyaGsYFirj4MksyWLI9hZSOMoHjKQqB7BxeDXaPbjqpyUBSsxHTWfbBw==', 
                            'timestamp': '2021-04-09T09:03:37.86117745Z', 
                            'validator_address': '627C9E0096F61C1A40A980781B38B3CFC7B32E93'
                        }
                    ]
                }
            }, 
            'block_id': {
                'hash': '3691B91FC6B22CB271AC0B20135200112716605F57F8C85C609FE4C2908011B1', 
                'parts': {
                    'hash': 'E2B3B49B97634892FD230C05F2D6116055070BF1043DE57C8C32C19647E19194', 
                    'total': 1
                }
            }
        }
    }, 
    'msg': 'success'
}
```



##### 3.5 查询用户通证

请求URL

> http://<host>:<port>/api/<version>/query/user/credit_balance

请求方式

> POST

输入参数（data字段下）

| 参数    | 类型   | 说明       |
| ------- | ------ | ---------- |
| userkey | string | 用户链地址 |

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
        "chain_addr": "contract1lanrvzxd99xy00zjgxfjm5pdqhs5jv6h5z9mex"
    }, 
    "timestamp": 1618295472, 
    "appid": "4fcf3871f4a023712bec9ed44ee4b709", 
    "sign_data": "MzI1YzE5ZWFkM2NmNTMzNjFiMWVmYTMwM2ZhZmU2MDQwMWU0NzJkM2QzMDA1OWM1YWI0ZjY5NjUwODQwMzg0ZA=="
}
```

返回示例

```json
{
    'code': 0, 
    'data': {
        'blcok': {
            'amount': '20',   /* 用户通证数量 */
            'denom': 'credit' /* 通证单位 */
        }
    }, 
    'msg': 'success'
}
```


