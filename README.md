# X
#### 开始
```
在main包下
//package main  main.go
import(
  _ "***/services"  //载入业务配置文件
)
//载入配置文件
x:=NewX("config.yaml")
//启动服务
x.Start()


//package services service.go
type TT struct{}

//业务服务执行入口
func (t *TT)Run(locker *sync.WaitGroup, name string){
	defer func() {
		locker.Done()
	}()
	//具体业务代码

}
//注册业务服务
func init(){
  global.Global.RegisterService("test",&TT{})
}

```

#### 编译
 通过条件编译的方式来达到按需加载包
 ```
go build -tags "需要的组件"
 如: go build -tags "mysql redis mongo"
```
 