package main

import (
	"fmt"
	"testing"
)

func TestDoUrlsMatchDomain(t *testing.T) {
	tests := []struct {
		name           string
		rawBaseUrl     string
		rawCurrentUrl  string
		expectedResult bool
	}{
		{
			name:           "urls belong to the same domain",
			rawBaseUrl:     "https://zietlow.io",
			rawCurrentUrl:  "https://zietlow.io/blog",
			expectedResult: true,
		},
		{
			name:           "urls are from a different domain",
			rawBaseUrl:     "https://zietlow.io",
			rawCurrentUrl:  "https://zietlow.cloud/",
			expectedResult: false,
		},
		{
			name:           "url is not valid",
			rawBaseUrl:     "https://zietlow.io",
			rawCurrentUrl:  "#",
			expectedResult: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := doUrlsMatchDomain(tc.rawBaseUrl, tc.rawCurrentUrl)
			fmt.Printf("\nresult: %v", result)
			if result != tc.expectedResult {
				t.Errorf("Test %d - %s Fail: Unexpected result when comparing %s to %s", i, tc.name, tc.rawBaseUrl, tc.rawCurrentUrl)
			}
		})
	}
}
