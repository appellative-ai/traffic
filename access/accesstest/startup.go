package host

const (
	egressLogOperatorNameFmt  = "fs/egress_logging_operators.json"
	ingressLogOperatorNameFmt = "fs/ingress_logging_operators.json"
)

/*
func Startup[E runtime.ErrorHandler, O runtime.OutputHandler](mux *http.ServeMux) (http.Handler, *runtime.Status) {
	var e E

	initOrigin()
	err := initLogging()
	if err != nil {
		return nil, e.Handle(nil, "/host/startup/logging", err)
	}
	status := initExtract[E]()
	if !status.OK() {
		return nil, status
	}
	errs := initControllers()
	if len(errs) > 0 {
		return nil, e.Handle(nil, "/host/startup/controllers", errs...)
	}
	initMux(mux)
	status = startupResources[E, O]()
	if !status.OK() {
		return mux, status
	}

	middleware2.ControllerWrapTransport(exchange.Client)
	return middleware2.ControllerHttpHostMetricsHandler(mux, ""), status
}


*/

/*
func Shutdown() {
	messaging.Shutdown()
}

func startupResources[E runtime.ErrorHandler, O runtime.OutputHandler]() *runtime.Status {
	return messaging.Startup[E, O](time.Second*5, nil)
}


*/

func initOrigin() {
	/*	shared.SetOrigin(shared.Origin{
			Region:     "Region",
			Zone:       "Zone",
			SubZone:    "SubZone",
			Service:    "Service",
			InstanceId: "InstanceId",
		})

	*/
}

/*
func initLogging() error {
	// Options that are defaulted to true for the statuses
	accesslog.SetIngressLogStatus(true)
	accesslog.SetEgressLogStatus(true)
	accesslog.SetPingLogStatus(true)

	// Enable logging function for access events middleware
	// middleware.SetLogFn(func(entry *data.Entry) {
	// log.Write[log.DebugOutputHandler, data.JsonFormatter](entry)
	//},
	//)
	controller.SetLogFn(func(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName string, timeout int, limit rate.Limit, burst int, retry, proxy, statusFlags string) {
		entry := accessdata.NewEntry(traffic, start, duration, req, resp, routeName, timeout, limit, burst, retry, proxy, statusFlags)
		accesslog.Write[accesslog.DebugOutputHandler, accessdata.JsonFormatter](entry)
	},
	)

	err := access.CreateIngressOperators(func() ([]byte, error) {
		return resource.ReadFile(ingressLogOperatorNameFmt)
	})
	if err == nil {
		err = accesslog.CreateEgressOperators(func() ([]byte, error) {
			return resource.ReadFile(egressLogOperatorNameFmt)
		})
	}
	return err
}

*/

/*
func initExtract[E runtime.ErrorHandler]() *runtime.Status {
	// middleware.SetExtractFn(func(req *http.Request) string {
	// 	return req.URL.Path
	// })
	//controller.SetExtractFn(func(req *http.Request) string {
	//	return req.URL.Path
	//})
	return accesslog2.Initialize[E]("http://localhost:8086/access-log", nil)
}


*/
