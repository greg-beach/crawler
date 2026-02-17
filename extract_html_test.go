package main 

import (
	"testing"	
)

func TestGetH1FromHTMLBasic(t *testing.T) {
	tests := []struct {
		name string
		inputBody string
		expected string
	}{
		{
		name: "get h1 basic",
		inputBody: "<html><body><h1>Test Title</h1></body></html>",
		expected: "Test Title",

		},
		{
		name: "no h1 tag",
		inputBody:  "<html><body></body></html>",
		expected: "",
		},
	}

	for i, tc := range(tests) {
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
		name string
		inputBody string
		expected string
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

	for i, tc := range(tests) {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test - %v - %s FAIL: expected: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}