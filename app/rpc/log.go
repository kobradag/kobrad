package rpc

import (
	"github.com/kobradag/kobrad/infrastructure/logger"
	"github.com/kobradag/kobrad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
