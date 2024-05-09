package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func GetConfig(confPath string, key string) *g.Var {
	var ctx = gctx.New()
	adapter, err := gcfg.NewAdapterFile(confPath)
	if err != nil {
		panic(err)
	}
	config := gcfg.NewWithAdapter(adapter)
	val, err := config.Get(ctx, key)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return val
}
func GetConfigs(key string) *g.Var {
	var ctx = gctx.New()
	val, err := g.Cfg().Get(ctx, key)
	if err != nil {
		fmt.Println("读取配置失败", err.Error())
		return nil
	}
	return val
}
