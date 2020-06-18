package components

import (
	_ "github.com/aiqiu506/x/components/core"
	"github.com/aiqiu506/x/global"
)

func Register() {
	for componentName,v:=range global.Global.Components{
		config:=global.Global.Config.Get(componentName)
		v.NewComponent(config)
	}
}
