package main

import (
	"fmt"

	"github.com/amilsil/edge/invoke"
)

// Contains base props of any assertion
type BaseAssertion struct {
	name string
}

// Contains the result of an assertion,
// along with any errors and message logs
// upon execution of them
type AssertionResult struct {
	result bool
}

func (assertion *BaseAssertion) validate() AssertionResult {
	return AssertionResult{}
}

// TODO: What would a header assertion contain, apart from a base?
type HeaderAssertion struct {
	BaseAssertion
}

// Validates a collection of assertions,
// by proxing to the workers,
// returns a collection of AssertionResults.
func validate(assertions []BaseAssertion) []AssertionResult {
	var results []AssertionResult
	// TODO: validate
	return results
}

const url string = "http://localhost:8080/ping"

func main() {
	body, _ := invoke.InvokeHttpCall(url)
	fmt.Println(body)
}
