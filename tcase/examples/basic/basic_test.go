package basic

import (
	"testing"

	"github.com/LNMMusic/gotools/tcase"
	"github.com/stretchr/testify/assert"
)

// Tests
func TestSum(t *testing.T) {
	// arrange (table test)
	type testcases struct {
		*tcase.TestConfig
	}
	cases := []testcases{
		{TestConfig: &tcase.TestConfig{
			Name: "Sum(1, 2) = 3",
			Params: tcase.NewArgs(1, 2),
			Expect: tcase.NewArgs(3),
		}},
		{TestConfig: &tcase.TestConfig{
			Name: "Sum(2, 2) = 4",
			Params: tcase.NewArgs(2, 2),
			Expect: tcase.NewArgs(4),
		}},
	}

	// act
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			// act
			result := Sum(c.Params.Int(0), c.Params.Int(1))

			// assert
			assert.Equal(t, c.Expect.Int(0), result)
		})
	}
}

func TestDiv(t *testing.T) {
	// arrange (table test)
	type testcases struct {
		*tcase.TestConfig
	}
	cases := []testcases{
		{TestConfig: &tcase.TestConfig{
			Name: "Div(1, 2) = 0.5",
			Params: tcase.NewArgs(1.0, 2.0),
			Expect: tcase.NewArgs(0.5, nil),
		}},
		{TestConfig: &tcase.TestConfig{
			Name: "Div(2, 2) = 1",
			Params: tcase.NewArgs(2.0, 2.0),
			Expect: tcase.NewArgs(1.0, nil),
		}},
		{TestConfig: &tcase.TestConfig{
			Name: "Div(1, 0) = error",
			Params: tcase.NewArgs(1.0, 0.0),
			Expect: tcase.NewArgs(0.0, ErrDivideByZero),
		}},
	}

	// act
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			// act
			result, err := Div(c.Params.Float64(0), c.Params.Float64(1))

			// assert
			assert.Equal(t, c.Expect.Float64(0), result)
			assert.Equal(t, c.Expect.Error(1), err)
		})
	}
}