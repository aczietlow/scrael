package main

import (
	"net/url"
	"reflect"
	"testing"
)

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
		expectedHeadings string
	}{
		{
			name:             "single heading",
			inputHtml:        page1,
			expectedHeadings: "Welcome to Boot.dev",
		},
		{
			name:             "no headings",
			inputHtml:        "<html><p>hi mom</p></html>",
			expectedHeadings: "",
		},
		{
			name:             "single heading",
			inputHtml:        "<h1>lonely heading</h1>",
			expectedHeadings: "lonely heading",
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
			}
		})

	}

}

func TestGetFirstaragraphFromHtml(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`

	tests := []struct {
		name                  string
		inputHtml             string
		expectedParagraphText string
	}{
		{
			name:                  "find first paragraph",
			inputHtml:             inputBody,
			expectedParagraphText: "Outside paragraph.",
		},
		{
			name:                  "no paragraphs",
			inputHtml:             "<html><h1>hi mom</h1></html>",
			expectedParagraphText: "",
		},
		{
			name:                  "single paragraph node",
			inputHtml:             "<p>lonely paragraph</p>",
			expectedParagraphText: "lonely paragraph",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputHtml)
			if err != nil {
				t.Errorf("Test %d - %s Fail: Unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expectedParagraphText {
				t.Errorf("Test %d - %s Fail: paragraph text did not match expected result: expected: %s Actual: %s", i, tc.name, tc.expectedParagraphText, actual)
			}
		})
	}
}

func TestGetURLsFromHtml(t *testing.T) {
	baseUrl, err := url.Parse("https://zietlow.io")
	if err != nil {
		t.Errorf("Failed parsing Url: %v", baseUrl)
	}
	tests := []struct {
		name         string
		inputHtml    string
		expectedUrls []string
	}{
		{
			name:         "find single url",
			inputHtml:    `<a href="https://zietlow.io/foo">bar</a>`,
			expectedUrls: []string{"https://zietlow.io/foo"},
		},
		{
			name: "multiple urls",
			inputHtml: `
			<body>
				<a href="/foo">bar</a>
				<p>some stuff happens</p>"
				<a href="https://zietlow.io/boo">ghosts</a>
			</body>`,
			expectedUrls: []string{"https://zietlow.io/foo", "https://zietlow.io/boo"},
		},
		{
			name:         "no urls present",
			inputHtml:    "<p>this is just a paragraph</p>",
			expectedUrls: []string{},
		},
		{
			name:         "Url with a fragment",
			inputHtml:    `<a href="#foo">Fragment Url</a>`,
			expectedUrls: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHtml(tc.inputHtml, baseUrl)
			if err != nil {
				t.Errorf("Test %d - %s Fail: Unexpected error: %v", i, tc.name, err)
				return
			}
			if len(actual) != len(tc.expectedUrls) {
				t.Errorf("Test %d - %s Fail: Returned a different number of URLs. Expected - %d Actual - %d", i, tc.name, len(tc.expectedUrls), len(actual))
				return
			}
			if !reflect.DeepEqual(actual, tc.expectedUrls) {
				t.Errorf("expected %v, got %v", tc.expectedUrls, actual)
			}
		})
	}
}

func TestGetImagesFromHtml(t *testing.T) {
	baseUrl, err := url.Parse("https://zietlow.io")
	if err != nil {
		t.Errorf("Failed parsing Url: %v", baseUrl)
	}

	tests := []struct {
		name      string
		inputHtml string
		expected  []string
	}{
		{
			name:      "fetch single image",
			inputHtml: `<html><body><img src="/logo.png" alt="logo"></body></html>`,
			expected:  []string{"https://zietlow.io/logo.png"},
		},
		{
			name: "multiple imgs",
			inputHtml: `
			<body>
				<img src="/logo.png" alt="logo">
				<img src="https://zietlow.io/face.png" alt="face">
				<img alt="false flag">
			</body>`,
			expected: []string{"https://zietlow.io/logo.png", "https://zietlow.io/face.png"},
		},
		{
			name:      "no imgs present",
			inputHtml: "<p>this is just a paragraph</p>",
			expected:  []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getImagesFromHtml(tc.inputHtml, baseUrl)
			if err != nil {
				t.Errorf("Test %d -  %s Fail: Unexpected Error: %v", i, tc.name, err)
			}
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %d - %s Fail: Returned a different number of URLs. Expected - %d Actual - %d", i, tc.name, len(tc.expected), len(actual))
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
