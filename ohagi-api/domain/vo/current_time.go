package vo

import (
	"context"
	"time"
)

type contextTimeKey string

const TimeKey contextTimeKey = "timeKey"

type CurrentTime struct {
	value time.Time
}

func NewCurrentTime(ctx context.Context) CurrentTime {
	v := ctx.Value(TimeKey)
	if t, ok := v.(time.Time); ok {
		return CurrentTime{value: t}
	}
	return CurrentTime{value: time.Now()}
}

func (t CurrentTime) Value() time.Time {
	return t.value
}
