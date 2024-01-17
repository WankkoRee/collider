package collider

import (
	"context"
	"sync"
	"time"
)

type Collider[Seed any] struct {
	/* 配置字段 */

	generator          func(finished context.Context, seedChan chan Seed) // 输入数据生成器
	checker            func(seed Seed) bool                               // 输入数据运算器
	speedometer        func(speedString string)                           // 定时测速器
	estimateCollisions float64                                            // 预估需要碰撞次数
	threads            int                                                // 线程数
	maxInputCache      uint64                                             // 输入通道最大缓存数
	meterInterval      time.Duration                                      // 测速间隔

	/* 任务通信字段 */

	finished   context.Context    // 是否已完成上下文
	finish     context.CancelFunc // 完成上下文方法, 调用后通知所有使用该上下文的线程
	colliders  sync.WaitGroup     // 线程池, 用于确保所有线程正常结束, 防止内存泄漏
	seedChan   chan Seed          // 输入通道, 用于提供碰撞的数据
	resultChan chan Seed          // 输出通道, 用于接收符合条件的碰撞结果

	/* 测速字段 */

	startTime    time.Time // 碰撞开始时间
	collided     uint64    // 已碰撞次数
	lastTime     time.Time // 上次测速时间
	lastCollided uint64    // 上次测速时已碰撞次数
}
