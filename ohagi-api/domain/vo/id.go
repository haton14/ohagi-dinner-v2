package vo

import (
	"math/rand"

	"github.com/oklog/ulid/v2"
)

type ID struct {
	value string
}

func NewID(currentTime CurrentTime) ID {
	t := currentTime.Value()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return ID{value: id.String()}
}

func (i ID) Value() string {
	return i.value
}
