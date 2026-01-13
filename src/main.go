package main

import (
	"fmt"
	"strings"
)

// NormalizePath takes a unix-like path and returns a normalized absolute path.
// Rules:
// - Treat empty input as "/"
// - Always return an absolute path starting with "/"
// - Collapse repeated "/" and remove "." segments
// - Resolve ".." segments but do not allow escaping above root (returns error)
func NormalizePath(p string) (string, error) {
	p = strings.TrimSpace(p)
	if p == "" {
		return "/", nil
	}

	// Make it absolute-ish: treat relative as rooted.
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	parts := strings.Split(p, "/")
	stack := make([]string, 0, len(parts))

	for _, seg := range parts {
		if seg == "" || seg == "." {
			continue
		}
		if seg == ".." {
			if len(stack) == 0 {
				return "", fmt.Errorf("path escapes root: %q", p)
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, seg)
	}

	if len(stack) == 0 {
		return "/", nil
	}
	return "/" + strings.Join(stack, "/"), nil
}
