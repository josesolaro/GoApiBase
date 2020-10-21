package providers

import (
	"ApiBase/src/api/utils"
)

func Logger() *utils.Logger{
	log := &utils.Logger{}
	log.Init()
	log.SetLogLevel("warn")
	return log
}