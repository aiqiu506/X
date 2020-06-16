package components

import (
	_ "x/components/core"
	"x/global"
)

func Register() {
	for componentName,v:=range global.Global.Components{
		config:=global.Global.Config.Get(componentName)
		v.NewComponent(config)
	}
}
