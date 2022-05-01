package vo_test

import (
	"context"
	"testing"
	"time"

	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
)

func TestCurrentTime(t *testing.T) {
	t.Run("現在時刻を順に生成する", func(t *testing.T) {
		beforeTime := time.Now()
		currentTime := vo.NewCurrentTime(context.Background())
		afterTime := time.Now()
		if currentTime.Value().Before(beforeTime) || currentTime.Value().After(afterTime) {
			t.Errorf("CurrentTime.Value() is invalid")
		}
	})

	t.Run("時間を指定して生成できる", func(t *testing.T) {
		now := time.Now()
		ctx := context.WithValue(context.Background(), vo.TimeKey, now)
		currentTime := vo.NewCurrentTime(ctx)
		if currentTime.Value() != now {
			t.Errorf("CurrentTime.Value() is invalid")
		}
	})
}
