package pkg

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQ1(t *testing.T) {
	type args struct {
		method string
		path   string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "when path is not /users",
			args: args{
				method: http.MethodGet,
				path:   "/",
			},
			want: http.StatusNotFound,
		},
		{
			name: "when method is not GET",
			args: args{
				method: http.MethodPost,
				path:   "/users",
			},
			want: http.StatusMethodNotAllowed,
		},
		{
			name: "happy case",
			args: args{
				method: http.MethodGet,
				path:   "/users",
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(
				tt.args.method,
				tt.args.path,
				bytes.NewReader([]byte{}),
			)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			HandleRequest(rec, req)

			if rec.Result().StatusCode != tt.want {
				t.Errorf("expected status: %d, got: %d", tt.want, rec.Result().StatusCode)
			}
		})
	}
}
