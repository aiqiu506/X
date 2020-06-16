package global

import (
	"sync"
	"x/conf"
)

type Component interface {
	NewComponent(config interface{})
}
type Service interface {
	Run(locker *sync.WaitGroup, name string)
}
type global struct {
	Components map[string]Component
	Services  map[string]Service
	Config * conf.ConfigEngine
}

func (g *global)RegisterService(name string ,service Service){
	if _, ok := g.Services[name]; !ok {
		g.Services[name] = service
	}
}

func (g *global)Register(name string,component Component)  {
	if _, ok := g.Components[name]; !ok {
		g.Components[name] = component
	}
}

func RunSync()  {
	lock := &sync.WaitGroup{}
	for k, v := range Global.Services {
		lock.Add(1)
		go v.Run(lock, k)
	}
	lock.Wait()

}

func Run()  {
	lock := &sync.WaitGroup{}
	for k, v := range Global.Services {
		lock.Add(1)
		v.Run(lock, k)
	}
	lock.Wait()
}


var Global global

func init()  {
	Global=global{
		Components:make(map[string]Component),
		Services:make(map[string]Service),
	}
}