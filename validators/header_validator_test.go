package validators

import (
	"net/http"
	"testing"

	"github.com/amilsil/edge/assertions"
	"github.com/stretchr/testify/assert"
)

func TestValidateHeaders(t *testing.T) {
	type args struct {
		headers    http.Header
		assertions []assertions.HeaderAssertion
	}
	tests := []struct {
		name string
		args args
		want []assertions.AssertionResult
	}{
		{
			name: "Must return empty AssertionResult set when no HeaderAssertions given",
			args: args{
				headers:    http.Header{},
				assertions: []assertions.HeaderAssertion{},
			},
			want: []assertions.AssertionResult{},
		},
		{
			name: "Must return a Failing AssertionResult when a matching header is missing",
			args: args{
				headers: http.Header{},
				assertions: []assertions.HeaderAssertion{
					{Name: "Content-Type"},
				},
			},
			want: []assertions.AssertionResult{
				{Result: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateHeaders(tt.args.headers, tt.args.assertions)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("ValidateHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
