## 模块划分



### Person

#### 人员信息 User
```
ID        // 用户id (UserID)
RecType   // Type = USER
Name      // 名称
UserType  // 用户类型: TRD, BANK, AH, DEL, ART, REV
Address   // 地址
Phone     // 电话
Email     // 邮件
Bank      // 银行
AccountNo // 银行账号
RegDate   // 注册日期
ChainAddr // 区块链用户地址(AccAddress)
Status    // 用户状态： ACTIVE, CLOSE, FORBID, WAIT
lastDate  // 最后修改日期
```



### Inventory

#### 物品信息 Item
```
ID             // 物品编号 (ItemID)
RecType        // ARTINV
ItemDesc       // 物品描述
ItemDetail     // 物品细节描述
ItemDate       // 物品出现年代
ItemType       // 类型：原作、复制品
ItemSubject    // 主题：古代、现代、风景、雕塑、人像，等等
ItemMedia      // 材质：石头、金属、瓷器、油画、素描、等等
ItemSize       // 尺寸
ItemImage      // 数字照片信息（存储在IPFS）
AESKey         // 物品密钥
ItemBasePrice  // 最近交易价格
CurrentOwnerID // 所有人信息 UserID
Status         // 状态: WAIT, OPEN, CLOSE, IN-BID 
lastDate       // 最后修改日期
```

#### 评价信息 Review

```
ID             // 评价编号
ItemID         // 物品编号
RecType        // REVIEW
ReviewerID     // 评论者 的 UserID
ReviewDetail   // 评价内容
ReviewDate     // 评论时间
Up             // 支持数量
Down           // 反对数量
Status         // 状态: WAIT, OPEN, CLOSE
lastDate       // 最后修改日期
```



### Auction

#### 拍卖请求信息 Request
```
ID              // 拍卖ID (AuctionID)
RecType         // AUCREQ
ItemID          // 物品编号
AuctionHouseID  // 拍卖行 USerID
SellerID        // 卖家UserID（须与物品所有者ID一致）
RequestDate     // 请求日期
ReservePrice    // 底价
Status          // 拍卖状态：INIT, OPEN, CLOSED  (由拍卖行设置)
OpenDate        // OPEN的日期 (由拍卖行设置)
CloseDate       // CLOSE的日期  (由拍卖行设置)
lastDate        // 最后修改日期
```

#### 出价信息 Bid
```
ID         // 出价ID (BidID)
AuctionID  // 拍卖ID
RecType    // BID
BidNo      // 出价编号
ItemID     // 物品编号
BuyerID    // 买家UserID（须与物品所有者UserID不同）
BidPrice   // 出价（须大于前一次出价）
BidTime    // 出价时间
Status     // 状态： ACCEPT, DENY
lastDate   // 最后修改日期
```





### Trans

#### 交易信息 transaction
```
ID           // 交易ID
AuctionID    // 拍卖ID
RecType      // POSTTRAN
ItemID       // 物品ID
TransType    // 交易类型：立即购买、竞价
UserId       // 买家ID
TransDate    // 成交时间
HammerTime   // 买家出价时间（成交价格出价）
HammerPrice  // 成交价格
Details      // 交易细节记录
Status       // POST_ACTION, SHIPPING, SUCCESS, WAIT
lastDate     // 最后修改日期
```







## 接口功能



#### 业务处理接口


| 序号 | 接口功能             | URI                        |
| :--: | -------------------- | -------------------------- |
|  1   | 注册用户             | /biz/user/register         |
|  2   | 修改用户信息         | /biz/user/modify           |
|  3   | 恢复已注册用户       | /biz/user/restore          |
|  4   | 新建物品             | /biz/item/new              |
|  5   | 修改物品信息         | /biz/item/modify           |
|  6   | 添加物品照片         | /biz/item/new_image        |
|  7   | 删除物品照片         | /biz/item/remove_image     |
|  8   | 修改物品所有人       | /biz/item/change_owner     |
|  9   | 生成物品NFT          | /biz/nft/new               |
|  10  | 添加物品评价         | /biz/review/new            |
|  11  | 修改物品评价         | /biz/review/modify         |
|  12  | 修改物品评价反馈信息 | /biz/review/feedback       |
|  13  | 发起拍卖             | /biz/auction/new           |
|  14  | 修改拍卖状态         | /biz/auction/change_status |
|  15  | 出价                 | /biz/auction/bid           |
|  16  | 建立成交交易         | /biz/transaction/new       |
|  17  | 审核用户             | /biz/audit/user            |
|  18  | 审核物品             | /biz/audit/item            |
|  19  | 审核照片             | /biz/audit/image           |
|  20  | 审核评价             | /biz/audit/review          |
|  21  | 审核拍卖请求         | /biz/audit/auction         |
|  22  | 审核成交交易         | /biz/audit/transaction     |



#### 查询接口

| 序号 | 接口功能             | URI                        |
| :--: | -------------------- | -------------------------- |
|  1   | 查询用户信息         | /query/user/info           |
|  2   | 验证用户身份         | /query/user/verify         |
|  3   | 查询物品清单         | /query/item/list           |
|  4   | 查询物品信息         | /query/item/info           |
|  5   | 验证物品NFT          | /query/nft/verify          |
|  6   | 查询物品评价清单     | /query/review/list         |
|  7   | 查询物品评价信息     | /query/review/info         |
|  8   | 查询拍卖行清单       | /query/auction_house/list  |
|  9   | 查询拍卖行信息       | /query/auction_house/info  |
|  10  | 查询拍卖清单         | /query/auction/list        |
|  11  | 查询拍卖信息         | /query/auction/info        |
|  12  | 查询出价信息         | /query/bid/info            |
|  13  | 查询最高出价         | /query/bid/highest         |
|  14  | 查询出价清单         | /query/bid/list            |
|  15  | 查询成交交易         | /query/transaction/list    |
|  16  | 查询成交交易信息     | /query/transaction/info    |
|  17  | 查询指定区块原始数据 | /query/block/rawdata       |
|  18  | 查询用户通证         | /query/user/credit_balance |



