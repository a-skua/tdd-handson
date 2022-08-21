package calculator

import (
	"testing"
)

func TestCalc(t *testing.T) {
	type test struct {
		name    string
		exp     string
		want    int
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calc(tt.exp)
			if tt.wantErr != (err != nil) {
				t.Errorf("want-error=%t, error=%v", tt.wantErr, err)
			}

			if result != tt.want {
				t.Errorf("want=%d, got=%d", tt.want, result)
			}
		})
	}

	tests := []*test{
		{
			name:    "1 + 2",
			exp:     "1 2 +",
			want:    3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
