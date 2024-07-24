package ascii

import "testing"

func TestMaps(t *testing.T) {
	tests := []struct {
		name   string
		args   int
		banner string
		want   string
	}{
		{"T", 470, "standard", " _______  "},
		{"X", 509, "standard", "  > <   "},
		{"[", 540, "standard", "      "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maps(tt.args, tt.banner); got != tt.want {
				t.Errorf("Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}
