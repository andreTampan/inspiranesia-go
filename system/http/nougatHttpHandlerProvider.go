package http

import (
	"fmt"
	"inspiranesia/system/config"
	"inspiranesia/system/logging"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/labstack/echo/v4"
)

type NougatHttpHandler interface {
	Start()
}

type NougatHttpStruct struct {
	config      *config.DefaultConfig
	logger      logging.NougatLoggingProvider
	httphandler interface{}
}

func NewNougat(config *config.DefaultConfig, logger logging.NougatLoggingProvider, httpHandler interface{}) *NougatHttpStruct {
	return &NougatHttpStruct{config: config, logger: logger, httphandler: httpHandler}
}

func ProvideHttpHandler(config *config.DefaultConfig, logger logging.NougatLoggingProvider) NougatHttpHandler {
	serverConfig := config.Server
	if serverConfig.HttpHandler == "echo" {
		var realEcho echo.Echo
		return NewNougat(config, logger, &realEcho)
	}
	var realIris iris.Application
	return NewNougat(config, logger, &realIris)

	// serverConfig := config.Server
	// if serverConfig.HttpHandler == "echo" {
	// 	return echo.New(serverConfig, logger)
	// }
	// return iris.New(serverConfig, logger)
}

func (n *NougatHttpStruct) Start() {
	fmt.Println("SETAN")
	serverConfig := n.config.Server
	logger := n.logger
	e, ok := n.httphandler.(*echo.Echo)
	if ok {
		fmt.Println("ECHO")
		e = echo.New()
		e.GET(serverConfig.RootName, func(c echo.Context) error {
			tytyd()
			return c.String(http.StatusOK, "This is default end point using echo :) kuda")
		})
		logger.Fatal(e.Start(":" + strconv.Itoa(serverConfig.Port)))
	}
	i, ok := n.httphandler.(*iris.Application)
	if ok {
		fmt.Println("IRIS")
		i = iris.Default()
		i.Get(serverConfig.RootName, func(ctx iris.Context) {
			tytyd()
			ctx.WriteString("This is default end point using iris :) sempak")
		})
		i.Listen(":" + strconv.Itoa(serverConfig.Port))
	}
	// go echo.New(serverConfig, logger).Start()
	// go falseIris.New(serverConfig, logger).Start()
}

func tytyd() {
	fmt.Println("IMPLEMENT YOUR BUSINESS LOGIC HERE")
}
