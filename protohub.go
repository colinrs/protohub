package main

import (
	"flag"
	"fmt"

	"github.com/colinrs/protohub/internal/config"
	"github.com/colinrs/protohub/internal/handler"
	"github.com/colinrs/protohub/internal/svc"
	"github.com/colinrs/protohub/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/protohub-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandlerCtx(response.ErrHandle)
	httpx.SetOkHandler(response.OKHandle)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
