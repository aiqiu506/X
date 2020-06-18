package X

import (
	"github.com/aiqiu506/x/components"
	"github.com/aiqiu506/x/conf"
	"github.com/aiqiu506/x/global"
	"log"
)

type X struct {
}

func NewX(path string) *X {

	global.Global.Config = &conf.ConfigEngine{}
	Err = global.Global.Config.Load(path)
	if Err != nil {
		log.Fatal(Err)
	}
	//注册组件
	components.Register()
	return &X{}
}

var (
	Err error
)

func (x *X) Start() {
	global.Run()
}

func (x *X) SyncStart() {
	global.RunSync()
}
