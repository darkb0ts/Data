package array

import (
	"reflect"
	"testing"
)

func TestLSearch(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		target  int
		want    int
		wantErr bool
	}{
		{
			name:    "Target found",
			arr:     []int{1, 2, 3, 4, 5},
			target:  3,
			want:    2,
			wantErr: false,
		},
		{
			name:    "Target not found",
			arr:     []int{1, 2, 3, 4, 5},
			target:  6,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Empty array",
			arr:     []int{},
			target:  3,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Target at beginning",
			arr:     []int{1, 2, 3, 4, 5},
			target:  1,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Target at end",
			arr:     []int{1, 2, 3, 4, 5},
			target:  5,
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LSearch(tt.arr, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("LSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSearch(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		target  int
		want    int
		wantErr bool
	}{
		{
			name:    "Target found",
			arr:     []int{1, 2, 3, 4, 5},
			target:  3,
			want:    2,
			wantErr: false,
		},
		{
			name:    "Target not found",
			arr:     []int{1, 2, 3, 4, 5},
			target:  6,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Empty array",
			arr:     []int{},
			target:  3,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Target at beginning",
			arr:     []int{1, 2, 3, 4, 5},
			target:  1,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Target at end",
			arr:     []int{1, 2, 3, 4, 5},
			target:  5,
			want:    4,
			wantErr: false,
		},
		{
			name:    "Target in middle",
			arr:     []int{1, 3, 5, 7, 9, 11, 13},
			target:  7,
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BSearch(tt.arr, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("BSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		want    []int
		wantErr bool
	}{
		{
			name:    "Normal case",
			arr:     []int{1, 2, 3, 4, 5},
			want:    []int{5, 4, 3, 2, 1},
			wantErr: false,
		},
		{
			name:    "Empty array",
			arr:     []int{},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "Single element",
			arr:     []int{1},
			want:    []int{1},
			wantErr: false,
		},
		{
			name:    "Even number of elements",
			arr:     []int{1, 2, 3, 4},
			want:    []int{4, 3, 2, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reverse(tt.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		min         int
		max         int
		noRepeat    []bool
		wantErr     bool
		checkUnique bool
	}{
		{
			name:        "Normal case",
			size:        5,
			min:         1,
			max:         10,
			wantErr:     false,
			checkUnique: false,
		},
		{
			name:        "Size zero",
			size:        0,
			min:         1,
			max:         10,
			wantErr:     true,
			checkUnique: false,
		},
		{
			name:        "Min greater than max",
			size:        5,
			min:         10,
			max:         1,
			wantErr:     true,
			checkUnique: false,
		},
		{
			name:        "Unique values",
			size:        5,
			min:         1,
			max:         10,
			noRepeat:    []bool{true},
			wantErr:     false,
			checkUnique: true,
		},
		{
			name:        "Unique values - size too large",
			size:        11,
			min:         1,
			max:         10,
			noRepeat:    []bool{true},
			wantErr:     true,
			checkUnique: false,
		},
		{
			name:        "Negative min and max",
			size:        5,
			min:         -10,
			max:         -1,
			wantErr:     false,
			checkUnique: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Random(tt.size, tt.min, tt.max, tt.noRepeat...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Random() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != tt.size {
					t.Errorf("Random() size = %v, want %v", len(got), tt.size)
				}
				for _, v := range got {
					if v < tt.min || v > tt.max {
						t.Errorf("Random() value %v out of range [%v, %v]", v, tt.min, tt.max)
					}
				}
				if tt.checkUnique {
					unique := make(map[int]bool)
					for _, v := range got {
						if unique[v] {
							t.Errorf("Random() generated duplicate value: %v", v)
						}
						unique[v] = true
					}
				}
			}
		})
	}
}