package models

import "time"

type ListenMsgGasNetStatus struct {
	Id                 int `orm:"pk;auto"`
	ReceiveBlockHeight int
	CreateTime         time.Time
	UpdateTime         time.Time
}

type ExpendMessages struct {
	Id                 int `orm:"pk;auto"`
	MessageId          string
	WalletId           string
	To                 string
	Epoch              string
	Gas                string
	BaseBurnFee        string
	OverEstimationBurn string
	Value              string
	//RewardValue        string
	Penalty string
	Method  uint64
	//Time       string
	CreateTime time.Time
}

type ExpendInfo struct {
	Id int `orm:"pk;auto"`
	//MinerId            string
	WalletId           string
	Epoch              string
	Gas                string
	BaseBurnFee        string
	OverEstimationBurn string
	Value              string
	//RewardValue        string
	Penalty    string
	Time       string
	UpdateTime time.Time
}

//reeward
type ListenRewardNetStatus struct {
	Id                 int `orm:"pk;auto"`
	ReceiveBlockHeight int
	CreateTime         time.Time
	UpdateTime         time.Time
}
type MinerInfo struct {
	Id           int `orm:"pk;auto"`
	MinerId      string
	QualityPower float64
	Pleage       float64
	CreateTime   time.Time
	UpdateTime   time.Time
}

//miner power status
type MinerPowerStatus struct {
	Id        int `orm:"pk;auto"`
	MinerId   string
	Power     float64
	Available float64
	PreCommit float64
	Vesting   float64
	Pleage    float64
	Time      string
}

type RewardInfo struct {
	Id      int `orm:"pk;auto"`
	MinerId string
	//	WalletId string
	Epoch      int
	Value      string
	Pledge     float64
	Power      float64
	BlockNum   int
	WinCounts  int64
	Time       string
	UpdateTime time.Time
}

type RewardInfoFormer struct {
	Id      int `orm:"pk;auto"`
	MinerId string
	//	WalletId string
	Epoch      int
	Reward     string
	Gas        string
	Penalty    string
	Value      string
	Power      float64
	Time       string
	UpdateTime time.Time
}

//出块记录
type MineBlocks struct {
	Id      int `orm:"pk;auto"`
	MinerId string
	//	WalletId string
	Epoch    int
	Reward   string
	Gas      string
	Penalty  string
	Value    string
	Power    float64
	WinCount int64
	//Time       string
	CreateTime time.Time
}

//打包的message记录
type MineMessages struct {
	Id         int `orm:"pk;auto"`
	MinerId    string
	MessageId  string
	Epoch      string
	Gas        string
	Penalty    string
	CreateTime time.Time
}

type MessageRewardInfo struct {
	Id      int `orm:"pk;auto"`
	MinerId string
	//	WalletId string
	Epoch string

	Value      string
	Time       string
	UpdateTime int64
}

type WalletBaseinfo struct {
	Id       int `orm:"pk;auto"`
	WalletId string
	//NodeId         string
	BalanceFil     string
	BalanceAttofil string
	CreateTime     int64
	UpdateTime     int64
	Status         string
}

type WalletHistoryData struct {
	Id       int `orm:"pk;auto"`
	WalletId string
	//NodeId         string
	BalanceFil     string
	BalanceAttofil string
	CreateTime     int64
	Status         string
}

type WalletProfitInfo struct {
	Id             int `orm:"pk;auto"`
	SettlementType string
	SettlementDate string
	WalletId       string
	//MinerId        string
	StartAmount string
	EndAmount   string
	Amount      string
	CreateTime  int64
	Status      string
}

type NetRunDataPro struct {
	Id                 int `orm:"pk;auto"`
	ReceiveBlockHeight int
	TotalShare         int
	AllShare           int
	CreateTime         int64
	UpdateTime         int64
}

type RewardMessages struct {
	Id        int `orm:"pk;auto"`
	MessageId string
	MinerId   string
	From      string
	Epoch     string

	Value string
	//RewardValue        string
	Method uint64

	Time       string
	CreateTime uint64
}

