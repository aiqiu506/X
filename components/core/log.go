package core

import (
	"github.com/aiqiu506/x/global"
	"github.com/aiqiu506/x/utils"
	"log"
	"os"
	"time"
)

type LogHandle struct {
	Path     string `map:"path"`
	FileName string `map:"fileName"`
	IsDaily  bool   `map:"isDaily"`
	NeedDir  bool   `map:"needDir"`
	File     *os.File
}

func (l *LogHandle) NewComponent(config interface{}) {
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf, l)
		if err != nil {
			log.Fatal(err)
		}
	}
	var name string
	if l.FileName != "" {
		name = l.FileName
	} else {
		name = "log"
	}
	l.makeFileName(name)
}

func (l *LogHandle) makeFileName(name string) {
	fileName := l.Path
	//每天生成
	if l.IsDaily {
		if l.NeedDir {
			fileName += "/" + time.Now().Format("20060102") + "/"
		} else {
			name = "/" + time.Now().Format("20060102") + "_" + name
		}
	} else {
		fileName += "/"
	}
	l.FileName = fileName + name + ".log"
}

func (l LogHandle) logWrite(content string, isExit bool) {
	l.File = utils.OpenFile(l.FileName)
	defer l.File.Close()
	logFile := log.New(l.File, "", log.LstdFlags)
	if isExit {
		logFile.Fatal(content)
	} else {
		logFile.Println(content)
	}
}

//重新修改日志文件名
func (l *LogHandle) SetFileName(name string) *LogHandle {
	L := &LogHandle{}
	L.makeFileName(name)
	return L
}

//调试输出
func (l LogHandle) Debug(content ...interface{}) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(utils.OutPutString(content...))
}

//日志记录到文件
func (l LogHandle) Info(content ...interface{}) {
	l.logWrite(utils.OutPutString(content...), false)
}

//记录日志，并结束程序
func (l LogHandle) Error(content ...interface{}) {
	l.logWrite(utils.OutPutString(content...), true)
}

var Log LogHandle

func init() {
	//注册日志组件
	global.Global.Register("log", &Log)
}
