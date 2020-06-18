package X

import (
	"github.com/aiqiu506/x/global"
	"sync"
	"testing"
)

type TT struct{}

func (t *TT) Run(locker *sync.WaitGroup, name string) {
	defer func() {
		locker.Done()
	}()
	//具体测试的业务代码

}
func TestNewX(t *testing.T) {
	global.Global.RegisterService("test", &TT{})
	x := NewX("config.yaml")

	x.Start()
}
