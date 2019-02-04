package exam

import "testing"

// TestExam is testing function.
func TestExam(t *testing.T) {
	for _, tc := range testCases {
		if observed := Exam(tc.A, tc.B); observed != tc.Expected {
			t.Fatalf("HelloWorld() = %v, want %v", observed, tc.Expected)
		}
	}
}
