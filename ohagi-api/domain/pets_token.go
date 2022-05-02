package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
)

type PetToken struct {
	value string
}

type PetTokenMaker interface {
	Create(vo.PetID, vo.Role, vo.CurrentTime) (PetToken, error)
	Verify(token string) (vo.PetID, vo.Role, error)
}

func NewPetTokenMaker(secretKey, publicKey string) PetTokenMaker {
	return petTokenKey{
		secretKey: secretKey,
		publicKey: publicKey,
	}
}

type petTokenKey struct {
	secretKey string
	publicKey string
}

func (ptk petTokenKey) Create(petID vo.PetID, role vo.Role, currentTime vo.CurrentTime) (PetToken, error) {
	t := currentTime.Value()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    "ohagi-api",
		"nbf":    t.Unix(),
		"exp":    t.Add(time.Minute * 10).Unix(),
		"pet_id": petID.Value(),
		"role":   role.Value(),
	})

	tokenString, err := token.SignedString(ptk.secretKey)
	if err != nil {
		return PetToken{}, err
	}
	return PetToken{value: tokenString}, nil
}

func (p PetToken) Value() string {
	return p.value
}

func (p PetToken) IsZero() bool {
	return p == PetToken{}
}

func (ptk petTokenKey) Verify(tokenString string) (vo.PetID, vo.Role, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(ptk.publicKey), nil
	})
	if err != nil {
		return vo.PetID{}, vo.Role{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token claims is not mapclaims")
	}
	nbf, ok := claims["nbf"].(int64)
	if !ok {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token claims nbf is not int64")
	}
	exp, ok := claims["exp"].(int64)
	if !ok {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token claims exp is not int64")
	}
	now := time.Now().Unix()
	if now < nbf {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token is not valid yet")
	}
	if now > exp {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token is expired")
	}

	petID := vo.NewPetID(claims["pet_id"].(string))
	role, err := vo.NewRole(claims["role"].(vo.RoleType))
	if err != nil {
		return vo.PetID{}, vo.Role{}, fmt.Errorf("token claims is invalid role")
	}
	return petID, role, nil
}
