package main

import (
	"fmt"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "empty input",
			input:       "",
			expected:    "/",
			expectError: false,
		},
		{
			name:        "whitespace only",
			input:       "   ",
			expected:    "/",
			expectError: false,
		},
		{
			name:        "root path",
			input:       "/",
			expected:    "/",
			expectError: false,
		},
		{
			name:        "simple absolute path",
			input:       "/home/user",
			expected:    "/home/user",
			expectError: false,
		},
		{
			name:        "relative path becomes absolute",
			input:       "home/user",
			expected:    "/home/user",
			expectError: false,
		},
		{
			name:        "path with trailing slash",
			input:       "/home/user/",
			expected:    "/home/user",
			expectError: false,
		},
		{
			name:        "path with multiple slashes",
			input:       "/home//user///docs",
			expected:    "/home/user/docs",
			expectError: false,
		},
		{
			name:        "path with current directory segments",
			input:       "/home/./user/./docs",
			expected:    "/home/user/docs",
			expectError: false,
		},
		{
			name:        "path with mixed slashes and dots",
			input:       "/home//./user/././docs/",
			expected:    "/home/user/docs",
			expectError: false,
		},
		{
			name:        "path with parent directory segments",
			input:       "/home/user/../docs",
			expected:    "/home/docs",
			expectError: false,
		},
		{
			name:        "complex path with multiple parent segments",
			input:       "/home/user/project/../..//docs/../notes",
			expected:    "/home/notes",
			expectError: false,
		},
		{
			name:        "relative path with parent segments",
			input:       "user/../docs",
			expected:    "/docs",
			expectError: false,
		},
		{
			name:        "path escaping root - single parent",
			input:       "/..",
			expected:    "",
			expectError: true,
		},
		{
			name:        "path escaping root - multiple parents",
			input:       "/../../etc/passwd",
			expected:    "",
			expectError: true,
		},
		{
			name:        "relative path escaping root",
			input:       "..",
			expected:    "",
			expectError: true,
		},
		{
			name:        "complex escape attempt",
			input:       "user/../../etc/passwd",
			expected:    "",
			expectError: true,
		},
		{
			name:        "path with only parent segments",
			input:       "../..",
			expected:    "",
			expectError: true,
		},
		{
			name:        "whitespace around valid path",
			input:       "  /home/user  ",
			expected:    "/home/user",
			expectError: false,
		},
		{
			name:        "whitespace around relative path",
			input:       "  home/user  ",
			expected:    "/home/user",
			expectError: false,
		},
		{
			name:        "whitespace around escape attempt",
			input:       "  ..  ",
			expected:    "",
			expectError: true,
		},
		{
			name:        "single dot segment",
			input:       ".",
			expected:    "/",
			expectError: false,
		},
		{
			name:        "multiple dot segments",
			input:       "././.",
			expected:    "/",
			expectError: false,
		},
		{
			name:        "dot segments with valid path",
			input:       "./home/./user/.",
			expected:    "/home/user",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizePath(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("NormalizePath(%q) expected error but got none, result: %q", tt.input, result)
				}
				// Verify that when there's an error, result is empty string as per function spec
				if result != "" {
					t.Errorf("NormalizePath(%q) expected empty result on error, got: %q", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("NormalizePath(%q) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("NormalizePath(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			}
		})
	}
}

// BenchmarkNormalizePath measures the performance of NormalizePath with various inputs
func BenchmarkNormalizePath(b *testing.B) {
	testCases := []string{
		"",
		"/",
		"/home/user/documents/file.txt",
		"/home//user///./../documents//file.txt",
		"relative/path/with/parent/../segments",
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("input_%s", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NormalizePath(tc)
			}
		})
	}
}
