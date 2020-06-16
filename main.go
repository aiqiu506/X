package X

import (
	"log"
	"x/components"
	"x/conf"
	"x/global"
)

type X struct {

}

func NewX(path string) *X {

	global.Global.Config=&conf.ConfigEngine{}
	Err=global.Global.Config.Load(path)
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

func (x *X)Start(){
	global.Run()
}

func (x *X)SyncStart(){
	global.RunSync()
}




