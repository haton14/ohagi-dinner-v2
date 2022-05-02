package vo

type Role struct {
	value RoleType
}

type RoleType string

const OwenrRole, WritableRole, ReadonlyRole RoleType = "owner", "writable", "readonly"

func NewRole(value RoleType) (Role, error) {
	if value != OwenrRole && value != WritableRole && value != ReadonlyRole {
		return Role{}, ErrEnum
	}
	return Role{value: value}, nil
}

func (r Role) Value() RoleType {
	return r.value
}

func (r Role) IsZero() bool {
	return r == Role{}
}
