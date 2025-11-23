package iteration

import "strings"

// Repeat returns `character` repeated `count` times.
//
// This version is optimized for performance:
//   - early exit for trivial cases
//   - preallocated strings.Builder capacity
//   - fast path for single-byte strings
func Repeat(character string, count int) string {
	if count <= 0 || len(character) == 0 {
		return ""
	}

	clen := len(character)

	var b strings.Builder
	b.Grow(count * clen)

	// Fast path: single-byte character
	if clen == 1 {
		ch := character[0]
		for i := 0; i < count; i++ {
			b.WriteByte(ch)
		}
		return b.String()
	}

	// General path: multi-byte string
	for i := 0; i < count; i++ {
		b.WriteString(character)
	}

	return b.String()
}
