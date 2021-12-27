package main

import "testing"

func TestDay11a(t *testing.T) {
	for _, pass := range []string{"abcdffaa", "ghjaabcc"} {
		t.Run("", func(t *testing.T) {
			if !passvalid(pass) {
				t.Fatalf("password should be valid: %q", pass)
			}
		})
	}
	for _, pass := range []string{"hijklmmn", "abbceffg", "abbcegjk"} {
		t.Run("", func(t *testing.T) {
			if passvalid(pass) {
				t.Fatalf("password should be invalid: %q", pass)
			}
		})
	}
}

func TestDay11a2(t *testing.T) {
	tt := []struct {
		input  string
		expect string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
		{"vzbxkghb", "vzbxxyzz"},
		{"vzbxxyzz", "vzcaabcc"},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			result := incpass(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %q, got %q", tc.expect, result)
			}
		})
	}
}
