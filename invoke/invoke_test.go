package invoke

import "testing"

func TestInvokeHttpCall(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "call mock api",
			args: args{
				url: "http://localhost:8080/ping",
			},
			want:    "Hello ping",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InvokeHttpCall(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvokeHttpCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InvokeHttpCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
