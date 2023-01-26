package agent

import (
	"sync"
	"time"

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
	interval := config.AgentTickerInterval * time.Second
	beat.ticker = time.NewTicker(interval)
	beat.Unlock()
}

func tickerReset() {
	beat.Lock()
	interval := config.AgentTickerInterval * time.Second
	beat.ticker.Reset(interval)
	beat.Unlock()
}

func tickerStop() {
	beat.Lock()
	beat.ticker.Stop()
	beat.ticker = nil
	beat.Unlock()
}
