package usecase

import (
	"context"

	"github.com/haton14/ohagi-dinner/ohagi-api/domain"
	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
)

type Auth interface {
	GenarateToken(req LoginRequest, res LoginResponse) error
}

func NewAuth(petTokenMaker domain.PetTokenMaker) Auth {
	return login{
		petTokenMaker: petTokenMaker,
	}
}

type login struct {
	petTokenMaker domain.PetTokenMaker
}

func (l login) GenarateToken(req LoginRequest, res LoginResponse) error {
	role, _ := vo.NewRole(vo.OwenrRole)
	petToken, _ := l.petTokenMaker.Create(vo.NewPetID("tmpID"), role, vo.NewCurrentTime(context.Background()))
	return res.RenderLoginResult(petToken.Value())
}

type LoginRequest struct {
	UserID string
}

type LoginResponse interface {
	RenderLoginResult(petsToken string) error
	RenderFailure(err error) error
}
