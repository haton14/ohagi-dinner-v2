package vo

type Role struct {
	value role
}

type role string

const OwenrRole, WritableRole, ReadonlyRole role = "owner", "writable", "readonly"

func NewRole(value role) (Role, error) {
	if value != OwenrRole && value != WritableRole && value != ReadonlyRole {
		return Role{}, ErrEnum
	}
	return Role{value: value}, nil
}

func (r Role) Value() role {
	return r.value
}

func (r Role) IsZero() bool {
	return r == Role{}
}
