package main

import (
	"github.com/micro/go-plugins/micro/cors/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
	"github.com/opentracing/opentracing-go"
	tracer "github.com/zhaojiaming110/openBank/plugins/tracer/jaeger"
	"github.com/zhaojiaming110/openBank/plugins/tracer/opentracing/stdhttp"
	"log"
)

func init() {
	plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
}

const name = "API gateway"

func main() {
	stdhttp.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	cmd.Init()
}
