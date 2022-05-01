package vo_test

import (
	"context"
	"testing"
	"time"

	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
)

func TestCurrentTime(t *testing.T) {
	beforeTime := time.Now()
	currentTime := vo.NewCurrentTime(context.Background())
	afterTime := time.Now()
	if currentTime.Value().Before(beforeTime) || currentTime.Value().After(afterTime) {
		t.Errorf("CurrentTime.Value() is invalid")
	}

	now := time.Now()
	ctx := context.WithValue(context.Background(), vo.TimeKey, now)
	currentTime = vo.NewCurrentTime(ctx)
	if currentTime.Value() != now {
		t.Errorf("CurrentTime.Value() is invalid")
	}
}
