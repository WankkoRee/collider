package collider

import (
	"fmt"
	"time"
)

func (c *Collider[Seed]) speedometerLoop() {
	ticker := time.NewTicker(c.meterInterval)
	for {
		select {
		case <-c.finished.Done():
			return
		case <-ticker.C:
			var collided = c.collided

			currentTime := time.Now()

			passTime := currentTime.Sub(c.lastTime).Milliseconds()
			passCollided := collided - c.lastCollided
			speed := float64(passCollided) / float64(passTime)

			cost := currentTime.Sub(c.startTime)
			estimateMill := (c.estimateCollisions - float64(collided)) / speed
			estimate := time.Duration(estimateMill) * time.Millisecond
			if estimateMill*1000*1000 > 1<<63-1 {
				estimate = time.Duration(1<<63 - 1)
			}

			c.lastTime = currentTime
			c.lastCollided = collided

			speedString := fmt.Sprintf(
				"collided: %d,\tspeed: %.3f kHash/s,\tcost: %s,\testimate: %s",
				collided,
				speed,
				cost.String(),
				estimate.String(),
			)

			c.speedometer(speedString)
		}
	}
}
