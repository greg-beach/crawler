package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "get h1 basic",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "no h1 tag",
			inputBody: "<html><body></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test - %v - %s FAIL: expected: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "outside / inside paragraph",
			inputBody: `<html><body>
			<p>Outside paragraph.</p>
			<main>
				<p>Main paragraph.</p>
			</main>
		</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "just outside paragraph",
			inputBody: `<html><body>
			<p>Outside paragraph.</p>
			<main>
			</main>
		</body></html>`,
			expected: "Outside paragraph.",
		},
		{
			name: "no paragraph",
			inputBody: `<html><body>
			<main>
			</main>
		</body></html>`,
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test - %v - %s FAIL: expected: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "single URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<a href="https://blog.boot.dev"><span>Boot.dev</span></a>
			</body></html>`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "multiple URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<a href="https://blog.boot.dev"><span>Boot.dev</span></a>
			<a href="https://blog.boot.dev/example"><span>Boot.dev/example</span></a>
			</body></html>`,
			expected: []string{"https://blog.boot.dev", "https://blog.boot.dev/example"},
		},
		{
			name:     "partial URL in href",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<a href="/example"><span>Boot.dev/example</span></a>
			</body></html>`,
			expected: []string{"https://blog.boot.dev/example"},
		},
		{
			name:     "additional anchor tag with no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<a name="section1"><span>Boot.dev</span></a>
			<a href="https://blog.boot.dev"><span>Boot.dev</span></a>
			</body></html>`,
			expected: []string{"https://blog.boot.dev"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test - %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test - %v - %s FAIL: expected: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "single image",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<img src="/logo.png" alt="Logo">
			</body></html>`,
			expected: []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:     "multiple images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<img src="/logo.png" alt="Logo">
			<img src="/cat&dog.png" alt="cat and dog">
			</body></html>`,
			expected: []string{"https://blog.boot.dev/logo.png", "https://blog.boot.dev/cat&dog.png"},
		},
		{
			name:     "no Image",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<a href="/example"><span>Boot.dev/example</span></a>
			</body></html>`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test - %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test - %v - %s FAIL: expected: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
