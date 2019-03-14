package invoke

import (
	"bytes"
	"net/http"
)

// Invoke an Http call
// returns the invocation response.
func InvokeHttpCall(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read the response body into a string
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	bodyAsString := buf.String()

	return bodyAsString, nil
}
