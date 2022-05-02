package vo

import (
	"math/rand"

	"github.com/oklog/ulid/v2"
)

type PetID struct {
	value string
}

func GeneratePetID(currentTime CurrentTime) PetID {
	return NewPetID(generateID(currentTime))
}

func NewPetID(id string) PetID {
	return PetID{value: id}
}

func (p PetID) Value() string {
	return p.value
}

func generateID(currentTime CurrentTime) string {
	t := currentTime.Value()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
