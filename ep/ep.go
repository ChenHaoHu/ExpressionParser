package ep

import (
	ep "github.com/ChenHaoHu/ExpressionParser/internal"
	log "github.com/sirupsen/logrus"
)

type EpEngine struct {
	E *ep.Engine
}

func init() {
	log.SetLevel(log.InfoLevel)
	//	log.SetReportCaller(true)
}

func NewEpEngine(rule string) (*EpEngine, error) {

	engine, err := ep.NewEngine(rule)

	if err != nil {
		return nil, err
	}

	epEngine := &EpEngine{
		E: engine,
	}
	return epEngine, nil
}

func (ep *EpEngine) Check(context map[string]string) bool {

	return ep.E.Calculate(context)
}

func EnableDebugLogLevel() {
	log.SetLevel(log.DebugLevel)
}
