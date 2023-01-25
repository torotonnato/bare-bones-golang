package agent

import (
	"time"

	"sync"

	"github.com/torotonnato/gobarebones/config"
)

var beat struct {
	ticker *time.Ticker
	sync.Mutex
}

func tickerStart() {
	beat.Lock()
	if beat.ticker != nil {
		beat.ticker.Stop()
	}
	beat.ticker = time.NewTicker(config.AgentTickerInterval * time.Second)
	beat.Unlock()
}

func tickerReset() {
	beat.Lock()
	beat.ticker.Reset(config.AgentTickerInterval * time.Second)
	beat.Unlock()
}

func tickerStop() {
	beat.Lock()
	beat.ticker.Stop()
	beat.ticker = nil
	beat.Unlock()
}
