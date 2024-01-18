package collider_test

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/wankkoree/collider"
	"math"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	ctx := context.Background()

	collide := collider.New(func(finished context.Context, seedChan chan string) {
		for {
			select {
			case <-finished.Done():
				return
			default:
				seedChan <- grand.S(16)
			}
		}
	}, func(input string) bool {
		hash := sha256.Sum256([]byte(input))
		hashStr := hex.EncodeToString(hash[:])
		return gstr.HasPrefix(hashStr, "000000")
	}, func(speedString string) {
		g.Log().Debug(ctx, speedString)
	}, math.Pow(16, 6), 0, 0xffffff, 1*time.Second)

	result := collide.Collide(ctx)
	if result != nil {
		g.Log().Info(ctx, *result)
	} else {
		g.Log().Info(ctx, "no result")
		t.FailNow()
	}
}

func TestRangeSuccess(t *testing.T) {
	ctx := context.Background()

	collide := collider.New(func(finished context.Context, seedChan chan uint32) {
		for i := uint32(0); i < 0xffffffff; i++ {
			select {
			case <-finished.Done():
				return
			default:
				seedChan <- i
			}
		}
	}, func(input uint32) bool {
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data, input)
		hash := sha256.Sum256(data)
		hashStr := hex.EncodeToString(hash[:])
		return gstr.HasPrefix(hashStr, "000000")
	}, func(speedString string) {
		g.Log().Debug(ctx, speedString)
	}, math.Pow(16, 6), 0, 0xffffff, 1*time.Second)

	result := collide.Collide(ctx)
	if result != nil {
		g.Log().Info(ctx, *result)
	} else {
		g.Log().Info(ctx, "no result")
		t.FailNow()
	}
}

func TestRangeFailed(t *testing.T) {
	ctx := context.Background()

	collide := collider.New(func(finished context.Context, seedChan chan uint32) {
		for i := uint32(0); i < 0x7fffff; i++ {
			select {
			case <-finished.Done():
				return
			default:
				seedChan <- i
			}
		}
	}, func(input uint32) bool {
		data := make([]byte, 4)
		binary.LittleEndian.PutUint32(data, input)
		hash := sha256.Sum256(data)
		hashStr := hex.EncodeToString(hash[:])
		return gstr.HasPrefix(hashStr, "000000")
	}, func(speedString string) {
		g.Log().Debug(ctx, speedString)
	}, math.Pow(16, 6), 0, 0xffffff, 1*time.Second)

	result := collide.Collide(ctx)
	if result != nil {
		g.Log().Info(ctx, *result)
		t.FailNow()
	} else {
		g.Log().Info(ctx, "no result")
	}
}
