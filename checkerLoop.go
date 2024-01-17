package collider

import "sync/atomic"

func (c *Collider[Seed]) checkerLoop() {
	defer c.colliders.Done()

	for {
		select {
		case <-c.finished.Done():
			return
		case seed, ok := <-c.seedChan:
			if !ok {
				return
			}
			if c.checker(seed) {
				c.finish()
				c.resultChan <- seed
				return
			}
			atomic.AddUint64(&c.collided, 1)
		}
	}
}
