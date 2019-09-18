package pimpdb

import (
	"github.com/badtheory/informer"
)

func (p *PimpDB) SetLoggerOptions(opt ...informer.Configuration) {

	var o informer.Configuration

	if len(opt) == 0 {
		o = informer.Configuration{}
	} else {
		o = opt[0]
	}

	err := informer.NewLogger(o, informer.InstanceZapLogger)
	if err != nil {
		informer.Fatalf("Could not instantiate log %s", err.Error())
	}

}

func LogDefault(id string, x interface{}, action string) {
	ctx := informer.WithFields(
		informer.Fields{
			"id":     id,
			"entity": x,
			"action": action,
		},
	)
	ctx.Infof("pimp_memory_action")
}