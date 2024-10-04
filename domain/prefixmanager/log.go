package prefixmanager

import (
	"github.com/kobradag/kobrad/infrastructure/logger"
	"github.com/kobradag/kobrad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
