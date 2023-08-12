package pkg

import (
	"net/http"
	"testing"
)

func TestPingSite(t *testing.T) {
	type args struct {
		site string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy case",
			args: args{
				"https://meet.google.com",
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PingSite(tt.args.site); got != tt.want {
				t.Errorf("PingSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
