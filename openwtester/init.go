package openwtester

import (
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/virtualeconomy-adapter/virtualeconomy"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(virtualeconomy.Symbol, virtualeconomy.NewWalletManager())
}
