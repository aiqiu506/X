package core

import (
	"log"
	"x/global"
	"x/utils"
)

type env struct {
	IsPro bool `map:"isPro"`
	Debug bool `map:"debug"`
}

func (e env) GetProEnv() bool {
	return e.IsPro
}
func (e env) GetDebug() bool {
	return e.Debug
}
func (e *env) SetDebug(debug bool) {
	e.Debug = debug
}
func (e *env) NewComponent(config interface{}) {

	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, e)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("env配置文件错误")
	}

}

var Env env

func init() {
	//注册日志组件
	global.Global.Register("env", &Env)
}
