package main

import "testing"

func TestGetH1FromHtml(t *testing.T) {

	page1 := `<html>
  <body>
    <h1>Welcome to Boot.dev</h1>
    <main>
      <p>Learn to code by building real projects.</p>
      <p>This is the second paragraph.</p>
    </main>
  </body>
</html>`

	tests := []struct {
		name             string
		inputHtml        string
		expectedHeadings []string
	}{
		{
			name:      "single heading",
			inputHtml: page1,
			expectedHeadings: []string{
				"Welcome to Boot.dev",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getH1FromHtml(tc.inputHtml)
			if err != nil {
				t.Errorf("Test %d - %s Fail: Unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expectedHeadings {
				t.Errorf("Test %d - %s Fail: Didn't find expected headings\nexpected: %v\nactual: %v", i, tc.name, tc.expectedHeadings, actual)
				return
			}
		})

	}

}

