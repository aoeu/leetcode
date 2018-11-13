package leetcode

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	testCases := []struct {
		emailAddresses []string
		expected       int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
	}

	for _, tc := range testCases {
		if n := numUniqueEmails(tc.emailAddresses); n != tc.expected {
			t.Errorf("Expected %v but received %v for %v", tc.expected, n, tc.emailAddresses)
		}
	}
}