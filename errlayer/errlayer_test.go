package errlayer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests
func TestErrorLayer_Wrap(t *testing.T) {
	// arrange
	err := New("error base")

	// act
	t.Run("Error Unlayer", func(t *testing.T) {
		// act
		msg := err.Error()

		// assert
		assert.Equal(t, "error base", msg)
	})

	t.Run("Error Layer", func(t *testing.T) {
		// arrange
		err := Wrap(err, "error wrap")

		// act
		msg := err.Error()

		// assert
		assert.Equal(t, "error wrap. error base", msg)
	})

	t.Run("Error Layer with format", func(t *testing.T) {
		// arrange
		err := Wrapf(err, "error wrap %s", "format")

		// act
		msg := err.Error()

		// assert
		assert.Equal(t, "error wrap format. error base", msg)
	})
}

func TestErrorLayer_Unwrap(t *testing.T) {
	t.Run("Unwrap error chain", func(t *testing.T) {
		// arrange
		err 	   := New("error base")
		wrapFirst  := Wrap(err, "error wrap")
		wrapSecond := Wrap(wrapFirst, "error wrap")
	
		// act
		unwrapFirst := Unwrap(wrapSecond)
		unwrapSecond:= Unwrap(unwrapFirst)
		unwrapBase  := Unwrap(unwrapSecond)
	
		// assert
		assert.EqualError(t, unwrapFirst, "error wrap. error base")
		assert.Equal(t, wrapFirst, unwrapFirst)
		assert.ErrorIs(t, unwrapFirst, wrapFirst)
	
		assert.EqualError(t, unwrapSecond, "error base")
		assert.Equal(t, err, unwrapSecond)
		assert.ErrorIs(t, unwrapSecond, err)
		
		assert.Nil(t, unwrapBase)
	})

	t.Run("Unwrap wrong type error", func(t *testing.T) {
		// arrange
		err := errors.New("external error")

		// act
		unwrap := Unwrap(err)

		// assert
		assert.Nil(t, unwrap)
	})

	t.Run("Unwrap external error", func(t *testing.T) {
		// arrange
		err := New("error base")
		wrapFirst := fmt.Errorf("%s. %w", "external error wrap", err)
		wrapSecond:= Wrap(wrapFirst, "last error wrap")

		// act
		unwrapFirst := Unwrap(wrapSecond)
		unwrapSecond:= Unwrap(unwrapFirst)

		// assert
		assert.EqualError(t, unwrapFirst, "external error wrap. error base")
		assert.Equal(t, wrapFirst, wrapFirst)
		assert.ErrorIs(t, unwrapFirst, wrapFirst)

		assert.Nil(t, unwrapSecond)
	})
}

func TestErrorLayer_Target(t *testing.T) {
	// arrange
	t.Run("Target is nil", func(t *testing.T) {
		// arrange
		err := New("error base")

		// act
		target := Target(err, nil)

		// assert
		assert.False(t, target)
	})

	t.Run("Error is nil", func(t *testing.T) {
		// act
		target := Target(nil, New("error base"))

		// assert
		assert.False(t, target)
	})

	t.Run("Target is not in chain", func(t *testing.T) {
		// arrange
		err := New("error base")

		// act
		target := Target(err, New("error target"))

		// assert
		assert.False(t, target)
	})

	t.Run("Target is in chain", func(t *testing.T) {
		// arrange
		err 	   := New("error base")
		wrapFirst  := Wrap(err, "error wrap")
		wrapSecond := Wrap(wrapFirst, "error wrap")

		// act
		t.Run("Target is second wrap", func(t *testing.T) {
			// act
			target := Target(wrapSecond, wrapSecond)

			// assert
			assert.True(t, target)
		})
		
		t.Run("Target is first wrap", func(t *testing.T) {
			// act
			target := Target(wrapSecond, wrapFirst)
			
			// assert
			assert.True(t, target)
		})
		
		t.Run("Target is base", func(t *testing.T) {
			// act
			target := Target(wrapSecond, err)

			// assert
			assert.True(t, target)
		})
	})

	t.Run("Target is in chain, middle err is external", func(t *testing.T) {
		// arrange
		err 	   := New("error base")
		wrapFirst  := fmt.Errorf("external error wrap: %w", err)
		wrapSecond := Wrap(wrapFirst, "error wrap second")

		t.Run("Target is second wrap", func(t *testing.T) {
			// act
			target := Target(wrapSecond, wrapSecond)

			// assert
			assert.True(t, target)
		})
		
		t.Run("Target is first wrap", func(t *testing.T) {
			// act
			target := Target(wrapSecond, wrapFirst)
			
			// assert
			assert.True(t, target)
		})
		
		t.Run("Target is base", func(t *testing.T) {
			// act
			target := Target(wrapSecond, err)

			// assert
			assert.False(t, target)
		})
	})
}