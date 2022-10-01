package lifegame

import (
	"reflect"
	"testing"
)

func TestCell_current(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.current(tt.x, tt.y)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "first",
			cell: NewCell(3, 3, []State{
				true, false, false,
				false, false, false,
				false, false, false,
			}),
			args: args{
				x: 0,
				y: 0,
			},
			want: true,
		},
		{
			name: "last",
			cell: NewCell(3, 3, []State{
				false, false, false,
				false, false, false,
				false, false, true,
			}),
			args: args{
				x: 2,
				y: 2,
			},
			want: true,
		},
		{
			name: "any",
			cell: NewCell(3, 3, []State{
				false, false, false,
				false, true, false,
				false, false, false,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: true,
		},
		{
			name: "outrange x",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, true,
				true, true, true,
			}),
			args: args{
				x: -1,
				y: 0,
			},
			want: false,
		},
		{
			name: "outrange x",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, true,
				true, true, true,
			}),
			args: args{
				x: 3,
				y: 0,
			},
			want: false,
		},
		{
			name: "outrange y",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, true,
				true, true, true,
			}),
			args: args{
				x: 0,
				y: -1,
			},
			want: false,
		},
		{
			name: "outrange y",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, true,
				true, true, true,
			}),
			args: args{
				x: 0,
				y: 3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCell_future(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.future(tt.x, tt.y)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "die to live",
			cell: NewCell(3, 3, []State{
				true, true, false,
				true, false, false,
				false, false, false,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: true,
		},
		{
			name: "live to live",
			cell: NewCell(3, 3, []State{
				true, true, false,
				true, false, false,
				false, false, false,
			}),
			args: args{
				x: 0,
				y: 1,
			},
			want: true,
		},
		{
			name: "live to live",
			cell: NewCell(3, 3, []State{
				true, true, false,
				true, true, false,
				false, false, false,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: true,
		},
		{
			name: "live to die",
			cell: NewCell(3, 3, []State{
				false, false, false,
				true, true, false,
				false, false, false,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: false,
		},
		{
			name: "live to die",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, false,
				false, false, false,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCell_resize(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want *Cell
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.resize(tt.x, tt.y)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "to large",
			cell: NewCell(3, 3, []State{
				true, true, true,
				true, true, true,
				true, true, true,
			}),
			args: args{
				x: 5,
				y: 6,
			},
			want: NewCell(5, 6, []State{
				true, true, true, false, false,
				true, true, true, false, false,
				true, true, true, false, false,
				false, false, false, false, false,
				false, false, false, false, false,
				false, false, false, false, false,
			}),
		},
		{
			name: "to small",
			cell: NewCell(5, 5, []State{
				false, false, false, true, true,
				false, false, false, true, true,
				true, true, true, true, true,
				true, true, true, true, true,
				true, true, true, true, true,
			}),
			args: args{
				x: 3,
				y: 2,
			},
			want: NewCell(3, 2, []State{
				false, false, false,
				false, false, false,
			}),
		},
		{
			name: "to same",
			cell: NewCell(3, 2, []State{
				false, true, false,
				true, false, true,
			}),
			args: args{
				x: 3,
				y: 2,
			},
			want: NewCell(3, 2, []State{
				false, true, false,
				true, false, true,
			}),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestLifegame_Next(t *testing.T) {
	lifegame := New(NewCell(3, 3, []State{
		false, true, false,
		false, true, false,
		false, true, false,
	}))

	lifegame.Next()
	want := New(NewCell(3, 3, []State{
		false, false, false,
		true, true, true,
		false, false, false,
	}))
	if !reflect.DeepEqual(want, lifegame) {
		t.Fatalf("want=%v, got=%v.", want, lifegame)
	}

	lifegame.Next()
	want = New(NewCell(3, 3, []State{
		false, true, false,
		false, true, false,
		false, true, false,
	}))
	if !reflect.DeepEqual(want, lifegame) {
		t.Fatalf("want=%v, got=%v.", want, lifegame)
	}
}

func TestLifegame_Table(t *testing.T) {
	type test struct {
		name     string
		lifegame *Lifegame
		want     [][]State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lifegame.Table()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "ok",
			lifegame: New(NewCell(3, 3, []State{
				false, true, true,
				false, true, true,
				false, true, true,
			})),
			want: [][]State{
				{false, true, true},
				{false, true, true},
				{false, true, true},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
