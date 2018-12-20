package myframework

import (
	"reflect"
)

type Handle func(*Context)

type AutoRouterInfo struct {
	rt   reflect.Type
	id   int
	name string
}

type Router struct {
	autoHandlerByPath map[string]AutoRouterInfo
}
