package s3

import (
	"testing"
)

// TestListNBuckets tests ListNBuckets(), this needs to be further improved so we can
func TestListNBuckets(t *testing.T) {
	tests := []struct {
		expect  string
		buckets int
	}{
		{
			expect:  success,
			buckets: 1,
		},
		{
			expect:  failure,
			buckets: -1,
		},
	}

	for _, test := range tests {
		got := ListNBuckets(test.buckets)
		if got != test.expect {
			t.Errorf("Expected %v, got %v", test.expect, got)
		}
	}
}
