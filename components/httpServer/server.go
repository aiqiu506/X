//+build httpServer !production

package httpServer

import (
	"log"
	"net/http"
	"x/global"
	"x/utils"
)

type HandlerFunc func(*Context)

type Server struct {
	host string
	port string
	router *Router
}

func New(host,port string) *Server {
	return &Server{
		host: host,
		port: port,
		router: newRouter()}
}

func (s *Server) addRoute(method string, pattern string, handler HandlerFunc) {
	s.router.addRoute(method, pattern, handler)
}

func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.addRoute("GET", pattern, handler)
}

func (s *Server) POST(pattern string, handler HandlerFunc) {
	s.addRoute("POST", pattern, handler)
}

func (s *Server) Run() (err error) {
	return http.ListenAndServe(s.host+":"+s.port, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	s.router.handle(c)
}


type Config struct {
	Port string `map:"port"`
	Host string `map:"host"`
}
type httpServer struct {
	Server *Server
}


func (h *httpServer) NewComponent(config interface{}) {
	params:=&Config{}
	if conf, ok := config.(map[interface{}]interface{}); ok {
		err := utils.MapToStruct(conf,params)
		if err != nil {
			log.Fatal(err)
		}
		h.Server=h.ServerRun(params)
	}else{
		log.Fatal("httpServer配置文件错误")
	}

}
func (h * httpServer)ServerRun(p *Config) *Server{
	return New(p.Host,p.Port)
}
var HttpServer httpServer

func init(){
	//注册日志组件
	global.Global.Register("httpServer", &HttpServer)
}
