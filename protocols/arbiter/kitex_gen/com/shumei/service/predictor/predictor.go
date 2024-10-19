// Code generated by Kitex v0.11.3. DO NOT EDIT.

package predictor

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	service "github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"predict": kitex.NewMethodInfo(
		predictHandler,
		newPredictorPredictArgs,
		newPredictorPredictResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"health": kitex.NewMethodInfo(
		healthHandler,
		newPredictorHealthArgs,
		newPredictorHealthResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	predictorServiceInfo                = NewServiceInfo()
	predictorServiceInfoForClient       = NewServiceInfoForClient()
	predictorServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return predictorServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return predictorServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return predictorServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Predictor"
	handlerType := (*service.Predictor)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "service",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func predictHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.PredictorPredictArgs)
	realResult := result.(*service.PredictorPredictResult)
	success, err := handler.(service.Predictor).Predict(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPredictorPredictArgs() interface{} {
	return service.NewPredictorPredictArgs()
}

func newPredictorPredictResult() interface{} {
	return service.NewPredictorPredictResult()
}

func healthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	_ = arg.(*service.PredictorHealthArgs)
	realResult := result.(*service.PredictorHealthResult)
	success, err := handler.(service.Predictor).Health(ctx)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newPredictorHealthArgs() interface{} {
	return service.NewPredictorHealthArgs()
}

func newPredictorHealthResult() interface{} {
	return service.NewPredictorHealthResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Predict(ctx context.Context, request *service.PredictRequest) (r *service.PredictResult_, err error) {
	var _args service.PredictorPredictArgs
	_args.Request = request
	var _result service.PredictorPredictResult
	if err = p.c.Call(ctx, "predict", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Health(ctx context.Context) (r bool, err error) {
	var _args service.PredictorHealthArgs
	var _result service.PredictorHealthResult
	if err = p.c.Call(ctx, "health", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
