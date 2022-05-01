package vo_test

import (
	"context"
	"testing"

	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
)

func TestID(t *testing.T) {
	t.Run("同一時刻のIDは等しい", func(t *testing.T) {
		currentTime := vo.NewCurrentTime(context.Background())
		id := vo.NewID(currentTime)
		id2 := vo.NewID(currentTime)
		if id.Value() != id2.Value() {
			t.Errorf("should be equal")
		}
	})

	t.Run("同一時刻のIDは等しい", func(t *testing.T) {
		currentTime := vo.NewCurrentTime(context.Background())
		currentTime2 := vo.NewCurrentTime(context.Background())
		currentTime3 := vo.NewCurrentTime(context.Background())
		id := vo.NewID(currentTime)
		id2 := vo.NewID(currentTime2)
		id3 := vo.NewID(currentTime3)
		if id.Value() < id2.Value() {
			t.Errorf("should be id < id2")
		}
		if id2.Value() < id3.Value() {
			t.Errorf("should be id2 < id3")
		}
	})

}
