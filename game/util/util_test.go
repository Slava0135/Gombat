package util

import "testing"

func TestRayTrace(t *testing.T) {
	type args struct {
		grid [][]bool
		from Position
		to   Position
	}
	tests := []struct {
		name   string
		args   args
		wantX  int
		wantY  int
		wantOk bool
	}{
		{
			name: "straight horizontal",
			args: args{
				[][]bool{
					{false, false, true, true, false},
				},
				Position{0.5, 0.5},
				Position{0.5, 4.5},
			},
			wantX:  0,
			wantY:  2,
			wantOk: true,
		},
		{
			name: "straight vertical",
			args: args{
				[][]bool{
					{false},
					{false},
					{false},
					{true},
					{true},
				},
				Position{0.5, 0.5},
				Position{4.5, 0.5},
			},
			wantX:  3,
			wantY:  0,
			wantOk: true,
		},
		{
			name: "3 by 3",
			args: args{
				[][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
				Position{0, 0},
				Position{2.5, 2.5},
			},
			wantX:  1,
			wantY:  1,
			wantOk: true,
		},
		{
			name: "3 by 3 (2)",
			args: args{
				[][]bool{
					{false, false, false},
					{true, true, false},
					{false, false, false},
				},
				Position{0, 0},
				Position{2.5, 1.5},
			},
			wantX:  1,
			wantY:  0,
			wantOk: true,
		},
		{
			name: "5 by 5",
			args: args{
				[][]bool{
					{false, false, false, false, false},
					{false, false, false, false, false},
					{false, true, true, false, false},
					{false, false, true, false, false},
					{false, false, false, false, false},
				},
				Position{0.5, 4.5},
				Position{4.5, 1.5},
			},
			wantX:  2,
			wantY:  2,
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotOk := RayTrace(tt.args.grid, tt.args.from, tt.args.to)
			if gotX != tt.wantX {
				t.Errorf("RayTrace() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("RayTrace() gotY = %v, want %v", gotY, tt.wantY)
			}
			if gotOk != tt.wantOk {
				t.Errorf("RayTrace() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
