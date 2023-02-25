package tcase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests: getters (static)
// -> nil values are not allowed
func TestArgs_Get(t *testing.T) {
	t.Run("Get - not nil", func(t *testing.T) {
		// arrange
		type Item struct {name string}
		item := Item{name: "test"}

		args := NewArgs(item)

		// act
		result := args.Get(0).(Item)
		
		// assert
		assert.Equal(t, result, item)
	})

	t.Run("Get - nil", func(t *testing.T) {
		// arrange
		args := NewArgs(nil)

		// act
		result := args.Get(0)

		// assert
		assert.Nil(t, result)
	})
}

func TestArgs_String(t *testing.T) {
	t.Run("String - not nil", func(t *testing.T) {
		// arrange
		args := NewArgs("test")

		// act
		result := args.String(0)

		// assert
		assert.Equal(t, result, "test")
	})

	t.Run("String - empty", func(t *testing.T) {
		// arrange
		args := NewArgs("")

		// act
		result := args.String(0)

		// assert
		assert.Equal(t, result, "")
	})
}

func TestArgs_Int(t *testing.T) {
	t.Run("Int - not nil", func(t *testing.T) {
		// arrange
		args := NewArgs(1)

		// act
		result := args.Int(0)

		// assert
		assert.Equal(t, result, 1)
	})

	t.Run("Int - zero", func(t *testing.T) {
		// arrange
		args := NewArgs(0)

		// act
		result := args.Int(0)

		// assert
		assert.Equal(t, result, 0)
	})
}

func TestArgs_Float64(t *testing.T) {
	t.Run("Float64 - not nil", func(t *testing.T) {
		// arrange
		args := NewArgs(1.0)

		// act
		result := args.Float64(0)

		// assert
		assert.Equal(t, result, 1.0)
	})

	t.Run("Float64 - zero", func(t *testing.T) {
		// arrange
		args := NewArgs(0.0)

		// act
		result := args.Float64(0)

		// assert
		assert.Equal(t, result, 0.0)
	})
}

func TestArgs_Bool(t *testing.T) {
	t.Run("Bool - true", func(t *testing.T) {
		// arrange
		args := NewArgs(true)

		// act
		result := args.Bool(0)

		// assert
		assert.Equal(t, result, true)
	})

	t.Run("Bool - false", func(t *testing.T) {
		// arrange
		args := NewArgs(false)

		// act
		result := args.Bool(0)

		// assert
		assert.Equal(t, result, false)
	})
}


// Tests: getters (pointer)
// -> nil values are allowed
func TestArgs_Error(t *testing.T) {
	t.Run("Error - not nil", func(t *testing.T) {
		// arrange
		err := errors.New("test error")
		
		args := NewArgs(err)

		// act
		e := args.Error(0)
		
		// assert
		assert.ErrorIs(t, e, err)
	})

	t.Run("Error - nil", func(t *testing.T) {
		// arrange
		args := NewArgs(nil)

		// act
		e := args.Error(0)

		// assert
		assert.Nil(t, e)
		assert.NoError(t, e)
	})
}