package collider

import (
	"context"
	"runtime"
	"time"
)

// New 生成新的碰撞器
func New[Seed any](
	generator func(finished context.Context, seedChan chan Seed),
	checker func(seed Seed) bool,
	speedometer func(speedString string),

	estimateCollisions float64,
	threads int,
	maxSeedCache uint64,
	meterInterval time.Duration,
) (collider Collider[Seed]) {
	collider = Collider[Seed]{}

	collider.generator = generator
	collider.checker = checker
	collider.speedometer = speedometer

	collider.estimateCollisions = estimateCollisions
	collider.threads = threads
	if collider.threads == 0 {
		collider.threads = runtime.GOMAXPROCS(0)
		if collider.threads > 1 {
			collider.threads-- // 为主线程保留性能
		}
	}
	collider.maxSeedCache = maxSeedCache
	collider.meterInterval = meterInterval

	return
}
