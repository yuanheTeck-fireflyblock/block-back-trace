package models

import (
	"github.com/filecoin-project/go-state-types/abi"
	"time"
)

type WalletInfoResp struct {
	Code          string
	Msg           string
	WalletBalance string
}

type WalletProfitResp struct {
	Code   string
	Msg    string
	Amount string
}

type RewardRespFormer struct {
	Code   string
	Msg    string
	Reward string
}

type PostResp struct {
	Code string
	Msg  string
}

type OrderInfoRequest struct {
	OrderId int
	UserId  int
	Share   int
	Power   float64
}

//--------------------------------
type RewardResp struct {
	Code           string
	Msg            string
	Reward         float64
	Pledge         float64
	Power          float64
	Gas            float64
	BlockNumber    int
	SectorsNumber  int64
	WinCount       int64
	TotalPower     float64
	TotalAvailable float64
	TotalPreCommit float64
	TotalPleage    float64
	TotalVesting   float64
	WindowPostGas  float64
	Penalty        float64
	Update         time.Time
}

type MessageGasTmp struct {
	Code string
	Msg  string
	Gas  float64
}

//-----------------------------
type GetBlockPercentageResp struct {
	Code            string
	Msg             string
	MinedPercentage string
	Mined           []BlockInfo
	Missed          []BlockInfo
}

type GetMinersLuckResp struct {
	Code       string
	Msg        string
	MinersLuck []MinerLuck
}

type MinerLuck struct {
	Miner       string
	Luck        string
	Power       float64
	BlockNumber int
	TotalValue  float64
}

type BlockInfo struct {
	Epoch abi.ChainEpoch
	Time  string
}
