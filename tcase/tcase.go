package tcase

// Args: constructor
func NewArgs(args ...interface{}) *Args {
	return &Args{args: args}
}

// Args: abstraction of an slice of arguments of type interface{}
//  - getters (static): not allowed nil values
//  - getters (pointer): allowed nil values (return nil with associated type)
type Args struct {
	args []interface{}
}
// getters (static) with type assertion
func (a *Args) Get(index int) interface{} {
	return a.args[index]
}
func (a *Args) String(index int) string {
	return a.args[index].(string)
}
func (a *Args) Int(index int) int {
	return a.args[index].(int)
}
func (a *Args) Float64(index int) float64 {
	return a.args[index].(float64)
}
func (a *Args) Bool(index int) bool {
	return a.args[index].(bool)
}
// getters (pointer) with type assertion
// -> in this case nil value is valid for an interface of type error but its not associated with the type
// -> by default is returned in this case an nil error, rather the direct type assertion
func (a *Args) Error(index int) error {
	var e error
	if a.args[index] == nil {
		return e
	}

	return a.args[index].(error)
}


// TestConfig: abstraction of a test configuration
//  - Name: name of the test
//  - Params: slice of arguments of type interface{}
//  - Expect: slice of arguments of type interface{}
type TestConfig struct {
	Name   string
	Params *Args
	Expect *Args
}