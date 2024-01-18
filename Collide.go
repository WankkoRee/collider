package collider

import (
	"context"
	"sync"
	"time"
)

func (c *Collider[Seed]) Collide(ctx context.Context) *Seed {
	c.finished, c.finish = context.WithCancel(ctx)
	c.colliders = sync.WaitGroup{}
	c.seedChan, c.resultChan = make(chan Seed, c.maxSeedCache), make(chan Seed, c.threads)

	c.startTime, c.lastTime = time.Now(), time.Now()

	go c.generatorLoop()
	go c.speedometerLoop()

	for i := 0; i < c.threads; i++ {
		c.colliders.Add(1)
		go c.checkerLoop()
	}

	go func() {
		c.colliders.Wait() // 可能有任务无解的情况
		c.finish()
	}()

	<-c.finished.Done()
	c.colliders.Wait()
	close(c.resultChan)

	if len(c.resultChan) > 0 {
		finalOutput := <-c.resultChan
		return &finalOutput
	} else {
		return nil
	}
}
