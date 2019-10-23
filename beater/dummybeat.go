package beater

import (
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb/module"
	"github.com/sreejithavg/dummybeat/config"
	"time"
)
type dummybeat struct {
	done         chan struct{}  // Channel used to initiate shutdown.
	modules      []staticModule // Active list of modules.
	config       config.Config
	client beat.Client
	// Options
	moduleOptions []module.Option
}
type staticModule struct {
	connector *module.Connector
	module    *module.Wrapper
}
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &dummybeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *dummybeat) Run(b *beat.Beat) error {
	logp.Info("dummybeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client,err = b.Publisher.Connect()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":  "dummybeat",
				"unity": counter,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}
func (bt *dummybeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
