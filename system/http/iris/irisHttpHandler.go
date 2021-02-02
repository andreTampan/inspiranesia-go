package iris

import (
	"inspiranesia/system/config"
	"inspiranesia/system/logging"
	"strconv"

	"github.com/kataras/iris/v12"
)

type IrisHttpHandler struct {
	port        int
	rootName    string
	httpHandler *iris.Application
	logger      logging.NougatLoggingProvider
}

func New(server config.Server, logger logging.NougatLoggingProvider) *IrisHttpHandler {
	return &IrisHttpHandler{
		port:        server.Port,
		rootName:    server.RootName,
		httpHandler: iris.Default(),
		logger:      logger,
	}
}

func (i IrisHttpHandler) Start() {
	i.httpHandler.Get(i.rootName, i.defaultRootEndpoint)
	i.httpHandler.Listen(":" + strconv.Itoa(i.port))
}

func (i IrisHttpHandler) defaultRootEndpoint(ctx iris.Context) {
	ctx.WriteString("This is default end point using iris :) sempak")
}
