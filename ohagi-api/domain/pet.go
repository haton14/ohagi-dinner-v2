package domain

import "github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"

type Pet struct {
	id   vo.PetID
	name vo.PetName
}

func NewPet(id vo.PetID, name vo.PetName) Pet {
	return Pet{
		id:   id,
		name: name,
	}
}

func (p Pet) ID() vo.PetID {
	return p.id
}

func (p Pet) Name() vo.PetName {
	return p.name
}

func (p Pet) IsEmpty() bool {
	return p == Pet{}
}
