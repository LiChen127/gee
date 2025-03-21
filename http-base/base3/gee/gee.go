package gee

import (
	"fmt"
	"log"
	"net/http"
)

/*
	HandleFunc 定义一个函数类型
*/
type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

// Engine实例化
func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}


func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	// key为路由规则
	key := method + "-" + pattern
	log.Printf("Route %s-%s\n", method, pattern)
	// 添加路由规则
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}


func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}