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
		{
			name:    "3 + 4",
			exp:     "3 4 +",
			want:    7,
			wantErr: false,
		},
		{
			name:    "3 * 4",
			exp:     "3 4 *",
			want:    12,
			wantErr: false,
		},
		{
			name:    "(1 + 2) * (3 + 4)",
			exp:     "1 2 + 3 4 + *",
			want:    21,
			wantErr: false,
		},
		{
			name:    "invalid statement: 1 2",
			exp:     "1 2",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid statement: 2 +",
			exp:     "2 +",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid statement: 2 *",
			exp:     "2 *",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid statement: 2 /",
			exp:     "2 /",
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid opecode: 1 2 mod",
			exp:     " 1 2 mod",
			want:    0,
			wantErr: true,
		},
		{
			name:    "(1 + 2) * (3 + 4) / 10",
			exp:     "1 2 + 3 4 + * 10 /",
			want:    2,
			wantErr: false,
		},
		// TODO 実装してみよう
		{
			name:    "10 - (1 + 2) * (3 + 4)",
			exp:     "10 1 2 + 3 4 + * -",
			want:    -11,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
