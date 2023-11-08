package gc

import "time"

func (gc *GC) Start() {
	ticker := time.NewTicker(gc.period)
	for {
		select {
		case <-gc.ctx.Done():
			break
		case <-ticker.C:
			gc.clear()
			gc.markBlobs()
		}
	}
}
