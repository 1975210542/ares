package common

import (
	"path"

	"github.com/yuin/gopher-lua"
)

const configDir = `./config`

var (
	cfg *lua.LTable
	ls  *lua.LState
)
var Config = make(map[string]map[string]string)

func init() {
	ls = lua.NewState()
	defer ls.Close()
}

func InitGlbConf() {
	LoadGlobalConfigs("ares.lua")

	cfg.ForEach(func(key, value lua.LValue) {
		ci := make(map[string]string)
		cfg.RawGetString(key.String()).(*lua.LTable).ForEach(func(k, v lua.LValue) {
			ci[k.String()] = v.String()
		})
		Config[key.String()] = ci
	})
}

func LoadGlobalConfigs(file string) {
	doFile(file)
	cfg = ls.Get(-1).(*lua.LTable)
	ls.Pop(1)
}

func doFile(file string) {
	err := ls.DoFile(path.Join(configDir, file))
	if err != nil {
		panic(err.Error())
	}
}