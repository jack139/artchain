package types

const (
	// 交易类型
	ActionRegister = "10" // 注册
	ActionContract = "11" // 签合同
	ActionDelivery = "12" // 合同验收

	// 通证奖励
	RewardRegister = "1credit" // 注册
	RewardContract = "2credit" // 签合同
	RewardDelivery = "3credit" // 合同验收
)

var FaucetAddress string
