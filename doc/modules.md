# 模块划分



## Person

### 人员信息 User
```
ID        // 用户id (UserID)
RecType   // Type = USER
Name      // 姓名
UserType  // 用户类型: TRD, BANK, AH, DEL, ART, REV
Address   // 地址
Phone     // 电话
Email     // 邮件
Bank      // 银行
AccountNo // 银行账号
Status    // 用户状态： ACTIVE, CLOSE, FORBID, WAIT
RegDate   // 注册日期
```



```shell
starport module create person
starport type user recType name userType address phone email bank accountNo status regDate --module person
```



## Inventory

### 物品信息 Item
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
```

### 评价信息 Review

```
ID             // 评价编号
ItemID         // 物品编号
RecType        // REVIEW
ReviewerID     // 评论者 的 UserID
ReviewDetail   // 评价内容
ReviewDate     // 评论时间
Up             // 支持数量
Down           // 反对数量
```



## Auction

### 拍卖请求信息 Request
```
ID              // 拍卖ID (AuctionID)
RecType         // AUCREQ
ItemID          // 物品编号
AuctionHouseID  // 拍卖行ID
SellerID        // 卖家ID（须与物品所有者ID一致）
RequestDate     // 请求日期
ReservePrice    // 底价
Status          // 拍卖状态：INIT, OPEN, CLOSED  (由拍卖行设置)
OpenDate        // OPEN的日期 (由拍卖行设置)
CloseDate       // CLOSE的日期  (由拍卖行设置)
```

### 出价信息 Bid
```
ID         // 出价ID (BidID)
AuctionID  // 拍卖ID
RecType    // BID
BidNo      // 出价编号
ItemID     // 物品编号
BuyerID    // 买家UserID（须与物品所有者UserID不同）
BidPrice   // 出价（须大于前一次出价）
BidTime    // 出价时间
```



## Trans

### 交易信息
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
```