package myframework

import (
	"net/http"
)

type Context struct {
	Input  *http.Request
	Output http.ResponseWriter
}