//分配至order和user版本使用 以下
type UserInfo struct {
	Id            int `orm:"pk;auto"`
	UserId        int
	Share         int
	Power         float64 //算力
	Available     float64 //可用余额
	TotalPleage   float64 //总质押
	AdvancePleage float64 //垫付质押
	Vesting       float64 //锁定金额
	Release       float64 //已释放
	Reward        float64 //总奖励
	Fee           float64 //总奖励
	UpdateTime    string
}

type UserBlockRewardInfo struct {
	Id         int `orm:"pk;auto"`
	UserId     int
	MinerId    string
	Reward     float64
	Power      float64
	Epoch      string
	CreateTime uint64
}

type UserDailyRewardInfo struct {
	Id               int `orm:"pk;auto"`
	UserId           int
	Reward           float64
	Power            float64
	Pledge           float64
	Fee              float64
	ImmediateRelease float64
	LinearRelease    float64
	Time             string
}

type OrderInfo struct {
	Id         int `orm:"pk;auto"`
	OrderId    int
	UserId     int
	Share      int
	Power      float64
	Reward     float64
	Epoch      int
	Time       string
	UpdateTime uint64
}

type OrderBlockRewardInfo struct {
	Id         int `orm:"pk;auto"`
	OrderId    int
	MinerId    string
	Reward     float64
	Power      float64
	Epoch      int
	CreateTime uint64
}

type OrderDailyRewardInfo struct {
	Id      int `orm:"pk;auto"`
	OrderId int
	//MinerId    string
	Reward     float64
	Pleage     float64
	Power      float64
	Fee        float64
	Epoch      int
	Time       string
	UpdateTime uint64
}

type SettlePlan struct {
	Id        int
	TreatyId  int
	ArticleId int
	Quantity  float64
	AddTime   string
}

type OrderGoods struct {
	Id         int
	ChannelId  int
	ArticleId  int
	OrderId    int
	GoodsId    int
	GoodsNo    string
	GoodsTitle string
	ImgUrl     string
	SpecText   string
	GoodsPrice float64
	RealPrice  float64
	Quantity   int
	Point      int
}

type Orders struct {
	Id               int
	SiteId           int
	OrderNo          string
	TradeNo          string
	UserId           int
	UserName         string
	PaymentId        int
	PaymentFee       float64
	PaymentStatus    int
	PaymentTime      string
	IdcFee           float64
	TreatyId         int
	TreatyStart_date int64
	TreatyEnd_date   int64
	TreatyQuantity   float64
	TreatyStatus     int
	TreatyTime       string
	ExpressId        int
	ExpressNo        string
	ExpressFee       float64
	ExpressStatus    int
	ExpressTime      string
	AcceptName       string
	PostCode         string
	Telphone         string
	Mobile           string
	Email            string
	Area             string
	Address          string
	Message          string
	Remark           string
	IsEsign          int
	IsIdc            int
	IsConfirm        int
	IsInvoice        int
	InvoiceTitle     string
	InvoiceTaxes     float64
	PayableAmount    float64
	RealAmount       float64
	OrderAmount      float64
	Point            int
	Status           int
	AddTime          string
	ConfirmTime      string
	CompleteTime     string
	PaymentVoucher   string
}

type OrderDailyCostInfo struct {
	Id          int `orm:"pk;auto"`
	OrderId     int
	UserId      int
	Expend      float64
	ValueReward float64
	Time        string
}

type VestingInfo struct {
	Id        int `orm:"pk;auto"`
	UserId    int
	Vesting   float64
	Release   float64
	Times     int32
	StartTime string
}

//转账记录表
type Transfer struct {
	Id            int `orm:"pk;auto"`
	From          string
	To            string
	ServiceCharge float64
	Value         float64
	Time          int64
}

//user 信息初始化表
type UserFilDaily struct {
	Id        int `orm:"pk;auto"`
	UserId    int
	Date      string
	FilAmount float64
	Type      int32
	Days      int32
	Remark    string
	AddTime   string
}

type UserFilPledge struct {
	Id         int `orm:"pk;auto"`
	UserId     int
	Date       string
	FilPledge  float64
	FilRelease float64
	Type       int32
	Days       int32
	Remark     string
	AddTime    string
}

type MinerAndWalletRelation struct {
	Id       int `orm:"pk;auto"`
	MinerId  string
	WalletId string
}
