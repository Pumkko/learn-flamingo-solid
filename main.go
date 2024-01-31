package main

import (
	helloworld "pumkko/learnFlamingo/src/helloWorld"
	solidfront "pumkko/learnFlamingo/src/solidFront"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/healthcheck"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/core/zap"
	"flamingo.me/flamingo/v3/framework/opencensus"
)

func main() {
	flamingo.App([]dingo.Module{
		new(healthcheck.Module),
		new(opencensus.Module),
		new(zap.Module),
		new(requestlogger.Module),
		new(gotemplate.Module),
		new(helloworld.Module),
		new(solidfront.Module),
	})
}
