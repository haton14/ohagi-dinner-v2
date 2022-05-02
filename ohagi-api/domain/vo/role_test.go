package vo_test

import (
	"testing"

	"github.com/haton14/ohagi-dinner/ohagi-api/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	t.Run("owner", func(t *testing.T) {
		role, err := vo.NewRole(vo.OwenrRole)
		assert.NoError(t, err)
		assert.False(t, role.IsZero())
		assert.Equal(t, vo.OwenrRole, role.Value())
	})

	t.Run("writable", func(t *testing.T) {
		role, err := vo.NewRole(vo.WritableRole)
		assert.NoError(t, err)
		assert.False(t, role.IsZero())
		assert.Equal(t, vo.WritableRole, role.Value())
	})

	t.Run("readonly", func(t *testing.T) {
		role, err := vo.NewRole(vo.ReadonlyRole)
		assert.NoError(t, err)
		assert.False(t, role.IsZero())
		assert.Equal(t, vo.ReadonlyRole, role.Value())
	})

	t.Run("invalid", func(t *testing.T) {
		role, err := vo.NewRole("invalid")
		assert.ErrorIs(t, err, vo.ErrEnum)
		assert.True(t, role.IsZero())
	})
}
