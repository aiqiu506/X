//+build jsonRpcClient  !production

package jsonRpcClient

import (
	"log"
	"net/rpc/jsonrpc"
	"x/global"
	"x/utils"
)

type JsonRpcStruct struct {
	Host string `map:"host"`
	Port string `map:"port"`
}

func (j *JsonRpcStruct) NewComponent(config interface{}) {
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, j)
		if err != nil {
			log.Fatal(err)
		}
	}else{
		log.Fatal("jsonRpcClient配置文件错误")
	}
}

var JsonRpc JsonRpcStruct

func init() {
	//注册组件
	global.Global.Register("jsonRpcClient", &JsonRpc)
}

func (j *JsonRpcStruct) Request(method string, jsonData interface{}) (interface{}, error) {
	var ret interface{}
	conn, err := jsonrpc.Dial("tcp", j.Host+":"+j.Port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	err = conn.Call(method, jsonData, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
