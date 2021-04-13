package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"profit-allocation/controllers"
	"profit-allocation/lotus"
	"profit-allocation/tool/log"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"profit-allocation/models"
	//_ "profit-allocation/routers"
	"syscall"
)

func main() {
	err := log.Init()
	if err != nil {
		fmt.Println("init log error:", err)
		return
	}
	if err := initDatabase(); err != nil {
		fmt.Println("init database error:", err)
		return
	}
	if err := models.InitData(); err != nil {
		fmt.Println("init data error:", err)
		return
	}

	//reward.TestMine()
	go lotus.Setup()
	var shutdownCh <-chan struct{}
	sigCh := make(chan os.Signal, 2)
	shutdownDone := make(chan struct{})
	go func() {
		select {
		case sig := <-sigCh:
			fmt.Println("received shutdown", "signal", sig)
		case <-shutdownCh:
			fmt.Println("received shutdown")
		}

		fmt.Println("Shutting down...")
		close(shutdownDone)
	}()
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	web.Router("/firefly/profit/total_reward_info", &controllers.RewardController{}, "get:GetRewardAndPledge")

	web.Router("/firefly/profit/total_messages_gas_info", &controllers.RewardController{}, "get:GetMessagesGas")

	web.Router("/firefly/profit/total_miner_info", &controllers.RewardController{}, "get:GetMinerInfo")
	web.Router("/firefly/profit/block", &controllers.BlockController{}, "get:GetMinerMineBlockPercentage")

	web.Run()
}

//初始化mysql
func initDatabase() error {
	// 注册数据库驱动
	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		return err
	}

	url, err := web.AppConfig.String("postgres")
	if err != nil {
		return err
	}
	maxIdle := 15
	maxConn := 15
	// 注册数据库
	if err := orm.RegisterDataBase("default", "postgres", url, orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn)); err != nil {
		return err
	}

	orm.RegisterModelWithPrefix("fly_",
		new(models.ListenMsgGasNetStatus),
		new(models.ListenRewardNetStatus),
		new(models.ExpendInfo),
		new(models.MinerInfo),
		new(models.ExpendMessages),
		new(models.PreAndProveMessages),
		//new(models.MineBlocks),
		new(models.MineMessages),
		new(models.MinerStatusAndDailyChange),
		new(models.MinerAndWalletRelation),
		new(models.MineBlockRight),
	)
	if err := orm.RunSyncdb("default", false, true); err != nil {
		return err
	}
	//lotus.InitMinerData()
	return nil
}
