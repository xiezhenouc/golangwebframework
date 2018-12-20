package myframework

import (
	"net/http"
	"reflect"
	"strings"
)

type Framework struct {
	r Router
}

type ControllerInterface interface {
	Init(context *Context)
}

func New() *Framework {
	return &Framework{}
}

func (fw *Framework) AddAutoRouter(prefix string, c ControllerInterface) {
	if fw.r.autoHandlerByPath == nil {
		fw.r.autoHandlerByPath = make(map[string]AutoRouterInfo)
	}

	reflectVal := reflect.ValueOf(c)

	for i := 0; i < reflectVal.NumMethod(); i++ {
		method := reflectVal.Method(i)
		name := reflectVal.Type().Method(i).Name

		var path string
		if prefix[len(prefix)-1] == '/' {
			path = strings.ToLower(prefix + name)
		} else {
			path = strings.ToLower(prefix + "/" + name)
		}
		if method.IsValid() {
			fw.r.autoHandlerByPath[path] = AutoRouterInfo{
				rt:   reflect.Indirect(reflectVal).Type(),
				id:   i,
				name: name,
			}
		}
	}
}

func (fw *Framework) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := &Context{
		Input:  r,
		Output: w,
	}

	path := r.URL.Path
	path = strings.ToLower(path)

	autoHandle, ok := fw.r.autoHandlerByPath[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	curController := reflect.New(autoHandle.rt)
	exeController, ok := curController.Interface().(ControllerInterface)

	if !ok {
		panic("curController.Interface().(ControllerInterface) error!")
		return
	}

	exeController.Init(context)
	method := curController.Method(autoHandle.id)

	// execute
	method.Call(nil)
}

func (fw *Framework) Run(server string) {
	http.ListenAndServe(server, fw)
}
