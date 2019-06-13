# 1 package
go get github.com/xiezhenouc/golangwebframework

# 2 使用示例
```golang
package main

import (
	"fmt"
	myframework "github.com/xiezhenouc/golangwebframework"
)

type TestController struct {
	Ctx *myframework.Context
}

func (t *TestController) Init(context *myframework.Context) {
	t.Ctx = context
}

func (t *TestController) SayHi() {
	fmt.Fprintln(t.Ctx.Output, "say hi ...")
}

func (t *TestController) SayYes() {
	fmt.Fprintln(t.Ctx.Output, "say yes ...")
}

func main() {
	fw := myframework.New()

	fw.AddAutoRouter("/test/", &TestController{})

	fw.Run(":8999")

}


```


