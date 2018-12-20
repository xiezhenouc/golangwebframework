# 1 package
$GOPATH环境下面增加一个新文件夹，文件夹命名为 myframework

# 2 使用示例
```go
package main

import (
    "fmt"
    "myframework"
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


