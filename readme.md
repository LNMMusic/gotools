# Go Tools

Package that provides a set of tools for Go.

<!-- Details -->
## ✅ T-Case
Its a module that provides a tool for testing table cases
> Args: slice of any type of values

> TestConfig: base config for test cases
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Name: name of test case
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Params: params for test case
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Expect: expected result for test

**EXAMPLE**

Controller: division function
```go
var (
    ErrDivideByZero = errors.New("error: divide by zero")
)

func Div(a, b float64) (result float64, err error) {
    if b == 0 {
        err = ErrDivideByZero
        return
    }

    result = a / b
    return
}
```

Test (table test)
```go
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
```

## ✅ Error Layer
Its a module that provides a tool for handling errors
> ErrorLayer: struct that contains a message and an error representing the error interface. It also contains a stack trace of the error.

> Wrap: allows to create a new ErrorLayer instance with an error wrapped in it.

> Wrapf: allows to create a new ErrorLayer instance with an error wrapped in it using a formatted message string.

> Unwrap: allows to get the error wrapped in the ErrorLayer instance.

> Target: returns a bool indicating if the error is included in the chain of errors.