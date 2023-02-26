# Go Tools

Package that provides a set of tools for Go.

<!-- Details -->
## T-Case
Its a module that provides a tool for testing table cases
> Args: slice of any type of values

> TestConfig: base config for test cases
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Name: name of test case
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Params: params for test case
>
> &nbsp;&nbsp;&nbsp;&nbsp;- Expect: expected result for test 

## Error Layer
Its a module that provides a tool for handling errors
> ErrorLayer: struct that contains a message and an error representing the error interface. It also contains a stack trace of the error.

> Wrap: allows to create a new ErrorLayer instance with an error wrapped in it.

> Wrapf: allows to create a new ErrorLayer instance with an error wrapped in it using a formatted message string.

> Unwrap: allows to get the error wrapped in the ErrorLayer instance.

> Target: returns a bool indicating if the error is included in the chain of errors.