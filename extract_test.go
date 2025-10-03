package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {

	inputURL := "https://zietlow.io"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	tests := []struct {
		name      string
		inputHtml string
		inputUrl  string
		expected  PageData
	}{
		{
			name:      "single values",
			inputHtml: inputBody,
			inputUrl:  inputURL,
			expected: PageData{
				Url:            "https://zietlow.io",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				outgoingLinks:  []string{"https://zietlow.io/link1"},
				imageUrl:       []string{"https://zietlow.io/image1.jpg"},
			},
		},
		{
			name: "multiple values",
			inputHtml: `<html>
			<body>
				<h1>Test Title</h1>
				<h1>wrong sub heading</h1>
				<p>This is the first paragraph.</p>
				<p>This is the second paragraph.</p>
				<a href="/link1">Link 1</a>
				<a href="/link2">Link 2</a>
				<img src="/image1.jpg" alt="Image 1">
				<img src="/image2.jpg" alt="Image 2">
			</body>
			</html>`,
			inputUrl: inputURL,
			expected: PageData{
				Url:            "https://zietlow.io",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				outgoingLinks:  []string{"https://zietlow.io/link1", "https://zietlow.io/link2"},
				imageUrl:       []string{"https://zietlow.io/image1.jpg", "https://zietlow.io/image2.jpg"},
			},
		},
		{
			name:      "subpage values",
			inputHtml: inputBody,
			inputUrl:  "https://zietlow.io/blog/",
			expected: PageData{
				Url:            "https://zietlow.io/blog/",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				outgoingLinks:  []string{"https://zietlow.io/link1"},
				imageUrl:       []string{"https://zietlow.io/image1.jpg"},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputHtml, tc.inputUrl)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Failure: %d - %s \nexpected: %+v\n actual: %+v\n", i, tc.name, tc.expected, actual)
			}
		})
	}
}
