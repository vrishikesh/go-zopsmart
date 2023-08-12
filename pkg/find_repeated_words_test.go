package pkg

import (
	"reflect"
	"testing"
)

func TestFindRepeatedWords(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `happy case`,
			args: args{
				input: `lorem ipsum dolor sit amet consectetur adipiscing lorem sit amet`,
			},
			want: []string{"lorem", "sit", "amet"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatedWords(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRepeatedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
