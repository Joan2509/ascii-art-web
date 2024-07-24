package ascii

import (
	"testing"
)

func TestFindHeadLine(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"%", '%', 47},
		{"V", 'V', 488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindHeadLine(tt.r); got != tt.want {
				t.Errorf("FindHeadLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine2(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want int
	}{
		{"%", '%', 48},
		{"V", 'V', 489},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line2(tt.r); got != tt.want {
				t.Errorf("Line2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine3(t *testing.T) {
	tests := []struct {
		name string
		n    rune
		want int
	}{
		{"%", '%', 49},
		{"V", 'V', 490},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line3(tt.n); got != tt.want {
				t.Errorf("Line3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine4(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want int
	}{
		{"%", '%', 50},
		{"V", 'V', 491},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line4(tt.args); got != tt.want {
				t.Errorf("Line4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine5(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want int
	}{
		{"%", '%', 51},
		{"V", 'V', 492},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line5(tt.args); got != tt.want {
				t.Errorf("Line5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine6(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want int
	}{
		{"%", '%', 52},
		{"V", 'V', 493},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line6(tt.args); got != tt.want {
				t.Errorf("Line6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine7(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want int
	}{
		{"%", '%', 53},
		{"V", 'V', 494},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line7(tt.args); got != tt.want {
				t.Errorf("Line7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine8(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want int
	}{
		{"%", '%', 54},
		{"V", 'V', 495},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Line8(tt.args); got != tt.want {
				t.Errorf("Line8() = %v, want %v", got, tt.want)
			}
		})
	}
}
