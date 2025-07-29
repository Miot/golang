package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:4", false},
		{"", false},
		{"11:22", false},
		{"aa:vv", false},
		{"5:22:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("ParseTime(%q) = %v, want nil", data.time, err)
		} else if !data.ok && err == nil {
			t.Errorf("ParseTime(%q) = nil, want error", data.time)
		}
	}
}
