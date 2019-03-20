package validators

import (
	"fmt"
	"net/http"

	"github.com/amilsil/edge/assertions"
)

// Validates a collection of http headers against given
// assertions. Returns a collection of AssertionResults.
func ValidateHeaders(
	headers http.Header,
	ha []assertions.HeaderAssertion) []assertions.AssertionResult {

	var results = make([]assertions.AssertionResult, 0)

	fmt.Println(len(ha))

	// Validate headers per each assertion.
	for _, assertion := range ha {
		var result = true
		relevantHeader := headers.Get(assertion.Name)
		if relevantHeader == "" {
			result = false
		}
		results = append(results, assertions.AssertionResult{result})
	}

	return results
}
