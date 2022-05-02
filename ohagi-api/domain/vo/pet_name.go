package vo

import "unicode/utf8"

type PetName struct {
	value string
}

func NewPetName(value string) (PetName, error) {
	if utf8.RuneCountInString(value) > 20 {
		return PetName{}, ErrMaxLength
	}
	if value == "" {
		return PetName{}, ErrEmpty
	}
	return PetName{value: value}, nil
}

func (p PetName) Value() string {
	return p.value
}

func (p PetName) IsZero() bool {
	return p == PetName{}
}
