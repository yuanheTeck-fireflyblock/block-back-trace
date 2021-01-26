package wallet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	lotusClient "github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"
	cbg "github.com/whyrusleeping/cbor-gen"
	"net/http"
	"profit-allocation/models"
	"reflect"
	"strconv"
	"time"
)

var procedureRates = 0.01
var walletLog = logging.Logger("wallet-log")

func Transfer(to string, value float64, uid int) {
	//验证余额

	userInfo := new(models.UserInfo)
	n, err := models.O.QueryTable("fly_user_info").Filter("user_id", uid).All(userInfo)
	if err != nil {
		walletLog.Error("Error query user :%+v info err:%+v ", uid, err)
		return
	}
	if n == 0 {
		walletLog.Error("Error no this user :%+v info ", uid)
		return
	}
	if userInfo.Available < value*(1+procedureRates) {
		walletLog.Error("Error user's available is not enough")
		return
	}

	requestHeader := http.Header{}
	ctx := context.Background()

	ndoeApi, closer, err := lotusClient.NewFullNodeRPC(ctx, models.LotusHost, requestHeader)
	if err != nil {
		walletLog.Error("Error transfer NewFullNodeRPC err:%+v", err)
		return
	}
	defer closer()
	//获取默认wallet地址
	fromAddr, err := ndoeApi.WalletDefaultAddress(ctx)
	if err != nil {
		walletLog.Error("Error transfer WalletDefaultAddress  err:%+v", err)
		return
	}
	balance, err := ndoeApi.WalletBalance(ctx, fromAddr)
	if err != nil {
		walletLog.Error("Error transfer get wallet :%+v balance   err:%+v", fromAddr, err)
		return
	}

	toAddr, err := address.NewFromString(to)
	if err != nil {
		walletLog.Error("Error transfer NewFromString toAddr err:%+v", err)
		return
	}
	valStr := strconv.FormatFloat(value, 'f', 18, 64)
	val, err := types.ParseFIL(valStr)
	if err != nil {
		walletLog.Error("Error transfer ParseFIL  err:%+v", err)
		return
	}

	if balance.Int64() < val.Int64() {
		walletLog.Error("Error walllet's :%+v balance :%+v is not enough", fromAddr, balance.Int64())
		return
	}
	gp, err := types.BigFromString("0")
	if err != nil {
		return
	}
	gfc, err := types.BigFromString("0")
	if err != nil {
		return
	}

	method := abi.MethodNum(0)

	var params []byte

	/*decparams, err := decodeTypedParams(ctx, api, toAddr, method, cctx.String("params-json"))
	if err != nil {
		return
	}
	params = decparams*/

	msg := &types.Message{
		From:       fromAddr,
		To:         toAddr,
		Value:      types.BigInt(val),
		GasPremium: gp,
		GasFeeCap:  gfc,
		GasLimit:   0,
		Method:     method,
		Params:     params,
	}

	sm, err := ndoeApi.MpoolPushMessage(ctx, msg, nil)
	if err != nil {
		return
	}
	userInfo.Available = userInfo.Available - value*(1+procedureRates)
	_, err = models.O.Update(userInfo)
	if err != nil {
		walletLog.Error("Error update user :%+v  info err:%+v", userInfo.UserId, err)
		return
	}
	transferInfo := models.Transfer{
		From:          fromAddr.String(),
		To:            to,
		ServiceCharge: value * procedureRates,
		Value:         value,
		Time:          time.Now().Unix(),
	}
	fmt.Println(sm.Cid())
	_, err = models.O.Insert(&transferInfo)
	if err != nil {
		walletLog.Error("Error insert transfer info  err:%+v", userInfo.UserId, err)
		return
	}
	return
}

func decodeTypedParams(ctx context.Context, fapi api.FullNode, to address.Address, method abi.MethodNum, paramstr string) ([]byte, error) {
	act, err := fapi.StateGetActor(ctx, to, types.EmptyTSK)
	if err != nil {
		return nil, err
	}

	methodMeta, found := stmgr.MethodsMap[act.Code][method]
	if !found {
		return nil, fmt.Errorf("method %d not found on actor %s", method, act.Code)
	}

	p := reflect.New(methodMeta.Params.Elem()).Interface().(cbg.CBORMarshaler)

	if err := json.Unmarshal([]byte(paramstr), p); err != nil {
		return nil, fmt.Errorf("unmarshaling input into params type: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := p.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
