package main

import (
	"fmt"

	"github.com/amilsil/edge/invoke"
	"github.com/amilsil/edge/validators"
)

func (assertion *BaseAssertion) validate() AssertionResult {
	return AssertionResult{}
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
	validators.ValidateHeaders()
	fmt.Println(body)
}
