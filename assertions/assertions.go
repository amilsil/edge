package assertions

// Contains base props of any assertion
type BaseAssertion struct {
	Name string
}

// TODO: What would a header assertion contain, apart from a base?
// TODO: How to inherit, create an interface?
type HeaderAssertion struct {
	Name string
}

// RESULTS -----------------------------------------

// Contains the result of an assertion,
// along with any errors and message logs
// upon execution of them
type AssertionResult struct {
	Result bool
}
